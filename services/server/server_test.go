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

package server

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/percona/pmm/api/serverpb"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/utils/testdb"
	"github.com/percona/pmm-managed/utils/tests"
)

func TestServer(t *testing.T) {
	sqlDB := testdb.Open(t, models.SkipFixtures, nil)
	defer func() {
		require.NoError(t, sqlDB.Close())
	}()
	r := new(mockSupervisordService)
	r.On("UpdateConfiguration", mock.Anything).Return(nil)

	mp := new(mockPrometheusService)
	mp.On("RequestConfigurationUpdate").Return(nil)

	newServer := func() *Server {
		s, err := NewServer(&ServerParams{
			DB:          reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf)),
			Supervisord: r,
			Prometheus:  mp,
		})
		require.NoError(t, err)
		return s
	}

	t.Run("UpdateSettingsFromEnv", func(t *testing.T) {
		t.Run("Typical", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DISABLE_UPDATES=true",
				"DISABLE_TELEMETRY=1",
				"METRICS_RESOLUTION_HR=1s",
				"METRICS_RESOLUTION_MR=2s",
				"METRICS_RESOLUTION_LR=3s",
				"DATA_RETENTION=240h",
			})
			require.Empty(t, errs)
			assert.Equal(t, true, s.envSettings.DisableUpdates)
			assert.Equal(t, true, s.envSettings.DisableTelemetry)
			assert.Equal(t, time.Second, s.envSettings.MetricsResolutions.HR)
			assert.Equal(t, 2*time.Second, s.envSettings.MetricsResolutions.MR)
			assert.Equal(t, 3*time.Second, s.envSettings.MetricsResolutions.LR)
			assert.Equal(t, 10*24*time.Hour, s.envSettings.DataRetention)
		})

		t.Run("Untypical", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DISABLE_TELEMETRY=TrUe",
				"METRICS_RESOLUTION=3S",
				"DATA_RETENTION=360H",
			})
			require.Empty(t, errs)
			assert.Equal(t, true, s.envSettings.DisableTelemetry)
			assert.Equal(t, 3*time.Second, s.envSettings.MetricsResolutions.HR)
			assert.Equal(t, 15*24*time.Hour, s.envSettings.DataRetention)
		})

		t.Run("NoValue", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DISABLE_TELEMETRY",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `failed to parse environment variable "DISABLE_TELEMETRY"`)
			assert.False(t, s.envSettings.DisableTelemetry)
		})

		t.Run("InvalidValue", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DISABLE_TELEMETRY=",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `invalid value "" for environment variable "DISABLE_TELEMETRY"`)
			assert.False(t, s.envSettings.DisableTelemetry)
		})

		t.Run("MetricsLessThenMin", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"METRICS_RESOLUTION=5ns",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `hr: minimal resolution is 1s`)
			assert.Zero(t, s.envSettings.MetricsResolutions.HR)
		})

		t.Run("DataRetentionLessThenMin", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DATA_RETENTION=12h",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `data_retention: minimal resolution is 24h`)
			assert.Zero(t, s.envSettings.DataRetention)
		})

		t.Run("Data retention is not a natural number of days", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DATA_RETENTION=30h",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `data_retention: should be a natural number of days`)
			assert.Zero(t, s.envSettings.DataRetention)
		})

		t.Run("Data retention without suffix", func(t *testing.T) {
			s := newServer()
			errs := s.UpdateSettingsFromEnv([]string{
				"DATA_RETENTION=30",
			})
			require.Len(t, errs, 1)
			require.EqualError(t, errs[0], `environment variable "DATA_RETENTION=30" has invalid duration 30`)
			assert.Zero(t, s.envSettings.DataRetention)
		})
	})

	t.Run("ValidateChangeSettingsRequest", func(t *testing.T) {
		s := newServer()

		ctx := context.TODO()

		tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Both alert_manager_rules and remove_alert_manager_rules are present."),
			s.validateChangeSettingsRequest(ctx, &serverpb.ChangeSettingsRequest{
				AlertManagerRules:       "something",
				RemoveAlertManagerRules: true,
			}))
	})

	t.Run("ValidateChangeSettingsSTT", func(t *testing.T) {
		s := newServer()

		ctx := context.TODO()

		tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Enable STT and disable STT cannot be both true"),
			s.validateChangeSettingsRequest(ctx, &serverpb.ChangeSettingsRequest{
				EnableStt:  true,
				DisableStt: true,
			}))

		tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Cannot enable STT while disabling telemetry"),
			s.validateChangeSettingsRequest(ctx, &serverpb.ChangeSettingsRequest{
				EnableStt:        true,
				DisableTelemetry: true,
			}))
	})

	t.Run("UpdateSettingsSTT", func(t *testing.T) {
		s := newServer()

		ctx := context.TODO()

		resp, err := s.ChangeSettings(ctx, &serverpb.ChangeSettingsRequest{
			EnableTelemetry: true,
		})
		assert.NoError(t, err)
		assert.True(t, resp.Settings.TelemetryEnabled)

		resp, err = s.ChangeSettings(ctx, &serverpb.ChangeSettingsRequest{
			EnableStt: true,
		})
		assert.NoError(t, err)
		assert.True(t, resp.Settings.SttEnabled)
	})

	t.Run("ValidateAlertManagerRules", func(t *testing.T) {
		s := newServer()

		t.Run("Valid", func(t *testing.T) {
			rules := strings.TrimSpace(`
groups:
- name: example
  rules:
  - alert: HighRequestLatency
    expr: job:request_latency_seconds:mean5m{job="myjob"} > 0.5
    for: 10m
    labels:
      severity: page
    annotations:
      summary: High request latency
			`) + "\n"
			err := s.validateAlertManagerRules(context.Background(), rules)
			assert.NoError(t, err)
		})

		t.Run("Zero", func(t *testing.T) {
			rules := strings.TrimSpace(`
groups:
- name: example
rules:
- alert: HighRequestLatency
expr: job:request_latency_seconds:mean5m{job="myjob"} > 0.5
for: 10m
labels:
severity: page
annotations:
summary: High request latency
			`) + "\n"
			err := s.validateAlertManagerRules(context.Background(), rules)
			tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Zero Alert Manager rules found."), err)
		})

		t.Run("Invalid", func(t *testing.T) {
			rules := strings.TrimSpace(`
groups:
- name: example
  rules:
  - alert: HighRequestLatency
			`) + "\n"
			err := s.validateAlertManagerRules(context.Background(), rules)
			tests.AssertGRPCError(t, status.New(codes.InvalidArgument, "Invalid Alert Manager rules."), err)
		})
	})
}
