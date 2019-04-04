// Copyright (C) 2019 Percona LLC
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

package grpc

import (
	"context"
	"fmt"

	inventorypb "github.com/percona/pmm/api/inventory"

	"github.com/percona/pmm-managed/services/inventory"
)

//nolint:unused
type agentsGrpcServer struct {
	s *inventory.AgentsService
}

// NewAgentsGrpcServer returns Inventory API handler for managing Agents.
func NewAgentsGrpcServer(s *inventory.AgentsService) inventorypb.AgentsServer {
	return &agentsGrpcServer{s}
}

// ListAgents returns a list of Agents for a given filters.
func (s *agentsGrpcServer) ListAgents(ctx context.Context, req *inventorypb.ListAgentsRequest) (*inventorypb.ListAgentsResponse, error) {
	filters := inventory.AgentFilters{
		PMMAgentID: req.GetPmmAgentId(),
		NodeID:     req.GetNodeId(),
		ServiceID:  req.GetServiceId(),
	}
	agents, err := s.s.List(ctx, filters)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.ListAgentsResponse)
	for _, agent := range agents {
		switch agent := agent.(type) {
		case *inventorypb.PMMAgent:
			res.PmmAgent = append(res.PmmAgent, agent)
		case *inventorypb.NodeExporter:
			res.NodeExporter = append(res.NodeExporter, agent)
		case *inventorypb.MySQLdExporter:
			res.MysqldExporter = append(res.MysqldExporter, agent)
		case *inventorypb.RDSExporter:
			res.RdsExporter = append(res.RdsExporter, agent)
		case *inventorypb.ExternalExporter:
			res.ExternalExporter = append(res.ExternalExporter, agent)
		case *inventorypb.MongoDBExporter:
			res.MongodbExporter = append(res.MongodbExporter, agent)
		case *inventorypb.QANMySQLPerfSchemaAgent:
			res.QanMysqlPerfschemaAgent = append(res.QanMysqlPerfschemaAgent, agent)
		case *inventorypb.PostgresExporter:
			res.PostgresExporter = append(res.PostgresExporter, agent)
		default:
			panic(fmt.Errorf("unhandled inventory Agent type %T", agent))
		}
	}
	return res, nil
}

// GetAgent returns a single Agent by ID.
func (s *agentsGrpcServer) GetAgent(ctx context.Context, req *inventorypb.GetAgentRequest) (*inventorypb.GetAgentResponse, error) {
	agent, err := s.s.Get(ctx, req.AgentId)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.GetAgentResponse)
	switch agent := agent.(type) {
	case *inventorypb.PMMAgent:
		res.Agent = &inventorypb.GetAgentResponse_PmmAgent{PmmAgent: agent}
	case *inventorypb.NodeExporter:
		res.Agent = &inventorypb.GetAgentResponse_NodeExporter{NodeExporter: agent}
	case *inventorypb.MySQLdExporter:
		res.Agent = &inventorypb.GetAgentResponse_MysqldExporter{MysqldExporter: agent}
	case *inventorypb.RDSExporter:
		res.Agent = &inventorypb.GetAgentResponse_RdsExporter{RdsExporter: agent}
	case *inventorypb.ExternalExporter:
		res.Agent = &inventorypb.GetAgentResponse_ExternalExporter{ExternalExporter: agent}
	case *inventorypb.MongoDBExporter:
		res.Agent = &inventorypb.GetAgentResponse_MongodbExporter{MongodbExporter: agent}
	case *inventorypb.QANMySQLPerfSchemaAgent:
		res.Agent = &inventorypb.GetAgentResponse_QanMysqlPerfschemaAgent{QanMysqlPerfschemaAgent: agent}
	case *inventorypb.PostgresExporter:
		res.Agent = &inventorypb.GetAgentResponse_PostgresExporter{PostgresExporter: agent}
	default:
		panic(fmt.Errorf("unhandled inventory Agent type %T", agent))
	}
	return res, nil

}

