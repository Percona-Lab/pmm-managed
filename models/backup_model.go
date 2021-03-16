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
	"time"

	"gopkg.in/reform.v1"
)

//go:generate reform

// DataModel represents a data model used for performing a backup.
type DataModel string

// DataModel types (in the same order as in backups.proto).
const (
	PhysicalDataModel DataModel = "physical"
	LogicalDataModel  DataModel = "logical"
)

// BackupStatus shows current status of Backup.
type BackupStatus string

const (
	PendingBackupStatus    BackupStatus = "pending"
	InProgressBackupStatus BackupStatus = "in_progress"
	PausedBackupStatus     BackupStatus = "paused"
	SuccessBackupStatus    BackupStatus = "success"
	ErrorBackupStatus      BackupStatus = "error"
)

// Backup represents destination for backup.
//reform:backups
type Backup struct {
	ID         string       `reform:"id,pk"`
	Name       string       `reform:"name"`
	Vendor     string       `reform:"vendor"`
	LocationID string       `reform:"location_id"`
	ServiceID  string       `reform:"service_id"`
	DataModel  DataModel    `reform:"data_model"`
	Status     BackupStatus `reform:"status"`
	CreatedAt  time.Time    `reform:"created_at"`
}

// BeforeInsert implements reform.BeforeInserter interface.
func (s *Backup) BeforeInsert() error {
	s.CreatedAt = Now()
	return nil
}

// BeforeUpdate implements reform.BeforeUpdater interface.
func (s *Backup) BeforeUpdate() error {
	return nil
}

// AfterFind implements reform.AfterFinder interface.
func (s *Backup) AfterFind() error {
	s.CreatedAt = s.CreatedAt.UTC()
	return nil
}

// check interfaces
var (
	_ reform.BeforeInserter = (*Backup)(nil)
	_ reform.BeforeUpdater  = (*Backup)(nil)
	_ reform.AfterFinder    = (*Backup)(nil)
)
