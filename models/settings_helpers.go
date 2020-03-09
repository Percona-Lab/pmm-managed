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

package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/utils/validators"
)

// GetSettings returns current PMM Server settings.
func GetSettings(q reform.DBTX) (*Settings, error) {
	var b []byte
	if err := q.QueryRow("SELECT settings FROM settings").Scan(&b); err != nil {
		return nil, errors.Wrap(err, "failed to select settings")
	}

	var s Settings
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal settings")
	}

	s.fillDefaults()
	return &s, nil
}

type ChangeSettingsParams struct {
	DisableTelemetry bool
	EnableTelemetry  bool

	MetricsResolutions MetricsResolutions

	DataRetention time.Duration

	AWSPartitions []string

	SSHKey string

	// not url.URL to keep username and password
	AlertManagerURL       string
	RemoveAlertManagerUrl bool
}

func UpdateSettings(q reform.DBTX, params ChangeSettingsParams) (*Settings, error) {
	err := ValidateSettings(params)
	if err != nil {
		return nil, err
	}
	settings, err := GetSettings(q)
	if err != nil {
		return nil, err
	}
	if params.DisableTelemetry {
		settings.Telemetry.Disabled = true
	}
	if params.EnableTelemetry {
		settings.Telemetry.Disabled = false
	}
	if params.MetricsResolutions.LR != 0 {
		settings.MetricsResolutions.LR = params.MetricsResolutions.LR
	}
	if params.MetricsResolutions.MR != 0 {
		settings.MetricsResolutions.MR = params.MetricsResolutions.MR
	}
	if params.MetricsResolutions.HR != 0 {
		settings.MetricsResolutions.HR = params.MetricsResolutions.HR
	}
	if params.DataRetention != 0 {
		settings.DataRetention = params.DataRetention
	}

	if len(params.AWSPartitions) != 0 {
		settings.AWSPartitions = deduplicateAWSPartitions(params.AWSPartitions)
	}
	if params.SSHKey != "" {
		settings.SSHKey = params.SSHKey
	}
	if params.AlertManagerURL != "" {
		settings.AlertManagerURL = params.AlertManagerURL
	}
	if params.RemoveAlertManagerUrl {
		settings.AlertManagerURL = ""
	}

	err = SaveSettings(q, settings)
	if err != nil {
		return nil, err
	}
	return settings, nil
}

func ValidateSettings(params ChangeSettingsParams) error {
	if params.EnableTelemetry && params.DisableTelemetry {
		return status.Error(codes.InvalidArgument, "Both enable_telemetry and disable_telemetry are present.")
	}

	checkCases := []struct {
		dur       time.Duration
		fieldName string
		validator func(time.Duration) (time.Duration, error)
	}{
		{params.MetricsResolutions.HR, "hr", validators.ValidateMetricResolution},
		{params.MetricsResolutions.MR, "mr", validators.ValidateMetricResolution},
		{params.MetricsResolutions.LR, "lr", validators.ValidateMetricResolution},
	}
	for _, v := range checkCases {
		if v.dur == 0 {
			continue
		}

		if _, err := v.validator(v.dur); err != nil {
			switch err.(type) {
			case validators.AliquotDurationError:
				return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: should be a natural number of seconds", v.fieldName))
			case validators.MinDurationError:
				return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: minimal resolution is 1s", v.fieldName))
			default:
				return status.Error(codes.InvalidArgument, fmt.Sprintf("%s: unknown error for", v.fieldName))
			}
		}
	}

	if params.DataRetention != 0 {
		if _, err := validators.ValidateDataRetention(params.DataRetention); err != nil {
			switch err.(type) {
			case validators.AliquotDurationError:
				return status.Error(codes.InvalidArgument, "data_retention: should be a natural number of days")
			case validators.MinDurationError:
				return status.Error(codes.InvalidArgument, "data_retention: minimal resolution is 24h")
			default:
				return status.Error(codes.InvalidArgument, "data_retention: unknown error")
			}
		}
	}

	var err error
	if err = validators.ValidateAWSPartitions(params.AWSPartitions); err != nil {
		return err
	}

	if params.AlertManagerURL != "" {
		if params.RemoveAlertManagerUrl {
			return status.Error(codes.InvalidArgument, "Both alert_manager_url and remove_alert_manager_url are present.")
		}

		// custom validation for typical error that is not handled well by url.Parse
		if !strings.Contains(params.AlertManagerURL, "//") {
			return status.Errorf(codes.InvalidArgument, "Invalid alert_manager_url: %s - missing protocol scheme.", params.AlertManagerURL)
		}
		u, err := url.Parse(params.AlertManagerURL)
		if err != nil {
			return status.Errorf(codes.InvalidArgument, "Invalid alert_manager_url: %s.", err)
		}
		if u.Host == "" {
			return status.Errorf(codes.InvalidArgument, "Invalid alert_manager_url: %s - missing host.", params.AlertManagerURL)
		}
	}

	return nil
}

// SaveSettings saves PMM Server settings.
// It may modify passed settings to fill defaults.
func SaveSettings(q reform.DBTX, s *Settings) error {
	s.fillDefaults()

	b, err := json.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to marshal settings")
	}

	_, err = q.Exec("UPDATE settings SET settings = $1", b)
	if err != nil {
		return errors.Wrap(err, "failed to update settings")
	}

	return nil
}

// deduplicateAWSPartitions deduplicates AWS partitions list.
func deduplicateAWSPartitions(partitions []string) []string {
	set := make(map[string]struct{})
	for _, p := range partitions {
		set[p] = struct{}{}
	}

	slice := make([]string, 0, len(set))
	for partition := range set {
		slice = append(slice, partition)
	}
	sort.Strings(slice)

	return slice
}
