// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package grafana provides facilities for working with Grafana.
package grafana

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm-managed/utils/irt"
	"github.com/percona/pmm-managed/utils/logger"
)

// Client represents a client for Grafana API.
type Client struct {
	addr string
	l    *logrus.Entry
	http *http.Client
	irtm prometheus.Collector
}

// NewClient creates a new client for given Grafana address.
func NewClient(addr string) *Client {
	var t http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	l := logrus.WithField("component", "grafana/client")
	if l.Logger.GetLevel() >= logrus.TraceLevel {
		t = irt.WithLogger(t, l.Tracef)
	}
	t, irtm := irt.WithMetrics(t, "grafana_client")

	return &Client{
		addr: addr,
		l:    l,
		http: &http.Client{
			Transport: t,
		},
		irtm: irtm,
	}
}

// Describe implements prometheus.Collector.
func (c *Client) Describe(ch chan<- *prometheus.Desc) {
	c.irtm.Describe(ch)
}

// Collect implements prometheus.Collector.
func (c *Client) Collect(ch chan<- prometheus.Metric) {
	c.irtm.Collect(ch)
}

type authError struct {
	code    int
	headers http.Header
}

func (a *authError) Error() string {
	return fmt.Sprintf("status code %d", a.code)
}

func (c *Client) get(ctx context.Context, path string, authHeaders http.Header, respBody interface{}) error {
	u := url.URL{
		Scheme: "http",
		Host:   c.addr,
		Path:   path,
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return errors.WithStack(err)
	}
	for k := range authHeaders {
		req.Header.Set(k, authHeaders.Get(k))
	}

	req = req.WithContext(ctx)
	resp, err := c.http.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close() //nolint:errcheck

	if resp.StatusCode != 200 {
		return &authError{
			code:    resp.StatusCode,
			headers: resp.Header,
		}
	}

	if err = json.NewDecoder(resp.Body).Decode(respBody); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// isGrafanaAdmin returns true if current authenticated user is Grafana (super) admin.
func (c *Client) isGrafanaAdmin(ctx context.Context, authHeaders http.Header) (bool, error) {
	// https://grafana.com/docs/http_api/user/#actual-user - work with any authentication

	var m map[string]interface{}
	if err := c.get(ctx, "/api/user", authHeaders, &m); err != nil {
		return false, err
	}

	a, _ := m["isGrafanaAdmin"].(bool)
	return a, nil
}

type role string

const (
	none   role = "none"
	viewer role = "viewer"
	editor role = "editor"
	admin  role = "admin"
)

func (c *Client) getRole(ctx context.Context, authHeaders http.Header) (role, error) {
	// https://grafana.com/docs/http_api/user/#organizations-of-the-actual-user

	var s []interface{}
	if err := c.get(ctx, "/api/user/orgs", authHeaders, &s); err != nil {
		return none, err
	}

	for _, el := range s {
		m, _ := el.(map[string]interface{})
		if m == nil {
			continue
		}

		// check only default organization (with ID 1)
		if id, _ := m["orgId"].(float64); id == 1 {
			role, _ := m["role"].(string)
			switch strings.ToLower(role) {
			case "viewer":
				return viewer, nil
			case "editor":
				return editor, nil
			case "admin":
				return admin, nil
			default:
				return none, nil
			}
		}
	}

	return none, nil
}

type annotation struct {
	Time time.Time `json:"-"`
	Tags []string  `json:"tags,omitempty"`
	Text string    `json:"text,omitempty"`

	TimeInt int64 `json:"time,omitempty"`
}

// encode annotation before sending request.
func (a *annotation) encode() {
	var t int64
	if !a.Time.IsZero() {
		t = a.Time.UnixNano() / int64(time.Millisecond)
	}
	a.TimeInt = t
}

// decode annotation after receiving response.
func (a *annotation) decode() {
	var t time.Time
	if a.TimeInt != 0 {
		t = time.Unix(0, a.TimeInt*int64(time.Millisecond))
	}
	a.Time = t
}

// CreateAnnotation creates annotation with given text and tags ("pmm_annotation" is added automatically)
// and returns Grafana's response text which is typically "Annotation added" or "Failed to save annotation".
func (c *Client) CreateAnnotation(ctx context.Context, tags []string, text string) (string, error) {
	// http://docs.grafana.org/http_api/annotations/#create-annotation

	request := &annotation{
		Tags: append([]string{"pmm_annotation"}, tags...),
		Text: text,
	}
	request.encode()
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return "", errors.Wrap(err, "failed to marhal request")
	}

	u := url.URL{
		Scheme: "http",
		Host:   c.addr,
		Path:   "/api/annotations",
	}
	resp, err := c.http.Post(u.String(), "application/json", &buf)
	if err != nil {
		return "", errors.Wrap(err, "failed to make request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Get(ctx).Warnf("Grafana responded with status %d.", resp.StatusCode)
	}

	var response struct {
		Message string `json:"message"`
	}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", errors.Wrap(err, "failed to decode JSON response")
	}
	return response.Message, nil
}

func (c *Client) findAnnotations(ctx context.Context, from, to time.Time) ([]annotation, error) {
	// http://docs.grafana.org/http_api/annotations/#find-annotations

	u := &url.URL{
		Scheme: "http",
		Host:   c.addr,
		Path:   "/api/annotations",
		RawQuery: url.Values{
			"from": []string{strconv.FormatInt(from.UnixNano()/int64(time.Millisecond), 10)},
			"to":   []string{strconv.FormatInt(to.UnixNano()/int64(time.Millisecond), 10)},
		}.Encode(),
	}
	resp, err := c.http.Get(u.String())
	if err != nil {
		return nil, errors.Wrap(err, "failed to make request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Get(ctx).Warnf("Grafana responded with status %d.", resp.StatusCode)
	}

	var response []annotation
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "failed to decode JSON response")
	}
	for i, r := range response {
		r.decode()
		response[i] = r
	}
	return response, nil
}

// check interfaces
var (
	_ prometheus.Collector = (*Client)(nil)
)
