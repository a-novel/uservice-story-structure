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

func TestUpdateBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *services.UpdateBeatRequest

		shouldCallUpdateBeatDAO bool
		updateBeatDAOResponse   *entities.Beat
		updateBeatDAOError      error

		expect    *services.UpdateBeatResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			shouldCallUpdateBeatDAO: true,
			updateBeatDAOResponse: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &services.UpdateBeatResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				CreatorID: "creator_id",
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
		},

		{
			name: "DAO/Error",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			shouldCallUpdateBeatDAO: true,
			updateBeatDAOError:      errors.New("dao error"),

			expectErr: services.ErrUpdateBeat,
		},

		{
			name: "InvalidRequest/NoName",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
		{
			name: "InvalidRequest/NoPrompt",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    "",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.UpdateBeatRequest{
				ID:        "",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.UpdateBeatRequest{
				ID:        "00000000x0000x0000x0000x000000000001",
				Name:      "name",
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
		{
			name: "InvalidRequest/NameTooLong",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      strings.Repeat("a", 65),
				Prompt:    "prompt",
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
		{
			name: "InvalidRequest/PromptTooLong",

			request: &services.UpdateBeatRequest{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "name",
				Prompt:    strings.Repeat("a", 1025),
				CreatorID: "creator_id",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},

		{
			name: "InvalidRequest/NoCreatorID",

			request: &services.UpdateBeatRequest{
				ID:     "00000000-0000-0000-0000-000000000001",
				Name:   "name",
				Prompt: "prompt",
			},

			expectErr: services.ErrInvalidUpdateBeatRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			updateBeatDAO := daomocks.NewMockUpdateBeat(t)

			if testCase.shouldCallUpdateBeatDAO {
				updateBeatDAO.
					On(
						"Exec",
						context.Background(),
						mock.MatchedBy(func(id uuid.UUID) bool { return id != uuid.Nil }),
						mock.MatchedBy(func(at time.Time) bool { return at.Unix() > 0 }),
						&dao.UpdateBeatRequest{
							Name:      testCase.request.Name,
							Prompt:    testCase.request.Prompt,
							CreatorID: testCase.request.CreatorID,
						},
					).
					Return(testCase.updateBeatDAOResponse, testCase.updateBeatDAOError)
			}

			service := services.NewUpdateBeat(updateBeatDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			updateBeatDAO.AssertExpectations(t)
		})
	}
}
