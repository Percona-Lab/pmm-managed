// Code generated by mockery v1.0.0. DO NOT EDIT.

package management

import agentpb "github.com/percona/pmm/api/agentpb"
import context "context"
import mock "github.com/stretchr/testify/mock"

// mockRegistry is an autogenerated mock type for the registry type
type mockRegistry struct {
	mock.Mock
}

// IsConnected provides a mock function with given fields: pmmAgentID
func (_m *mockRegistry) IsConnected(pmmAgentID string) bool {
	ret := _m.Called(pmmAgentID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(pmmAgentID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Kick provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockRegistry) Kick(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}

// SendRequest provides a mock function with given fields: ctx, pmmAgentID, payload
func (_m *mockRegistry) SendRequest(ctx context.Context, pmmAgentID string, payload agentpb.ServerRequestPayload) (agentpb.AgentResponsePayload, error) {
	ret := _m.Called(ctx, pmmAgentID, payload)

	var r0 agentpb.AgentResponsePayload
	if rf, ok := ret.Get(0).(func(context.Context, string, agentpb.ServerRequestPayload) agentpb.AgentResponsePayload); ok {
		r0 = rf(ctx, pmmAgentID, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(agentpb.AgentResponsePayload)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, agentpb.ServerRequestPayload) error); ok {
		r1 = rf(ctx, pmmAgentID, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendSetStateRequest provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockRegistry) SendSetStateRequest(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}
