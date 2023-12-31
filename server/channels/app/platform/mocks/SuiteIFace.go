// Code generated by mockery v2.32.2. DO NOT EDIT.

// Regenerate this file using `make platform-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost/server/public/model"
	mock "github.com/stretchr/testify/mock"
)

// SuiteIFace is an autogenerated mock type for the SuiteIFace type
type SuiteIFace struct {
	mock.Mock
}

// GetSession provides a mock function with given fields: token
func (_m *SuiteIFace) GetSession(token string) (*model.Session, *model.AppError) {
	ret := _m.Called(token)

	var r0 *model.Session
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string) (*model.Session, *model.AppError)); ok {
		return rf(token)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Session); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Session)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *model.AppError); ok {
		r1 = rf(token)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// RolesGrantPermission provides a mock function with given fields: roleNames, permissionId
func (_m *SuiteIFace) RolesGrantPermission(roleNames []string, permissionId string) bool {
	ret := _m.Called(roleNames, permissionId)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]string, string) bool); ok {
		r0 = rf(roleNames, permissionId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// UserCanSeeOtherUser provides a mock function with given fields: userID, otherUserId
func (_m *SuiteIFace) UserCanSeeOtherUser(userID string, otherUserId string) (bool, *model.AppError) {
	ret := _m.Called(userID, otherUserId)

	var r0 bool
	var r1 *model.AppError
	if rf, ok := ret.Get(0).(func(string, string) (bool, *model.AppError)); ok {
		return rf(userID, otherUserId)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(userID, otherUserId)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) *model.AppError); ok {
		r1 = rf(userID, otherUserId)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.AppError)
		}
	}

	return r0, r1
}

// NewSuiteIFace creates a new instance of SuiteIFace. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSuiteIFace(t interface {
	mock.TestingT
	Cleanup(func())
}) *SuiteIFace {
	mock := &SuiteIFace{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
