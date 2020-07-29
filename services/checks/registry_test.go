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

package checks

import (
	"testing"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/percona-platform/saas/pkg/check"
	"github.com/percona/pmm/api/alertmanager/ammodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegistry(t *testing.T) {
	t.Run("Create and Collect Alerts", func(t *testing.T) {
		r := newRegistry()

		nowValue := time.Now().UTC().Round(0) // strip a monotonic clock reading
		r.nowF = func() time.Time { return nowValue }
		alertTTL := resolveTimeoutFactor * defaultResendInterval
		checkResults := []sttCheckResult{
			{
				checkName: "name",
				target: target{
					agentID:   "/agent_id/123",
					serviceID: "/service_id/123",
					labels: map[string]string{
						"foo": "bar",
					},
				},
				result: check.Result{
					Summary:     "check summary",
					Description: "check description",
					Severity:    check.Warning,
					Labels: map[string]string{
						"baz": "qux",
					},
				},
			},
		}

		r.set(checkResults)

		expectedAlert := &ammodels.PostableAlert{
			Annotations: map[string]string{
				"summary":     "check summary",
				"description": "check description",
			},
			EndsAt: strfmt.DateTime(nowValue.Add(alertTTL)),
			Alert: ammodels.Alert{
				Labels: map[string]string{
					"alert_id":  "/stt/3293a1002b32c425dbda851406bf1428dced74cd",
					"alertname": "name",
					"baz":       "qux",
					"foo":       "bar",
					"severity":  "warning",
					"stt_check": "1",
				},
			},
		}

		collectedAlerts := r.collect(alertTTL)
		require.Len(t, collectedAlerts, 1)
		require.Equal(t, 1, cap(collectedAlerts))
		assert.Equal(t, expectedAlert, collectedAlerts[0])
	})
}
