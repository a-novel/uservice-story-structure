// Code generated by mockery v2.46.3. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/a-novel/uservice-story-structure/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockGetPlotPoint is an autogenerated mock type for the GetPlotPoint type
type MockGetPlotPoint struct {
	mock.Mock
}

type MockGetPlotPoint_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetPlotPoint) EXPECT() *MockGetPlotPoint_Expecter {
	return &MockGetPlotPoint_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, id
func (_m *MockGetPlotPoint) Exec(ctx context.Context, id uuid.UUID) (*entities.PlotPoint, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *entities.PlotPoint
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*entities.PlotPoint, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *entities.PlotPoint); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.PlotPoint)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetPlotPoint_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockGetPlotPoint_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockGetPlotPoint_Expecter) Exec(ctx interface{}, id interface{}) *MockGetPlotPoint_Exec_Call {
	return &MockGetPlotPoint_Exec_Call{Call: _e.mock.On("Exec", ctx, id)}
}

func (_c *MockGetPlotPoint_Exec_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockGetPlotPoint_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockGetPlotPoint_Exec_Call) Return(_a0 *entities.PlotPoint, _a1 error) *MockGetPlotPoint_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetPlotPoint_Exec_Call) RunAndReturn(run func(context.Context, uuid.UUID) (*entities.PlotPoint, error)) *MockGetPlotPoint_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetPlotPoint creates a new instance of MockGetPlotPoint. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetPlotPoint(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetPlotPoint {
	mock := &MockGetPlotPoint{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
