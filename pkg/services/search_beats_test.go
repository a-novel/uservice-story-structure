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

func TestSearchBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *services.SearchBeatsRequest

		shouldCallSearchBeatsDAO bool
		searchBeatsDAOResponse   uuid.UUIDs
		searchBeatsDAOError      error

		expect    *services.SearchBeatsResponse
		expectErr error
	}{
		{
			name: "OK/All",

			request: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        0,
				Sort:          "name",
				SortDirection: "asc",
			},

			shouldCallSearchBeatsDAO: true,

			searchBeatsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchBeatsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},
		{
			name: "OK/Defaults",

			request: &services.SearchBeatsRequest{
				Limit: 10,
			},

			shouldCallSearchBeatsDAO: true,

			searchBeatsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchBeatsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},
		{
			name: "OK/CreatorIDs",

			request: &services.SearchBeatsRequest{
				Limit:         10,
				Offset:        0,
				Sort:          "name",
				SortDirection: "asc",
				CreatorIDs:    []string{"creator_id_1", "creator_id_2"},
			},

			shouldCallSearchBeatsDAO: true,

			searchBeatsDAOResponse: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},

			expect: &services.SearchBeatsResponse{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},
		},

		{
			name: "DAO/Error",

			request: &services.SearchBeatsRequest{
				Limit: 10,
			},

			shouldCallSearchBeatsDAO: true,
			searchBeatsDAOError:      errors.New("dao error"),

			expectErr: services.ErrSearchBeats,
		},
		{
			name: "InvalidRequest/InvalidSort",

			request: &services.SearchBeatsRequest{
				Limit: 10,
				Sort:  "invalid",
			},

			expectErr: services.ErrInvalidSearchBeatsRequest,
		},
		{
			name: "InvalidRequest/InvalidSortDirection",

			request: &services.SearchBeatsRequest{
				Limit:         10,
				SortDirection: "invalid",
			},

			expectErr: services.ErrInvalidSearchBeatsRequest,
		},
		{
			name: "InvalidRequest/LimitTooLow",

			request: &services.SearchBeatsRequest{
				Limit: 0,
			},

			expectErr: services.ErrInvalidSearchBeatsRequest,
		},
		{
			name: "InvalidRequest/LimitTooHigh",

			request: &services.SearchBeatsRequest{
				Limit: 129,
			},

			expectErr: services.ErrInvalidSearchBeatsRequest,
		},
		{
			name: "InvalidRequest/OffsetTooLow",

			request: &services.SearchBeatsRequest{
				Limit:  10,
				Offset: -1,
			},

			expectErr: services.ErrInvalidSearchBeatsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			searchBeatsDAO := daomocks.NewMockSearchBeats(t)

			if testCase.shouldCallSearchBeatsDAO {
				searchBeatsDAO.
					On("Exec", context.Background(), &dao.SearchBeatsRequest{
						Limit:         testCase.request.Limit,
						Offset:        testCase.request.Offset,
						Sort:          testCase.request.Sort,
						SortDirection: testCase.request.SortDirection,
						CreatorIDs:    testCase.request.CreatorIDs,
					}).
					Return(testCase.searchBeatsDAOResponse, testCase.searchBeatsDAOError)
			}

			service := services.NewSearchBeats(searchBeatsDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			searchBeatsDAO.AssertExpectations(t)
		})
	}
}
