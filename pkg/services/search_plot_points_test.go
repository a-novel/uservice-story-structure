package services_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	daomocks "github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestSearchPlotPoints(t *testing.T) {
	testCases := []struct {
		name string

		request *services.SearchPlotPointsRequest

		shouldCallSearchPlotPointsDAO bool
		searchPlotPointsDAOResponse   uuid.UUIDs
		searchPlotPointsDAOError      error

		expect    *services.SearchPlotPointsResponse
		expectErr error
	}{
		{
			name: "OK/All",

			request: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        0,
				Sort:          "name",
				SortDirection: "asc",
			},

			shouldCallSearchPlotPointsDAO: true,

			searchPlotPointsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchPlotPointsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},
		{
			name: "OK/Defaults",

			request: &services.SearchPlotPointsRequest{
				Limit: 10,
			},

			shouldCallSearchPlotPointsDAO: true,

			searchPlotPointsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchPlotPointsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},
		{
			name: "OK/CreatorIDs",

			request: &services.SearchPlotPointsRequest{
				Limit:         10,
				Offset:        0,
				Sort:          "name",
				SortDirection: "asc",
				CreatorIDs:    []string{"creator_id_1", "creator_id_2"},
			},

			shouldCallSearchPlotPointsDAO: true,

			searchPlotPointsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchPlotPointsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},

		{
			name: "DAO/Error",

			request: &services.SearchPlotPointsRequest{
				Limit: 10,
			},

			shouldCallSearchPlotPointsDAO: true,
			searchPlotPointsDAOError:      errors.New("dao error"),

			expectErr: services.ErrSearchPlotPoints,
		},
		{
			name: "InvalidRequest/InvalidSort",

			request: &services.SearchPlotPointsRequest{
				Limit: 10,
				Sort:  "invalid",
			},

			expectErr: services.ErrInvalidSearchPlotPointsRequest,
		},
		{
			name: "InvalidRequest/InvalidSortDirection",

			request: &services.SearchPlotPointsRequest{
				Limit:         10,
				SortDirection: "invalid",
			},

			expectErr: services.ErrInvalidSearchPlotPointsRequest,
		},
		{
			name: "InvalidRequest/LimitTooLow",

			request: &services.SearchPlotPointsRequest{
				Limit: 0,
			},

			expectErr: services.ErrInvalidSearchPlotPointsRequest,
		},
		{
			name: "InvalidRequest/LimitTooHigh",

			request: &services.SearchPlotPointsRequest{
				Limit: 129,
			},

			expectErr: services.ErrInvalidSearchPlotPointsRequest,
		},
		{
			name: "InvalidRequest/OffsetTooLow",

			request: &services.SearchPlotPointsRequest{
				Limit:  10,
				Offset: -1,
			},

			expectErr: services.ErrInvalidSearchPlotPointsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			searchPlotPointsDAO := daomocks.NewMockSearchPlotPoints(t)

			if testCase.shouldCallSearchPlotPointsDAO {
				searchPlotPointsDAO.
					On("Exec", context.Background(), &dao.SearchPlotPointsRequest{
						Limit:         testCase.request.Limit,
						Offset:        testCase.request.Offset,
						Sort:          testCase.request.Sort,
						SortDirection: testCase.request.SortDirection,
						CreatorIDs:    testCase.request.CreatorIDs,
					}).
					Return(testCase.searchPlotPointsDAOResponse, testCase.searchPlotPointsDAOError)
			}

			service := services.NewSearchPlotPoints(searchPlotPointsDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			searchPlotPointsDAO.AssertExpectations(t)
		})
	}
}
