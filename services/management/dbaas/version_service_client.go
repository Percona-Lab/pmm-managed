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

package dbaas

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	goversion "github.com/hashicorp/go-version"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"

	"github.com/percona/pmm-managed/utils/irt"
)

const (
	psmdbOperator = "psmdb-operator"
	pxcOperator   = "pxc-operator"
	zeroVersion   = "0.0.0"
)

// ErrZeroLatestVersion indicates that there is no version greater than zero version, hence we failed to get latest version.
var ErrZeroLatestVersion error = errors.New("no version greater than zero version")

// componentVersion contains info about exact component version.
type componentVersion struct {
	ImagePath string `json:"imagePath"`
	ImageHash string `json:"imageHash"`
	Status    string `json:"status"`
	Critical  bool   `json:"critical"`
}

type matrix struct {
	Mongod       map[string]componentVersion `json:"mongod"`
	Pxc          map[string]componentVersion `json:"pxc"`
	Pmm          map[string]componentVersion `json:"pmm"`
	Proxysql     map[string]componentVersion `json:"proxysql"`
	Haproxy      map[string]componentVersion `json:"haproxy"`
	Backup       map[string]componentVersion `json:"backup"`
	Operator     map[string]componentVersion `json:"operator"`
	LogCollector map[string]componentVersion `json:"logCollector"`
}

// LatestVersions contains latest version of PXC and PSMDB operator.
type LatestVersions struct {
	PSMDBOperator string `json:"psmdbOperator"`
	PXCOperator   string `json:"pxcOperator"`
}

// VersionServiceResponse represents response from version service API.
type VersionServiceResponse struct {
	Versions []struct {
		Product  string `json:"product"`
		Operator string `json:"operator"`
		Matrix   matrix `json:"matrix"`
	} `json:"versions"`
}

// componentsParams contains params to filter components in version service API.
type componentsParams struct {
	operator        string
	operatorVersion string
	dbVersion       string
}

// VersionServiceClient represents a client for Version Service API.
type VersionServiceClient struct {
	url  string
	http *http.Client
	irtm prom.Collector
}

// NewVersionServiceClient creates a new client for given version service URL.
func NewVersionServiceClient(url string) *VersionServiceClient {
	var t http.RoundTripper = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   3 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          50,
		IdleConnTimeout:       90 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	if logrus.GetLevel() >= logrus.TraceLevel {
		t = irt.WithLogger(t, logrus.WithField("component", "versionService/client").Tracef)
	}
	t, irtm := irt.WithMetrics(t, "versionService_client")

	return &VersionServiceClient{
		url: url,
		http: &http.Client{
			Transport: t,
		},
		irtm: irtm,
	}
}

// Describe implements prometheus.Collector.
func (c *VersionServiceClient) Describe(ch chan<- *prom.Desc) {
	c.irtm.Describe(ch)
}

// Collect implements prometheus.Collector.
func (c *VersionServiceClient) Collect(ch chan<- prom.Metric) {
	c.irtm.Collect(ch)
}

// Matrix calls version service with given params and returns components matrix.
func (c *VersionServiceClient) Matrix(ctx context.Context, params componentsParams) (*VersionServiceResponse, error) {
	paths := []string{c.url, params.operator}
	if params.operatorVersion != "" {
		paths = append(paths, params.operatorVersion)
		if params.dbVersion != "" {
			paths = append(paths, params.dbVersion)
		}
	}
	url := strings.Join(paths, "/")
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var vsResponse VersionServiceResponse
	err = json.Unmarshal(body, &vsResponse)
	if err != nil {
		return nil, err
	}

	return &vsResponse, nil
}

// GetLatestOperatorVersion returns the latest operator version for a given operator type.
func (c *VersionServiceClient) GetLatestOperatorVersion(ctx context.Context, operatorType string) (string, error) {
	versions, err := c.Matrix(ctx, componentsParams{operator: operatorType})
	if err != nil {
		return "", err
	}
	latestVersion, _ := goversion.NewVersion(zeroVersion)
	for _, version := range versions.Versions {
		operatorVersion, err := goversion.NewVersion(version.Operator)
		if err != nil {
			return "", err
		}
		if operatorVersion.GreaterThan(latestVersion) {
			latestVersion = operatorVersion
		}
	}
	if latestVersion.String() == zeroVersion {
		return "", ErrZeroLatestVersion
	}
	return latestVersion.String(), nil
}
