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
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

//go:generate reform

const (
	// maximum time for connecting to the database
	sqlDialTimeout = 5 * time.Second
)

// AgentType represents Agent type as stored in database.
type AgentType string

// Agent types.
const (
	PMMAgentType            AgentType = "pmm-agent"
	NodeExporterAgentType   AgentType = "node_exporter"
	MySQLdExporterAgentType AgentType = "mysqld_exporter"

	PostgresExporterAgentType AgentType = "postgres_exporter"
	RDSExporterAgentType      AgentType = "rds_exporter"
)

// AgentRow represents Agent as stored in database.
//reform:agents
type AgentRow struct {
	ID           string    `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID string    `reform:"runs_on_node_id"`
	Disabled     bool      `reform:"disabled"`
	// CreatedAt    time.Time `reform:"created_at"`
	// UpdatedAt    time.Time `reform:"updated_at"`

	Version         *string `reform:"version"`
	ListenPort      *uint16 `reform:"listen_port"`
	ServiceUsername *string `reform:"service_username"`
	ServicePassword *string `reform:"service_password"`
}

// BeforeInsert implements reform.BeforeInserter interface.
//nolint:unparam
func (ar *AgentRow) BeforeInsert() error {
	// now := time.Now().Truncate(time.Microsecond).UTC()
	// ar.CreatedAt = now
	// ar.UpdatedAt = now
	return nil
}

// BeforeUpdate implements reform.BeforeUpdater interface.
//nolint:unparam
func (ar *AgentRow) BeforeUpdate() error {
	// now := time.Now().Truncate(time.Microsecond).UTC()
	// ar.UpdatedAt = now
	return nil
}

// AfterFind implements reform.AfterFinder interface.
//nolint:unparam
func (ar *AgentRow) AfterFind() error {
	// ar.CreatedAt = ar.CreatedAt.UTC()
	// ar.UpdatedAt = ar.UpdatedAt.UTC()
	return nil
}

// check interfaces
var (
	_ reform.BeforeInserter = (*AgentRow)(nil)
	_ reform.BeforeUpdater  = (*AgentRow)(nil)
	_ reform.AfterFinder    = (*AgentRow)(nil)
)

// TODO remove code below

//reform:agents
type Agent struct {
	ID           string    `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID string    `reform:"runs_on_node_id"`
	Disabled     bool      `reform:"disabled"`

	// TODO Does it really belong there? Remove when we have agent without one.
	ListenPort *uint16 `reform:"listen_port"`
}

// NameForSupervisor returns a name of agent for supervisor.
func NameForSupervisor(typ AgentType, listenPort uint16) string {
	return fmt.Sprintf("pmm-%s-%d", typ, listenPort)
}

//reform:agents
type MySQLdExporter struct {
	ID           string    `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID string    `reform:"runs_on_node_id"`
	Disabled     bool      `reform:"disabled"`

	ServiceUsername        *string `reform:"service_username"`
	ServicePassword        *string `reform:"service_password"`
	ListenPort             *uint16 `reform:"listen_port"`
	MySQLDisableTablestats *bool   `reform:"mysql_disable_tablestats"`
}

func (m *MySQLdExporter) DSN(service *MySQLService) string {
	cfg := mysql.NewConfig()
	cfg.User = *m.ServiceUsername
	cfg.Passwd = *m.ServicePassword

	cfg.Net = "tcp"
	cfg.Addr = net.JoinHostPort(*service.Address, strconv.Itoa(int(*service.Port)))

	cfg.Timeout = sqlDialTimeout

	// TODO TLSConfig: "true", https://jira.percona.com/browse/PMM-1727
	// TODO Other parameters?
	return cfg.FormatDSN()
}

// binary name is postgres_exporter, that's why PostgresExporter below is not PostgreSQLExporter

//reform:agents
// PostgresExporter exports PostgreSQL metrics.
type PostgresExporter struct {
	ID           string    `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID string    `reform:"runs_on_node_id"`
	Disabled     bool      `reform:"disabled"`

	ServiceUsername *string `reform:"service_username"`
	ServicePassword *string `reform:"service_password"`
	ListenPort      *uint16 `reform:"listen_port"`
}

// DSN returns DSN for PostgreSQL service.
func (p *PostgresExporter) DSN(service *PostgreSQLService) string {
	q := make(url.Values)
	q.Set("sslmode", "disable") // TODO https://jira.percona.com/browse/PMM-1727
	q.Set("connect_timeout", strconv.Itoa(int(sqlDialTimeout.Seconds())))

	address := net.JoinHostPort(*service.Address, strconv.Itoa(int(*service.Port)))
	uri := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(*p.ServiceUsername, *p.ServicePassword),
		Host:     address,
		Path:     "postgres",
		RawQuery: q.Encode(),
	}
	return uri.String()
}

//reform:agents
type RDSExporter struct {
	ID           string    `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID string    `reform:"runs_on_node_id"`
	Disabled     bool      `reform:"disabled"`

	ListenPort *uint16 `reform:"listen_port"`
}

// AgentFilters represents filters for agents list.
type AgentFilters struct {
	ServiceID *string
}

// AgentsByFilters returns agents providing insights for a given filters.
func AgentsByFilters(q *reform.Querier, filters AgentFilters) ([]*AgentRow, error) {
	var agentIDs []interface{}
	var structs []reform.Struct
	var err error
	if filters.ServiceID != nil {
		agentServices, err := q.SelectAllFrom(AgentServiceView, "WHERE service_id = ?", *filters.ServiceID)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		for _, str := range agentServices {
			agentIDs = append(agentIDs, str.(*AgentService).AgentID)
		}

		if len(agentIDs) == 0 {
			return []*AgentRow{}, nil
		}

		structs, err = q.FindAllFrom(AgentRowTable, "id", agentIDs...)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	} else {
		structs, err = q.SelectAllFrom(AgentRowTable, "ORDER BY ID")
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}
	agents := make([]*AgentRow, len(structs))
	for i, str := range structs {
		agents[i] = str.(*AgentRow)
	}
	return agents, nil
}
