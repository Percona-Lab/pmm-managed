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

// Package qan contains business logic of working with QAN.
package qan

import (
	"context"
	"time"

	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/api/qanpb"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// Client represents qan-api client for data collection.
type Client struct {
	c  qanCollectorClient
	db *reform.DB
	l  *logrus.Entry
}

// NewClient returns new client for given gRPC connection.
func NewClient(cc *grpc.ClientConn, db *reform.DB) *Client {
	return &Client{
		c:  qanpb.NewCollectorClient(cc),
		db: db,
		l:  logrus.WithField("component", "qan"),
	}
}

func mergeLabels(node *models.Node, service *models.Service, agent *models.Agent) (map[string]string, error) {
	res := make(map[string]string, 16)

	labels, err := node.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[name] = value
	}

	labels, err = service.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[name] = value
	}

	labels, err = agent.UnifiedLabels()
	if err != nil {
		return nil, err
	}
	for name, value := range labels {
		res[name] = value
	}

	return res, nil
}

// Collect adds labels to the data from pmm-agent and sends it to qan-api.
func (c *Client) Collect(ctx context.Context, req *agentpb.CollectRequest) error {
	// TODO That code is simple, but performance will be very bad for any non-trivial load.
	// https://jira.percona.com/browse/PMM-3894

	start := time.Now()
	defer func() {
		if dur := time.Since(start); dur > time.Second {
			c.l.Warnf("Collect for %d buckets took %s.", len(req.MetricsBucket), dur)
		}
	}()
	metricBuckets := make([]*qanpb.MetricsBucket, len(req.MetricsBucket))

	for i, m := range req.MetricsBucket {
		if m.Common.AgentId == "" {
			c.l.Errorf("Empty agent_id for bucket with query_id %q, can't add labels.", m.Common.Queryid)
			continue
		}

		// get agent
		agent, err := models.FindAgentByID(c.db.Querier, m.Common.AgentId)
		if err != nil {
			c.l.Error(err)
			continue
		}

		// get service
		services, err := models.ServicesForAgent(c.db.Querier, m.Common.AgentId)
		if err != nil {
			c.l.Error(err)
			continue
		}
		if len(services) != 1 {
			c.l.Errorf("Expected 1 Service, got %d.", len(services))
			continue
		}
		service := services[0]

		// get node for that service (not for that agent)
		node, err := models.FindNodeByID(c.db.Querier, service.NodeID)
		if err != nil {
			c.l.Error(err)
			continue
		}

		labels, err := mergeLabels(node, service, agent)
		if err != nil {
			c.l.Error(err)
			continue
		}
		mb := &qanpb.MetricsBucket{
			Queryid:              m.Common.Queryid,
			Fingerprint:          m.Common.Fingerprint,
			Database:             m.Common.Database,
			Schema:               m.Common.Schema,
			Tables:               m.Common.Tables,
			Username:             m.Common.Username,
			ClientHost:           m.Common.ClientHost,
			AgentId:              m.Common.AgentId,
			AgentType:            m.Common.AgentType,
			PeriodStartUnixSecs:  m.Common.PeriodStartUnixSecs,
			PeriodLengthSecs:     m.Common.PeriodLengthSecs,
			Example:              m.Common.Example,
			ExampleFormat:        m.Common.ExampleFormat,
			IsTruncated:          m.Common.IsTruncated,
			ExampleType:          m.Common.ExampleType,
			ExampleMetrics:       m.Common.ExampleMetrics,
			NumQueriesWithErrors: m.Common.NumQueriesWithErrors,
			Errors:               m.Common.Errors,
			NumQueries:           m.Common.NumQueries,
			MQueryTimeCnt:        m.Common.MQueryTimeCnt,
			MQueryTimeSum:        m.Common.MQueryTimeSum,
			MQueryTimeMin:        m.Common.MQueryTimeMin,
			MQueryTimeMax:        m.Common.MQueryTimeMax,
			MQueryTimeP99:        m.Common.MQueryTimeP99,
		}

		switch {
		case m.Mysql != nil:
			fillMySQL(mb, m.Mysql)
		case m.Mongodb != nil:
			fillMongoDB(mb, m.Mongodb)
		case m.Postgresql != nil:
			fillPostgreSQL(mb, m.Postgresql)
		}

		// in order of fields in MetricsBucket
		for labelName, field := range map[string]*string{
			"service_name":    &mb.ServiceName,
			"replication_set": &mb.ReplicationSet,
			"cluster":         &mb.Cluster,
			"service_type":    &mb.ServiceType,
			"environment":     &mb.Environment,
			"az":              &mb.Az,
			"region":          &mb.Region,
			"node_model":      &mb.NodeModel,
			"container_name":  &mb.ContainerName,
			"agent_id":        &mb.AgentId,
		} {
			*field = labels[labelName]
			delete(labels, labelName)
		}

		mb.Labels = labels

		metricBuckets[i] = mb
	}

	qanReq := &qanpb.CollectRequest{
		MetricsBucket: metricBuckets,
	}

	c.l.Debugf("%+v", qanReq)
	res, err := c.c.Collect(ctx, qanReq)
	if err != nil {
		return errors.Wrap(err, "failed to sent CollectRequest to QAN")
	}
	c.l.Debugf("%+v", res)
	return nil
}

