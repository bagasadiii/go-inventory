// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	context "context"
	model "inventory-app/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// AccRepoMethod is an autogenerated mock type for the AccRepoMethod type
type AccRepoMethod struct {
	mock.Mock
}

// CreateAcc provides a mock function with given fields: ctx, acc
func (_m *AccRepoMethod) CreateAcc(ctx context.Context, acc *model.Account) error {
	ret := _m.Called(ctx, acc)

	if len(ret) == 0 {
		panic("no return value specified for CreateAcc")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) error); ok {
		r0 = rf(ctx, acc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAcc provides a mock function with given fields: ctx, id
func (_m *AccRepoMethod) DeleteAcc(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAcc")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByUsername provides a mock function with given fields: ctx, username
func (_m *AccRepoMethod) FindByUsername(ctx context.Context, username string) (*model.Account, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for FindByUsername")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Account, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Account); ok {
		r0 = rf(ctx, username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAcc provides a mock function with given fields: ctx, acc
func (_m *AccRepoMethod) UpdateAcc(ctx context.Context, acc *model.Account) (*model.Account, error) {
	ret := _m.Called(ctx, acc)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAcc")
	}

	var r0 *model.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) (*model.Account, error)); ok {
		return rf(ctx, acc)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Account) *model.Account); ok {
		r0 = rf(ctx, acc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Account) error); ok {
		r1 = rf(ctx, acc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccRepoMethod creates a new instance of AccRepoMethod. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccRepoMethod(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccRepoMethod {
	mock := &AccRepoMethod{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
