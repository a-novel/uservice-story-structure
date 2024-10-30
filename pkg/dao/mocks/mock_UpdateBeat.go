// Code generated by mockery v2.46.3. DO NOT EDIT.

package daomocks

import (
	context "context"

	dao "github.com/a-novel/uservice-story-structure/pkg/dao"
	entities "github.com/a-novel/uservice-story-structure/pkg/entities"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
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

// Exec provides a mock function with given fields: ctx, id, now, data
func (_m *MockUpdateBeat) Exec(ctx context.Context, id uuid.UUID, now time.Time, data *dao.UpdateBeatRequest) (*entities.Beat, error) {
	ret := _m.Called(ctx, id, now, data)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 *entities.Beat
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, *dao.UpdateBeatRequest) (*entities.Beat, error)); ok {
		return rf(ctx, id, now, data)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, time.Time, *dao.UpdateBeatRequest) *entities.Beat); ok {
		r0 = rf(ctx, id, now, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Beat)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, time.Time, *dao.UpdateBeatRequest) error); ok {
		r1 = rf(ctx, id, now, data)
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
//   - id uuid.UUID
//   - now time.Time
//   - data *dao.UpdateBeatRequest
func (_e *MockUpdateBeat_Expecter) Exec(ctx interface{}, id interface{}, now interface{}, data interface{}) *MockUpdateBeat_Exec_Call {
	return &MockUpdateBeat_Exec_Call{Call: _e.mock.On("Exec", ctx, id, now, data)}
}

func (_c *MockUpdateBeat_Exec_Call) Run(run func(ctx context.Context, id uuid.UUID, now time.Time, data *dao.UpdateBeatRequest)) *MockUpdateBeat_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID), args[2].(time.Time), args[3].(*dao.UpdateBeatRequest))
	})
	return _c
}

func (_c *MockUpdateBeat_Exec_Call) Return(_a0 *entities.Beat, _a1 error) *MockUpdateBeat_Exec_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUpdateBeat_Exec_Call) RunAndReturn(run func(context.Context, uuid.UUID, time.Time, *dao.UpdateBeatRequest) (*entities.Beat, error)) *MockUpdateBeat_Exec_Call {
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