func fillMySQL(mb *qanpb.MetricsBucket, bm *agentpb.MetricsBucket_MySQL) {
	mb.MLockTimeCnt = bm.MLockTimeCnt
	mb.MLockTimeSum = bm.MLockTimeSum
	mb.MLockTimeMin = bm.MLockTimeMin
	mb.MLockTimeMax = bm.MLockTimeMax
	mb.MLockTimeP99 = bm.MLockTimeP99

	mb.MRowsSentCnt = bm.MRowsSentCnt
	mb.MRowsSentSum = bm.MRowsSentSum
	mb.MRowsSentMin = bm.MRowsSentMin
	mb.MRowsSentMax = bm.MRowsSentMax
	mb.MRowsSentP99 = bm.MRowsSentP99

	mb.MRowsExaminedCnt = bm.MRowsExaminedCnt
	mb.MRowsExaminedSum = bm.MRowsExaminedSum
	mb.MRowsExaminedMin = bm.MRowsExaminedMin
	mb.MRowsExaminedMax = bm.MRowsExaminedMax
	mb.MRowsExaminedP99 = bm.MRowsExaminedP99

	mb.MRowsAffectedCnt = bm.MRowsAffectedCnt
	mb.MRowsAffectedSum = bm.MRowsAffectedSum
	mb.MRowsAffectedMin = bm.MRowsAffectedMin
	mb.MRowsAffectedMax = bm.MRowsAffectedMax
	mb.MRowsAffectedP99 = bm.MRowsAffectedP99

	mb.MRowsReadCnt = bm.MRowsReadCnt
	mb.MRowsReadSum = bm.MRowsReadSum
	mb.MRowsReadMin = bm.MRowsReadMin
	mb.MRowsReadMax = bm.MRowsReadMax
	mb.MRowsReadP99 = bm.MRowsReadP99

	mb.MMergePassesCnt = bm.MMergePassesCnt
	mb.MMergePassesSum = bm.MMergePassesSum
	mb.MMergePassesMin = bm.MMergePassesMin
	mb.MMergePassesMax = bm.MMergePassesMax
	mb.MMergePassesP99 = bm.MMergePassesP99

	mb.MInnodbIoROpsCnt = bm.MInnodbIoROpsCnt
	mb.MInnodbIoROpsSum = bm.MInnodbIoROpsSum
	mb.MInnodbIoROpsMin = bm.MInnodbIoROpsMin
	mb.MInnodbIoROpsMax = bm.MInnodbIoROpsMax
	mb.MInnodbIoROpsP99 = bm.MInnodbIoROpsP99

	mb.MInnodbIoRBytesCnt = bm.MInnodbIoRBytesCnt
	mb.MInnodbIoRBytesSum = bm.MInnodbIoRBytesSum
	mb.MInnodbIoRBytesMin = bm.MInnodbIoRBytesMin
	mb.MInnodbIoRBytesMax = bm.MInnodbIoRBytesMax
	mb.MInnodbIoRBytesP99 = bm.MInnodbIoRBytesP99

	mb.MInnodbIoRWaitCnt = bm.MInnodbIoRWaitCnt
	mb.MInnodbIoRWaitSum = bm.MInnodbIoRWaitSum
	mb.MInnodbIoRWaitMin = bm.MInnodbIoRWaitMin
	mb.MInnodbIoRWaitMax = bm.MInnodbIoRWaitMax
	mb.MInnodbIoRWaitP99 = bm.MInnodbIoRWaitP99

	mb.MInnodbRecLockWaitCnt = bm.MInnodbRecLockWaitCnt
	mb.MInnodbRecLockWaitSum = bm.MInnodbRecLockWaitSum
	mb.MInnodbRecLockWaitMin = bm.MInnodbRecLockWaitMin
	mb.MInnodbRecLockWaitMax = bm.MInnodbRecLockWaitMax
	mb.MInnodbRecLockWaitP99 = bm.MInnodbRecLockWaitP99

	mb.MInnodbQueueWaitCnt = bm.MInnodbQueueWaitCnt
	mb.MInnodbQueueWaitSum = bm.MInnodbQueueWaitSum
	mb.MInnodbQueueWaitMin = bm.MInnodbQueueWaitMin
	mb.MInnodbQueueWaitMax = bm.MInnodbQueueWaitMax
	mb.MInnodbQueueWaitP99 = bm.MInnodbQueueWaitP99

	mb.MInnodbPagesDistinctCnt = bm.MInnodbPagesDistinctCnt
	mb.MInnodbPagesDistinctSum = bm.MInnodbPagesDistinctSum
	mb.MInnodbPagesDistinctMin = bm.MInnodbPagesDistinctMin
	mb.MInnodbPagesDistinctMax = bm.MInnodbPagesDistinctMax
	mb.MInnodbPagesDistinctP99 = bm.MInnodbPagesDistinctP99

	mb.MQueryLengthCnt = bm.MQueryLengthCnt
	mb.MQueryLengthSum = bm.MQueryLengthSum
	mb.MQueryLengthMin = bm.MQueryLengthMin
	mb.MQueryLengthMax = bm.MQueryLengthMax
	mb.MQueryLengthP99 = bm.MQueryLengthP99

	mb.MBytesSentCnt = bm.MBytesSentCnt
	mb.MBytesSentSum = bm.MBytesSentSum
	mb.MBytesSentMin = bm.MBytesSentMin
	mb.MBytesSentMax = bm.MBytesSentMax
	mb.MBytesSentP99 = bm.MBytesSentP99

	mb.MTmpTablesCnt = bm.MTmpTablesCnt
	mb.MTmpTablesSum = bm.MTmpTablesSum
	mb.MTmpTablesMin = bm.MTmpTablesMin
	mb.MTmpTablesMax = bm.MTmpTablesMax
	mb.MTmpTablesP99 = bm.MTmpTablesP99

	mb.MTmpDiskTablesCnt = bm.MTmpDiskTablesCnt
	mb.MTmpDiskTablesSum = bm.MTmpDiskTablesSum
	mb.MTmpDiskTablesMin = bm.MTmpDiskTablesMin
	mb.MTmpDiskTablesMax = bm.MTmpDiskTablesMax
	mb.MTmpDiskTablesP99 = bm.MTmpDiskTablesP99

	mb.MTmpTableSizesCnt = bm.MTmpTableSizesCnt
	mb.MTmpTableSizesSum = bm.MTmpTableSizesSum
	mb.MTmpTableSizesMin = bm.MTmpTableSizesMin
	mb.MTmpTableSizesMax = bm.MTmpTableSizesMax
	mb.MTmpTableSizesP99 = bm.MTmpTableSizesP99

	mb.MQcHitCnt = bm.MQcHitCnt
	mb.MQcHitSum = bm.MQcHitSum

	mb.MFullScanCnt = bm.MFullScanCnt
	mb.MFullScanSum = bm.MFullScanSum

	mb.MFullJoinCnt = bm.MFullJoinCnt
	mb.MFullJoinSum = bm.MFullJoinSum

	mb.MTmpTableCnt = bm.MTmpTableCnt
	mb.MTmpTableSum = bm.MTmpTableSum

	mb.MTmpTableOnDiskCnt = bm.MTmpTableOnDiskCnt
	mb.MTmpTableOnDiskSum = bm.MTmpTableOnDiskSum

	mb.MFilesortCnt = bm.MFilesortCnt
	mb.MFilesortSum = bm.MFilesortSum

	mb.MFilesortOnDiskCnt = bm.MFilesortOnDiskCnt
	mb.MFilesortOnDiskSum = bm.MFilesortOnDiskSum

	mb.MSelectFullRangeJoinCnt = bm.MSelectFullRangeJoinCnt
	mb.MSelectFullRangeJoinSum = bm.MSelectFullRangeJoinSum

	mb.MSelectRangeCnt = bm.MSelectRangeCnt
	mb.MSelectRangeSum = bm.MSelectRangeSum

	mb.MSelectRangeCheckCnt = bm.MSelectRangeCheckCnt
	mb.MSelectRangeCheckSum = bm.MSelectRangeCheckSum

	mb.MSortRangeCnt = bm.MSortRangeCnt
	mb.MSortRangeSum = bm.MSortRangeSum

	mb.MSortRowsCnt = bm.MSortRowsCnt
	mb.MSortRowsSum = bm.MSortRowsSum

	mb.MSortScanCnt = bm.MSortScanCnt
	mb.MSortScanSum = bm.MSortScanSum

	mb.MNoIndexUsedCnt = bm.MNoIndexUsedCnt
	mb.MNoIndexUsedSum = bm.MNoIndexUsedSum

	mb.MNoGoodIndexUsedCnt = bm.MNoGoodIndexUsedCnt
	mb.MNoGoodIndexUsedSum = bm.MNoGoodIndexUsedSum
}

