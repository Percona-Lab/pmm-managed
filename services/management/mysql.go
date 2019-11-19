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
	"github.com/percona/pmm-managed/services"
)

// MySQLService MySQL Management Service.
type MySQLService struct {
	db       *reform.DB
	registry agentsRegistry
}

// NewMySQLService creates new MySQL Management Service.
func NewMySQLService(db *reform.DB, registry agentsRegistry) *MySQLService {
	return &MySQLService{db, registry}
}

// Add adds "MySQL Service", "MySQL Exporter Agent" and "QAN MySQL PerfSchema Agent".
func (s *MySQLService) Add(ctx context.Context, req *managementpb.AddMySQLRequest) (*managementpb.AddMySQLResponse, error) {
	res := new(managementpb.AddMySQLResponse)

	if e := s.db.InTransaction(func(tx *reform.TX) error {
		// tweak according to API docs
		if req.MaxSlowlogFileSize == 0 {
			req.MaxSlowlogFileSize = 1 << 30 // 1 GB
		}
		if req.MaxSlowlogFileSize < 0 {
			req.MaxSlowlogFileSize = 0
		}

		nodeID, err := nodeID(tx, req.NodeId, req.NodeName, req.AddNode, req.Address)
		if err != nil {
			return err
		}
		service, err := models.AddNewService(tx.Querier, models.MySQLServiceType, &models.AddDBMSServiceParams{
			ServiceName:    req.ServiceName,
			NodeID:         nodeID,
			Environment:    req.Environment,
			Cluster:        req.Cluster,
			ReplicationSet: req.ReplicationSet,
			Address:        pointer.ToStringOrNil(req.Address),
			Port:           pointer.ToUint16OrNil(uint16(req.Port)),
			CustomLabels:   req.CustomLabels,
		})
		if err != nil {
			return err
		}

		invService, err := services.ToAPIService(service)
		if err != nil {
			return err
		}
		res.Service = invService.(*inventorypb.MySQLService)

		row, err := models.CreateAgent(tx.Querier, models.MySQLdExporterType, &models.CreateAgentParams{
			PMMAgentID:    req.PmmAgentId,
			ServiceID:     service.ServiceID,
			Username:      req.Username,
			Password:      req.Password,
			TLS:           req.Tls,
			TLSSkipVerify: req.TlsSkipVerify,
			// MaxTableNumber: int64(req.MaxNumberOfTables),
		})
		if err != nil {
			return err
		}
		if !req.SkipConnectionCheck {
			if err = s.registry.CheckConnectionToService(ctx, tx.Querier, service, row); err != nil {
				return err
			}
		}

		agent, err := services.ToAPIAgent(tx.Querier, row)
		if err != nil {
			return err
		}
		res.MysqldExporter = agent.(*inventorypb.MySQLdExporter)

		if req.QanMysqlPerfschema {
			row, err = models.CreateAgent(tx.Querier, models.QANMySQLPerfSchemaAgentType, &models.CreateAgentParams{
				PMMAgentID:            req.PmmAgentId,
				ServiceID:             service.ServiceID,
				Username:              req.Username,
				Password:              req.Password,
				TLS:                   req.Tls,
				TLSSkipVerify:         req.TlsSkipVerify,
				QueryExamplesDisabled: req.DisableQueryExamples,
			})
			if err != nil {
				return err
			}

			agent, err = services.ToAPIAgent(tx.Querier, row)
			if err != nil {
				return err
			}
			res.QanMysqlPerfschema = agent.(*inventorypb.QANMySQLPerfSchemaAgent)
		}

		if req.QanMysqlSlowlog {
			row, err = models.CreateAgent(tx.Querier, models.QANMySQLSlowlogAgentType, &models.CreateAgentParams{
				PMMAgentID:            req.PmmAgentId,
				ServiceID:             service.ServiceID,
				Username:              req.Username,
				Password:              req.Password,
				TLS:                   req.Tls,
				TLSSkipVerify:         req.TlsSkipVerify,
				QueryExamplesDisabled: req.DisableQueryExamples,
				MaxQueryLogSize:       req.MaxSlowlogFileSize,
			})
			if err != nil {
				return err
			}

			agent, err = services.ToAPIAgent(tx.Querier, row)
			if err != nil {
				return err
			}
			res.QanMysqlSlowlog = agent.(*inventorypb.QANMySQLSlowlogAgent)
		}

		return nil
	}); e != nil {
		return nil, e
	}

	s.registry.SendSetStateRequest(ctx, req.PmmAgentId)
	return res, nil
}
