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

package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

// rules maps original URL prefix to minimal required role.
var rules = map[string]role{
	// TODO https://jira.percona.com/browse/PMM-4420
	"/agent.Agent/Connect": none,

	"/inventory.":                     admin,
	"/management.":                    admin,
	"/management.Actions/":            viewer,
	"/server.Server/CheckUpdates":     viewer,
	"/server.Server/UpdateStatus":     none, // special token-based auth
	"/server.Server/AWSInstanceCheck": none, // special case for mustSetup
	"/server.":                        admin,

	"/v1/inventory/":          admin,
	"/v1/management/":         admin,
	"/v1/management/Actions/": viewer,
	"/v1/Updates/Check":       viewer,
	"/v1/Updates/Status":      none, // special token-based auth
	"/v1/AWSInstanceCheck":    none, // special case for mustSetup
	"/v1/Updates/":            admin,
	"/v1/Settings/":           admin,

	// must be available without authentication for health checking
	"/v1/readyz": none,
	"/ping":      none, // PMM 1.x variant

	// must not be available without authentication as it can leak data
	"/v1/version":         viewer,
	"/managed/v1/version": viewer, // PMM 1.x variant

	"/v0/qan/": viewer,

	"/prometheus/": admin,
	"/graph/":      none,
	"/qan/":        none,
	"/swagger/":    none,

	// "/auth_request" and "/setup" have auth_request disabled in nginx config

	// "/" is a special case in this code
}

// nginx auth_request directive supports only 401 and 403 - every other code results in 500.
// Our APIs can return codes.PermissionDenied which maps to 403 / http.StatusForbidden.
// Our APIs MUST NOT return codes.Unauthenticated which maps to 401 / http.StatusUnauthorized
// as this code is reserved for auth_request.
const authenticationErrorCode = 401

// clientError contains authentication error response details.
type authError struct {
	code    codes.Code // error code for API client; not mapped to HTTP status code
	message string
}

// AuthServer authenticates incoming requests via Grafana API.
type AuthServer struct {
	c       *Client
	checker awsInstanceChecker
	l       *logrus.Entry

	// TODO server metrics should be provided by middleware https://jira.percona.com/browse/PMM-4326
}

// NewAuthServer creates new AuthServer.
func NewAuthServer(c *Client, checker awsInstanceChecker) *AuthServer {
	return &AuthServer{
		c:       c,
		checker: checker,
		l:       logrus.WithField("component", "grafana/auth"),
	}
}

// ServeHTTP serves internal location /auth_request for both authentication subrequests
// and subsequent normal requests.
func (s *AuthServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if s.l.Logger.GetLevel() >= logrus.DebugLevel {
		b, err := httputil.DumpRequest(req, true)
		if err != nil {
			s.l.Errorf("Failed to dump request: %v.", err)
		}
		s.l.Debugf("Request:\n%s", b)
	}

	origMethod, origURI := req.Header.Get("X-Original-Method"), req.Header.Get("X-Original-Uri")
	if origMethod == "" {
		s.l.Panic("X-Original-Method")
	}
	if origURI == "" {
		s.l.Panic("Empty X-Original-Uri.")
	}
	req.Method = origMethod
	req.URL.Path = origURI
	l := s.l.WithField("req", fmt.Sprintf("%s %s", origMethod, origURI))

	// TODO l := logger.Get(ctx) once we have it after https://jira.percona.com/browse/PMM-4326

	if s.mustSetup(rw, req, l) {
		return
	}

	// fail-safe
	ctx, cancel := context.WithTimeout(req.Context(), 3*time.Second)
	defer cancel()

	if err := s.authenticate(ctx, req, l); err != nil {
		// nginx completely ignores auth_request subrequest response body.
		// We respond with 401 (authenticationErrorCode); our nginx configuration then sends
		// the same request as a normal request to the same location and returns response body to the client.

		// copy grpc-gateway behavior: set correct codes, set both "error" and "message"
		m := map[string]interface{}{
			"code":    int(err.code),
			"error":   err.message,
			"message": err.message,
		}
		rw.Header().Set("Content-Type", "application/json")

		rw.WriteHeader(authenticationErrorCode)
		if err := json.NewEncoder(rw).Encode(m); err != nil {
			l.Warnf("%s", err)
		}
	}
}

