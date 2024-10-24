// Code generated by mockery v2.46.0. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	services "github.com/a-novel/uservice-story-structure/pkg/services"
	mock "github.com/stretchr/testify/mock"
)

// MockUpdateBeat is an autogenerated mock type for the UpdateBeat type
type MockUpdateBeat struct {
	mock.Mock
}

type MockUpdateBeat_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdateBeat) EXPECT() *MockUpdateBeat_Expecter {
	return &MockUpdateBeat_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, data
func (_m *MockUpdateBeat) Exec(ctx context.Context, data *services.UpdateBeatRequest) (*services.UpdateBeatResponse, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *services.UpdateBeatResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *services.UpdateBeatRequest) (*services.UpdateBeatResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *services.UpdateBeatRequest) *services.UpdateBeatResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.UpdateBeatResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *services.UpdateBeatRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdateBeat_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockUpdateBeat_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - data *services.UpdateBeatRequest
func (_e *MockUpdateBeat_Expecter) Exec(ctx interface{}, data interface{}) *MockUpdateBeat_Exec_Call {
	return &MockUpdateBeat_Exec_Call{Call: _e.mock.On("Exec", ctx, data)}
}

func (_c *MockUpdateBeat_Exec_Call) Run(run func(ctx context.Context, data *services.UpdateBeatRequest)) *MockUpdateBeat_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*services.UpdateBeatRequest))
	})
	return _c
}

func (_c *MockUpdateBeat_Exec_Call) Return(_a0 *services.UpdateBeatResponse, _a1 error) *MockUpdateBeat_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdateBeat_Exec_Call) RunAndReturn(run func(context.Context, *services.UpdateBeatRequest) (*services.UpdateBeatResponse, error)) *MockUpdateBeat_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdateBeat creates a new instance of MockUpdateBeat. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdateBeat(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdateBeat {
	mock := &MockUpdateBeat{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
