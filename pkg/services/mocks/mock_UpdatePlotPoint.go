// Code generated by mockery v2.46.3. DO NOT EDIT.

package servicesmocks

import (
	context "context"

	services "github.com/a-novel/uservice-story-structure/pkg/services"
	mock "github.com/stretchr/testify/mock"
)

// MockUpdatePlotPoint is an autogenerated mock type for the UpdatePlotPoint type
type MockUpdatePlotPoint struct {
	mock.Mock
}

type MockUpdatePlotPoint_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUpdatePlotPoint) EXPECT() *MockUpdatePlotPoint_Expecter {
	return &MockUpdatePlotPoint_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, data
func (_m *MockUpdatePlotPoint) Exec(ctx context.Context, data *services.UpdatePlotPointRequest) (*services.UpdatePlotPointResponse, error) {
	ret := _m.Called(ctx, data)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *services.UpdatePlotPointResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *services.UpdatePlotPointRequest) (*services.UpdatePlotPointResponse, error)); ok {
		return rf(ctx, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *services.UpdatePlotPointRequest) *services.UpdatePlotPointResponse); ok {
		r0 = rf(ctx, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*services.UpdatePlotPointResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *services.UpdatePlotPointRequest) error); ok {
		r1 = rf(ctx, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUpdatePlotPoint_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockUpdatePlotPoint_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - data *services.UpdatePlotPointRequest
func (_e *MockUpdatePlotPoint_Expecter) Exec(ctx interface{}, data interface{}) *MockUpdatePlotPoint_Exec_Call {
	return &MockUpdatePlotPoint_Exec_Call{Call: _e.mock.On("Exec", ctx, data)}
}

func (_c *MockUpdatePlotPoint_Exec_Call) Run(run func(ctx context.Context, data *services.UpdatePlotPointRequest)) *MockUpdatePlotPoint_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*services.UpdatePlotPointRequest))
	})
	return _c
}

func (_c *MockUpdatePlotPoint_Exec_Call) Return(_a0 *services.UpdatePlotPointResponse, _a1 error) *MockUpdatePlotPoint_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdatePlotPoint_Exec_Call) RunAndReturn(run func(context.Context, *services.UpdatePlotPointRequest) (*services.UpdatePlotPointResponse, error)) *MockUpdatePlotPoint_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUpdatePlotPoint creates a new instance of MockUpdatePlotPoint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUpdatePlotPoint(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUpdatePlotPoint {
	mock := &MockUpdatePlotPoint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
