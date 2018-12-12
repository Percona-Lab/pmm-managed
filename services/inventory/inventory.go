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

// Package inventory contains inventory business logic: Nodes, Services, Agents.
package inventory

import (
	"github.com/google/uuid"
)

// TODO Decide about transactions.

// FIXME un-export MakeID when we remove old services (mysql, postgresql, remote)

// MakeID generates new random ID for Node, Service, or Agent.
func MakeID() string {
	return "gen:" + uuid.New().String()
}

var makeID = MakeID
