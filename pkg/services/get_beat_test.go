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

func TestGetBeat(t *testing.T) {
	testCases := []struct {
		name string

		request *services.GetBeatRequest

		shouldCallGetBeatDAO bool
		getBeatDAOResponse   *entities.Beat
		getBeatDAOError      error

		expect    *services.GetBeatResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.GetBeatRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallGetBeatDAO: true,
			getBeatDAOResponse: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Name:      "Beat 1",
				Prompt:    "Prompt 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},

			expect: &services.GetBeatResponse{
				ID:        "00000000-0000-0000-0000-000000000001",
				Name:      "Beat 1",
				Prompt:    "Prompt 1",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
			},
		},
		{
			name: "DAO/Error",

			request: &services.GetBeatRequest{
				ID: "00000000-0000-0000-0000-000000000001",
			},

			shouldCallGetBeatDAO: true,
			getBeatDAOError:      errors.New("dao error"),

			expectErr: services.ErrGetBeat,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.GetBeatRequest{},

			expectErr: services.ErrInvalidGetBeatRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.GetBeatRequest{
				ID: "00000000x0000x0000x0000x000000000001",
			},

			expectErr: services.ErrInvalidGetBeatRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			getBeatDAO := daomocks.NewMockGetBeat(t)

			if testCase.shouldCallGetBeatDAO {
				getBeatDAO.
					On(
						"Exec",
						context.Background(),
						uuid.MustParse(testCase.request.ID),
					).
					Return(testCase.getBeatDAOResponse, testCase.getBeatDAOError)
			}

			service := services.NewGetBeat(getBeatDAO)
			response, err := service.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			getBeatDAO.AssertExpectations(t)
		})
	}
}
