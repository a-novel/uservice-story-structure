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

func TestListBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *services.ListBeatsRequest

		shouldCallListBeatsDAO bool
		listBeatsDAOResponse   []*entities.Beat
		listBeatsDAOError      error

		expect    *services.ListBeatsResponse
		expectErr error
	}{
		{
			name: "OK",

			request: &services.ListBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListBeatsDAO: true,
			listBeatsDAOResponse: []*entities.Beat{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					CreatorID: "creator_id_1",
					Name:      "Beat 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					CreatorID: "creator_id_2",
					Name:      "Beat 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},

			expect: &services.ListBeatsResponse{
				Beats: []*services.ListBeatsResponseBeat{
					{
						ID:        "00000000-0000-0000-0000-000000000001",
						CreatorID: "creator_id_1",
						Name:      "Beat 1",
						Prompt:    "Prompt 1",
						CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
					},
					{
						ID:        "00000000-0000-0000-0000-000000000002",
						CreatorID: "creator_id_2",
						Name:      "Beat 2",
						Prompt:    "Prompt 2",
						CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.ListBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListBeatsDAO: true,
			listBeatsDAOResponse:   []*entities.Beat{},

			expect: &services.ListBeatsResponse{
				Beats: []*services.ListBeatsResponseBeat{},
			},
		},
		{
			name: "DAO/Error",

			request: &services.ListBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallListBeatsDAO: true,
			listBeatsDAOError:      errors.New("dao error"),

			expectErr: services.ErrListBeats,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.ListBeatsRequest{},

			expectErr: services.ErrInvalidListBeatsRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.ListBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000x0000x0000x0000x000000000002",
				},
			},

			expectErr: services.ErrInvalidListBeatsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			listBeatsDAO := daomocks.NewMockListBeats(t)

			if testCase.shouldCallListBeatsDAO {
				listBeatsDAO.
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
					Return(testCase.listBeatsDAOResponse, testCase.listBeatsDAOError)
			}

			service := services.NewListBeats(listBeatsDAO)
			response, err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, response)

			listBeatsDAO.AssertExpectations(t)
		})
	}
}
