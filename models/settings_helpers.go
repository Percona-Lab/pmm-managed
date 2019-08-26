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
	"encoding/json"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
)

// GetSettings returns current PMM Server settings.
func GetSettings(q reform.DBTX) (*Settings, error) {
	var b []byte
	if err := q.QueryRow("SELECT settings FROM settings").Scan(&b); err != nil {
		return nil, errors.Wrap(err, "failed to select settings")
	}

	var s Settings
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal settings")
	}

	s.fillDefaults()
	return &s, nil
}

// SaveSettings saves PMM Server settings.
// It may modify passed settings to fill defaults.
func SaveSettings(q reform.DBTX, s *Settings) error {
	s.fillDefaults()

	for _, pair := range []struct {
		dur  time.Duration
		name string
	}{
		{dur: s.MetricsResolutions.HR, name: "hr"},
		{dur: s.MetricsResolutions.MR, name: "mr"},
		{dur: s.MetricsResolutions.LR, name: "lr"},
	} {
		if pair.dur < time.Second {
			return status.Error(codes.InvalidArgument, pair.name+": minimal resolution is 1s")
		}
		if pair.dur.Truncate(time.Second) != pair.dur {
			return status.Error(codes.InvalidArgument, pair.name+": should a natural number of seconds")
		}
	}

	if s.QAN.DataRetention < 24*time.Hour {
		return status.Error(codes.InvalidArgument, "data_retention: minimal resolution is 24h")
	}
	if s.QAN.DataRetention.Truncate(24*time.Hour) != s.QAN.DataRetention {
		return status.Error(codes.InvalidArgument, "data_retention: should a natural number of days")
	}

	b, err := json.Marshal(s)
	if err != nil {
		return errors.Wrap(err, "failed to marshal settings")
	}

	_, err = q.Exec("UPDATE settings SET settings = $1", b)
	if err != nil {
		return errors.Wrap(err, "failed to update settings")
	}

	return nil
}