// AddPMMAgent adds pmm-agent Agent.
func (s *agentsGrpcServer) AddPMMAgent(ctx context.Context, req *inventorypb.AddPMMAgentRequest) (*inventorypb.AddPMMAgentResponse, error) {
	agent, err := s.s.AddPMMAgent(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddPMMAgentResponse{
		PmmAgent: agent,
	}
	return res, nil
}

func (s *agentsGrpcServer) ChangePMMAgent(context.Context, *inventorypb.ChangePMMAgentRequest) (*inventorypb.ChangePMMAgentResponse, error) {
	panic("not implemented")
}

// AddNodeExporter adds node_exporter Agent.
func (s *agentsGrpcServer) AddNodeExporter(ctx context.Context, req *inventorypb.AddNodeExporterRequest) (*inventorypb.AddNodeExporterResponse, error) {
	agent, err := s.s.AddNodeExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddNodeExporterResponse{
		NodeExporter: agent,
	}
	return res, nil
}

// ChangeNodeExporter changes disabled flag and custom labels of node_exporter Agent.
func (s *agentsGrpcServer) ChangeNodeExporter(ctx context.Context, req *inventorypb.ChangeNodeExporterRequest) (*inventorypb.ChangeNodeExporterResponse, error) {
	agent, err := s.s.ChangeNodeExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.ChangeNodeExporterResponse{
		NodeExporter: agent,
	}
	return res, nil
}

// AddMySQLdExporter adds mysqld_exporter Agent.
func (s *agentsGrpcServer) AddMySQLdExporter(ctx context.Context, req *inventorypb.AddMySQLdExporterRequest) (*inventorypb.AddMySQLdExporterResponse, error) {
	agent, err := s.s.AddMySQLdExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddMySQLdExporterResponse{
		MysqldExporter: agent,
	}
	return res, nil
}

// ChangeMySQLdExporter changes disabled flag and custom labels of mysqld_exporter Agent.
func (s *agentsGrpcServer) ChangeMySQLdExporter(ctx context.Context, req *inventorypb.ChangeMySQLdExporterRequest) (*inventorypb.ChangeMySQLdExporterResponse, error) {
	agent, err := s.s.ChangeMySQLdExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.ChangeMySQLdExporterResponse{
		MysqldExporter: agent,
	}
	return res, nil
}

// AddRDSExporter adds rds_exporter Agent.
func (s *agentsGrpcServer) AddRDSExporter(ctx context.Context, req *inventorypb.AddRDSExporterRequest) (*inventorypb.AddRDSExporterResponse, error) {
	panic("not implemented yet")
}

func (s *agentsGrpcServer) ChangeRDSExporter(context.Context, *inventorypb.ChangeRDSExporterRequest) (*inventorypb.ChangeRDSExporterResponse, error) {
	panic("not implemented")
}

// AddExternalExporter adds external Agent.
func (s *agentsGrpcServer) AddExternalExporter(ctx context.Context, req *inventorypb.AddExternalExporterRequest) (*inventorypb.AddExternalExporterResponse, error) {
	panic("not implemented yet")
}

func (s *agentsGrpcServer) ChangeExternalExporter(context.Context, *inventorypb.ChangeExternalExporterRequest) (*inventorypb.ChangeExternalExporterResponse, error) {
	panic("not implemented")
}

// AddMongoDBExporter adds mongodb_exporter Agent.
func (s *agentsGrpcServer) AddMongoDBExporter(ctx context.Context, req *inventorypb.AddMongoDBExporterRequest) (*inventorypb.AddMongoDBExporterResponse, error) {
	agent, err := s.s.AddMongoDBExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddMongoDBExporterResponse{
		MongodbExporter: agent,
	}
	return res, nil
}

// ChangeMongoDBExporter changes disabled flag and custom labels of mongo_exporter Agent.
func (s *agentsGrpcServer) ChangeMongoDBExporter(ctx context.Context, req *inventorypb.ChangeMongoDBExporterRequest) (*inventorypb.ChangeMongoDBExporterResponse, error) {
	agent, err := s.s.ChangeMongoDBExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.ChangeMongoDBExporterResponse{
		MongodbExporter: agent,
	}
	return res, nil
}

// AddQANMySQLPerfSchemaAgent adds MySQL PerfSchema QAN Agent.
//nolint:lll
func (s *agentsGrpcServer) AddQANMySQLPerfSchemaAgent(ctx context.Context, req *inventorypb.AddQANMySQLPerfSchemaAgentRequest) (*inventorypb.AddQANMySQLPerfSchemaAgentResponse, error) {
	agent, err := s.s.AddQANMySQLPerfSchemaAgent(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddQANMySQLPerfSchemaAgentResponse{
		QanMysqlPerfschemaAgent: agent,
	}
	return res, nil
}

// ChangeQANMySQLPerfSchemaAgent changes disabled flag and custom labels of MySQL PerfSchema QAN Agent.
//nolint:lll
func (s *agentsGrpcServer) ChangeQANMySQLPerfSchemaAgent(ctx context.Context, req *inventorypb.ChangeQANMySQLPerfSchemaAgentRequest) (*inventorypb.ChangeQANMySQLPerfSchemaAgentResponse, error) {
	agent, err := s.s.ChangeQANMySQLPerfSchemaAgent(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.ChangeQANMySQLPerfSchemaAgentResponse{
		QanMysqlPerfschemaAgent: agent,
	}
	return res, nil
}

// AddPostgresExporter adds postgres_exporter Agent.
func (s *agentsGrpcServer) AddPostgresExporter(ctx context.Context, req *inventorypb.AddPostgresExporterRequest) (*inventorypb.AddPostgresExporterResponse, error) {
	agent, err := s.s.AddPostgresExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddPostgresExporterResponse{
		PostgresExporter: agent,
	}
	return res, nil

}

// ChangePostgresExporter changes disabled flag and custom labels of postgres_exporter Agent.
//nolint:lll
func (s *agentsGrpcServer) ChangePostgresExporter(ctx context.Context, req *inventorypb.ChangePostgresExporterRequest) (*inventorypb.ChangePostgresExporterResponse, error) {
	agent, err := s.s.ChangePostgresExporter(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.ChangePostgresExporterResponse{
		PostgresExporter: agent,
	}
	return res, nil
}

// RemoveAgent removes Agent.
func (s *agentsGrpcServer) RemoveAgent(ctx context.Context, req *inventorypb.RemoveAgentRequest) (*inventorypb.RemoveAgentResponse, error) {
	if err := s.s.Remove(ctx, req.AgentId); err != nil {
		return nil, err
	}

	return new(inventorypb.RemoveAgentResponse), nil
}
