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

package logger

import (
	"strings"
	"time"

	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"gopkg.in/reform.v1"
)

const (
	namespace = "go_sql"
	subsystem = "reform"
)

// TODO https://jira.percona.com/browse/PMM-5302 Move to percona/pmm utils and use in pmm-agent and qan-api2.
type Reform struct {
	l          *reform.PrintfLogger
	mRequests  *prom.CounterVec
	mResponses *prom.SummaryVec
}

func NewReform(driver, dbName string, l *logrus.Entry) *Reform {
	constLabels := prom.Labels{
		"driver": driver,
		"db":     dbName,
	}

	return &Reform{
		l: reform.NewPrintfLogger(l.Tracef),
		mRequests: prom.NewCounterVec(prom.CounterOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        "requests_total",
			Help:        "Total number of queries started.",
			ConstLabels: constLabels,
		}, []string{"statement"}),
		mResponses: prom.NewSummaryVec(prom.SummaryOpts{
			Namespace:   namespace,
			Subsystem:   subsystem,
			Name:        "response_seconds",
			Help:        "Response durations in seconds.",
			ConstLabels: constLabels,
			Objectives:  map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}, []string{"statement", "error"}),
	}
}

func statement(query string) string {
	query = strings.ToLower(strings.TrimSpace(query))
	parts := strings.SplitN(query, " ", 2)
	if len(parts) != 2 {
		return query
	}
	return parts[0]
}

// Before implements reform.Logger.
func (r *Reform) Before(query string, args []interface{}) {
	r.l.Before(query, args)

	r.mRequests.WithLabelValues(statement(query)).Inc()
}

// After implements reform.Logger.
func (r *Reform) After(query string, args []interface{}, d time.Duration, err error) {
	r.l.After(query, args, d, err)

	e := "0"
	if err != nil {
		e = "1"
	}
	r.mResponses.WithLabelValues(statement(query), e).Observe(d.Seconds())
}

// Describe implements prom.Collector.
func (r *Reform) Describe(ch chan<- *prom.Desc) {
	r.mRequests.Describe(ch)
	r.mResponses.Describe(ch)
}

// Collect implements prom.Collector.
func (r *Reform) Collect(ch chan<- prom.Metric) {
	r.mRequests.Collect(ch)
	r.mResponses.Collect(ch)
}

// check interfaces
var (
	_ reform.Logger  = (*Reform)(nil)
	_ prom.Collector = (*Reform)(nil)
)
