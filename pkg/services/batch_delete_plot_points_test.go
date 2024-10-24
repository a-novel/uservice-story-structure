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

func TestBatchDeletePlotPoints(t *testing.T) {
	testCases := []struct {
		name string

		request *services.BatchDeletePlotPointsRequest

		shouldCallBatchDeletePlotPointsDAO bool
		batchDeletePlotPointsDAOResponse   []*entities.PlotPoint
		batchDeletePlotPointsDAOError      error

		expectErr error
	}{
		{
			name: "OK",

			request: &services.BatchDeletePlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeletePlotPointsDAO: true,
			batchDeletePlotPointsDAOResponse: []*entities.PlotPoint{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name:      "PlotPoint 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Name:      "PlotPoint 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.BatchDeletePlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeletePlotPointsDAO: true,
		},
		{
			name: "DAO/Error",

			request: &services.BatchDeletePlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeletePlotPointsDAO: true,
			batchDeletePlotPointsDAOError:      errors.New("dao error"),

			expectErr: services.ErrBatchDeletePlotPoints,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.BatchDeletePlotPointsRequest{},

			expectErr: services.ErrInvalidBatchDeletePlotPointsRequest,
		},
		{
			name: "InvalidRequest/TooManyIDs",

			request: &services.BatchDeletePlotPointsRequest{
				IDs: make([]string, 129),
			},

			expectErr: services.ErrInvalidBatchDeletePlotPointsRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.BatchDeletePlotPointsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000x0000x0000x0000x000000000002",
				},
			},

			expectErr: services.ErrInvalidBatchDeletePlotPointsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			batchDeletePlotPointsDAO := daomocks.NewMockBatchDeletePlotPoints(t)

			if testCase.shouldCallBatchDeletePlotPointsDAO {
				batchDeletePlotPointsDAO.
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
					Return(testCase.batchDeletePlotPointsDAOResponse, testCase.batchDeletePlotPointsDAOError)
			}

			service := services.NewBatchDeletePlotPoints(batchDeletePlotPointsDAO)
			err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)

			batchDeletePlotPointsDAO.AssertExpectations(t)
		})
	}
}
