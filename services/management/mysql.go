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

type agentStateRequestSender interface {
	SendSetStateRequest(ctx context.Context, pmmAgentID string)
	IsConnected(pmmAgentID string) bool
}

// MySQLService MySQL Management Service.
type MySQLService struct {
	db   *reform.DB
	asrs agentStateRequestSender
}

// NewMySQLService creates new MySQL Management Service.
func NewMySQLService(db *reform.DB, asrs agentStateRequestSender) *MySQLService {
	return &MySQLService{db, asrs}
}

// Add adds "MySQL Service", "MySQL Exporter Agent" and "QAN MySQL PerfSchema Agent".
func (s *MySQLService) Add(ctx context.Context, req *managementpb.AddMySQLRequest) (res *managementpb.AddMySQLResponse, err error) {
	res = &managementpb.AddMySQLResponse{}

	if e := s.db.InTransaction(func(tx *reform.TX) error {

		service, err := models.AddNewService(tx.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
			ServiceName: req.ServiceName,
			NodeID:      req.NodeId,
			Address:     pointer.ToStringOrNil(req.Address),
			Port:        pointer.ToUint16OrNil(uint16(req.Port)),
		})

		if err != nil {
			return err
		}

		invService, err := inventory.ToInventoryService(service)
		if err != nil {
			return err
		}

		res.Service = invService.(*inventorypb.MySQLService)

		if req.MysqldExporter {

			params := &models.AddExporterAgentParams{
				PMMAgentID: req.PmmAgentId,
				ServiceID:  invService.ID(),
				Username:   req.Username,
				Password:   req.Password,
			}
			row, err := models.AgentAddExporter(tx.Querier, models.MySQLdExporterType, params)
			if err != nil {
				return err
			}

			agent, err := inventory.ToInventoryAgent(tx.Querier, row, s.asrs)
			if err != nil {
				return err
			}

			res.MysqldExporter = agent.(*inventorypb.MySQLdExporter)
		}

		if req.QanMysqlPerfschema {

			params := &models.AddExporterAgentParams{
				PMMAgentID: req.PmmAgentId,
				ServiceID:  invService.ID(),
				Username:   req.QanUsername,
				Password:   req.QanPassword,
			}

			row, err := models.AgentAddExporter(tx.Querier, models.QANMySQLPerfSchemaAgentType, params)
			if err != nil {
				return err
			}

			qAgent, err := inventory.ToInventoryAgent(tx.Querier, row, s.asrs)
			if err != nil {
				return err
			}

			res.QanMysqlPerfschema = qAgent.(*inventorypb.QANMySQLPerfSchemaAgent)
		}

		if req.QanMysqlSlowlog {
			params := &models.AddExporterAgentParams{
				PMMAgentID: req.PmmAgentId,
				ServiceID:  invService.ID(),
				Username:   req.QanUsername,
				Password:   req.QanPassword,
			}

			row, err := models.AgentAddExporter(tx.Querier, models.QANMySQLSlowlogAgentType, params)
			if err != nil {
				return err
			}

			qAgent, err := inventory.ToInventoryAgent(tx.Querier, row, s.asrs)
			if err != nil {
				return err
			}

			res.QanMysqlSlowlog = qAgent.(*inventorypb.QANMySQLSlowlogAgent)
		}

		return nil
	}); e != nil {
		return nil, e
	}

	s.asrs.SendSetStateRequest(ctx, req.PmmAgentId)

	return res, nil
}
