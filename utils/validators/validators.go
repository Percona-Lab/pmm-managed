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

// Package validators contains environment variables validator.
package validators

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// MetricsResolutions contains standard Prometheus metrics resolutions.
type MetricsResolutions struct {
	HR time.Duration
	MR time.Duration
	LR time.Duration
}

// EnvSettings contains PMM Server settings.
type EnvSettings struct {
	DisableUpdates     bool
	DisableTelemetry   bool
	MetricsResolutions MetricsResolutions
	DataRetention      time.Duration
}

const (
	// MetricsResolutionMin is the smallest value metric resolution can accept.
	MetricsResolutionMin = time.Second
	// MetricsResolutionMultipleOf is value metrics resolution should be multiple of.
	MetricsResolutionMultipleOf = time.Second
	// DataRetentionMin is the smallest value data retention can accept.
	DataRetentionMin = 24 * time.Hour
	// DataRetentionMultipleOf is a value of data retention should be multiple of.
	DataRetentionMultipleOf = 24 * time.Hour
)

// InvalidDurationError invalid duration error.
type InvalidDurationError string

func (e InvalidDurationError) Error() string { return string(e) }

// MinDurationError minimum allowed duration error.
type MinDurationError string

func (e MinDurationError) Error() string { return string(e) }

// AliquotDurationError multiple of duration allowed error.
type AliquotDurationError string

func (e AliquotDurationError) Error() string { return string(e) }

// ValidateEnvVars validates given environment variables.
//
// Returns valid setting and two lists with errors and warnings.
// This function is mainly used in pmm-managed-init to early validate passed
// environment variables, and provide user warnings about unknown variables.
// In case of error, the docker run terminates.
// Short description of environment variables:
//  - PATH, HOSTNAME, TERM, HOME are default environment variables that will be ignored;
//  - DISABLE_UPDATES is a boolean flag to enable or disable pmm-server update;
//  - DISABLE_TELEMETRY is a boolean flag to enable or disable pmm telemetry;
//  - METRICS_RESOLUTION, METRICS_RESOLUTION, METRICS_RESOLUTION_HR,
// METRICS_RESOLUTION_LR are durations of metrics resolution;
//  - DATA_RETENTION is the duration of how long keep time-series data in ClickHouse;
//  - the environment variables prefixed with GF_ passed as related to Grafana.
func ValidateEnvVars(envs []string) (envSettings EnvSettings, errs []error, warns []string) {
	for _, env := range envs {
		p := strings.SplitN(env, "=", 2)
		if len(p) != 2 {
			errs = append(errs, fmt.Errorf("failed to parse environment variable %q", env))
			continue
		}

		var err error

		k, v := strings.ToUpper(p[0]), strings.ToLower(p[1])
		switch k {
		// Skip default environment variables.
		case "PATH", "HOSTNAME", "TERM", "HOME", "PWD", "SHLVL", "_":
		case "DISABLE_UPDATES":
			envSettings.DisableUpdates, err = strconv.ParseBool(v)
			if err != nil {
				err = fmt.Errorf("invalid environment variable %q", env)
			}
		case "DISABLE_TELEMETRY":
			envSettings.DisableTelemetry, err = strconv.ParseBool(v)
			if err != nil {
				err = fmt.Errorf("invalid environment variable %q", env)
			}
		case "METRICS_RESOLUTION", "METRICS_RESOLUTION_HR":
			if envSettings.MetricsResolutions.HR, err = ValidateStringMetricResolution(v); err != nil {
				err = formatEnvVariableError(err, env, v, MetricsResolutionMin, MetricsResolutionMultipleOf)
			}
		case "METRICS_RESOLUTION_MR":
			if envSettings.MetricsResolutions.MR, err = ValidateStringMetricResolution(v); err != nil {
				err = formatEnvVariableError(err, env, v, MetricsResolutionMin, MetricsResolutionMultipleOf)
			}
		case "METRICS_RESOLUTION_LR":
			if envSettings.MetricsResolutions.LR, err = ValidateStringMetricResolution(v); err != nil {
				err = formatEnvVariableError(err, env, v, MetricsResolutionMin, MetricsResolutionMultipleOf)
			}
		case "DATA_RETENTION":
			if envSettings.DataRetention, err = ValidateStringDataRetention(v); err != nil {
				err = formatEnvVariableError(err, env, v, DataRetentionMin, DataRetentionMultipleOf)
			}
		default:
			if !strings.HasPrefix(k, "GF_") {
				warns = append(warns, fmt.Sprintf("unknown environment variable %q", env))
			}
		}
		if err != nil {
			errs = append(errs, err)
		}
	}
	return envSettings, errs, warns
}

func formatEnvVariableError(err error, env, value string, min, multipleOf time.Duration) error {
	switch e := err.(type) {
	case InvalidDurationError:
		return fmt.Errorf("environment variable %q has invalid duration %s", env, value)
	case MinDurationError:
		return fmt.Errorf("environment variable %q cannot be less then %s", env, min)
	case AliquotDurationError:
		fmt.Printf("")
		return fmt.Errorf("environment variable %q should be a multiple of %s", env, multipleOf)
	default:
		return errors.Wrap(e, "unknown error")
	}
}

// ValidateStringDuration validate duration as string value.
func ValidateStringDuration(value string, min, multipleOf time.Duration) (time.Duration, error) {
	d, err := time.ParseDuration(value)
	if err != nil {
		return d, InvalidDurationError("invalid duration error")
	}

	return ValidateDuration(d, min, multipleOf)
}

// ValidateDuration validate duration.
func ValidateDuration(d, min, multipleOf time.Duration) (time.Duration, error) {
	if d < min {
		return d, MinDurationError("min duration error")
	}

	if d.Truncate(multipleOf) != d {
		return d, AliquotDurationError("aliquot	duration error")
	}
	return d, nil
}

// ValidateStringMetricResolution validate metric resolution.
func ValidateStringMetricResolution(value string) (time.Duration, error) {
	return ValidateStringDuration(value, MetricsResolutionMin, MetricsResolutionMultipleOf)
}

// ValidateMetricResolution validate metric resolution.
func ValidateMetricResolution(value time.Duration) (time.Duration, error) {
	return ValidateDuration(value, MetricsResolutionMin, MetricsResolutionMultipleOf)
}

// ValidateStringDataRetention validate metric resolution.
func ValidateStringDataRetention(value string) (time.Duration, error) {
	return ValidateStringDuration(value, DataRetentionMin, DataRetentionMultipleOf)
}

// ValidateDataRetention validate metric resolution.
func ValidateDataRetention(value time.Duration) (time.Duration, error) {
	return ValidateDuration(value, DataRetentionMin, DataRetentionMultipleOf)
}
