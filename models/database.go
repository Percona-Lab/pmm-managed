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
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

var initialCurrentTime = Now().Format(time.RFC3339)

// databaseSchema maps schema version from schema_migrations table (id column) to a slice of DDL queries.
var databaseSchema = [][]string{
	1: {
		`CREATE TABLE schema_migrations (
			id INT NOT NULL,
			PRIMARY KEY (id)
		)`,

		`CREATE TABLE nodes (
			-- common
			node_id VARCHAR NOT NULL,
			node_type VARCHAR NOT NULL,
			node_name VARCHAR NOT NULL,
			machine_id VARCHAR,
			custom_labels TEXT,
			address VARCHAR,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,

			-- Generic
			distro VARCHAR,
			distro_version VARCHAR,

			-- Container
			docker_container_id VARCHAR,
			docker_container_name VARCHAR,

			-- RemoteAmazonRDS
			-- RDS instance is stored in address
			region VARCHAR,

			PRIMARY KEY (node_id),
			UNIQUE (node_name),
			UNIQUE (machine_id),
			UNIQUE (docker_container_id),
			UNIQUE (address, region),

			CHECK (node_type <> ''),
			CHECK (node_name <> ''),
			CHECK (machine_id <> ''),
			CHECK (address <> ''),
			CHECK (distro <> ''),
			CHECK (distro_version <> ''),
			CHECK (docker_container_id <> ''),
			CHECK (docker_container_name <> ''),
			CHECK (region <> '')
		)`,

		fmt.Sprintf(`INSERT INTO nodes (node_id, node_type,	node_name, created_at, updated_at) VALUES ('%s', '%s', 'PMM Server', '%s', '%s')`, //nolint:gosec
			PMMServerNodeID, GenericNodeType, initialCurrentTime, initialCurrentTime), //nolint:gosec

		`CREATE TABLE services (
			-- common
			service_id VARCHAR NOT NULL,
			service_type VARCHAR NOT NULL,
			service_name VARCHAR NOT NULL,
			node_id VARCHAR NOT NULL,
			custom_labels TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,

			address VARCHAR(255),
			port INTEGER,

			PRIMARY KEY (service_id),
			UNIQUE (service_name),
			FOREIGN KEY (node_id) REFERENCES nodes (node_id),

			CHECK (service_type <> ''),
			CHECK (service_name <> ''),
			CHECK (node_id <> ''),
			CHECK (address <> '')
		)`,

		`CREATE TABLE agents (
			-- common
			agent_id VARCHAR NOT NULL,
			agent_type VARCHAR NOT NULL,
			runs_on_node_id VARCHAR,
			pmm_agent_id VARCHAR,
			custom_labels TEXT,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,

			-- state
			status VARCHAR NOT NULL,
			listen_port INTEGER,
			version VARCHAR,

			-- Credentials to access service
			username VARCHAR,
			password VARCHAR,
			metrics_url VARCHAR,

			PRIMARY KEY (agent_id),
			FOREIGN KEY (runs_on_node_id) REFERENCES nodes (node_id),
			FOREIGN KEY (pmm_agent_id) REFERENCES agents (agent_id),

			CHECK (agent_type <> ''),
			CHECK (runs_on_node_id <> ''),
			CHECK (pmm_agent_id <> ''),
			CHECK (version <> ''),
			CHECK (username <> ''),
			CHECK (password <> ''),
			CHECK (metrics_url <> '')
		)`,

		`CREATE TABLE agent_nodes (
			agent_id VARCHAR NOT NULL,
			node_id VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL,

			FOREIGN KEY (agent_id) REFERENCES agents (agent_id),
			FOREIGN KEY (node_id) REFERENCES nodes (node_id),
			UNIQUE (agent_id, node_id)
		)`,

		`CREATE TABLE agent_services (
			agent_id VARCHAR NOT NULL,
			service_id VARCHAR NOT NULL,
			created_at TIMESTAMP NOT NULL,

			FOREIGN KEY (agent_id) REFERENCES agents (agent_id),
			FOREIGN KEY (service_id) REFERENCES services (service_id),
			UNIQUE (agent_id, service_id)
		)`,

		`CREATE TABLE telemetry (
  			uuid VARCHAR PRIMARY KEY,
  			created_at TIMESTAMP NOT NULL
		)`,
	},
}

// OpenDB opens connection to PostgreSQL database and runs migrations.
func OpenDB(name, username, password string, logf reform.Printf) (*sql.DB, error) {
	q := make(url.Values)
	q.Set("sslmode", "disable")

	address := "127.0.0.1:5432"
	uri := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(username, password),
		Host:     address,
		Path:     name,
		RawQuery: q.Encode(),
	}
	if uri.Path == "" {
		uri.Path = "postgres"
	}
	dsn := uri.String()

	db, err := sql.Open("postgres", dsn)
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
		db.SetConnMaxLifetime(0)
		err = db.Ping()
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to PostgreSQL.")
	}

	if name == "" {
		return db, nil
	}

	latestVersion := len(databaseSchema) - 1 // skip item 0
	var currentVersion int
	err = db.QueryRow("SELECT id FROM schema_migrations ORDER BY id DESC LIMIT 1").Scan(&currentVersion)
	if pErr, ok := err.(*pq.Error); ok && pErr.Code == "42P01" { // undefined_table (see https://www.postgresql.org/docs/current/errcodes-appendix.html)
		err = nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get current version.")
	}
	logf("Current database schema version: %d. Latest version: %d.", currentVersion, latestVersion)

	for version := currentVersion + 1; version <= latestVersion; version++ {
		logf("Migrating database to schema version %d ...", version)
		queries := databaseSchema[version]
		queries = append(queries, fmt.Sprintf(`INSERT INTO schema_migrations (id) VALUES (%d)`, version))
		for _, q := range queries {
			q = strings.TrimSpace(q)
			logf("\n%s\n", q)
			if _, err = db.Exec(q); err != nil {
				return nil, errors.Wrapf(err, "Failed to execute statement:\n%s.", q)
			}
		}
	}

	return db, nil
}
