// Code generated by mockery v2.46.3. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/a-novel/uservice-story-structure/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockListBeats is an autogenerated mock type for the ListBeats type
type MockListBeats struct {
	mock.Mock
}

type MockListBeats_Expecter struct {
	mock *mock.Mock
}

func (_m *MockListBeats) EXPECT() *MockListBeats_Expecter {
	return &MockListBeats_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, ids
func (_m *MockListBeats) Exec(ctx context.Context, ids []uuid.UUID) ([]*entities.Beat, error) {
	ret := _m.Called(ctx, ids)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 []*entities.Beat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) ([]*entities.Beat, error)); ok {
		return rf(ctx, ids)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []uuid.UUID) []*entities.Beat); ok {
		r0 = rf(ctx, ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []uuid.UUID) error); ok {
		r1 = rf(ctx, ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockListBeats_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockListBeats_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - ids []uuid.UUID
func (_e *MockListBeats_Expecter) Exec(ctx interface{}, ids interface{}) *MockListBeats_Exec_Call {
	return &MockListBeats_Exec_Call{Call: _e.mock.On("Exec", ctx, ids)}
}

func (_c *MockListBeats_Exec_Call) Run(run func(ctx context.Context, ids []uuid.UUID)) *MockListBeats_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]uuid.UUID))
	})
	return _c
}

func (_c *MockListBeats_Exec_Call) Return(_a0 []*entities.Beat, _a1 error) *MockListBeats_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListBeats_Exec_Call) RunAndReturn(run func(context.Context, []uuid.UUID) ([]*entities.Beat, error)) *MockListBeats_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockListBeats creates a new instance of MockListBeats. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockListBeats(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockListBeats {
	mock := &MockListBeats{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
