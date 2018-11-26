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

package handlers

import (
	"context"
	"fmt"

	api "github.com/percona/pmm/api/inventory"

	"github.com/percona/pmm-managed/services/inventory"
)

// ServicesServer handles Inventory API requests to manage Services.
type ServicesServer struct {
	Services *inventory.ServicesService
}

// ListServices returns a list of all Services.
func (s *ServicesServer) ListServices(ctx context.Context, req *api.ListServicesRequest) (*api.ListServicesResponse, error) {
	services, err := s.Services.List(ctx)
	if err != nil {
		return nil, err
	}

	res := new(api.ListServicesResponse)
	for _, service := range services {
		switch service := service.(type) {
		case *api.MySQLService:
			res.Mysql = append(res.Mysql, service)
		default:
			panic(fmt.Errorf("unhandled inventory Service type %T", service))
		}
	}
	return res, nil
}

// GetService returns a single Service by ID.
func (s *ServicesServer) GetService(ctx context.Context, req *api.GetServiceRequest) (*api.GetServiceResponse, error) {
	service, err := s.Services.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	res := new(api.GetServiceResponse)
	switch service := service.(type) {
	case *api.MySQLService:
		res.Service = &api.GetServiceResponse_Mysql{Mysql: service}
	default:
		panic(fmt.Errorf("unhandled inventory Service type %T", service))
	}
	return res, nil
}

// AddMySQLService adds MySQL Service.
func (s *ServicesServer) AddMySQLService(ctx context.Context, req *api.AddMySQLServiceRequest) (*api.AddMySQLServiceResponse, error) {
	panic("not implemented")
}

// ChangeMySQLService changes MySQL Service.
func (s *ServicesServer) ChangeMySQLService(ctx context.Context, req *api.ChangeMySQLServiceRequest) (*api.ChangeMySQLServiceResponse, error) {
	panic("not implemented")
}

// RemoveService removes Service without any Agents.
func (s *ServicesServer) RemoveService(ctx context.Context, req *api.RemoveServiceRequest) (*api.RemoveServiceResponse, error) {
	if err := s.Services.Remove(ctx, req.Id); err != nil {
		return nil, err
	}

	return new(api.RemoveServiceResponse), nil
}

// check interfaces
var (
	_ api.ServicesServer = (*ServicesServer)(nil)
)
