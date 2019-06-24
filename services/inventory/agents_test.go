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
	"reflect"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"github.com/percona/pmm/api/inventorypb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/utils/logger"
	"github.com/percona/pmm-managed/utils/testdb"
	"github.com/percona/pmm-managed/utils/tests"
)

func TestAgents(t *testing.T) {
	var ctx context.Context

	setup := func(t *testing.T) (ss *ServicesService, as *AgentsService, teardown func(t *testing.T)) {
		ctx = logger.Set(context.Background(), t.Name())
		uuid.SetRand(new(tests.IDReader))

		sqlDB := testdb.Open(t, models.SetupFixtures)
		db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(t.Logf))

		r := new(mockAgentsRegistry)
		r.Test(t)

		teardown = func(t *testing.T) {
			r.AssertExpectations(t)
			require.NoError(t, sqlDB.Close())
		}
		ss = NewServicesService(db, r)
		as = NewAgentsService(db, r)

		return
	}

	t.Run("Basic", func(t *testing.T) {
		// FIXME split this test into several smaller

		ss, as, teardown := setup(t)
		defer teardown(t)

		as.r.(*mockAgentsRegistry).On("IsConnected", models.PMMServerAgentID).Return(true)
		actualAgents, err := as.List(ctx, AgentFilters{})
		require.NoError(t, err)
		require.Len(t, actualAgents, 4) // PMM Server's pmm-agent, node_exporter, postgres_exporter, PostgreSQL QAN

		as.r.(*mockAgentsRegistry).On("IsConnected", "/agent_id/00000000-0000-4000-8000-000000000005").Return(true)
		as.r.(*mockAgentsRegistry).On("SendSetStateRequest", ctx, "/agent_id/00000000-0000-4000-8000-000000000005")
		as.r.(*mockAgentsRegistry).On("CheckConnectionToService", ctx,
			mock.AnythingOfType(reflect.TypeOf(&models.Service{}).Name()),
			mock.AnythingOfType(reflect.TypeOf(&models.Agent{}).Name())).Return(nil)
		pmmAgent, err := as.AddPMMAgent(ctx, &inventorypb.AddPMMAgentRequest{
			RunsOnNodeId: models.PMMServerNodeID,
		})
		require.NoError(t, err)
		expectedPMMAgent := &inventorypb.PMMAgent{
			AgentId:      "/agent_id/00000000-0000-4000-8000-000000000005",
			RunsOnNodeId: models.PMMServerNodeID,
			Connected:    true,
		}
		assert.Equal(t, expectedPMMAgent, pmmAgent)

		actualNodeExporter, err := as.AddNodeExporter(ctx, &inventorypb.AddNodeExporterRequest{
			PmmAgentId: pmmAgent.AgentId,
		})
		require.NoError(t, err)
		expectedNodeExporter := &inventorypb.NodeExporter{
			AgentId:    "/agent_id/00000000-0000-4000-8000-000000000006",
			PmmAgentId: "/agent_id/00000000-0000-4000-8000-000000000005",
		}
		assert.Equal(t, expectedNodeExporter, actualNodeExporter)

		actualNodeExporter, err = as.ChangeNodeExporter(ctx, &inventorypb.ChangeNodeExporterRequest{
			AgentId: "/agent_id/00000000-0000-4000-8000-000000000006",
			Common: &inventorypb.ChangeCommonAgentParams{
				ChangeDisabled: &inventorypb.ChangeCommonAgentParams_Disabled{
					Disabled: true,
				},
			},
		})
		require.NoError(t, err)
		expectedNodeExporter = &inventorypb.NodeExporter{
			AgentId:    "/agent_id/00000000-0000-4000-8000-000000000006",
			PmmAgentId: "/agent_id/00000000-0000-4000-8000-000000000005",
			Disabled:   true,
		}
		assert.Equal(t, expectedNodeExporter, actualNodeExporter)

		actualAgent, err := as.Get(ctx, "/agent_id/00000000-0000-4000-8000-000000000006")
		require.NoError(t, err)
		assert.Equal(t, expectedNodeExporter, actualAgent)

		s, err := ss.AddMySQL(ctx, &models.AddDBMSServiceParams{
			ServiceName: "test-mysql",
			NodeID:      models.PMMServerNodeID,
			Address:     pointer.ToString("127.0.0.1"),
			Port:        pointer.ToUint16(3306),
		})
		require.NoError(t, err)

		actualAgent, err = as.AddMySQLdExporter(ctx, &inventorypb.AddMySQLdExporterRequest{
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  s.ServiceId,
			Username:   "username",
		})
		require.NoError(t, err)
		expectedMySQLdExporter := &inventorypb.MySQLdExporter{
			AgentId:    "/agent_id/00000000-0000-4000-8000-000000000008",
			PmmAgentId: "/agent_id/00000000-0000-4000-8000-000000000005",
			ServiceId:  s.ServiceId,
			Username:   "username",
		}
		assert.Equal(t, expectedMySQLdExporter, actualAgent)

		actualAgent, err = as.Get(ctx, "/agent_id/00000000-0000-4000-8000-000000000008")
		require.NoError(t, err)
		assert.Equal(t, expectedMySQLdExporter, actualAgent)

		ms, err := ss.AddMongoDB(ctx, &models.AddDBMSServiceParams{
			ServiceName: "test-mongo",
			NodeID:      models.PMMServerNodeID,
			Address:     pointer.ToString("127.0.0.1"),
			Port:        pointer.ToUint16(27017),
		})
		require.NoError(t, err)

		actualAgent, err = as.AddMongoDBExporter(ctx, &inventorypb.AddMongoDBExporterRequest{
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  ms.ServiceId,
			Username:   "username",
		})
		require.NoError(t, err)
		expectedMongoDBExporter := &inventorypb.MongoDBExporter{
			AgentId:    "/agent_id/00000000-0000-4000-8000-00000000000a",
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  ms.ServiceId,
			Username:   "username",
		}
		assert.Equal(t, expectedMongoDBExporter, actualAgent)

		actualAgent, err = as.Get(ctx, "/agent_id/00000000-0000-4000-8000-00000000000a")
		require.NoError(t, err)
		assert.Equal(t, expectedMongoDBExporter, actualAgent)

		actualAgent, err = as.AddQANMySQLSlowlogAgent(ctx, &inventorypb.AddQANMySQLSlowlogAgentRequest{
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  s.ServiceId,
			Username:   "username",
		})
		require.NoError(t, err)
		expectedQANMySQLSlowlogAgent := &inventorypb.QANMySQLSlowlogAgent{
			AgentId:    "/agent_id/00000000-0000-4000-8000-00000000000b",
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  s.ServiceId,
			Username:   "username",
		}
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgent)

		actualAgent, err = as.Get(ctx, "/agent_id/00000000-0000-4000-8000-00000000000b")
		require.NoError(t, err)
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgent)

		ps, err := ss.AddPostgreSQL(ctx, &models.AddDBMSServiceParams{
			ServiceName: "test-postgres",
			NodeID:      models.PMMServerNodeID,
			Address:     pointer.ToString("127.0.0.1"),
			Port:        pointer.ToUint16(5432),
		})
		require.NoError(t, err)

		actualAgent, err = as.AddPostgresExporter(ctx, &inventorypb.AddPostgresExporterRequest{
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  ps.ServiceId,
			Username:   "username",
		})
		require.NoError(t, err)
		expectedPostgresExporter := &inventorypb.PostgresExporter{
			AgentId:    "/agent_id/00000000-0000-4000-8000-00000000000d",
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  ps.ServiceId,
			Username:   "username",
		}
		assert.Equal(t, expectedPostgresExporter, actualAgent)

		actualAgent, err = as.Get(ctx, "/agent_id/00000000-0000-4000-8000-00000000000d")
		require.NoError(t, err)
		assert.Equal(t, expectedPostgresExporter, actualAgent)

		actualAgents, err = as.List(ctx, AgentFilters{})
		require.NoError(t, err)
		for i, a := range actualAgents {
			t.Logf("%d: %T %s", i, a, a)
		}
		require.Len(t, actualAgents, 10)
		assert.Equal(t, pmmAgent, actualAgents[3])
		assert.Equal(t, expectedNodeExporter, actualAgents[4])
		assert.Equal(t, expectedMySQLdExporter, actualAgents[5])
		assert.Equal(t, expectedMongoDBExporter, actualAgents[6])
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgents[7])
		assert.Equal(t, expectedPostgresExporter, actualAgents[8])

		// filter by service ID
		actualAgents, err = as.List(ctx, AgentFilters{ServiceID: s.ServiceId})
		require.NoError(t, err)
		require.Len(t, actualAgents, 2)
		assert.Equal(t, expectedMySQLdExporter, actualAgents[0])
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgents[1])

		actualAgents, err = as.List(ctx, AgentFilters{PMMAgentID: pmmAgent.AgentId})
		require.NoError(t, err)
		require.Len(t, actualAgents, 5)
		assert.Equal(t, expectedNodeExporter, actualAgents[0])
		assert.Equal(t, expectedMySQLdExporter, actualAgents[1])
		assert.Equal(t, expectedMongoDBExporter, actualAgents[2])
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgents[3])
		assert.Equal(t, expectedPostgresExporter, actualAgents[4])

		actualAgents, err = as.List(ctx, AgentFilters{PMMAgentID: pmmAgent.AgentId, NodeID: models.PMMServerNodeID})
		require.NoError(t, err)
		require.Len(t, actualAgents, 5)
		assert.Equal(t, expectedNodeExporter, actualAgents[0])
		assert.Equal(t, expectedMySQLdExporter, actualAgents[1])
		assert.Equal(t, expectedMongoDBExporter, actualAgents[2])
		assert.Equal(t, expectedQANMySQLSlowlogAgent, actualAgents[3])
		assert.Equal(t, expectedPostgresExporter, actualAgents[4])

		actualAgents, err = as.List(ctx, AgentFilters{NodeID: models.PMMServerNodeID})
		require.NoError(t, err)
		require.Len(t, actualAgents, 2)
		assert.Equal(t, expectedNodeExporter, actualAgents[1])

		as.r.(*mockAgentsRegistry).On("Kick", ctx, "/agent_id/00000000-0000-4000-8000-000000000005").Return(true)
		err = as.Remove(ctx, "/agent_id/00000000-0000-4000-8000-000000000005", true)
		require.NoError(t, err)
		actualAgent, err = as.Get(ctx, "/agent_id/00000000-0000-4000-8000-000000000005")
		tests.AssertGRPCError(t, status.New(codes.NotFound, `Agent with ID "/agent_id/00000000-0000-4000-8000-000000000005" not found.`), err)
		assert.Nil(t, actualAgent)

		actualAgents, err = as.List(ctx, AgentFilters{})
		require.NoError(t, err)
		require.Len(t, actualAgents, 4) // PMM Server's pmm-agent, node_exporter, postgres_exporter, PostgreSQL QAN
	})

	t.Run("GetEmptyID", func(t *testing.T) {
		_, as, teardown := setup(t)
		defer teardown(t)

		actualNode, err := as.Get(ctx, "")
		tests.AssertGRPCError(t, status.New(codes.InvalidArgument, `Empty Agent ID.`), err)
		assert.Nil(t, actualNode)
	})

	t.Run("AddPMMAgent", func(t *testing.T) {
		_, as, teardown := setup(t)
		defer teardown(t)

		as.r.(*mockAgentsRegistry).On("IsConnected", "/agent_id/00000000-0000-4000-8000-000000000005").Return(false)
		actualAgent, err := as.AddPMMAgent(ctx, &inventorypb.AddPMMAgentRequest{
			RunsOnNodeId: models.PMMServerNodeID,
		})
		require.NoError(t, err)
		expectedPMMAgent := &inventorypb.PMMAgent{
			AgentId:      "/agent_id/00000000-0000-4000-8000-000000000005",
			RunsOnNodeId: models.PMMServerNodeID,
			Connected:    false,
		}
		assert.Equal(t, expectedPMMAgent, actualAgent)

		as.r.(*mockAgentsRegistry).On("IsConnected", "/agent_id/00000000-0000-4000-8000-000000000006").Return(true)
		actualAgent, err = as.AddPMMAgent(ctx, &inventorypb.AddPMMAgentRequest{
			RunsOnNodeId: models.PMMServerNodeID,
		})
		require.NoError(t, err)
		expectedPMMAgent = &inventorypb.PMMAgent{
			AgentId:      "/agent_id/00000000-0000-4000-8000-000000000006",
			RunsOnNodeId: models.PMMServerNodeID,
			Connected:    true,
		}
		assert.Equal(t, expectedPMMAgent, actualAgent)
	})

	t.Run("AddPmmAgentNotFound", func(t *testing.T) {
		_, as, teardown := setup(t)
		defer teardown(t)

		_, err := as.AddNodeExporter(ctx, &inventorypb.AddNodeExporterRequest{
			PmmAgentId: "no-such-id",
		})
		tests.AssertGRPCError(t, status.New(codes.NotFound, `Agent with ID "no-such-id" not found.`), err)
	})

	t.Run("AddServiceNotFound", func(t *testing.T) {
		_, as, teardown := setup(t)
		defer teardown(t)

		as.r.(*mockAgentsRegistry).On("IsConnected", "/agent_id/00000000-0000-4000-8000-000000000005").Return(true)
		pmmAgent, err := as.AddPMMAgent(ctx, &inventorypb.AddPMMAgentRequest{
			RunsOnNodeId: models.PMMServerNodeID,
		})
		require.NoError(t, err)

		_, err = as.AddMySQLdExporter(ctx, &inventorypb.AddMySQLdExporterRequest{
			PmmAgentId: pmmAgent.AgentId,
			ServiceId:  "no-such-id",
		})
		tests.AssertGRPCError(t, status.New(codes.NotFound, `Service with ID "no-such-id" not found.`), err)
	})

	t.Run("RemoveNotFound", func(t *testing.T) {
		_, as, teardown := setup(t)
		defer teardown(t)

		err := as.Remove(ctx, "no-such-id", false)
		tests.AssertGRPCError(t, status.New(codes.NotFound, `Agent with ID "no-such-id" not found.`), err)
	})
}
