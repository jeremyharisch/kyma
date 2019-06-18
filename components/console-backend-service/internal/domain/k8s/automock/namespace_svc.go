// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import gqlschema "github.com/kyma-project/kyma/components/console-backend-service/internal/gqlschema"

import mock "github.com/stretchr/testify/mock"
import v1 "k8s.io/api/core/v1"

// namespaceSvc is an autogenerated mock type for the namespaceSvc type
type namespaceSvc struct {
	mock.Mock
}

// Create provides a mock function with given fields: name, labels
func (_m *namespaceSvc) Create(name string, labels gqlschema.Labels) (*v1.Namespace, error) {
	ret := _m.Called(name, labels)

	var r0 *v1.Namespace
	if rf, ok := ret.Get(0).(func(string, gqlschema.Labels) *v1.Namespace); ok {
		r0 = rf(name, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, gqlschema.Labels) error); ok {
		r1 = rf(name, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name
func (_m *namespaceSvc) Delete(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Find provides a mock function with given fields: name
func (_m *namespaceSvc) Find(name string) (*v1.Namespace, error) {
	ret := _m.Called(name)

	var r0 *v1.Namespace
	if rf, ok := ret.Get(0).(func(string) *v1.Namespace); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields:
func (_m *namespaceSvc) List() ([]*v1.Namespace, error) {
	ret := _m.Called()

	var r0 []*v1.Namespace
	if rf, ok := ret.Get(0).(func() []*v1.Namespace); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Namespace)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
