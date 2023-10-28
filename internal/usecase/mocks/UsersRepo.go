// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/romankravchuk/pastebin/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// UsersRepo is an autogenerated mock type for the UsersRepo type
type UsersRepo struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, u
func (_m *UsersRepo) Create(ctx context.Context, u *entity.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: ctx, email
func (_m *UsersRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	ret := _m.Called(ctx, email)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*entity.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *entity.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUsersRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsersRepo creates a new instance of UsersRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsersRepo(t mockConstructorTestingTNewUsersRepo) *UsersRepo {
	mock := &UsersRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
