package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	daomocks "github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestListPlotPoints(t *testing.T) {
	testCases := []struct {
		name string

		request *services.ListPlotPointsRequest

		shouldCallListPlotPointsDAO bool
		listPlotPointsDAOResponse   []*entities.PlotPoint
		listPlotPointsDAOError      error

		expect    *services.ListPlotPointsResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.ListPlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListPlotPointsDAO: true,
			listPlotPointsDAOResponse: []*entities.PlotPoint{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					CreatorID: "creator_id_1",
					Name:      "Plot Point 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					CreatorID: "creator_id_2",
					Name:      "Plot Point 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},

			expect: &services.ListPlotPointsResponse{
				PlotPoints: []*services.ListPlotPointsResponsePlotPoint{
					{
						ID:        "00000000-0000-0000-0000-000000000001",
						CreatorID: "creator_id_1",
						Name:      "Plot Point 1",
						Prompt:    "Prompt 1",
						CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						ID:        "00000000-0000-0000-0000-000000000002",
						CreatorID: "creator_id_2",
						Name:      "Plot Point 2",
						Prompt:    "Prompt 2",
						CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.ListPlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListPlotPointsDAO: true,
			listPlotPointsDAOResponse:   []*entities.PlotPoint{},

			expect: &services.ListPlotPointsResponse{
				PlotPoints: []*services.ListPlotPointsResponsePlotPoint{},
			},
		},
		{
			name: "DAO/Error",

			request: &services.ListPlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListPlotPointsDAO: true,
			listPlotPointsDAOError:      errors.New("dao error"),

			expectErr: services.ErrListPlotPoints,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.ListPlotPointsRequest{},

			expectErr: services.ErrInvalidListPlotPointsRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.ListPlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000x0000x0000x0000x000000000002",
				},
			},

			expectErr: services.ErrInvalidListPlotPointsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			listPlotPointsDAO := daomocks.NewMockListPlotPoints(t)

			if testCase.shouldCallListPlotPointsDAO {
				listPlotPointsDAO.
					On(
						"Exec",
						context.Background(),
						mock.MatchedBy(func(ids uuid.UUIDs) bool {
							strIDs := ids.Strings()
							if len(strIDs) != len(testCase.request.IDs) {
								return false
							}

							for index, id := range testCase.request.IDs {
								if id != strIDs[index] {
									return false
								}
							}

							return true
						}),
					).
					Return(testCase.listPlotPointsDAOResponse, testCase.listPlotPointsDAOError)
			}

			service := services.NewListPlotPoints(listPlotPointsDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			listPlotPointsDAO.AssertExpectations(t)
		})
	}
}
