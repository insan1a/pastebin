// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	entity "github.com/romankravchuk/pastebin/internal/entity"
	mock "github.com/stretchr/testify/mock"

	oauth2 "golang.org/x/oauth2"
)

// OAuthWebAPI is an autogenerated mock type for the OAuthWebAPI type
type OAuthWebAPI struct {
	mock.Mock
}

// GetToken provides a mock function with given fields: code
func (_m *OAuthWebAPI) GetToken(code string) (*oauth2.Token, error) {
	ret := _m.Called(code)

	var r0 *oauth2.Token
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*oauth2.Token, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *oauth2.Token); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*oauth2.Token)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserInfo provides a mock function with given fields: token
func (_m *OAuthWebAPI) GetUserInfo(token *oauth2.Token) (*entity.User, error) {
	ret := _m.Called(token)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*oauth2.Token) (*entity.User, error)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(*oauth2.Token) *entity.User); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*oauth2.Token) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOAuthWebAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewOAuthWebAPI creates a new instance of OAuthWebAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOAuthWebAPI(t mockConstructorTestingTNewOAuthWebAPI) *OAuthWebAPI {
	mock := &OAuthWebAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
