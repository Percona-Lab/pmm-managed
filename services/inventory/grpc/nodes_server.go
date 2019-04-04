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

package grpc

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	inventorypb "github.com/percona/pmm/api/inventory"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/inventory"
)

type nodesGrpcServer struct {
	db *reform.DB
}

// NewNodesGrpcServer returns Inventory API handler for managing Nodes.
func NewNodesGrpcServer(db *reform.DB) inventorypb.NodesServer {
	return &nodesGrpcServer{
		db: db,
	}
}

// ListNodes returns a list of all Nodes.
func (s *nodesGrpcServer) ListNodes(ctx context.Context, req *inventorypb.ListNodesRequest) (*inventorypb.ListNodesResponse, error) {

	allNodes := make([]*models.Node, 0)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		allNodes, err = models.FindAllNodes(tx.Querier)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	nodes, err := inventory.ToInventoryNodes(allNodes)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.ListNodesResponse)
	for _, node := range nodes {
		switch node := node.(type) {
		case *inventorypb.GenericNode:
			res.Generic = append(res.Generic, node)
		case *inventorypb.ContainerNode:
			res.Container = append(res.Container, node)
		case *inventorypb.RemoteNode:
			res.Remote = append(res.Remote, node)
		case *inventorypb.RemoteAmazonRDSNode:
			res.RemoteAmazonRds = append(res.RemoteAmazonRds, node)
		default:
			panic(fmt.Errorf("unhandled inventory Node type %T", node))
		}
	}
	return res, nil
}

// GetNode returns a single Node by ID.
func (s *nodesGrpcServer) GetNode(ctx context.Context, req *inventorypb.GetNodeRequest) (*inventorypb.GetNodeResponse, error) {
	modelNode := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		modelNode, err = models.FindNodeByID(tx.Querier, req.NodeId)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	node, err := inventory.ToInventoryNode(modelNode)
	if err != nil {
		return nil, err
	}

	res := new(inventorypb.GetNodeResponse)
	switch node := node.(type) {
	case *inventorypb.GenericNode:
		res.Node = &inventorypb.GetNodeResponse_Generic{Generic: node}
	case *inventorypb.ContainerNode:
		res.Node = &inventorypb.GetNodeResponse_Container{Container: node}
	case *inventorypb.RemoteNode:
		res.Node = &inventorypb.GetNodeResponse_Remote{Remote: node}
	case *inventorypb.RemoteAmazonRDSNode:
		res.Node = &inventorypb.GetNodeResponse_RemoteAmazonRds{RemoteAmazonRds: node}
	default:
		panic(fmt.Errorf("unhandled inventory Node type %T", node))
	}
	return res, nil
}

// AddGenericNode adds Generic Node.
func (s *nodesGrpcServer) AddGenericNode(ctx context.Context, req *inventorypb.AddGenericNodeRequest) (*inventorypb.AddGenericNodeResponse, error) {
	params := &models.AddNodeParams{
		NodeName:      req.NodeName,
		MachineID:     pointer.ToStringOrNil(req.MachineId),
		Distro:        pointer.ToStringOrNil(req.Distro),
		DistroVersion: pointer.ToStringOrNil(req.DistroVersion),
		CustomLabels:  req.CustomLabels,
		Address:       pointer.ToStringOrNil(req.Address),
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.AddNode(tx.Querier, models.GenericNodeType, params)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := inventory.ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddGenericNodeResponse{
		Generic: invNode.(*inventorypb.GenericNode),
	}
	return res, nil
}

// AddContainerNode adds Container Node.
func (s *nodesGrpcServer) AddContainerNode(ctx context.Context, req *inventorypb.AddContainerNodeRequest) (*inventorypb.AddContainerNodeResponse, error) {
	params := &models.AddNodeParams{
		NodeName:            req.NodeName,
		MachineID:           pointer.ToStringOrNil(req.MachineId),
		DockerContainerID:   pointer.ToStringOrNil(req.DockerContainerId),
		DockerContainerName: pointer.ToStringOrNil(req.DockerContainerName),
		CustomLabels:        req.CustomLabels,
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.AddNode(tx.Querier, models.ContainerNodeType, params)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := inventory.ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddContainerNodeResponse{
		Container: invNode.(*inventorypb.ContainerNode),
	}
	return res, nil
}

// AddRemoteNode adds Remote Node.
func (s *nodesGrpcServer) AddRemoteNode(ctx context.Context, req *inventorypb.AddRemoteNodeRequest) (*inventorypb.AddRemoteNodeResponse, error) {
	params := &models.AddNodeParams{
		NodeName:     req.NodeName,
		CustomLabels: req.CustomLabels,
	}

	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// No hostname for Container, etc.
	node := new(models.Node)
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		node, err = models.AddNode(tx.Querier, models.RemoteNodeType, params)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := inventory.ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddRemoteNodeResponse{
		Remote: invNode.(*inventorypb.RemoteNode),
	}
	return res, nil
}

// AddRemoteAmazonRDSNode adds Amazon (AWS) RDS remote Node.
//nolint:lll
func (s *nodesGrpcServer) AddRemoteAmazonRDSNode(ctx context.Context, req *inventorypb.AddRemoteAmazonRDSNodeRequest) (*inventorypb.AddRemoteAmazonRDSNodeResponse, error) {
	params := &models.AddNodeParams{
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
		node, err = models.AddNode(tx.Querier, models.RemoteAmazonRDSNodeType, params)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	invNode, err := inventory.ToInventoryNode(node)
	if err != nil {
		return nil, err
	}

	res := &inventorypb.AddRemoteAmazonRDSNodeResponse{
		RemoteAmazonRds: invNode.(*inventorypb.RemoteAmazonRDSNode),
	}
	return res, nil
}

// RemoveNode removes Node without any Agents and Services.
func (s *nodesGrpcServer) RemoveNode(ctx context.Context, req *inventorypb.RemoveNodeRequest) (*inventorypb.RemoveNodeResponse, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0.

	// TODO check absence of Services and Agents

	e := s.db.InTransaction(func(tx *reform.TX) error {
		err := models.RemoveNode(tx.Querier, req.NodeId)
		if err != nil {
			return err // TODO: Convert to gRPC errors
		}
		return nil
	})
	if e != nil {
		return nil, e
	}

	return new(inventorypb.RemoveNodeResponse), nil
}
