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
	"github.com/percona/pmm/api/inventory"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/agents"
)

// AgentsService works with inventory API Agents.
type AgentsService struct {
	q *reform.Querier
	r *agents.Registry
}

func NewAgentsService(q *reform.Querier, r *agents.Registry) *AgentsService {
	return &AgentsService{
		q: q,
		r: r,
	}
}

// makeAgent converts database row to Inventory API Agent.
func (as *AgentsService) makeAgent(ctx context.Context, row *models.AgentRow) (inventory.Agent, error) {
	switch row.Type {
	case models.PMMAgentType:
		return &inventory.PMMAgent{
			Id: row.ID,
			HostNodeInfo: &inventory.HostNodeInfo{
				NodeId:            row.RunsOnNodeID,
				ContainerId:       "TODO",
				ContainerName:     "TODO",
				KubernetesPodUid:  "TODO",
				KubernetesPodName: "TODO",
			},
			Running: !row.Disabled,
		}, nil

	case models.NodeExporterAgentType:
		return &inventory.NodeExporter{
			Id: row.ID,
			HostNodeInfo: &inventory.HostNodeInfo{
				NodeId:            row.RunsOnNodeID,
				ContainerId:       "TODO",
				ContainerName:     "TODO",
				KubernetesPodUid:  "TODO",
				KubernetesPodName: "TODO",
			},
			Disabled: row.Disabled,
		}, nil

	case models.MySQLdExporterAgentType:
		var agentService models.AgentService
		if err := as.q.FindOneTo(&agentService, "agent_id", row.ID); err != nil {
			return nil, errors.WithStack(err)
		}

		return &inventory.MySQLdExporter{
			Id: row.ID,
			HostNodeInfo: &inventory.HostNodeInfo{
				NodeId:            row.RunsOnNodeID,
				ContainerId:       "TODO",
				ContainerName:     "TODO",
				KubernetesPodUid:  "TODO",
				KubernetesPodName: "TODO",
			},
			Disabled:  row.Disabled,
			ServiceId: agentService.ServiceID,
			Username:  pointer.GetString(row.ServiceUsername),
		}, nil

	default:
		panic(fmt.Errorf("unhandled AgentRow type %s", row.Type))
	}
}

func (as *AgentsService) get(ctx context.Context, id string) (*models.AgentRow, error) {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty Agent ID.")
	}

	row := &models.AgentRow{ID: id}
	switch err := as.q.Reload(row); err {
	case nil:
		return row, nil
	case reform.ErrNoRows:
		return nil, status.Errorf(codes.NotFound, "Agent with ID %q not found.", id)
	default:
		return nil, errors.WithStack(err)
	}
}

// List selects all Agents in a stable order for a given service.
func (as *AgentsService) List(ctx context.Context, serviceID string) ([]inventory.Agent, error) {
	filters := models.AgentFilters{}
	if serviceID != "" {
		filters.ServiceID = &serviceID
	}
	agentRows, err := models.AgentsByFilters(as.q, filters)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// TODO That loop makes len(agentRows) SELECTs, that can be slow. Optimize when needed.
	res := make([]inventory.Agent, len(agentRows))
	for i, row := range agentRows {
		agent, err := as.makeAgent(ctx, row)
		if err != nil {
			return nil, err
		}
		res[i] = agent
	}
	return res, nil
}

// Get selects a single Agent by ID.
func (as *AgentsService) Get(ctx context.Context, id string) (inventory.Agent, error) {
	row, err := as.get(ctx, id)
	if err != nil {
		return nil, err
	}
	return as.makeAgent(ctx, row)
}

