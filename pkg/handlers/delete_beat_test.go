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

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestDeleteBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.DeleteBeatServiceExecRequest

		serviceErr error

		expectCode codes.Code
	}{
		{
			name: "OK",

			request: &storystructurev1.DeleteBeatServiceExecRequest{Id: "id", CreatorId: "creator_id"},

			expectCode: codes.OK,
		},
		{
			name: "InvalidArgument",

			request: &storystructurev1.DeleteBeatServiceExecRequest{Id: "id", CreatorId: "creator_id"},

			serviceErr: services.ErrInvalidDeleteBeatRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "NotFound",

			request: &storystructurev1.DeleteBeatServiceExecRequest{Id: "id", CreatorId: "creator_id"},

			serviceErr: dao.ErrBeatNotFound,

			expectCode: codes.NotFound,
		},
		{
			name: "InternalError",

			request: &storystructurev1.DeleteBeatServiceExecRequest{Id: "id", CreatorId: "creator_id"},

			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockDeleteBeat(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), &services.DeleteBeatRequest{
					ID:        testCase.request.GetId(),
					CreatorID: testCase.request.GetCreatorId(),
				}).
				Return(testCase.serviceErr)

			logger.On("Report", handlers.DeleteBeatServiceName, mock.Anything)

			handler := handlers.NewDeleteBeat(service, logger)

			_, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
