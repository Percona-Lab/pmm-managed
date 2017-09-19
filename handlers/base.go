// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Package handlers implements gRPC API of pmm-managed.
package handlers

import (
	"golang.org/x/net/context"

	"github.com/percona/pmm-managed/api"
)

type BaseServer struct {
	PMMVersion string
}

func (s *BaseServer) Version(context.Context, *api.BaseVersionRequest) (*api.BaseVersionResponse, error) {
	return &api.BaseVersionResponse{
		Version: s.PMMVersion,
	}, nil
}

// check interface
var _ api.BaseServer = (*BaseServer)(nil)
