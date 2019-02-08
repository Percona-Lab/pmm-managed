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

	"github.com/go-sql-driver/mysql" // register SQL driver
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gopkg.in/reform.v1"
)

// FIXME Re-add created_at/updated_at: https://jira.percona.com/browse/PMM-3350

// databaseSchema maps schema version from schema_migrations table (id column) to a slice of DDL queries.
//
// Initial AUTO_INCREMENT values are spaced to prevent programming errors, or at least make them more visible.
// It does not imply that one can have at most 1000 nodes, etc.
var databaseSchema = [][]string{
	1: {
		`CREATE TABLE schema_migrations (
			id INT NOT NULL,
			PRIMARY KEY (id)
		)`,

		`CREATE TABLE nodes (
			id VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			-- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

			hostname VARCHAR(255),
			region VARCHAR(255),

			PRIMARY KEY (id),
			UNIQUE (name),
			UNIQUE (hostname, region)
		)`,

		`INSERT INTO nodes (id, type, name) VALUES ('` + PMMServerNodeID + `', '` + string(PMMServerNodeType) + `', 'PMM Server')`,

		`CREATE TABLE services (
			id VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			name VARCHAR(255) NOT NULL,
			node_id VARCHAR(255) NOT NULL,
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			-- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

			address VARCHAR(255),
			port SMALLINT UNSIGNED,
			unix_socket VARCHAR(255),

			aws_access_key VARCHAR(255),
			aws_secret_key VARCHAR(255),
			engine VARCHAR(255),
			engine_version VARCHAR(255),

			PRIMARY KEY (id),
			UNIQUE (name),
			FOREIGN KEY (node_id) REFERENCES nodes (id)
		)`,

		`CREATE TABLE agents (
			id VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			runs_on_node_id VARCHAR(255) NOT NULL,
			disabled BOOL NOT NULL,
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			-- updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

			version VARCHAR(255),
			listen_port SMALLINT UNSIGNED,
			service_username VARCHAR(255),
			service_password VARCHAR(255),

			mysql_disable_tablestats TINYINT(1),

			PRIMARY KEY (id),
			FOREIGN KEY (runs_on_node_id) REFERENCES nodes (id)
		)`,

		`CREATE TABLE agent_nodes (
			agent_id VARCHAR(255) NOT NULL,
			node_id VARCHAR(255) NOT NULL,
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			container_id VARCHAR(255),
			container_name VARCHAR(255),
			kubernetes_pod_uid VARCHAR(255),
			kubernetes_pod_name VARCHAR(255),

			FOREIGN KEY (agent_id) REFERENCES agents (id),
			FOREIGN KEY (node_id) REFERENCES nodes (id),
			UNIQUE (agent_id, node_id)
		)`,

		`CREATE TABLE agent_services (
			agent_id VARCHAR(255) NOT NULL,
			service_id VARCHAR(255) NOT NULL,
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			container_id VARCHAR(255),
			container_name VARCHAR(255),
			kubernetes_pod_uid VARCHAR(255),
			kubernetes_pod_name VARCHAR(255),

			FOREIGN KEY (agent_id) REFERENCES agents (id),
			FOREIGN KEY (service_id) REFERENCES services (id),
			UNIQUE (agent_id, service_id)
		)`,
	},
}

func OpenDB(name, username, password string, logf reform.Printf) (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = username
	cfg.Passwd = password
	cfg.DBName = name

	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"

	// required for reform
	cfg.ClientFoundRows = true
	cfg.ParseTime = true

	dsn := cfg.FormatDSN()
	db, err := sql.Open("mysql", dsn)
	if err == nil {
		db.SetMaxIdleConns(10)
		db.SetMaxOpenConns(10)
		db.SetConnMaxLifetime(0)
		err = db.Ping()
	}
	if err != nil {
		return nil, errors.Wrap(err, "Failed to connect to MySQL.")
	}

	if name == "" {
		return db, nil
	}

	latestVersion := len(databaseSchema) - 1 // skip item 0
	var currentVersion int
	err = db.QueryRow("SELECT id FROM schema_migrations ORDER BY id DESC LIMIT 1").Scan(&currentVersion)
	if myErr, ok := err.(*mysql.MySQLError); ok && myErr.Number == 0x47a { // 1046 table doesn't exist
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

// postgresDatabaseSchema maps schema version from schema_migrations table (id column) to a slice of DDL queries.
var postgresDatabaseSchema = [][]string{
	1: {
		`CREATE TABLE schema_migrations (
			id INT NOT NULL,
			PRIMARY KEY (id)
		)`,

		`CREATE TABLE key_values (
			key VARCHAR(255) NOT NULL PRIMARY KEY,
			value TEXT NOT NULL
			-- created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)`,
	},
}

func OpenPostgresDB(name, username, password string, logf reform.Printf) (*sql.DB, error) {

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
	if myErr, ok := err.(*pq.Error); ok && myErr.Code == "42P01" { // 1046 table doesn't exist
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
