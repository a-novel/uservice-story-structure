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

func TestDeletePlotPoint(t *testing.T) {
	testCases := []struct {
		name string

		request *services.DeletePlotPointRequest

		shouldCallDeletePlotPointDAO bool
		deletePlotPointDAOResponse   *entities.PlotPoint
		deletePlotPointDAOError      error

		expectErr error
	}{
		{
			name: "OK",

			request: &services.DeletePlotPointRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeletePlotPointDAO: true,
			deletePlotPointDAOResponse: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "OK/CreatorID",

			request: &services.DeletePlotPointRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
			},

			shouldCallDeletePlotPointDAO: true,
			deletePlotPointDAOResponse: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.DeletePlotPointRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeletePlotPointDAO: true,
		},
		{
			name: "DAO/Error",

			request: &services.DeletePlotPointRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeletePlotPointDAO: true,
			deletePlotPointDAOError:      errors.New("dao error"),

			expectErr: services.ErrDeletePlotPoint,
		},

		{
			name: "InvalidRequest/NoID",

			request: &services.DeletePlotPointRequest{},

			expectErr: services.ErrInvalidDeletePlotPointRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.DeletePlotPointRequest{
				ID: "00000000x0000x0000x0000x000000000001",
			},

			expectErr: services.ErrInvalidDeletePlotPointRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			deletePlotPointDAO := daomocks.NewMockDeletePlotPoint(t)

			if testCase.shouldCallDeletePlotPointDAO {
				deletePlotPointDAO.
					On(
						"Exec",
						context.Background(),
						uuid.MustParse(testCase.request.ID),
						testCase.request.CreatorID,
					).
					Return(testCase.deletePlotPointDAOResponse, testCase.deletePlotPointDAOError)
			}

			service := services.NewDeletePlotPoint(deletePlotPointDAO)
			err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)

			deletePlotPointDAO.AssertExpectations(t)
		})
	}
}
