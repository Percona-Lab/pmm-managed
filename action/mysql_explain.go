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

package action

import (
	"context"

	"gopkg.in/reform.v1"
)

type MySQLExplain struct {
	ID         string
	ServiceID  string
	PMMAgentID string
	Dsn        string
	Query      string

	q *reform.Querier
	s MySQLExplainStarter
	r PMMAgentIDResolver
}

func NewMySQLExplain(q *reform.Querier, s MySQLExplainStarter, r PMMAgentIDResolver) *MySQLExplain {
	return &MySQLExplain{
		ID: createActionID(),
		q:  q,
		s:  s,
		r:  r,
	}
}

func (exp *MySQLExplain) Prepare(serviceID, pmmAgentID, query string) error {
	var err error
	exp.Query = query
	exp.ServiceID = serviceID
	exp.PMMAgentID = pmmAgentID

	exp.PMMAgentID, err = exp.r.ResolveByServiceID(exp.ServiceID, exp.PMMAgentID)
	if err != nil {
		return err
	}
	// TODO: Add DSN string: findAgentForService(MYSQLD_EXPORTER, req.ServiceID)

	return nil
}

func (exp *MySQLExplain) Start(ctx context.Context) error {
	return exp.s.StartMySQLExplainAction(ctx, exp)
}
