package services_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	daomocks "github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestGetPlotPoint(t *testing.T) {
	testCases := []struct {
		name string

		request *services.GetPlotPointRequest

		shouldCallGetPlotPointDAO bool
		getPlotPointDAOResponse   *entities.PlotPoint
		getPlotPointDAOError      error

		expect    *services.GetPlotPointResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.GetPlotPointRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallGetPlotPointDAO: true,
			getPlotPointDAOResponse: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				CreatorID: "creator_id",
				Name:      "PlotPoint 1",
				Prompt:    "Prompt 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &services.GetPlotPointResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
				Name:      "PlotPoint 1",
				Prompt:    "Prompt 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "DAO/Error",

			request: &services.GetPlotPointRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallGetPlotPointDAO: true,
			getPlotPointDAOError:      errors.New("dao error"),

			expectErr: services.ErrGetPlotPoint,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.GetPlotPointRequest{},

			expectErr: services.ErrInvalidGetPlotPointRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.GetPlotPointRequest{
				ID: "00000000x0000x0000x0000x000000000001",
			},

			expectErr: services.ErrInvalidGetPlotPointRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			getPlotPointDAO := daomocks.NewMockGetPlotPoint(t)

			if testCase.shouldCallGetPlotPointDAO {
				getPlotPointDAO.
					On(
						"Exec",
						context.Background(),
						uuid.MustParse(testCase.request.ID),
					).
					Return(testCase.getPlotPointDAOResponse, testCase.getPlotPointDAOError)
			}

			service := services.NewGetPlotPoint(getPlotPointDAO)
			response, err := service.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			getPlotPointDAO.AssertExpectations(t)
		})
	}
}
