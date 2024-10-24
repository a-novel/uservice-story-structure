package services_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	daomocks "github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestUpdatePlotPoint(t *testing.T) {
	testCases := []struct {
		name string

		request *services.UpdatePlotPointRequest

		shouldCallUpdatePlotPointDAO bool
		updatePlotPointDAOResponse   *entities.PlotPoint
		updatePlotPointDAOError      error

		expect    *services.UpdatePlotPointResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "name",
				Prompt: "prompt",
			},

			shouldCallUpdatePlotPointDAO: true,
			updatePlotPointDAOResponse: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &services.UpdatePlotPointResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
		},

		{
			name: "DAO/Error",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "name",
				Prompt: "prompt",
			},

			shouldCallUpdatePlotPointDAO: true,
			updatePlotPointDAOError:      errors.New("dao error"),

			expectErr: services.ErrUpdatePlotPoint,
		},

		{
			name: "InvalidRequest/NoName",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NoPrompt",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "name",
				Prompt: "",
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.UpdatePlotPointRequest{
				ID:     "",
				Name:   "name",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000x0000x0000x0000x000000000001",
				Name:   "name",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NameTooLong",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   strings.Repeat("a", 65),
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
		{
			name: "InvalidRequest/PromptTooLong",

			request: &services.UpdatePlotPointRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "name",
				Prompt: strings.Repeat("a", 1025),
			},

			expectErr: services.ErrInvalidUpdatePlotPointRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			updatePlotPointDAO := daomocks.NewMockUpdatePlotPoint(t)

			if testCase.shouldCallUpdatePlotPointDAO {
				updatePlotPointDAO.
					On(
						"Exec",
						context.Background(),
						mock.MatchedBy(func(id uuid.UUID) bool { return id != uuid.Nil }),
						mock.MatchedBy(func(at time.Time) bool { return at.Unix() > 0 }),
						&dao.UpdatePlotPointRequest{
							Name:   testCase.request.Name,
							Prompt: testCase.request.Prompt,
						},
					).
					Return(testCase.updatePlotPointDAOResponse, testCase.updatePlotPointDAOError)
			}

			service := services.NewUpdatePlotPoint(updatePlotPointDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			updatePlotPointDAO.AssertExpectations(t)
		})
	}
}
