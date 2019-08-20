// Code generated by mockery v1.0.0. DO NOT EDIT.

package management

import (
	context "context"

	models "github.com/percona/pmm-managed/models"
	mock "github.com/stretchr/testify/mock"
)

// mockAgentsRegistry is an autogenerated mock type for the agentsRegistry type
type mockAgentsRegistry struct {
	mock.Mock
}

// CheckConnectionToService provides a mock function with given fields: ctx, service, agent
func (_m *mockAgentsRegistry) CheckConnectionToService(ctx context.Context, service *models.Service, agent *models.Agent) error {
	ret := _m.Called(ctx, service, agent)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Service, *models.Agent) error); ok {
		r0 = rf(ctx, service, agent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsConnected provides a mock function with given fields: pmmAgentID
func (_m *mockAgentsRegistry) IsConnected(pmmAgentID string) bool {
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
func (_m *mockAgentsRegistry) Kick(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}

// SendSetStateRequest provides a mock function with given fields: ctx, pmmAgentID
func (_m *mockAgentsRegistry) SendSetStateRequest(ctx context.Context, pmmAgentID string) {
	_m.Called(ctx, pmmAgentID)
}