// mustSetup returns true if AWS instance ID must be checked.
func (s *AuthServer) mustSetup(rw http.ResponseWriter, req *http.Request, l *logrus.Entry) bool {
	// /setup page uses this API.
	if req.URL.Path == "/server.Server/AWSInstanceCheck" || req.URL.Path == "/v1/AWSInstanceCheck" {
		return false
	}

	// This header is used to pass information that setup is required from auth_request subrequest
	// to normal request to return redirect with location - something that auth_request can't do.
	const mustSetupHeader = "X-Must-Setup"

	// Redirect to /setup page.
	if req.Header.Get(mustSetupHeader) != "" {
		const redirectCode = 303 // temporary, not cacheable, always GET
		l.Warnf("AWS instance ID must be checked, returning %d with Location.", redirectCode)
		rw.Header().Set("Location", "/setup")
		rw.WriteHeader(redirectCode)
		return true
	}

	// Use X-Test-Must-Setup header for testing.
	// There is no way to skip check, only to enforce it.
	mustCheck := s.checker.MustCheck()
	if req.Header.Get("X-Test-Must-Setup") != "" {
		l.Debug("X-Test-Must-Setup is present, enforcing AWS instance ID check.")
		mustCheck = true
	}

	if mustCheck {
		l.Warnf("AWS instance ID must be checked, returning %d with %s.", authenticationErrorCode, mustSetupHeader)
		rw.Header().Set(mustSetupHeader, "1") // any non-empty value is ok
		rw.WriteHeader(authenticationErrorCode)
		return true
	}

	return false
}

// nextPrefix returns path's prefix, stopping on slashes and dots:
// /foo.Bar/Baz -> /foo.Bar/ -> /foo. -> /
// That works for both gRPC and JSON URLs.
func nextPrefix(path string) string {
	path = strings.TrimRight(path, "/.")
	if i := strings.LastIndexAny(path, "/."); i != -1 {
		return path[:i+1]
	}
	return path
}

func (s *AuthServer) authenticate(ctx context.Context, req *http.Request, l *logrus.Entry) *authError {
	// find the longest prefix present in rules, stopping on slashes and dots:
	// /foo.Bar/Baz -> /foo.Bar/ -> /foo. -> /
	prefix := req.URL.Path
	for prefix != "/" {
		if _, ok := rules[prefix]; ok {
			break
		}
		prefix = nextPrefix(prefix)
	}

	// fallback to Grafana admin if there is no explicit rule
	minRole, ok := rules[prefix]
	if ok {
		l = l.WithField("prefix", prefix)
	} else {
		l.Warn("No explicit rule, falling back to Grafana admin.")
		minRole = grafanaAdmin
	}

	if minRole == none {
		l.Debugf("Minimal required role is %q, granting access without checking Grafana.", minRole)
		return nil
	}

	// check Grafana with some headers from request
	authHeaders := make(http.Header)
	for _, k := range []string{
		"Authorization",
		"Cookie",
	} {
		if v := req.Header.Get(k); v != "" {
			authHeaders.Set(k, v)
		}
	}
	role, err := s.c.getRole(ctx, authHeaders)
	if err != nil {
		l.Warnf("%s", err)
		if cErr, ok := errors.Cause(err).(*clientError); ok {
			code := codes.Internal
			if cErr.Code == 401 || cErr.Code == 403 {
				code = codes.Unauthenticated
			}
			return &authError{code: code, message: cErr.ErrorMessage}
		}
		return &authError{code: codes.Internal, message: "Internal server error."}
	}
	l = l.WithField("role", role.String())

	if role == grafanaAdmin {
		l.Debugf("Grafana admin, allowing access.")
		return nil
	}

	if minRole <= role {
		l.Debugf("Minimal required role is %q, granting access.", minRole)
		return nil
	}

	l.Warnf("Minimal required role is %q.", minRole)
	return &authError{code: codes.PermissionDenied, message: "Access denied."}
}
