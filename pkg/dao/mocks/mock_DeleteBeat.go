// Code generated by mockery v2.46.3. DO NOT EDIT.

package daomocks

import (
	context "context"

	entities "github.com/a-novel/uservice-story-structure/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockDeleteBeat is an autogenerated mock type for the DeleteBeat type
type MockDeleteBeat struct {
	mock.Mock
}

type MockDeleteBeat_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDeleteBeat) EXPECT() *MockDeleteBeat_Expecter {
	return &MockDeleteBeat_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, id, creatorID
func (_m *MockDeleteBeat) Exec(ctx context.Context, id uuid.UUID, creatorID string) (*entities.Beat, error) {
	ret := _m.Called(ctx, id, creatorID)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *entities.Beat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) (*entities.Beat, error)); ok {
		return rf(ctx, id, creatorID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) *entities.Beat); ok {
		r0 = rf(ctx, id, creatorID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, id, creatorID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDeleteBeat_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockDeleteBeat_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
//   - creatorID string
func (_e *MockDeleteBeat_Expecter) Exec(ctx interface{}, id interface{}, creatorID interface{}) *MockDeleteBeat_Exec_Call {
	return &MockDeleteBeat_Exec_Call{Call: _e.mock.On("Exec", ctx, id, creatorID)}
}

func (_c *MockDeleteBeat_Exec_Call) Run(run func(ctx context.Context, id uuid.UUID, creatorID string)) *MockDeleteBeat_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *MockDeleteBeat_Exec_Call) Return(_a0 *entities.Beat, _a1 error) *MockDeleteBeat_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDeleteBeat_Exec_Call) RunAndReturn(run func(context.Context, uuid.UUID, string) (*entities.Beat, error)) *MockDeleteBeat_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDeleteBeat creates a new instance of MockDeleteBeat. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDeleteBeat(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDeleteBeat {
	mock := &MockDeleteBeat{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
