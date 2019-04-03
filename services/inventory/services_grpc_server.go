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

package inventory

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	inventorypb "github.com/percona/pmm/api/inventory"

	"github.com/percona/pmm-managed/models"
)

type servicesGrpcServer struct {
	s *ServicesService
}

// NewServicesGrpcServer returns Inventory API handler for managing Services.
func NewServicesGrpcServer(s *ServicesService) inventorypb.ServicesServer {
	return &servicesGrpcServer{s}
}

// ListServices returns a list of all Services.
func (s *servicesGrpcServer) ListServices(ctx context.Context, req *inventorypb.ListServicesRequest) (*inventorypb.ListServicesResponse, error) {
	services, err := s.s.List(ctx)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.ListServicesResponse)
	for _, service := range services {
		switch service := service.(type) {
		case *inventorypb.MySQLService:
			res.Mysql = append(res.Mysql, service)
		case *inventorypb.AmazonRDSMySQLService:
			res.AmazonRdsMysql = append(res.AmazonRdsMysql, service)
		case *inventorypb.MongoDBService:
			res.Mongodb = append(res.Mongodb, service)
		default:
			panic(fmt.Errorf("unhandled inventory Service type %T", service))
		}
	}
	return res, nil
}

// GetService returns a single Service by ID.
func (s *servicesGrpcServer) GetService(ctx context.Context, req *inventorypb.GetServiceRequest) (*inventorypb.GetServiceResponse, error) {
	service, err := s.s.Get(ctx, req.ServiceId)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.GetServiceResponse)
	switch service := service.(type) {
	case *inventorypb.MySQLService:
		res.Service = &inventorypb.GetServiceResponse_Mysql{Mysql: service}
	case *inventorypb.AmazonRDSMySQLService:
		res.Service = &inventorypb.GetServiceResponse_AmazonRdsMysql{AmazonRdsMysql: service}
	case *inventorypb.MongoDBService:
		res.Service = &inventorypb.GetServiceResponse_Mongodb{Mongodb: service}
	default:
		panic(fmt.Errorf("unhandled inventory Service type %T", service))
	}
	return res, nil
}

// AddMySQLService adds MySQL Service.
func (s *servicesGrpcServer) AddMySQLService(ctx context.Context, req *inventorypb.AddMySQLServiceRequest) (*inventorypb.AddMySQLServiceResponse, error) {
	service, err := s.s.AddMySQL(ctx, &models.AddDBMSServiceParams{
		ServiceName:  req.ServiceName,
		NodeID:       req.NodeId,
		Address:      pointer.ToStringOrNil(req.Address),
		Port:         pointer.ToUint16OrNil(uint16(req.Port)),
		CustomLabels: req.CustomLabels,
	})
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddMySQLServiceResponse{
		Mysql: service,
	}
	return res, nil
}

// AddAmazonRDSMySQLService adds AmazonRDSMySQL Service.
func (s *servicesGrpcServer) AddAmazonRDSMySQLService(ctx context.Context, req *inventorypb.AddAmazonRDSMySQLServiceRequest) (*inventorypb.AddAmazonRDSMySQLServiceResponse, error) {
	panic("not implemented yet")
}

func (s *servicesGrpcServer) AddMongoDBService(ctx context.Context, req *inventorypb.AddMongoDBServiceRequest) (*inventorypb.AddMongoDBServiceResponse, error) {
	service, err := s.s.AddMongoDB(ctx, &models.AddDBMSServiceParams{
		ServiceName:  req.ServiceName,
		NodeID:       req.NodeId,
		Address:      pointer.ToStringOrNil(req.Address),
		Port:         pointer.ToUint16OrNil(uint16(req.Port)),
		CustomLabels: req.CustomLabels,
	})
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddMongoDBServiceResponse{
		Mongodb: service,
	}
	return res, nil
}

func (s *servicesGrpcServer) AddPostgreSQLService(ctx context.Context, req *inventorypb.AddPostgreSQLServiceRequest) (*inventorypb.AddPostgreSQLServiceResponse, error) {
	service, err := s.s.AddPostgreSQL(ctx, &models.AddDBMSServiceParams{
		ServiceName:  req.ServiceName,
		NodeID:       req.NodeId,
		Address:      pointer.ToStringOrNil(req.Address),
		Port:         pointer.ToUint16OrNil(uint16(req.Port)),
		CustomLabels: req.CustomLabels,
	})
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddPostgreSQLServiceResponse{
		Postgresql: service,
	}
	return res, nil
}

// RemoveService removes Service without any Agents.
func (s *servicesGrpcServer) RemoveService(ctx context.Context, req *inventorypb.RemoveServiceRequest) (*inventorypb.RemoveServiceResponse, error) {
	if err := s.s.Remove(ctx, req.ServiceId); err != nil {
		return nil, err
	}

	return new(inventorypb.RemoveServiceResponse), nil
}
