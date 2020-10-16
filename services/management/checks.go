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

package management

import (
	"context"

	"github.com/percona/pmm/api/managementpb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/percona/pmm-managed/services"
)

// ChecksAPIService represents security checks service API.
type ChecksAPIService struct {
	checksService checksService
	l             *logrus.Entry
}

// NewChecksAPIService creates new Checks API Service.
func NewChecksAPIService(checksService checksService) *ChecksAPIService {
	return &ChecksAPIService{
		checksService: checksService,
		l:             logrus.WithField("component", "management/checks"),
	}
}

// StartSecurityChecks starts STT checks execution.
func (s *ChecksAPIService) StartSecurityChecks(ctx context.Context) (*managementpb.StartSecurityChecksResponse, error) {
	err := s.checksService.StartChecks(ctx)
	if err != nil {
		if err == services.ErrSTTDisabled {
			return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
		}

		return nil, errors.Wrap(err, "failed to start security checks")
	}

	return &managementpb.StartSecurityChecksResponse{}, nil
}

// GetSecurityCheckResults returns the results of the STT checks that were run.
func (s *ChecksAPIService) GetSecurityCheckResults() (*managementpb.GetSecurityCheckResultsResponse, error) {
	results, err := s.checksService.GetSecurityCheckResults()
	if err != nil {
		if err == services.ErrSTTDisabled {
			return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
		}

		return nil, errors.Wrap(err, "failed to get security check results")
	}

	checkResults := make([]*managementpb.SecurityCheckResult, 0, len(results))
	for _, result := range results {
		checkResults = append(checkResults, &managementpb.SecurityCheckResult{
			Summary:     result.Summary,
			Description: result.Description,
			Severity:    managementpb.Severity(result.Severity),
			Labels:      result.Labels,
		})
	}

	return &managementpb.GetSecurityCheckResultsResponse{Results: checkResults}, nil
}

// ListSecurityChecks returns all available STT checks.
func (s *ChecksAPIService) ListSecurityChecks() (*managementpb.ListSecurityChecksResponse, error) {
	disChecks, err := s.checksService.GetDisabledChecks()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get disabled checks list")
	}

	m := make(map[string]struct{}, len(disChecks))
	for _, c := range disChecks {
		m[c] = struct{}{}
	}

	checks := s.checksService.GetAllChecks()
	res := make([]*managementpb.SecurityCheck, 0, len(checks))
	for _, c := range checks {
		_, disabled := m[c.Name]
		res = append(res, &managementpb.SecurityCheck{Name: c.Name, Disabled: disabled})
	}

	return &managementpb.ListSecurityChecksResponse{Checks: res}, nil
}

// ChangeSecurityCheck allows to change STT checks state.
func (s *ChecksAPIService) ChangeSecurityCheck(req *managementpb.ChangeSecurityCheckRequest) (*managementpb.ChangeSecurityCheckResponse, error) {
	var enableChecks, disableChecks []string
	for _, check := range req.Params {
		if check.Enable && check.Disable {
			return nil, status.Errorf(codes.InvalidArgument, "Check %s has enable and disable parameters set to the true.", check.Name)
		}

		if check.Enable {
			enableChecks = append(enableChecks, check.Name)
		}

		if check.Disable {
			disableChecks = append(disableChecks, check.Name)
		}
	}

	err := s.checksService.EnableChecks(enableChecks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to enable disabled security checks")
	}

	err = s.checksService.DisableChecks(disableChecks)
	if err != nil {
		return nil, errors.Wrap(err, "failed to disable security checks")
	}

	return &managementpb.ChangeSecurityCheckResponse{}, nil
}
