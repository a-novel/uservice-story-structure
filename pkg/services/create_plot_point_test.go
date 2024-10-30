package services_test

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	daomocks "github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestCreatePlotPoint(t *testing.T) {
	testCases := []struct {
		name string

		request *services.CreatePlotPointRequest

		shouldCallCreatePlotPointDAO bool
		createPlotPointDAOResponse   *entities.PlotPoint
		createPlotPointDAOError      error

		expect    *services.CreatePlotPointResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.CreatePlotPointRequest{
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			shouldCallCreatePlotPointDAO: true,
			createPlotPointDAOResponse: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},

			expect: &services.CreatePlotPointResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},

		{
			name: "DAO/Error",

			request: &services.CreatePlotPointRequest{
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			shouldCallCreatePlotPointDAO: true,
			createPlotPointDAOError:      errors.New("dao error"),

			expectErr: services.ErrCreatePlotPoint,
		},

		{
			name: "InvalidRequest/NoName",

			request: &services.CreatePlotPointRequest{
				Name:      "",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NoPrompt",

			request: &services.CreatePlotPointRequest{
				Name:      "name",
				Prompt:    "",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NameTooLong",

			request: &services.CreatePlotPointRequest{
				Name:      strings.Repeat("a", 65),
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
		{
			name: "InvalidRequest/PromptTooLong",

			request: &services.CreatePlotPointRequest{
				Name:      "name",
				Prompt:    strings.Repeat("a", 1025),
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
		{
			name: "InvalidRequest/NoCreatorID",

			request: &services.CreatePlotPointRequest{
				Name:   "name",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
		{
			name: "InvalidRequest/CreatorIDTooLong",

			request: &services.CreatePlotPointRequest{
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: strings.Repeat("a", 129),
			},

			expectErr: services.ErrInvalidCreatePlotPointRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			createPlotPointDAO := daomocks.NewMockCreatePlotPoint(t)

			if testCase.shouldCallCreatePlotPointDAO {
				createPlotPointDAO.
					On(
						"Exec",
						context.Background(),
						mock.MatchedBy(func(id uuid.UUID) bool { return id != uuid.Nil }),
						mock.MatchedBy(func(at time.Time) bool { return at.Unix() > 0 }),
						&dao.CreatePlotPointRequest{
							Name:      testCase.request.Name,
							Prompt:    testCase.request.Prompt,
							CreatorID: testCase.request.CreatorID,
						},
					).
					Return(testCase.createPlotPointDAOResponse, testCase.createPlotPointDAOError)
			}

			service := services.NewCreatePlotPoint(createPlotPointDAO)
			response, err := service.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			createPlotPointDAO.AssertExpectations(t)
		})
	}
}
