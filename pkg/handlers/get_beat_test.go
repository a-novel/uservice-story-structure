package handlers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	adaptersmocks "github.com/a-novel/golib/loggers/adapters/mocks"
	"github.com/a-novel/golib/testutils"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestGetBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.GetBeatServiceExecRequest

		serviceResp *services.GetBeatResponse
		serviceErr  error

		expect     *storystructurev1.GetBeatServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.GetBeatServiceExecRequest{Id: "id"},

			serviceResp: &services.GetBeatResponse{
				ID:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &storystructurev1.GetBeatServiceExecResponse{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
				CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: timestamppb.New(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/NoUpdateTimestamp",

			request: &storystructurev1.GetBeatServiceExecRequest{Id: "id"},

			serviceResp: &services.GetBeatResponse{
				ID:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},

			expect: &storystructurev1.GetBeatServiceExecResponse{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
				CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.GetBeatServiceExecRequest{Id: "id"},

			serviceErr: services.ErrInvalidGetBeatRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "NotFound",

			request: &storystructurev1.GetBeatServiceExecRequest{Id: "id"},

			serviceErr: dao.ErrBeatNotFound,

			expectCode: codes.NotFound,
		},
		{
			name: "Internal",

			request: &storystructurev1.GetBeatServiceExecRequest{Id: "id"},

			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockGetBeat(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.GetBeatRequest{ID: testCase.request.GetId()}).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.GetBeatServiceName, mock.Anything)

			handler := handlers.NewGetBeat(service, logger)

			resp, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
