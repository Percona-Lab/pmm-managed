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

	inventorypb "github.com/percona/pmm/api/inventory"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// ServicesService works with inventory API Services.
type ServicesService struct {
	db *reform.DB
	r  registry
}

// NewServicesService creates new ServicesService
func NewServicesService(db *reform.DB, r registry) *ServicesService {
	return &ServicesService{
		db: db,
		r:  r,
	}
}

// List selects all Services in a stable order.
//nolint:unparam
func (ss *ServicesService) List(ctx context.Context) ([]inventorypb.Service, error) {
	services := make([]*models.Service, 0)
	e := ss.db.InTransaction(func(tx *reform.TX) error {
		var err error
		services, err = models.FindAllServices(tx.Querier)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	return models.ToInventoryServices(services)
}

// Get selects a single Service by ID.
//nolint:unparam
func (ss *ServicesService) Get(ctx context.Context, id string) (inventorypb.Service, error) {
	service := new(models.Service)
	e := ss.db.InTransaction(func(tx *reform.TX) error {
		var err error
		service, err = models.FindServiceByID(tx.Querier, id)
		if err != nil {
			return err
		}
		return nil
	})

	if e != nil {
		return nil, e
	}

	return models.ToInventoryService(service)
}

// AddMySQL inserts MySQL Service with given parameters.
//nolint:dupl,unparam
func (ss *ServicesService) AddMySQL(ctx context.Context, params *models.AddDBMSServiceParams) (*inventorypb.MySQLService, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// Both address and socket can't be empty, etc.

	service := new(models.Service)
	e := ss.db.InTransaction(func(tx *reform.TX) error {
		var err error
		service, err = models.AddNewService(tx.Querier, models.MySQLServiceType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	res, err := models.ToInventoryService(service)
	if err != nil {
		return nil, err
	}
	return res.(*inventorypb.MySQLService), nil
}

// AddMongoDB inserts MongoDB Service with given parameters.
//nolint:dupl,unparam
func (ss *ServicesService) AddMongoDB(ctx context.Context, params *models.AddDBMSServiceParams) (*inventorypb.MongoDBService, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416

	service := new(models.Service)
	e := ss.db.InTransaction(func(tx *reform.TX) error {
		var err error
		service, err = models.AddNewService(tx.Querier, models.MongoDBServiceType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	res, err := models.ToInventoryService(service)
	if err != nil {
		return nil, err
	}
	return res.(*inventorypb.MongoDBService), nil
}

// AddPostgreSQL inserts PostgreSQL Service with given parameters.
//nolint:dupl,unparam
func (ss *ServicesService) AddPostgreSQL(ctx context.Context, params *models.AddDBMSServiceParams) (*inventorypb.PostgreSQLService, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// Both address and socket can't be empty, etc.

	service := new(models.Service)
	e := ss.db.InTransaction(func(tx *reform.TX) error {
		var err error
		service, err = models.AddNewService(tx.Querier, models.PostgreSQLServiceType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	res, err := models.ToInventoryService(service)
	if err != nil {
		return nil, err
	}
	return res.(*inventorypb.PostgreSQLService), nil
}

// Remove deletes Service by ID.
//nolint:unparam
func (ss *ServicesService) Remove(ctx context.Context, id string) error {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0.

	// TODO check absence of Agents

	e := ss.db.InTransaction(func(tx *reform.TX) error {
		err := models.RemoveService(tx.Querier, id)
		if err != nil {
			return err
		}
		return nil
	})

	return e
}
