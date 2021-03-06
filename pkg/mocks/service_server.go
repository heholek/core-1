// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import v1 "github.com/open-integration/core/pkg/api/v1"

// ServiceServer is an autogenerated mock type for the ServiceServer type
type ServiceServer struct {
	mock.Mock
}

// Call provides a mock function with given fields: _a0, _a1
func (_m *ServiceServer) Call(_a0 context.Context, _a1 *v1.CallRequest) (*v1.CallResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.CallResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.CallRequest) *v1.CallResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.CallResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.CallRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields: _a0, _a1
func (_m *ServiceServer) Init(_a0 context.Context, _a1 *v1.InitRequest) (*v1.InitResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1.InitResponse
	if rf, ok := ret.Get(0).(func(context.Context, *v1.InitRequest) *v1.InitResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.InitResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *v1.InitRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
