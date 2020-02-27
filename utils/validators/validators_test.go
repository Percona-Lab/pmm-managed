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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEnvVarValidator(t *testing.T) {
	t.Run("Valid env variables", func(t *testing.T) {
		envs := []string{
			"DISABLE_UPDATES=True",
			"DISABLE_TELEMETRY=False",
			"METRICS_RESOLUTION=5m",
			"METRICS_RESOLUTION_MR=5s",
			"METRICS_RESOLUTION_LR=1h",
			"DATA_RETENTION=72h",
		}
		expectedEnvVars := EnvSettings{
			DataRetention:    72 * time.Hour,
			DisableTelemetry: false,
			DisableUpdates:   true,
			MetricsResolutions: MetricsResolutions{
				HR: 5 * time.Minute,
				MR: 5 * time.Second,
				LR: time.Hour,
			},
		}

		gotEnvVars, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotEnvVars, expectedEnvVars)
		assert.Nil(t, gotErrs)
		assert.Nil(t, gotWarns)
	})

	t.Run("Unknown env variables", func(t *testing.T) {
		envs := []string{"UNKNOWN_VAR=VAL", "ANOTHER_UNKNOWN_VAR=VAL"}
		expectedEnvVars := EnvSettings{}
		expectedWarns := []string{
			`unknown environment variable "UNKNOWN_VAR=VAL"`,
			`unknown environment variable "ANOTHER_UNKNOWN_VAR=VAL"`,
		}

		gotEnvVars, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotEnvVars, expectedEnvVars)
		assert.Nil(t, gotErrs)
		assert.Equal(t, expectedWarns, gotWarns)
	})

	t.Run("Default env vars", func(t *testing.T) {
		envs := []string{
			"PATH=/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
			"HOSTNAME=host",
			"TERM=xterm-256color",
			"HOME=/home/user/",
		}
		expectedEnvVars := EnvSettings{}

		gotEnvVars, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotEnvVars, expectedEnvVars)
		assert.Nil(t, gotErrs)
		assert.Nil(t, gotWarns)
	})

	t.Run("Invalid env variables values", func(t *testing.T) {
		envs := []string{
			"DISABLE_UPDATES=5",
			"DISABLE_TELEMETRY=X",
			"METRICS_RESOLUTION=5f",
			"METRICS_RESOLUTION_MR=s5",
			"METRICS_RESOLUTION_LR=1hour",
			"DATA_RETENTION=keep one week",
		}
		expectedEnvVars := EnvSettings{}
		expectedErrs := []error{
			fmt.Errorf("invalid environment variable %q", "DISABLE_UPDATES=5"),
			fmt.Errorf("invalid environment variable %q", "DISABLE_TELEMETRY=X"),
			fmt.Errorf("environment variable %q has invalid duration 5f", "METRICS_RESOLUTION=5f"),
			fmt.Errorf("environment variable %q has invalid duration s5", "METRICS_RESOLUTION_MR=s5"),
			fmt.Errorf("environment variable %q has invalid duration 1hour", "METRICS_RESOLUTION_LR=1hour"),
			fmt.Errorf("environment variable %q has invalid duration keep one week", "DATA_RETENTION=keep one week"),
		}

		gotEnvVars, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotEnvVars, expectedEnvVars)
		assert.Equal(t, gotErrs, expectedErrs)
		assert.Nil(t, gotWarns)
	})

	t.Run("Grafana env vars", func(t *testing.T) {
		envs := []string{
			`GF_AUTH_GENERIC_OAUTH_ALLOWED_DOMAINS='example.com'`,
			`GF_AUTH_GENERIC_OAUTH_ENABLED='true'`,
			`GF_PATHS_CONFIG="/etc/grafana/grafana.ini"`,
			`GF_PATHS_DATA="/var/lib/grafana"`,
			`GF_PATHS_HOME="/usr/share/grafana"`,
			`GF_PATHS_LOGS="/var/log/grafana"`,
			`GF_PATHS_PLUGINS="/var/lib/grafana/plugins"`,
			`GF_PATHS_PROVISIONING="/etc/grafana/provisioning"`,
		}
		expectedEnvVars := EnvSettings{}

		gotEnvVars, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotEnvVars, expectedEnvVars)
		assert.Nil(t, gotErrs)
		assert.Nil(t, gotWarns)
	})

	t.Run("Data retention less then a day", func(t *testing.T) {
		envs := []string{
			"DATA_RETENTION=1h",
		}
		expectedErrs := []error{
			fmt.Errorf("environment variable %q cannot be less then 24h0m0s", "DATA_RETENTION=1h"),
		}
		_, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotErrs, expectedErrs)
		assert.Nil(t, gotWarns)
	})

	t.Run("Data retention is not a natural number of days", func(t *testing.T) {
		envs := []string{
			"DATA_RETENTION=30h",
		}
		expectedErrs := []error{
			fmt.Errorf("environment variable %q should be a multiple of 24h0m0s", "DATA_RETENTION=30h"),
		}
		_, gotErrs, gotWarns := ValidateEnvVars(envs)
		assert.Equal(t, gotErrs, expectedErrs)
		assert.Nil(t, gotWarns)
	})
}
