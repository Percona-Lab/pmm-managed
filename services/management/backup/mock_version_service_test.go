// Code generated by mockery v1.0.0. DO NOT EDIT.

package backup

import mock "github.com/stretchr/testify/mock"

// mockVersionService is an autogenerated mock type for the versionService type
type mockVersionService struct {
	mock.Mock
}

// GetLocalMySQLVersion provides a mock function with given fields: agentID
func (_m *mockVersionService) GetLocalMySQLVersion(agentID string) (string, error) {
	ret := _m.Called(agentID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(agentID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(agentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetXtrabackupVersion provides a mock function with given fields: agentID
func (_m *mockVersionService) GetXtrabackupVersion(agentID string) (string, error) {
	ret := _m.Called(agentID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(agentID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(agentID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
