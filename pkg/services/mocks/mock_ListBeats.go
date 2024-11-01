// Code generated by mockery v2.46.3. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	services "github.com/a-novel/uservice-story-structure/pkg/services"
	mock "github.com/stretchr/testify/mock"
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

// Exec provides a mock function with given fields: ctx, data
func (_m *MockListBeats) Exec(ctx context.Context, data *services.ListBeatsRequest) (*services.ListBeatsResponse, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *services.ListBeatsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *services.ListBeatsRequest) (*services.ListBeatsResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *services.ListBeatsRequest) *services.ListBeatsResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.ListBeatsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *services.ListBeatsRequest) error); ok {
		r1 = rf(ctx, data)
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
//   - data *services.ListBeatsRequest
func (_e *MockListBeats_Expecter) Exec(ctx interface{}, data interface{}) *MockListBeats_Exec_Call {
	return &MockListBeats_Exec_Call{Call: _e.mock.On("Exec", ctx, data)}
}

func (_c *MockListBeats_Exec_Call) Run(run func(ctx context.Context, data *services.ListBeatsRequest)) *MockListBeats_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*services.ListBeatsRequest))
	})
	return _c
}

func (_c *MockListBeats_Exec_Call) Return(_a0 *services.ListBeatsResponse, _a1 error) *MockListBeats_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockListBeats_Exec_Call) RunAndReturn(run func(context.Context, *services.ListBeatsRequest) (*services.ListBeatsResponse, error)) *MockListBeats_Exec_Call {
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
