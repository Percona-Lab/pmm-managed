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
	"github.com/google/uuid"
	inventorypb "github.com/percona/pmm/api/inventory"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// ServicesService works with inventory API Services.
type ServicesService struct {
	q  *reform.Querier
	r  registry
	ns *NodesService
}

// NewServicesService creates new ServicesService
func NewServicesService(q *reform.Querier, r registry, ns *NodesService) *ServicesService {
	return &ServicesService{
		q:  q,
		r:  r,
		ns: ns,
	}
}

// makeService converts database row to Inventory API Service.
func makeService(row *models.Service) (inventorypb.Service, error) {
	labels, err := row.GetCustomLabels()
	if err != nil {
		return nil, err
	}

	switch row.ServiceType {
	case models.MySQLServiceType:
		return &inventorypb.MySQLService{
			ServiceId:    row.ServiceID,
			ServiceName:  row.ServiceName,
			NodeId:       row.NodeID,
			Address:      pointer.GetString(row.Address),
			Port:         uint32(pointer.GetUint16(row.Port)),
			CustomLabels: labels,
		}, nil
	case models.MongoDBServiceType:
		return &inventorypb.MongoDBService{
			ServiceId:    row.ServiceID,
			ServiceName:  row.ServiceName,
			NodeId:       row.NodeID,
			Address:      pointer.GetString(row.Address),
			Port:         uint32(pointer.GetUint16(row.Port)),
			CustomLabels: labels,
		}, nil

	default:
		panic(fmt.Errorf("unhandled Service type %s", row.ServiceType))
	}
}

func (ss *ServicesService) get(ctx context.Context, id string) (*models.Service, error) {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty Service ID.")
	}

	row := &models.Service{ServiceID: id}
	switch err := ss.q.Reload(row); err {
	case nil:
		return row, nil
	case reform.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "Service with ID %q not found.", id)
	default:
		return nil, errors.WithStack(err)
	}
}

//nolint:unparam
func (ss *ServicesService) checkUniqueID(ctx context.Context, id string, external *reform.Querier) error {
	if id == "" {
		panic("empty Service ID")
	}

	querier := ss.q
	if external != nil {
		querier = external
	}

	row := &models.Service{ServiceID: id}
	switch err := querier.Reload(row); err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "Service with ID %q already exists.", id)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}

//nolint:unparam
func (ss *ServicesService) checkUniqueName(ctx context.Context, name string, external *reform.Querier) error {
	querier := ss.q
	if external != nil {
		querier = external
	}

	_, err := querier.FindOneFrom(models.ServiceTable, "service_name", name)
	switch err {
	case nil:
		return status.Errorf(codes.AlreadyExists, "Service with name %q already exists.", name)
	case reform.ErrNoRows:
		return nil
	default:
		return errors.WithStack(err)
	}
}

// List selects all Services in a stable order.
func (ss *ServicesService) List(ctx context.Context) ([]inventorypb.Service, error) {
	structs, err := ss.q.SelectAllFrom(models.ServiceTable, "ORDER BY service_id")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	res := make([]inventorypb.Service, len(structs))
	for i, str := range structs {
		row := str.(*models.Service)
		res[i], err = makeService(row)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

// Get selects a single Service by ID.
func (ss *ServicesService) Get(ctx context.Context, id string) (inventorypb.Service, error) {
	row, err := ss.get(ctx, id)
	if err != nil {
		return nil, err
	}
	return makeService(row)
}

// AddMySQL inserts MySQL Service with given parameters.
func (ss *ServicesService) AddMySQL(
	ctx context.Context,
	name,
	nodeID string,
	address *string,
	port *uint16,
	external *reform.Querier) (*inventorypb.MySQLService, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// Both address and socket can't be empty, etc.

	querier := ss.q
	if external != nil {
		querier = external
	}

	id := "/service_id/" + uuid.New().String()
	if err := ss.checkUniqueID(ctx, id, querier); err != nil {
		return nil, err
	}
	if err := ss.checkUniqueName(ctx, name, querier); err != nil {
		return nil, err
	}

	if _, err := ss.ns.Get(ctx, nodeID, querier); err != nil {
		return nil, err
	}

	row := &models.Service{
		ServiceID:   id,
		ServiceType: models.MySQLServiceType,
		ServiceName: name,
		NodeID:      nodeID,
		Address:     address,
		Port:        port,
	}
	if err := querier.Insert(row); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := makeService(row)
	if err != nil {
		return nil, err
	}
	return res.(*inventorypb.MySQLService), nil
}

// AddMongoDB inserts MongoDB Service with given parameters.
func (ss *ServicesService) AddMongoDB(ctx context.Context, name, nodeID string, address *string, port *uint16) (*inventorypb.MongoDBService, error) {

	id := "/service_id/" + uuid.New().String()
	if err := ss.checkUniqueID(ctx, id, nil); err != nil {
		return nil, err
	}
	if err := ss.checkUniqueName(ctx, name, nil); err != nil {
		return nil, err
	}

	if _, err := ss.ns.Get(ctx, nodeID, nil); err != nil {
		return nil, err
	}

	row := &models.Service{
		ServiceID:   id,
		ServiceType: models.MongoDBServiceType,
		ServiceName: name,
		NodeID:      nodeID,
		Address:     address,
		Port:        port,
	}
	if err := ss.q.Insert(row); err != nil {
		return nil, errors.WithStack(err)
	}
	res, err := makeService(row)
	if err != nil {
		return nil, err
	}
	return res.(*inventorypb.MongoDBService), nil
}

// Change updates Service by ID.
func (ss *ServicesService) Change(ctx context.Context, id string, name string) (inventorypb.Service, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0, name is not empty and valid.

	if err := ss.checkUniqueName(ctx, name, nil); err != nil {
		return nil, err
	}

	row, err := ss.get(ctx, id)
	if err != nil {
		return nil, err
	}

	row.ServiceName = name
	if err = ss.q.Update(row); err != nil {
		return nil, errors.WithStack(err)
	}
	return makeService(row)
}

// Remove deletes Service by ID.
func (ss *ServicesService) Remove(ctx context.Context, id string) error {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0.

	// TODO check absence of Agents

	err := ss.q.Delete(&models.Service{ServiceID: id})
	if err == reform.ErrNoRows {
		return status.Errorf(codes.NotFound, "Service with ID %q not found.", id)
	}
	return errors.WithStack(err)
}