func fillMongoDB(mb *qanpb.MetricsBucket, bm *agentpb.MetricsBucket_MongoDB) {
	mb.MDocsReturnedCnt = bm.MDocsReturnedCnt
	mb.MDocsReturnedSum = bm.MDocsReturnedSum
	mb.MDocsReturnedMin = bm.MDocsReturnedMin
	mb.MDocsReturnedMax = bm.MDocsReturnedMax
	mb.MDocsReturnedP99 = bm.MDocsReturnedP99

	mb.MResponseLengthCnt = bm.MResponseLengthCnt
	mb.MResponseLengthSum = bm.MResponseLengthSum
	mb.MResponseLengthMin = bm.MResponseLengthMin
	mb.MResponseLengthMax = bm.MResponseLengthMax
	mb.MResponseLengthP99 = bm.MResponseLengthP99

	mb.MDocsScannedCnt = bm.MDocsScannedCnt
	mb.MDocsScannedSum = bm.MDocsScannedSum
	mb.MDocsScannedMin = bm.MDocsScannedMin
	mb.MDocsScannedMax = bm.MDocsScannedMax
	mb.MDocsScannedP99 = bm.MDocsScannedP99
}

func fillPostgreSQL(mb *qanpb.MetricsBucket, bp *agentpb.MetricsBucket_PostgreSQL) {
	mb.MRowsSentCnt = bp.MRowsCnt
	mb.MRowsSentSum = bp.MRowsSum

	mb.MSharedBlksHitCnt = bp.MSharedBlksHitCnt
	mb.MSharedBlksHitSum = bp.MSharedBlksHitSum
	mb.MSharedBlksReadCnt = bp.MSharedBlksReadCnt
	mb.MSharedBlksReadSum = bp.MSharedBlksReadSum
	mb.MSharedBlksDirtiedCnt = bp.MSharedBlksDirtiedCnt
	mb.MSharedBlksDirtiedSum = bp.MSharedBlksDirtiedSum
	mb.MSharedBlksWrittenCnt = bp.MSharedBlksWrittenCnt
	mb.MSharedBlksWrittenSum = bp.MSharedBlksWrittenSum

	mb.MLocalBlksHitCnt = bp.MLocalBlksHitCnt
	mb.MLocalBlksHitSum = bp.MLocalBlksHitSum
	mb.MLocalBlksReadCnt = bp.MLocalBlksReadCnt
	mb.MLocalBlksReadSum = bp.MLocalBlksReadSum
	mb.MLocalBlksDirtiedCnt = bp.MLocalBlksDirtiedCnt
	mb.MLocalBlksDirtiedSum = bp.MLocalBlksDirtiedSum
	mb.MLocalBlksWrittenCnt = bp.MLocalBlksWrittenCnt
	mb.MLocalBlksWrittenSum = bp.MLocalBlksWrittenSum

	mb.MTempBlksReadCnt = bp.MTempBlksReadCnt
	mb.MTempBlksReadSum = bp.MTempBlksReadSum
	mb.MTempBlksWrittenCnt = bp.MTempBlksWrittenCnt
	mb.MTempBlksWrittenSum = bp.MTempBlksWrittenSum

	mb.MBlkReadTimeCnt = bp.MBlkReadTimeCnt
	mb.MBlkReadTimeSum = bp.MBlkReadTimeSum
}
