package handlers_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	commonv1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/common/v1"
	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	adaptersmocks "github.com/a-novel/golib/loggers/adapters/mocks"
	"github.com/a-novel/golib/testutils"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/handlers"
	"github.com/a-novel/uservice-story-structure/pkg/services"
	servicesmocks "github.com/a-novel/uservice-story-structure/pkg/services/mocks"
)

func TestSearchBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.SearchBeatsServiceExecRequest

		shouldCallServiceWith *services.SearchBeatsRequest
		serviceResp           *services.SearchBeatsResponse
		serviceErr            error

		expect     *storystructurev1.SearchBeatsServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK/Default",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/CreatorIDs",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				CreatorIds: []string{"creator-1", "creator-2"},
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionNone,
				CreatorIDs:    []string{"creator-1", "creator-2"},
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByAsc",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination:     &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderDirection: commonv1.SortDirection_SORT_DIRECTION_ASC,
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionAsc,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByDesc",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination:     &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderDirection: commonv1.SortDirection_SORT_DIRECTION_DESC,
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionDesc,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByName",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortBeats_SORT_BEATS_BY_NAME,
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatName,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByCreatedAt",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortBeats_SORT_BEATS_BY_CREATED_AT,
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatCreatedAt,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByUpdatedAt",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortBeats_SORT_BEATS_BY_UPDATED_AT,
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatUpdatedAt,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchBeatsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchBeatsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},

		{
			name: "InvalidArgument",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceErr: services.ErrInvalidSearchBeatsRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",

			request: &storystructurev1.SearchBeatsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortBeatNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockSearchBeats(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), testCase.shouldCallServiceWith).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.SearchBeatsServiceName, mock.Anything)

			handler := handlers.NewSearchBeats(service, logger)

			resp, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
