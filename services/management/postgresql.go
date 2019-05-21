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

package management

import (
	"context"

	"github.com/AlekSi/pointer"
	"github.com/percona/pmm/api/inventorypb"
	"github.com/percona/pmm/api/managementpb"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"

	// FIXME Refactor, as service shouldn't depend on other service in one abstraction level.
	// https://jira.percona.com/browse/PMM-3541
	// See also main_test.go
	"github.com/percona/pmm-managed/services/inventory"
)

// PostgreSQLService PostgreSQL Management Service.
type PostgreSQLService struct {
	db       *reform.DB
	registry registry
}

// NewPostgreSQLService creates new PostgreSQL Management Service.
func NewPostgreSQLService(db *reform.DB, registry registry) *PostgreSQLService {
	return &PostgreSQLService{db, registry}
}

// Add adds "PostgreSQL Service", "PostgreSQL Exporter Agent" and "QAN PostgreSQL PerfSchema Agent".
func (s *PostgreSQLService) Add(ctx context.Context, req *managementpb.AddPostgreSQLRequest) (*managementpb.AddPostgreSQLResponse, error) {
	res := new(managementpb.AddPostgreSQLResponse)

	if e := s.db.InTransaction(func(tx *reform.TX) error {
		service, err := models.AddNewService(tx.Querier, models.PostgreSQLServiceType, &models.AddDBMSServiceParams{
			ServiceName:  req.ServiceName,
			NodeID:       req.NodeId,
			Environment:  req.Environment,
			Address:      pointer.ToStringOrNil(req.Address),
			Port:         pointer.ToUint16OrNil(uint16(req.Port)),
			CustomLabels: req.CustomLabels,
		})
		if err != nil {
			return err
		}

		invService, err := inventory.ToInventoryService(service)
		if err != nil {
			return err
		}
		res.Service = invService.(*inventorypb.PostgreSQLService)

		row, err := models.AgentAddExporter(tx.Querier, models.PostgresExporterType, &models.AddExporterAgentParams{
			PMMAgentID: req.PmmAgentId,
			ServiceID:  service.ServiceID,
			Username:   req.Username,
			Password:   req.Password,
		})
		if err != nil {
			return err
		}

		if !req.SkipConnectionCheck {
			err = s.registry.CheckConnectionToService(ctx, service, row)
			if err != nil {
				return err
			}
		}

		agent, err := inventory.ToInventoryAgent(tx.Querier, row, s.registry)
		if err != nil {
			return err
		}
		res.PostgresExporter = agent.(*inventorypb.PostgresExporter)

		return nil
	}); e != nil {
		return nil, e
	}

	s.registry.SendSetStateRequest(ctx, req.PmmAgentId)
	return res, nil
}
