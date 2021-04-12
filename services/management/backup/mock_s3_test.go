// Code generated by mockery v1.0.0. DO NOT EDIT.

package backup

import mock "github.com/stretchr/testify/mock"

// MockS3 is an autogenerated mock type for the S3 type
type MockS3 struct {
	mock.Mock
}

// GetBucketLocation provides a mock function with given fields: host, secure, accessKey, secretKey, name
func (_m *MockS3) GetBucketLocation(host string, secure bool, accessKey string, secretKey string, name string) (string, error) {
	ret := _m.Called(host, secure, accessKey, secretKey, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, bool, string, string, string) string); ok {
		r0 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, string, string, string) error); ok {
		r1 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}