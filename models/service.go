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
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

// ServicesForAgent returns all Services for which Agent with given ID provides insights.
func ServicesForAgent(q *reform.Querier, agentID string) ([]*Service, error) {
	structs, err := q.FindAllFrom(AgentServiceView, "agent_id", agentID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select Service IDs")
	}

	serviceIDs := make([]interface{}, len(structs))
	for i, s := range structs {
		serviceIDs[i] = s.(*AgentService).ServiceID
	}
	if len(serviceIDs) == 0 {
		return []*Service{}, nil
	}

	p := strings.Join(q.Placeholders(1, len(serviceIDs)), ", ")
	tail := fmt.Sprintf("WHERE service_id IN (%s) ORDER BY service_id", p) //nolint:gosec
	structs, err = q.SelectAllFrom(ServiceTable, tail, serviceIDs...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select Services")
	}

	res := make([]*Service, len(structs))
	for i, s := range structs {
		res[i] = s.(*Service)
	}
	return res, nil
}

//go:generate reform

// ServiceType represents Service type as stored in database.
type ServiceType string

// Service types.
const (
	MySQLServiceType          ServiceType = "mysql"
	AmazonRDSMySQLServiceType ServiceType = "amazon-rds-mysql"
)

// Service represents Service as stored in database.
//reform:services
type Service struct {
	ServiceID   string      `reform:"service_id,pk"`
	ServiceType ServiceType `reform:"service_type"`
	ServiceName string      `reform:"service_name"`
	NodeID      string      `reform:"node_id"`
	CreatedAt   time.Time   `reform:"created_at"`
	// UpdatedAt time.Time   `reform:"updated_at"`

	Address    *string `reform:"address"`
	Port       *uint16 `reform:"port"`
	UnixSocket *string `reform:"unix_socket"`
}

// BeforeInsert implements reform.BeforeInserter interface.
//nolint:unparam
func (s *Service) BeforeInsert() error {
	now := Now()
	s.CreatedAt = now
	// s.UpdatedAt = now
	return nil
}

// BeforeUpdate implements reform.BeforeUpdater interface.
//nolint:unparam
func (s *Service) BeforeUpdate() error {
	// now := Now()
	// s.UpdatedAt = now
	return nil
}

// AfterFind implements reform.AfterFinder interface.
//nolint:unparam
func (s *Service) AfterFind() error {
	s.CreatedAt = s.CreatedAt.UTC()
	// s.UpdatedAt = s.UpdatedAt.UTC()
	return nil
}

// check interfaces
var (
	_ reform.BeforeInserter = (*Service)(nil)
	_ reform.BeforeUpdater  = (*Service)(nil)
	_ reform.AfterFinder    = (*Service)(nil)
)
