package handlers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	adaptersmocks "github.com/a-novel/golib/loggers/adapters/mocks"
	"github.com/a-novel/golib/testutils"

	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestCreateBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.CreateBeatServiceExecRequest

		serviceResp *services.CreateBeatResponse
		serviceErr  error

		expect     *storystructurev1.CreateBeatServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.CreateBeatServiceExecRequest{
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceResp: &services.CreateBeatResponse{
				ID:        "id",
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},

			expect: &storystructurev1.CreateBeatServiceExecResponse{
				Id:        "id",
				CreatorId: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: timestamppb.New(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
				UpdatedAt: nil,
			},
			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.CreateBeatServiceExecRequest{
				Name:      "name",
				Prompt:    "prompt",
				CreatorId: "creator_id",
			},

			serviceErr: services.ErrInvalidCreateBeatRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",

			request: &storystructurev1.CreateBeatServiceExecRequest{
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
			service := servicesmocks.NewMockCreateBeat(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.CreateBeatRequest{
					Name:      testCase.request.GetName(),
					Prompt:    testCase.request.GetPrompt(),
					CreatorID: testCase.request.GetCreatorId(),
				}).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.CreateBeatServiceName, mock.Anything)

			handler := handlers.NewCreateBeat(service, logger)
			resp, err := handler.Exec(context.Background(), testCase.request)

			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
