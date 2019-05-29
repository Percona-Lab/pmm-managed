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

	"github.com/percona/pmm/api/managementpb"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/action"
	"github.com/percona/pmm-managed/services/agents"
)

type actionsServer struct {
	r  *agents.Registry
	db *reform.DB
}

// NewActionsServer creates Management Actions Server.
func NewActionsServer(r *agents.Registry, db *reform.DB) managementpb.ActionsServer {
	return &actionsServer{r, db}
}

// GetAction gets an action result.
//nolint:lll
func (s *actionsServer) GetAction(ctx context.Context, req *managementpb.GetActionRequest) (*managementpb.GetActionResponse, error) {
	res, err := models.FindActionResultByID(s.db.Querier, req.ActionId)
	if err != nil {
		return nil, err
	}

	return &managementpb.GetActionResponse{
		ActionId:   res.ID,
		PmmAgentId: res.PmmAgentID,
		Done:       res.Done,
		Error:      res.Error,
		Output:     res.Output,
	}, nil
}

// StartPTSummaryAction starts pt-summary action.
//nolint:lll,dupl
func (s *actionsServer) StartPTSummaryAction(ctx context.Context, req *managementpb.StartPTSummaryActionRequest) (*managementpb.StartPTSummaryActionResponse, error) {
	res, err := models.CreateActionResult(s.db.Querier, req.PmmAgentId)
	if err != nil {
		return nil, err
	}

	a := &action.PtSummary{
		ID:         res.ID,
		NodeID:     req.NodeId,
		PMMAgentID: req.PmmAgentId,
		Args:       []string{},
	}

	ag, err := models.FindPMMAgentsForNode(s.db.Querier, a.NodeID)
	if err != nil {
		return nil, err
	}

	a.PMMAgentID, err = models.FindPmmAgentIDToRunAction(a.PMMAgentID, ag)
	if err != nil {
		return nil, err
	}

	err = s.r.StartPTSummaryAction(ctx, a)
	if err != nil {
		return nil, err
	}

	return &managementpb.StartPTSummaryActionResponse{
		PmmAgentId: a.PMMAgentID,
		ActionId:   a.ID,
	}, nil
}

// StartPTMySQLSummaryAction starts pt-mysql-summary action.
//nolint:lll,dupl
func (s *actionsServer) StartPTMySQLSummaryAction(ctx context.Context, req *managementpb.StartPTMySQLSummaryActionRequest) (*managementpb.StartPTMySQLSummaryActionResponse, error) {
	res, err := models.CreateActionResult(s.db.Querier, req.PmmAgentId)
	if err != nil {
		return nil, err
	}

	a := &action.PtMySQLSummary{
		ID:         res.ID,
		ServiceID:  req.ServiceId,
		PMMAgentID: req.PmmAgentId,
		Args:       []string{},
	}

	ag, err := models.FindPMMAgentsForService(s.db.Querier, a.ServiceID)
	if err != nil {
		return nil, err
	}

	a.PMMAgentID, err = models.FindPmmAgentIDToRunAction(a.PMMAgentID, ag)
	if err != nil {
		return nil, err
	}

	err = s.r.StartPTMySQLSummaryAction(ctx, a)
	if err != nil {
		return nil, err
	}

	return &managementpb.StartPTMySQLSummaryActionResponse{
		PmmAgentId: a.PMMAgentID,
		ActionId:   a.ID,
	}, nil
}

// StartMySQLExplainAction starts mysql-explain action.
//nolint:lll,dupl
func (s *actionsServer) StartMySQLExplainAction(ctx context.Context, req *managementpb.StartMySQLExplainActionRequest) (*managementpb.StartMySQLExplainActionResponse, error) {
	res, err := models.CreateActionResult(s.db.Querier, req.PmmAgentId)
	if err != nil {
		return nil, err
	}

	a := &action.MySQLExplain{
		ID:         res.ID,
		ServiceID:  req.ServiceId,
		PMMAgentID: req.PmmAgentId,
		Query:      req.Query,
	}

	ag, err := models.FindPMMAgentsForService(s.db.Querier, a.ServiceID)
	if err != nil {
		return nil, err
	}

	a.PMMAgentID, err = models.FindPmmAgentIDToRunAction(a.PMMAgentID, ag)
	if err != nil {
		return nil, err
	}

	a.Dsn, err = models.FindDSNByServiceID(s.db.Querier, a.ServiceID, req.Database)
	if err != nil {
		return nil, err
	}

	err = s.r.StartMySQLExplainAction(ctx, a)
	if err != nil {
		return nil, err
	}

	return &managementpb.StartMySQLExplainActionResponse{
		PmmAgentId: a.PMMAgentID,
		ActionId:   a.ID,
	}, nil
}

// StartMySQLExplainJSONAction starts mysql-explain json action.
//nolint:lll,dupl
func (s *actionsServer) StartMySQLExplainJSONAction(ctx context.Context, req *managementpb.StartMySQLExplainJSONActionRequest) (*managementpb.StartMySQLExplainJSONActionResponse, error) {
	res, err := models.CreateActionResult(s.db.Querier, req.PmmAgentId)
	if err != nil {
		return nil, err
	}

	a := &action.MySQLExplainJSON{
		ID:         res.ID,
		ServiceID:  req.ServiceId,
		PMMAgentID: req.PmmAgentId,
		Query:      req.Query,
	}

	ag, err := models.FindPMMAgentsForService(s.db.Querier, a.ServiceID)
	if err != nil {
		return nil, err
	}

	a.PMMAgentID, err = models.FindPmmAgentIDToRunAction(a.PMMAgentID, ag)
	if err != nil {
		return nil, err
	}

	a.Dsn, err = models.FindDSNByServiceID(s.db.Querier, a.ServiceID, req.Database)
	if err != nil {
		return nil, err
	}

	err = s.r.StartMySQLExplainJSONAction(ctx, a)
	if err != nil {
		return nil, err
	}

	return &managementpb.StartMySQLExplainJSONActionResponse{
		PmmAgentId: a.PMMAgentID,
		ActionId:   a.ID,
	}, nil
}

// CancelAction stops an Action.
//nolint:lll
func (s *actionsServer) CancelAction(ctx context.Context, req *managementpb.CancelActionRequest) (*managementpb.CancelActionResponse, error) {
	ar, err := models.FindActionResultByID(s.db.Querier, req.ActionId)
	if err != nil {
		return nil, err
	}

	err = s.r.StopAction(ctx, ar.ID)
	if err != nil {
		return nil, err
	}

	return &managementpb.CancelActionResponse{}, nil
}
