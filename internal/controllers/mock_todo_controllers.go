// Code generated by mockery v2.42.0. DO NOT EDIT.

package controllers

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// MockTodoControllers is an autogenerated mock type for the TodoControllers type
type MockTodoControllers struct {
	mock.Mock
}

type MockTodoControllers_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTodoControllers) EXPECT() *MockTodoControllers_Expecter {
	return &MockTodoControllers_Expecter{mock: &_m.Mock}
}

// AddTodoController provides a mock function with given fields: c
func (_m *MockTodoControllers) AddTodoController(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for AddTodoController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTodoControllers_AddTodoController_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTodoController'
type MockTodoControllers_AddTodoController_Call struct {
	*mock.Call
}

// AddTodoController is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockTodoControllers_Expecter) AddTodoController(c interface{}) *MockTodoControllers_AddTodoController_Call {
	return &MockTodoControllers_AddTodoController_Call{Call: _e.mock.On("AddTodoController", c)}
}

func (_c *MockTodoControllers_AddTodoController_Call) Run(run func(c echo.Context)) *MockTodoControllers_AddTodoController_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockTodoControllers_AddTodoController_Call) Return(_a0 error) *MockTodoControllers_AddTodoController_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoControllers_AddTodoController_Call) RunAndReturn(run func(echo.Context) error) *MockTodoControllers_AddTodoController_Call {
	_c.Call.Return(run)
	return _c
}

// ChangeTodoStatusController provides a mock function with given fields: c
func (_m *MockTodoControllers) ChangeTodoStatusController(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ChangeTodoStatusController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTodoControllers_ChangeTodoStatusController_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeTodoStatusController'
type MockTodoControllers_ChangeTodoStatusController_Call struct {
	*mock.Call
}

// ChangeTodoStatusController is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockTodoControllers_Expecter) ChangeTodoStatusController(c interface{}) *MockTodoControllers_ChangeTodoStatusController_Call {
	return &MockTodoControllers_ChangeTodoStatusController_Call{Call: _e.mock.On("ChangeTodoStatusController", c)}
}

func (_c *MockTodoControllers_ChangeTodoStatusController_Call) Run(run func(c echo.Context)) *MockTodoControllers_ChangeTodoStatusController_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockTodoControllers_ChangeTodoStatusController_Call) Return(_a0 error) *MockTodoControllers_ChangeTodoStatusController_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoControllers_ChangeTodoStatusController_Call) RunAndReturn(run func(echo.Context) error) *MockTodoControllers_ChangeTodoStatusController_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTodoController provides a mock function with given fields: c
func (_m *MockTodoControllers) DeleteTodoController(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodoController")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockTodoControllers_DeleteTodoController_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTodoController'
type MockTodoControllers_DeleteTodoController_Call struct {
	*mock.Call
}

// DeleteTodoController is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockTodoControllers_Expecter) DeleteTodoController(c interface{}) *MockTodoControllers_DeleteTodoController_Call {
	return &MockTodoControllers_DeleteTodoController_Call{Call: _e.mock.On("DeleteTodoController", c)}
}

func (_c *MockTodoControllers_DeleteTodoController_Call) Run(run func(c echo.Context)) *MockTodoControllers_DeleteTodoController_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockTodoControllers_DeleteTodoController_Call) Return(_a0 error) *MockTodoControllers_DeleteTodoController_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoControllers_DeleteTodoController_Call) RunAndReturn(run func(echo.Context) error) *MockTodoControllers_DeleteTodoController_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTodoControllers creates a new instance of MockTodoControllers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTodoControllers(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTodoControllers {
	mock := &MockTodoControllers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
