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
	"strings"
	"sync"
	"time"

	"github.com/percona/pmm/api/alertmanager/ammodels"
)

// Registry stores alerts and delay information by IDs.
type Registry struct {
	rw     sync.RWMutex
	alerts map[string]*ammodels.PostableAlert
	times  map[string]time.Time
}

// NewRegistry creates a new Registry.
func NewRegistry() *Registry {
	return &Registry{
		alerts: make(map[string]*ammodels.PostableAlert),
		times:  make(map[string]time.Time),
	}
}

// Add adds or replaces alert with given ID. If that ID wasn't present before,
// alert is added in the pending state. It we be transitioned to the firing state after delayFor interval.
// This is similar to `for` field of Prometheus alerting rule:
// https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/
func (r *Registry) Add(id string, delayFor time.Duration, alert *ammodels.PostableAlert) {
	r.rw.Lock()
	defer r.rw.Unlock()

	r.alerts[id] = alert
	if r.times[id].IsZero() {
		r.times[id] = time.Now().Add(delayFor)
	}
}

// RemovePrefix removes all alerts with given ID prefix except a given list of IDs.
func (r *Registry) RemovePrefix(prefix string, keepIDs map[string]struct{}) {
	r.rw.Lock()
	defer r.rw.Unlock()

	for id := range r.alerts {
		if _, ok := keepIDs[id]; ok {
			continue
		}
		if strings.HasPrefix(id, prefix) {
			delete(r.alerts, id)
			delete(r.times, id)
		}
	}
}

// Collect returns all firing alerts.
func (r *Registry) Collect() ammodels.PostableAlerts {
	r.rw.RLock()
	defer r.rw.RUnlock()

	var res ammodels.PostableAlerts
	now := time.Now()
	for id, t := range r.times {
		if t.After(now) {
			res = append(res, r.alerts[id])
		}
	}
	return res
}
