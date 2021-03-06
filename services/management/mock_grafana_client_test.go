// Code generated by mockery v1.0.0. DO NOT EDIT.

package management

import (
	context "context"
	time "time"

	mock "github.com/stretchr/testify/mock"
)

// mockGrafanaClient is an autogenerated mock type for the grafanaClient type
type mockGrafanaClient struct {
	mock.Mock
}

// CreateAnnotation provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *mockGrafanaClient) CreateAnnotation(_a0 context.Context, _a1 []string, _a2 time.Time, _a3 string, _a4 string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, []string, time.Time, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, time.Time, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
