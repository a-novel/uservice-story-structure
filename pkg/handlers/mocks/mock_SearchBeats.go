// Code generated by mockery v2.46.3. DO NOT EDIT.

package handlersmocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"
)

// MockSearchBeats is an autogenerated mock type for the SearchBeats type
type MockSearchBeats struct {
	mock.Mock
}

type MockSearchBeats_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSearchBeats) EXPECT() *MockSearchBeats_Expecter {
	return &MockSearchBeats_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: _a0, _a1
func (_m *MockSearchBeats) Exec(_a0 context.Context, _a1 *storystructurev1.SearchBeatsServiceExecRequest) (*storystructurev1.SearchBeatsServiceExecResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *storystructurev1.SearchBeatsServiceExecResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *storystructurev1.SearchBeatsServiceExecRequest) (*storystructurev1.SearchBeatsServiceExecResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *storystructurev1.SearchBeatsServiceExecRequest) *storystructurev1.SearchBeatsServiceExecResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*storystructurev1.SearchBeatsServiceExecResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *storystructurev1.SearchBeatsServiceExecRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSearchBeats_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSearchBeats_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *storystructurev1.SearchBeatsServiceExecRequest
func (_e *MockSearchBeats_Expecter) Exec(_a0 interface{}, _a1 interface{}) *MockSearchBeats_Exec_Call {
	return &MockSearchBeats_Exec_Call{Call: _e.mock.On("Exec", _a0, _a1)}
}

func (_c *MockSearchBeats_Exec_Call) Run(run func(_a0 context.Context, _a1 *storystructurev1.SearchBeatsServiceExecRequest)) *MockSearchBeats_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*storystructurev1.SearchBeatsServiceExecRequest))
	})
	return _c
}

func (_c *MockSearchBeats_Exec_Call) Return(_a0 *storystructurev1.SearchBeatsServiceExecResponse, _a1 error) *MockSearchBeats_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSearchBeats_Exec_Call) RunAndReturn(run func(context.Context, *storystructurev1.SearchBeatsServiceExecRequest) (*storystructurev1.SearchBeatsServiceExecResponse, error)) *MockSearchBeats_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSearchBeats creates a new instance of MockSearchBeats. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSearchBeats(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSearchBeats {
	mock := &MockSearchBeats{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
