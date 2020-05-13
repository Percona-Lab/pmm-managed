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

package alertmanager

import (
	"testing"
	"time"

	"github.com/percona/pmm/api/alertmanager/ammodels"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/percona/pmm-managed/services"
)

func TestRegistry(t *testing.T) {
	nowValue, origNow := now(), now
	now = func() time.Time {
		return nowValue
	}
	defer func() {
		now = origNow
	}()

	t.Run("DelayFor", func(t *testing.T) {
		r := NewRegistry()
		alertParams := new(services.AlertParams)
		r.CreateAlert(alertParams, time.Minute)
		assert.Empty(t, r.collect())

		// 1 second before
		nowValue = nowValue.Add(59 * time.Second)
		assert.Empty(t, r.collect())

		// exactly that time
		nowValue = nowValue.Add(time.Second)
		assert.Empty(t, r.collect())

		expected := &ammodels.PostableAlert{
			Annotations: alertParams.Annotations,
			Alert: ammodels.Alert{
				Labels: alertParams.Labels,
			},
		}

		// 1 second after
		nowValue = nowValue.Add(time.Second)
		alerts := r.collect()
		require.Len(t, alerts, 1)
		assert.Equal(t, expected, alerts[0])
	})
}
