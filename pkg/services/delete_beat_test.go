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

func TestDeleteBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *services.DeleteBeatRequest

		shouldCallDeleteBeatDAO bool
		deleteBeatDAOResponse   *entities.Beat
		deleteBeatDAOError      error

		expectErr error
	}{
		{
			name: "OK",

			request: &services.DeleteBeatRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeleteBeatDAO: true,
			deleteBeatDAOResponse: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:      "name",
				Prompt:    "prompt",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.DeleteBeatRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeleteBeatDAO: true,
		},
		{
			name: "DAO/Error",

			request: &services.DeleteBeatRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallDeleteBeatDAO: true,
			deleteBeatDAOError:      errors.New("dao error"),

			expectErr: services.ErrDeleteBeat,
		},

		{
			name: "InvalidRequest/NoID",

			request: &services.DeleteBeatRequest{},

			expectErr: services.ErrInvalidDeleteBeatRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.DeleteBeatRequest{
				ID: "00000000x0000x0000x0000x000000000001",
			},

			expectErr: services.ErrInvalidDeleteBeatRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			deleteBeatDAO := daomocks.NewMockDeleteBeat(t)

			if testCase.shouldCallDeleteBeatDAO {
				deleteBeatDAO.
					On(
						"Exec",
						context.Background(),
						uuid.MustParse(testCase.request.ID),
					).
					Return(testCase.deleteBeatDAOResponse, testCase.deleteBeatDAOError)
			}

			service := services.NewDeleteBeat(deleteBeatDAO)
			err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)

			deleteBeatDAO.AssertExpectations(t)
		})
	}
}
