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

	"github.com/percona/pmm/api/managementpb"
)

//nolint:unused
type nodeGrpcServer struct {
	svc *NodeService
}

// NewManagementNodeGrpcServer creates Management Node Server.
func NewManagementNodeGrpcServer(s *NodeService) managementpb.NodeServer {
	return &nodeGrpcServer{svc: s}
}

// Register do registration of new Node.
func (s *nodeGrpcServer) Register(ctx context.Context, req *managementpb.RegisterNodeRequest) (res *managementpb.RegisterNodeResponse, err error) {
	return s.svc.Register(ctx, req)
}
