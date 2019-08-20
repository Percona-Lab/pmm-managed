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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	t.Run("ParseEnv", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			s := NewServer(nil, nil, []string{
				"DISABLE_TELEMETRY=1",
				"METRICS_RESOLUTION=2",
				"DATA_RETENTION=240h",
			})
			assert.Equal(t, true, s.envDisableTelemetry)
			assert.Equal(t, 2*time.Second, s.envMetricsResolution)
			assert.Equal(t, 10*24*time.Hour, s.envDataRetention)

			s = NewServer(nil, nil, []string{
				"DISABLE_TELEMETRY=TrUe",
				"METRICS_RESOLUTION=3S",
				"DATA_RETENTION=360h",
			})
			assert.Equal(t, true, s.envDisableTelemetry)
			assert.Equal(t, 3*time.Second, s.envMetricsResolution)
			assert.Equal(t, 15*24*time.Hour, s.envDataRetention)
		})

		t.Run("Invalid", func(t *testing.T) {
			s := NewServer(nil, nil, []string{
				"DISABLE_TELEMETRY=YES",
				"METRICS_RESOLUTION=0.1s",
				"DATA_RETENTION=15d",
			})
			assert.Equal(t, false, s.envDisableTelemetry)
			assert.Equal(t, time.Duration(0), s.envMetricsResolution)
			assert.Equal(t, time.Duration(0), s.envDataRetention)

			s = NewServer(nil, nil, []string{
				"DISABLE_TELEMETRY=on",
				"METRICS_RESOLUTION=-1",
				"DATA_RETENTION=-10",
			})
			assert.Equal(t, false, s.envDisableTelemetry)
			assert.Equal(t, time.Duration(0), s.envMetricsResolution)
			assert.Equal(t, time.Duration(0), s.envDataRetention)
		})
	})
}
