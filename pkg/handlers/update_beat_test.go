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

func TestUpdateBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.UpdateBeatServiceExecRequest

		serviceResp *services.UpdateBeatResponse
		serviceErr  error

		expect     *storystructurev1.UpdateBeatServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.UpdateBeatServiceExecRequest{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceResp: &services.UpdateBeatResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &storystructurev1.UpdateBeatServiceExecResponse{
				Id:        "00000000-0000-0000-0000-000000000001",
				CreatorId: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: timestamppb.New(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/NoUpdate",

			request: &storystructurev1.UpdateBeatServiceExecRequest{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceResp: &services.UpdateBeatResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},

			expect: &storystructurev1.UpdateBeatServiceExecResponse{
				Id:        "00000000-0000-0000-0000-000000000001",
				CreatorId: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.UpdateBeatServiceExecRequest{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceErr: services.ErrInvalidUpdateBeatRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "NotFound",

			request: &storystructurev1.UpdateBeatServiceExecRequest{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceErr: dao.ErrBeatNotFound,

			expectCode: codes.NotFound,
		},
		{
			name: "Internal",

			request: &storystructurev1.UpdateBeatServiceExecRequest{
				Id:        "id",
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockUpdateBeat(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.UpdateBeatRequest{
					ID:        testCase.request.GetId(),
					Name:      testCase.request.GetName(),
					Prompt:    testCase.request.GetPrompt(),
					CreatorID: testCase.request.GetCreatorId(),
				}).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.UpdateBeatServiceName, mock.Anything)

			handler := handlers.NewUpdateBeat(service, logger)

			resp, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