// AddPMMAgent inserts pmm-agent Agent with given parameters.
func (as *AgentsService) AddPMMAgent(ctx context.Context, nodeID string) (inventory.Agent, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// TODO Check runs-on Node: it must be BM, VM, DC (i.e. not remote, AWS RDS, etc.)

	ns := NewNodesService(as.q)
	if _, err := ns.get(ctx, nodeID); err != nil {
		return nil, err
	}

	row := &models.AgentRow{
		ID:           makeID(),
		Type:         models.PMMAgentType,
		RunsOnNodeID: nodeID,
		Disabled:     true, // not connected when we create it
	}
	if err := as.q.Insert(row); err != nil {
		return nil, errors.WithStack(err)
	}

	err := as.q.Insert(&models.AgentNode{
		AgentID: row.ID,
		NodeID:  nodeID,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	agent, err := as.makeAgent(ctx, row)
	return agent, err
}

// AddNodeExporter inserts node_exporter Agent with given parameters.
func (as *AgentsService) AddNodeExporter(ctx context.Context, nodeID string, disabled bool) (inventory.Agent, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// TODO Check runs-on Node: it must be BM, VM, DC (i.e. not remote, AWS RDS, etc.)

	ns := NewNodesService(as.q)
	if _, err := ns.get(ctx, nodeID); err != nil {
		return nil, err
	}

	row := &models.AgentRow{
		ID:           makeID(),
		Type:         models.NodeExporterAgentType,
		RunsOnNodeID: nodeID,
		Disabled:     disabled,
	}
	if err := as.q.Insert(row); err != nil {
		return nil, errors.WithStack(err)
	}

	err := as.q.Insert(&models.AgentNode{
		AgentID: row.ID,
		NodeID:  nodeID,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return as.makeAgent(ctx, row)
}

// AddMySQLdExporter inserts mysqld_exporter Agent with given parameters.
func (as *AgentsService) AddMySQLdExporter(ctx context.Context, nodeID string, disabled bool, serviceID string, username, password *string) (inventory.Agent, error) {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// TODO Check runs-on Node: it must be BM, VM, DC (i.e. not remote, AWS RDS, etc.)

	ns := NewNodesService(as.q)
	if _, err := ns.get(ctx, nodeID); err != nil {
		return nil, err
	}

	ss := NewServicesService(as.q)
	if _, err := ss.get(ctx, serviceID); err != nil {
		return nil, err
	}

	row := &models.AgentRow{
		ID:              makeID(),
		Type:            models.MySQLdExporterAgentType,
		RunsOnNodeID:    nodeID,
		Disabled:        disabled,
		ServiceUsername: username,
		ServicePassword: password,
	}
	if err := as.q.Insert(row); err != nil {
		return nil, errors.WithStack(err)
	}

	err := as.q.Insert(&models.AgentNode{
		AgentID: row.ID,
		NodeID:  nodeID,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	err = as.q.Insert(&models.AgentService{
		AgentID:   row.ID,
		ServiceID: serviceID,
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return as.makeAgent(ctx, row)
}

// SetDisabled enables or disables Agent by ID.
func (as *AgentsService) SetDisabled(ctx context.Context, id string, disabled bool) error {
	row, err := as.get(ctx, id)
	if err != nil {
		return err
	}

	row.Disabled = disabled
	err = as.q.Update(row)
	return errors.WithStack(err)
}

// Remove deletes Agent by ID.
func (as *AgentsService) Remove(ctx context.Context, id string) error {
	// TODO Decide about validation. https://jira.percona.com/browse/PMM-1416
	// ID is not 0.

	row, err := as.get(ctx, id)
	if err != nil {
		return err
	}

	if _, err = as.q.DeleteFrom(models.AgentServiceView, "WHERE agent_id = "+as.q.Placeholder(1), id); err != nil {
		return errors.WithStack(err)
	}
	if _, err = as.q.DeleteFrom(models.AgentNodeView, "WHERE agent_id = "+as.q.Placeholder(1), id); err != nil {
		return errors.WithStack(err)
	}

	if err = as.q.Delete(row); err != nil {
		return errors.WithStack(err)
	}

	if row.Type == models.PMMAgentType {
		as.r.Kick(ctx, id)
	}
	return nil
}
