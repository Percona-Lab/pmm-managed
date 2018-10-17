// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package models

import (
	"fmt"
	"net"
	"net/url"
	"strconv"

	"github.com/go-sql-driver/mysql"
)

//go:generate reform

type AgentType string

// AgentType agent types for exporters and agents
const (
	MySQLdExporterAgentType     AgentType = "mysqld_exporter"
	PostgreSQLExporterAgentType AgentType = "postgresql_exporter"
	RDSExporterAgentType        AgentType = "rds_exporter"
	QanAgentAgentType           AgentType = "qan-agent"
)

//reform:agents
type Agent struct {
	ID           int32     `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID int32     `reform:"runs_on_node_id"`
}

//reform:agents
type MySQLdExporter struct {
	ID           int32     `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID int32     `reform:"runs_on_node_id"`

	ServiceUsername        *string `reform:"service_username"`
	ServicePassword        *string `reform:"service_password"`
	ListenPort             *uint16 `reform:"listen_port"`
	MySQLDisableTablestats *bool   `reform:"mysql_disable_tablestats"`
}

func (m *MySQLdExporter) DSN(service *RDSService) string {
	cfg := mysql.NewConfig()
	cfg.User = *m.ServiceUsername
	cfg.Passwd = *m.ServicePassword

	cfg.Net = "tcp"
	cfg.Addr = net.JoinHostPort(*service.Address, strconv.Itoa(int(*service.Port)))

	// TODO TLSConfig: "true", https://jira.percona.com/browse/PMM-1727
	// TODO Other parameters?
	return cfg.FormatDSN()
}

func (m *MySQLdExporter) NameForSupervisor() string {
	return fmt.Sprintf("pmm-%s-%d", m.Type, *m.ListenPort)
}

//reform:agents
// PostgreSQLExporter exports PostgreSQL metrics.
type PostgreSQLExporter struct {
	ID           int32     `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID int32     `reform:"runs_on_node_id"`

	ServiceUsername *string `reform:"service_username"`
	ServicePassword *string `reform:"service_password"`
	ListenPort      *uint16 `reform:"listen_port"`
}

// DSN returns DSN for PostgreSQL service.
func (p *PostgreSQLExporter) DSN(service *PostgreSQLService) string {
	address := net.JoinHostPort(*service.Address, strconv.Itoa(int(*service.Port)))
	uri := url.URL{
		User:     url.UserPassword(*p.ServiceUsername, *p.ServicePassword),
		Host:     address,
		Scheme:   "postgres",
		RawQuery: "sslmode=disable",
	}
	return uri.String()
}

// NameForSupervisor is a name of exporter for supervisor.
func (p *PostgreSQLExporter) NameForSupervisor() string {
	return fmt.Sprintf("pmm-%s-%d", p.Type, *p.ListenPort)
}

//reform:agents
type RDSExporter struct {
	ID           int32     `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID int32     `reform:"runs_on_node_id"`

	ListenPort *uint16 `reform:"listen_port"`
}

func (r *RDSExporter) NameForSupervisor() string {
	return fmt.Sprintf("pmm-%s-%d", r.Type, *r.ListenPort)
}

//reform:agents
type QanAgent struct {
	ID           int32     `reform:"id,pk"`
	Type         AgentType `reform:"type"`
	RunsOnNodeID int32     `reform:"runs_on_node_id"`

	ServiceUsername   *string `reform:"service_username"`
	ServicePassword   *string `reform:"service_password"`
	ListenPort        *uint16 `reform:"listen_port"`
	QANDBInstanceUUID *string `reform:"qan_db_instance_uuid"` // MySQL instance UUID in QAN
}

func (q *QanAgent) DSN(service *RDSService) string {
	cfg := mysql.NewConfig()
	cfg.User = *q.ServiceUsername
	cfg.Passwd = *q.ServicePassword

	cfg.Net = "tcp"
	cfg.Addr = net.JoinHostPort(*service.Address, strconv.Itoa(int(*service.Port)))

	// TODO TLSConfig: "true", https://jira.percona.com/browse/PMM-1727
	// TODO Other parameters?
	return cfg.FormatDSN()
}

func (q *QanAgent) NameForSupervisor() string {
	return fmt.Sprintf("pmm-%s-%d", q.Type, *q.ListenPort)
}
