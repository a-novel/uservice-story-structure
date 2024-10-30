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

	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestListBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.ListBeatsServiceExecRequest

		serviceResp *services.ListBeatsResponse
		serviceErr  error

		expect     *storystructurev1.ListBeatsServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.ListBeatsServiceExecRequest{
				Ids: []string{"id-1", "id-2", "id-3"},
			},

			serviceResp: &services.ListBeatsResponse{
				Beats: []*services.ListBeatsResponseBeat{
					{
						ID:        "00000000-0000-0000-0000-000000000001",
						Name:      "Beat 1",
						Prompt:    "Prompt 1",
						CreatorID: "creator-1",
						CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						ID:        "00000000-0000-0000-0000-000000000002",
						Name:      "Beat 2",
						Prompt:    "Prompt 2",
						CreatorID: "creator-2",
						CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					},
				},
			},

			expect: &storystructurev1.ListBeatsServiceExecResponse{
				Beats: []*storystructurev1.ListBeatsServiceExecResponseElement{
					{
						Id:        "00000000-0000-0000-0000-000000000001",
						Name:      "Beat 1",
						Prompt:    "Prompt 1",
						CreatorId: "creator-1",
						CreatedAt: timestamppb.New(time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC)),
						UpdatedAt: timestamppb.New(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						Id:        "00000000-0000-0000-0000-000000000002",
						Name:      "Beat 2",
						Prompt:    "Prompt 2",
						CreatorId: "creator-2",
						CreatedAt: timestamppb.New(time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.ListBeatsServiceExecRequest{
				Ids: []string{"id-1", "id-2", "id-3"},
			},

			serviceErr: services.ErrInvalidListBeatsRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",

			request: &storystructurev1.ListBeatsServiceExecRequest{
				Ids: []string{"id-1", "id-2", "id-3"},
			},

			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockListBeats(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.ListBeatsRequest{IDs: testCase.request.GetIds()}).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.ListBeatsServiceName, mock.Anything)

			handler := handlers.NewListBeats(service, logger)

			resp, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
