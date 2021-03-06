// Code generated by mockery v2.1.0. DO NOT EDIT.

package project

import (
	context "context"

	models "github.com/goharbor/harbor/src/common/models"
	mock "github.com/stretchr/testify/mock"

	project "github.com/goharbor/harbor/src/controller/project"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, projectID, options
func (_m *Controller) Get(ctx context.Context, projectID int64, options ...project.Option) (*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, int64, ...project.Option) *models.Project); ok {
		r0 = rf(ctx, projectID, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, ...project.Option) error); ok {
		r1 = rf(ctx, projectID, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByName provides a mock function with given fields: ctx, projectName, options
func (_m *Controller) GetByName(ctx context.Context, projectName string, options ...project.Option) (*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, projectName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(context.Context, string, ...project.Option) *models.Project); ok {
		r0 = rf(ctx, projectName, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...project.Option) error); ok {
		r1 = rf(ctx, projectName, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, query, options
func (_m *Controller) List(ctx context.Context, query *models.ProjectQueryParam, options ...project.Option) ([]*models.Project, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*models.Project
	if rf, ok := ret.Get(0).(func(context.Context, *models.ProjectQueryParam, ...project.Option) []*models.Project); ok {
		r0 = rf(ctx, query, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.ProjectQueryParam, ...project.Option) error); ok {
		r1 = rf(ctx, query, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
