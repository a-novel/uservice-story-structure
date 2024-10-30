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

func TestSearchPlotPoints(t *testing.T) {
	testCases := []struct {
		name string

		request *storystructurev1.SearchPlotPointsServiceExecRequest

		shouldCallServiceWith *services.SearchPlotPointsRequest
		serviceResp           *services.SearchPlotPointsResponse
		serviceErr            error

		expect     *storystructurev1.SearchPlotPointsServiceExecResponse
		expectCode codes.Code
	}{
		{
			name: "OK/Default",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/CreatorIDs",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				CreatorIds: []string{"creator-1", "creator-2"},
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionNone,
				CreatorIDs:    []string{"creator-1", "creator-2"},
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByAsc",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination:     &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderDirection: commonv1.SortDirection_SORT_DIRECTION_ASC,
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionAsc,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByDesc",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination:     &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderDirection: commonv1.SortDirection_SORT_DIRECTION_DESC,
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionDesc,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByName",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_NAME,
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointName,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByCreatedAt",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_CREATED_AT,
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointCreatedAt,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},
		{
			name: "OK/SortByUpdatedAt",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
				OrderBy:    storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_UPDATED_AT,
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointUpdatedAt,
				SortDirection: entities.SortDirectionNone,
			},
			serviceResp: &services.SearchPlotPointsResponse{
				IDs: []string{"id-1", "id-2", "id-3"},
			},

			expect: &storystructurev1.SearchPlotPointsServiceExecResponse{
				Ids: []string{"id-1", "id-2", "id-3"},
			},
			expectCode: codes.OK,
		},

		{
			name: "InvalidArgument",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceErr: services.ErrInvalidSearchPlotPointsRequest,

			expectCode: codes.InvalidArgument,
		},
		{
			name: "Internal",

			request: &storystructurev1.SearchPlotPointsServiceExecRequest{
				Pagination: &commonv1.Pagination{Limit: 10, Offset: 20},
			},

			shouldCallServiceWith: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        20,
				Sort:          entities.SortPlotPointNone,
				SortDirection: entities.SortDirectionNone,
			},
			serviceErr: errors.New("uwups"),

			expectCode: codes.Internal,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			service := servicesmocks.NewMockSearchPlotPoints(t)
			logger := adaptersmocks.NewMockGRPC(t)

			service.
				On("Exec", context.Background(), testCase.shouldCallServiceWith).
				Return(testCase.serviceResp, testCase.serviceErr)

			logger.On("Report", handlers.SearchPlotPointsServiceName, mock.Anything)

			handler := handlers.NewSearchPlotPoints(service, logger)

			resp, err := handler.Exec(context.Background(), testCase.request)
			testutils.RequireGRPCCodesEqual(t, err, testCase.expectCode)
			require.Equal(t, testCase.expect, resp)

			service.AssertExpectations(t)
			logger.AssertExpectations(t)
		})
	}
}
