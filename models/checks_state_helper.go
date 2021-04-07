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

package models

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

// FindCheckStates returns all CheckStates stored in the table.
func FindCheckStates(q *reform.Querier) (map[string]Interval, error) {
	rows, err := q.SelectAllFrom(ChecksStateTable, "")
	switch err {
	case nil:
		cs := make(map[string]Interval)
		for _, r := range rows {
			state := r.(*ChecksState)
			cs[state.Name] = state.Interval
		}
		return cs, nil
	case reform.ErrNoRows:
		return nil, err
	default:
		return nil, errors.WithStack(err)
	}
}

// FindCheckStateByName finds ChecksState by check name.
func FindCheckStateByName(q *reform.Querier, name string) (*ChecksState, error) {
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "Empty Check name.")
	}

	cs := &ChecksState{Name: name}
	switch err := q.Reload(cs); err {
	case nil:
		return cs, nil
	case reform.ErrNoRows:
		return nil, err
	default:
		return nil, errors.WithStack(err)
	}
}

// CreateCheckState persists ChecksState.
func CreateCheckState(q *reform.Querier, name string, interval Interval) (*ChecksState, error) {
	row := &ChecksState{
		Name:     name,
		Interval: interval,
	}

	if err := q.Insert(row); err != nil {
		return nil, errors.Wrap(err, "failed to create checks state")
	}

	return row, nil
}

// ChangeCheckState updates the interval of a check state if already present.
func ChangeCheckState(q *reform.Querier, name string, interval Interval) (*ChecksState, error) {
	row, err := FindCheckStateByName(q, name)
	if err != nil {
		return nil, err
	}

	row.Interval = interval

	if err := q.Update(row); err != nil {
		return nil, errors.Wrap(err, "failed to update checks state")
	}

	return row, nil
}
