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
	"github.com/percona/pmm/api/inventorypb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

type NodesService struct {
	db *reform.DB
}

// NewNodesService returns Inventory API handler for managing Nodes.
func NewNodesService(db *reform.DB) *NodesService {
	return &NodesService{
		db: db,
	}
}

// List returns a list of all Nodes.
//nolint:unparam
func (s *NodesService) List(ctx context.Context, req *inventorypb.ListNodesRequest) ([]inventorypb.Node, error) {
	allNodes := make([]*models.Node, 0)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		allNodes, err = models.FindAllNodes(tx.Querier)
		if err != nil {
			return status.Error(codes.NotFound, "nodes wasn't found")
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	nodes, err := ToInventoryNodes(allNodes)
	if err != nil {
		return nil, err
	}

	return nodes, nil
}

// Get returns a single Node by ID.
//nolint:unparam
func (s *NodesService) Get(ctx context.Context, req *inventorypb.GetNodeRequest) (inventorypb.Node, error) {
	modelNode := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		modelNode, err = models.FindNodeByID(tx.Querier, req.NodeId)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	node, err := ToInventoryNode(modelNode)
	if err != nil {
		return nil, err
	}

	return node, nil
}

// AddGenericNode adds Generic Node.
//nolint:unparam
func (s *NodesService) AddGenericNode(ctx context.Context, req *inventorypb.AddGenericNodeRequest) (*inventorypb.GenericNode, error) {
	params := &models.CreateNodeParams{
		NodeName:     req.NodeName,
		MachineID:    pointer.ToStringOrNil(req.MachineId),
		Distro:       pointer.ToStringOrNil(req.Distro),
		CustomLabels: req.CustomLabels,
		Address:      pointer.ToStringOrNil(req.Address),
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.CreateNode(tx.Querier, models.GenericNodeType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	return invNode.(*inventorypb.GenericNode), nil
}

// AddContainerNode adds Container Node.
//nolint:unparam
func (s *NodesService) AddContainerNode(ctx context.Context, req *inventorypb.AddContainerNodeRequest) (*inventorypb.ContainerNode, error) {
	params := &models.CreateNodeParams{
		NodeName:      req.NodeName,
		MachineID:     pointer.ToStringOrNil(req.MachineId),
		ContainerID:   pointer.ToStringOrNil(req.ContainerId),
		ContainerName: pointer.ToStringOrNil(req.ContainerName),
		CustomLabels:  req.CustomLabels,
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.CreateNode(tx.Querier, models.ContainerNodeType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	return invNode.(*inventorypb.ContainerNode), nil
}

// AddRemoteNode adds Remote Node.
//nolint:unparam
func (s *NodesService) AddRemoteNode(ctx context.Context, req *inventorypb.AddRemoteNodeRequest) (*inventorypb.RemoteNode, error) {
	params := &models.CreateNodeParams{
		NodeName:     req.NodeName,
		CustomLabels: req.CustomLabels,
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.CreateNode(tx.Querier, models.RemoteNodeType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	return invNode.(*inventorypb.RemoteNode), nil
}

// AddRemoteAmazonRDSNode adds Amazon (AWS) RDS remote Node.
//nolint:lll,unparam
func (s *NodesService) AddRemoteAmazonRDSNode(ctx context.Context, req *inventorypb.AddRemoteAmazonRDSNodeRequest) (*inventorypb.RemoteAmazonRDSNode, error) {
	params := &models.CreateNodeParams{
		NodeName:     req.NodeName,
		Address:      &req.Instance,
		Region:       &req.Region,
		CustomLabels: req.CustomLabels,
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.CreateNode(tx.Querier, models.RemoteAmazonRDSNodeType, params)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	return invNode.(*inventorypb.RemoteAmazonRDSNode), nil
}

// Remove removes Node without any Agents and Services.
//nolint:unparam
func (s *NodesService) Remove(ctx context.Context, req *inventorypb.RemoveNodeRequest) (*inventorypb.RemoveNodeResponse, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0.

	// TODO check absence of Services and Agents

	e := s.db.InTransaction(func(tx *reform.TX) error {
		err := models.RemoveNode(tx.Querier, req.NodeId)
		if err != nil {
			return err
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	return new(inventorypb.RemoveNodeResponse), nil
}

// ToInventoryNode converts database row to Inventory API Node.
func ToInventoryNode(row *models.Node) (inventorypb.Node, error) {
	labels, err := row.GetCustomLabels()
	if err != nil {
		return nil, err
	}

	switch row.NodeType {
	case models.GenericNodeType:
		return &inventorypb.GenericNode{
			NodeId:       row.NodeID,
			NodeName:     row.NodeName,
			MachineId:    pointer.GetString(row.MachineID),
			Distro:       pointer.GetString(row.Distro),
			CustomLabels: labels,
			Address:      pointer.GetString(row.Address),
		}, nil

	case models.ContainerNodeType:
		return &inventorypb.ContainerNode{
			NodeId:        row.NodeID,
			NodeName:      row.NodeName,
			MachineId:     pointer.GetString(row.MachineID),
			ContainerId:   pointer.GetString(row.ContainerID),
			ContainerName: pointer.GetString(row.ContainerName),
			CustomLabels:  labels,
			Address:       pointer.GetString(row.Address),
		}, nil

	case models.RemoteNodeType:
		return &inventorypb.RemoteNode{
			NodeId:       row.NodeID,
			NodeName:     row.NodeName,
			CustomLabels: labels,
		}, nil

	case models.RemoteAmazonRDSNodeType:
		return &inventorypb.RemoteAmazonRDSNode{
			NodeId:       row.NodeID,
			NodeName:     row.NodeName,
			Instance:     pointer.GetString(row.Address),
			Region:       pointer.GetString(row.Region),
			CustomLabels: labels,
		}, nil

	default:
		panic(fmt.Errorf("unhandled Node type %s", row.NodeType))
	}
}

// ToInventoryNodes converts database rows to Inventory API Nodes.
func ToInventoryNodes(nodes []*models.Node) ([]inventorypb.Node, error) {
	var err error
	res := make([]inventorypb.Node, len(nodes))
	for i, n := range nodes {
		res[i], err = ToInventoryNode(n)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
