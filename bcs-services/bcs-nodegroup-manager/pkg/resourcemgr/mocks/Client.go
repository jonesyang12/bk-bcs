// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	resourcemgr "github.com/Tencent/bk-bcs/bcs-services/bcs-nodegroup-manager/pkg/resourcemgr"
	storage "github.com/Tencent/bk-bcs/bcs-services/bcs-nodegroup-manager/pkg/storage"
	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// GetResourcePool provides a mock function with given fields: poolID, option
func (_m *Client) GetResourcePool(poolID string, option *resourcemgr.GetOptions) (*storage.ResourcePool, error) {
	ret := _m.Called(poolID, option)

	var r0 *storage.ResourcePool
	if rf, ok := ret.Get(0).(func(string, *resourcemgr.GetOptions) *storage.ResourcePool); ok {
		r0 = rf(poolID, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ResourcePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *resourcemgr.GetOptions) error); ok {
		r1 = rf(poolID, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcePoolByCondition provides a mock function with given fields: poolID, consumerID, deviceRecord, option
func (_m *Client) GetResourcePoolByCondition(poolID string, consumerID string, deviceRecord string, option *resourcemgr.GetOptions) (*storage.ResourcePool, error) {
	ret := _m.Called(poolID, consumerID, deviceRecord, option)

	var r0 *storage.ResourcePool
	if rf, ok := ret.Get(0).(func(string, string, string, *resourcemgr.GetOptions) *storage.ResourcePool); ok {
		r0 = rf(poolID, consumerID, deviceRecord, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storage.ResourcePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, *resourcemgr.GetOptions) error); ok {
		r1 = rf(poolID, consumerID, deviceRecord, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTask provides a mock function with given fields: poolID, consumerID, option
func (_m *Client) ListTasks(poolID string, consumerID string, option *resourcemgr.ListOptions) ([]*storage.ScaleDownTask, error) {
	ret := _m.Called(poolID, consumerID, option)

	var r0 []*storage.ScaleDownTask
	if rf, ok := ret.Get(0).(func(string, string, *resourcemgr.ListOptions) []*storage.ScaleDownTask); ok {
		r0 = rf(poolID, consumerID, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.ScaleDownTask)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *resourcemgr.ListOptions) error); ok {
		r1 = rf(poolID, consumerID, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Client) GetTaskByID(recordID string, opt *resourcemgr.GetOptions) (*storage.ScaleDownTask, error){
	panic("not impleted")
}

// ListResourcePools provides a mock function with given fields: option
func (_m *Client) ListResourcePools(option *resourcemgr.ListOptions) ([]*storage.ResourcePool, error) {
	ret := _m.Called(option)

	var r0 []*storage.ResourcePool
	if rf, ok := ret.Get(0).(func(*resourcemgr.ListOptions) []*storage.ResourcePool); ok {
		r0 = rf(option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*storage.ResourcePool)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*resourcemgr.ListOptions) error); ok {
		r1 = rf(option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewClient creates a new instance of Client. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
