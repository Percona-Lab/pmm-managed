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
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

// FindDSNByServiceIDandPMMAgentID resolves DSN by service id.
func FindDSNByServiceIDandPMMAgentID(q *reform.Querier, serviceID, pmmAgentID, db string) (string, error) {
	// FIXME This function is problematic:
	//
	// * it will return error in case we run multiple exporters for the same service with different credentials;
	//
	// * MySQLdExporter's DSN does not use ParseTime that could be helpful for query actions,
	//   but we can't change if for mysqld_exporter for compatibility reasons.
	//
	// rewrite logic to use agent_id instead of service_id?

	svc, err := FindServiceByID(q, serviceID)
	if err != nil {
		return "", err
	}

	agents, err := FindAgents(q, AgentFilters{
		ServiceID:  serviceID,
		PMMAgentID: pmmAgentID,
	})
	if err != nil {
		return "", err
	}

	types := make(map[AgentType]*Agent)
	for _, a := range agents {
		if _, ok := types[a.AgentType]; ok {
			return "", status.Errorf(codes.FailedPrecondition, "Couldn't resolve dsn, as there should be only one agent of %s type", a.AgentType)
		}

		types[a.AgentType] = a
	}

	switch svc.ServiceType {
	case MySQLServiceType:
		if _, ok := types[QANMySQLSlowlogAgentType]; ok {
			return types[QANMySQLSlowlogAgentType].DSN(svc, time.Second, db), nil
		}
		if _, ok := types[QANMySQLPerfSchemaAgentType]; ok {
			return types[QANMySQLPerfSchemaAgentType].DSN(svc, time.Second, db), nil
		}
		if _, ok := types[MySQLdExporterType]; ok {
			return types[MySQLdExporterType].DSN(svc, time.Second, db), nil
		}
	case MongoDBServiceType:
		if _, ok := types[QANMongoDBProfilerAgentType]; ok {
			return types[QANMongoDBProfilerAgentType].DSN(svc, time.Second, db), nil
		}
		if _, ok := types[MongoDBExporterType]; ok {
			return types[MongoDBExporterType].DSN(svc, time.Second, db), nil
		}
	case PostgreSQLServiceType:
		if _, ok := types[QANPostgreSQLPgStatementsAgentType]; ok {
			return types[QANPostgreSQLPgStatementsAgentType].DSN(svc, time.Second, db), nil
		}
		if _, ok := types[MongoDBExporterType]; ok {
			return types[MongoDBExporterType].DSN(svc, time.Second, db), nil
		}
	}

	if len(types) > 0 {
		for _, t := range types {
			return t.DSN(svc, time.Second, db), nil
		}
	}

	return "", status.Errorf(codes.FailedPrecondition, "Couldn't resolve dsn, as service is unsupported")
}
