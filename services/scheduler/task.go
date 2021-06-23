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

package scheduler

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/percona/pmm-managed/models"
)

// Task represents task which will be run inside scheduler.
type Task interface {
	Run(ctx context.Context) error
	Type() models.ScheduledTaskType
	Data() models.ScheduledTaskData
	ID() string
	SetID(string)
}

// common implementation for all tasks.
type common struct {
	id string
}

func (c *common) ID() string {
	return c.id
}

func (c *common) SetID(id string) {
	c.id = id
}

// PrintTask implements Task for logging mesage.
type PrintTask struct {
	*common
	Message string
}

// NewPrintTask creates new task which prints message.
func NewPrintTask(message string) *PrintTask {
	return &PrintTask{
		common:  &common{},
		Message: message,
	}
}

// Run starts task.
func (j *PrintTask) Run(ctx context.Context) error {
	logrus.Info(j.Message)
	return nil
}

// Type returns task type.
func (j *PrintTask) Type() models.ScheduledTaskType {
	return models.ScheduledPrintTask
}

// Data returns data needed for running a task.
func (j *PrintTask) Data() models.ScheduledTaskData {
	return models.ScheduledTaskData{
		Print: &models.PrintTaskData{
			Message: j.Message,
		},
	}
}

// BackupRetryData holds common data for backup retrying.
type BackupRetryData struct {
	Retries  uint
	Interval time.Duration
}

type mySQLBackupTask struct {
	*common
	retry               BackupRetryData
	backupsLogicService backupsLogicService
	ServiceID           string
	LocationID          string
	Name                string
	Description         string
}

func NewMySQLBackupTask(backupsLogicService backupsLogicService, serviceID, locationID, name, description string, retry BackupRetryData) *mySQLBackupTask {
	return &mySQLBackupTask{
		common:              &common{},
		backupsLogicService: backupsLogicService,
		ServiceID:           serviceID,
		LocationID:          locationID,
		Name:                name,
		Description:         description,
		retry:               retry,
	}
}

func (t *mySQLBackupTask) Run(ctx context.Context) error {
	name := t.Name + "_" + time.Now().Format(time.RFC3339)
	_, err := t.backupsLogicService.PerformBackup(ctx, t.ServiceID, t.LocationID, name, t.ID())
	return err
}

func (t *mySQLBackupTask) Type() models.ScheduledTaskType {
	return models.ScheduledMySQLBackupTask
}

func (t *mySQLBackupTask) Data() models.ScheduledTaskData {
	return models.ScheduledTaskData{
		MySQLBackupTask: &models.MySQLBackupTaskData{
			ServiceID:   t.ServiceID,
			LocationID:  t.LocationID,
			Name:        t.Name,
			Description: t.Description,
			Retry: models.BackupRetryData{
				Retries:  t.retry.Retries,
				Interval: t.retry.Interval,
			},
		},
	}
}

type mongoBackupTask struct {
	*common
	retry               BackupRetryData
	backupsLogicService backupsLogicService
	ServiceID           string
	LocationID          string
	Name                string
	Description         string
}

func NewMongoBackupTask(backupsLogicService backupsLogicService, serviceID, locationID, name, description string, retry BackupRetryData) *mongoBackupTask {
	return &mongoBackupTask{
		common:              &common{},
		backupsLogicService: backupsLogicService,
		ServiceID:           serviceID,
		LocationID:          locationID,
		Name:                name,
		Description:         description,
		retry:               retry,
	}
}

func (t *mongoBackupTask) Run(ctx context.Context) error {
	name := t.Name + "_" + time.Now().Format(time.RFC3339)
	_, err := t.backupsLogicService.PerformBackup(ctx, t.ServiceID, t.LocationID, name, t.ID())
	return err
}

func (t *mongoBackupTask) Type() models.ScheduledTaskType {
	return models.ScheduledMongoBackupTask
}

func (t *mongoBackupTask) Data() models.ScheduledTaskData {
	return models.ScheduledTaskData{
		MongoBackupTask: &models.MongoBackupTaskData{
			ServiceID:   t.ServiceID,
			LocationID:  t.LocationID,
			Name:        t.Name,
			Description: t.Description,
			Retry: models.BackupRetryData{
				Retries:  t.retry.Retries,
				Interval: t.retry.Interval,
			},
		},
	}
}
