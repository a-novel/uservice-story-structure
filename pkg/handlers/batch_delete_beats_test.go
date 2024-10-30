package handlers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc/codes"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	adaptersmocks "github.com/a-novel/golib/loggers/adapters/mocks"
	"github.com/a-novel/golib/testutils"

	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestBatchDeleteBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.BatchDeleteBeatsServiceExecRequest

		serviceErr error

		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.BatchDeleteBeatsServiceExecRequest{
				Ids:       []string{"id-1", "id-2", "id-3"},
				CreatorId: "creator-id",
			},

			serviceErr: nil,

			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.BatchDeleteBeatsServiceExecRequest{
				Ids:       []string{"id-1", "id-2", "id-3"},
				CreatorId: "creator-id",
			},

			serviceErr: services.ErrInvalidBatchDeleteBeatsRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",

			request: &storystructurev1.BatchDeleteBeatsServiceExecRequest{
				Ids:       []string{"id-1", "id-2", "id-3"},
				CreatorId: "creator-id",
			},

			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockBatchDeleteBeats(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.BatchDeleteBeatsRequest{
					IDs:       testCase.request.GetIds(),
					CreatorID: testCase.request.GetCreatorId(),
				}).
				Return(testCase.serviceErr)

			logger.On("Report", handlers.BatchDeleteBeatsServiceName, mock.Anything)

			handler := handlers.NewBatchDeleteBeats(service, logger)
			_, err := handler.Exec(context.Background(), testCase.request)

			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
