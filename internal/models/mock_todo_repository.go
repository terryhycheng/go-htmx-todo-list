// Code generated by mockery v2.42.0. DO NOT EDIT.

package models

import mock "github.com/stretchr/testify/mock"

// MockTodoRepository is an autogenerated mock type for the TodoRepository type
type MockTodoRepository struct {
	mock.Mock
}

type MockTodoRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockTodoRepository) EXPECT() *MockTodoRepository_Expecter {
	return &MockTodoRepository_Expecter{mock: &_m.Mock}
}

// AddTodo provides a mock function with given fields: todo
func (_m *MockTodoRepository) AddTodo(todo TodoGorm) *TodoGorm {
	ret := _m.Called(todo)

	if len(ret) == 0 {
		panic("no return value specified for AddTodo")
	}

	var r0 *TodoGorm
	if rf, ok := ret.Get(0).(func(TodoGorm) *TodoGorm); ok {
		r0 = rf(todo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TodoGorm)
		}
	}

	return r0
}

// MockTodoRepository_AddTodo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddTodo'
type MockTodoRepository_AddTodo_Call struct {
	*mock.Call
}

// AddTodo is a helper method to define mock.On call
//   - todo TodoGorm
func (_e *MockTodoRepository_Expecter) AddTodo(todo interface{}) *MockTodoRepository_AddTodo_Call {
	return &MockTodoRepository_AddTodo_Call{Call: _e.mock.On("AddTodo", todo)}
}

func (_c *MockTodoRepository_AddTodo_Call) Run(run func(todo TodoGorm)) *MockTodoRepository_AddTodo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(TodoGorm))
	})
	return _c
}

func (_c *MockTodoRepository_AddTodo_Call) Return(_a0 *TodoGorm) *MockTodoRepository_AddTodo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_AddTodo_Call) RunAndReturn(run func(TodoGorm) *TodoGorm) *MockTodoRepository_AddTodo_Call {
	_c.Call.Return(run)
	return _c
}

// ChangeStatus provides a mock function with given fields: id
func (_m *MockTodoRepository) ChangeStatus(id int64) *TodoGorm {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for ChangeStatus")
	}

	var r0 *TodoGorm
	if rf, ok := ret.Get(0).(func(int64) *TodoGorm); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TodoGorm)
		}
	}

	return r0
}

// MockTodoRepository_ChangeStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChangeStatus'
type MockTodoRepository_ChangeStatus_Call struct {
	*mock.Call
}

// ChangeStatus is a helper method to define mock.On call
//   - id int64
func (_e *MockTodoRepository_Expecter) ChangeStatus(id interface{}) *MockTodoRepository_ChangeStatus_Call {
	return &MockTodoRepository_ChangeStatus_Call{Call: _e.mock.On("ChangeStatus", id)}
}

func (_c *MockTodoRepository_ChangeStatus_Call) Run(run func(id int64)) *MockTodoRepository_ChangeStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockTodoRepository_ChangeStatus_Call) Return(_a0 *TodoGorm) *MockTodoRepository_ChangeStatus_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_ChangeStatus_Call) RunAndReturn(run func(int64) *TodoGorm) *MockTodoRepository_ChangeStatus_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteTodoById provides a mock function with given fields: id
func (_m *MockTodoRepository) DeleteTodoById(id int64) TodoGorm {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteTodoById")
	}

	var r0 TodoGorm
	if rf, ok := ret.Get(0).(func(int64) TodoGorm); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(TodoGorm)
	}

	return r0
}

// MockTodoRepository_DeleteTodoById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTodoById'
type MockTodoRepository_DeleteTodoById_Call struct {
	*mock.Call
}

// DeleteTodoById is a helper method to define mock.On call
//   - id int64
func (_e *MockTodoRepository_Expecter) DeleteTodoById(id interface{}) *MockTodoRepository_DeleteTodoById_Call {
	return &MockTodoRepository_DeleteTodoById_Call{Call: _e.mock.On("DeleteTodoById", id)}
}

func (_c *MockTodoRepository_DeleteTodoById_Call) Run(run func(id int64)) *MockTodoRepository_DeleteTodoById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockTodoRepository_DeleteTodoById_Call) Return(_a0 TodoGorm) *MockTodoRepository_DeleteTodoById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_DeleteTodoById_Call) RunAndReturn(run func(int64) TodoGorm) *MockTodoRepository_DeleteTodoById_Call {
	_c.Call.Return(run)
	return _c
}

// GetTodoById provides a mock function with given fields: id
func (_m *MockTodoRepository) GetTodoById(id int64) *TodoGorm {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetTodoById")
	}

	var r0 *TodoGorm
	if rf, ok := ret.Get(0).(func(int64) *TodoGorm); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*TodoGorm)
		}
	}

	return r0
}

// MockTodoRepository_GetTodoById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTodoById'
type MockTodoRepository_GetTodoById_Call struct {
	*mock.Call
}

// GetTodoById is a helper method to define mock.On call
//   - id int64
func (_e *MockTodoRepository_Expecter) GetTodoById(id interface{}) *MockTodoRepository_GetTodoById_Call {
	return &MockTodoRepository_GetTodoById_Call{Call: _e.mock.On("GetTodoById", id)}
}

func (_c *MockTodoRepository_GetTodoById_Call) Run(run func(id int64)) *MockTodoRepository_GetTodoById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockTodoRepository_GetTodoById_Call) Return(_a0 *TodoGorm) *MockTodoRepository_GetTodoById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_GetTodoById_Call) RunAndReturn(run func(int64) *TodoGorm) *MockTodoRepository_GetTodoById_Call {
	_c.Call.Return(run)
	return _c
}

// GetTodos provides a mock function with given fields:
func (_m *MockTodoRepository) GetTodos() []TodoGorm {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetTodos")
	}

	var r0 []TodoGorm
	if rf, ok := ret.Get(0).(func() []TodoGorm); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]TodoGorm)
		}
	}

	return r0
}

// MockTodoRepository_GetTodos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTodos'
type MockTodoRepository_GetTodos_Call struct {
	*mock.Call
}

// GetTodos is a helper method to define mock.On call
func (_e *MockTodoRepository_Expecter) GetTodos() *MockTodoRepository_GetTodos_Call {
	return &MockTodoRepository_GetTodos_Call{Call: _e.mock.On("GetTodos")}
}

func (_c *MockTodoRepository_GetTodos_Call) Run(run func()) *MockTodoRepository_GetTodos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockTodoRepository_GetTodos_Call) Return(_a0 []TodoGorm) *MockTodoRepository_GetTodos_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockTodoRepository_GetTodos_Call) RunAndReturn(run func() []TodoGorm) *MockTodoRepository_GetTodos_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockTodoRepository creates a new instance of MockTodoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockTodoRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockTodoRepository {
	mock := &MockTodoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}