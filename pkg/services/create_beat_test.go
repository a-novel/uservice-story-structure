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
	"github.com/a-novel/uservice-story-structure/pkg/dao/mocks"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
	"github.com/a-novel/uservice-story-structure/pkg/services"
)

func TestCreateBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *services.CreateBeatRequest

		shouldCallCreateBeatDAO bool
		createBeatDAOResponse   *entities.Beat
		createBeatDAOError      error

		expect    *services.CreateBeatResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.CreateBeatRequest{
				Name:   "name",
				Prompt: "prompt",
			},

			shouldCallCreateBeatDAO: true,
			createBeatDAOResponse: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},

			expect: &services.CreateBeatResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},

		{
			name: "DAO/Error",

			request: &services.CreateBeatRequest{
				Name:   "name",
				Prompt: "prompt",
			},

			shouldCallCreateBeatDAO: true,
			createBeatDAOError:      errors.New("dao error"),

			expectErr: services.ErrCreateBeat,
		},

		{
			name: "InvalidRequest/NoName",

			request: &services.CreateBeatRequest{
				Name:   "",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidCreateBeatRequest,
		},
		{
			name: "InvalidRequest/NoPrompt",

			request: &services.CreateBeatRequest{
				Name:   "name",
				Prompt: "",
			},

			expectErr: services.ErrInvalidCreateBeatRequest,
		},
		{
			name: "InvalidRequest/NameTooLong",

			request: &services.CreateBeatRequest{
				Name:   strings.Repeat("a", 65),
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidCreateBeatRequest,
		},
		{
			name: "InvalidRequest/PromptTooLong",

			request: &services.CreateBeatRequest{
				Name:   "name",
				Prompt: strings.Repeat("a", 1025),
			},

			expectErr: services.ErrInvalidCreateBeatRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			createBeatDAO := daomocks.NewMockCreateBeat(t)

			if testCase.shouldCallCreateBeatDAO {
				createBeatDAO.
					On(
						"Exec",
						context.Background(),
						mock.MatchedBy(func(id uuid.UUID) bool { return id != uuid.Nil }),
						mock.MatchedBy(func(at time.Time) bool { return at.Unix() > 0 }),
						&dao.CreateBeatRequest{
							Name:   testCase.request.Name,
							Prompt: testCase.request.Prompt,
						},
					).
					Return(testCase.createBeatDAOResponse, testCase.createBeatDAOError)
			}

			service := services.NewCreateBeat(createBeatDAO)
			response, err := service.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			createBeatDAO.AssertExpectations(t)
		})
	}
}
