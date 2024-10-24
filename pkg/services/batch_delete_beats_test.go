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

func TestBatchDeleteBeats(t *testing.T) {
	testCases := []struct {
		name string

		request *services.BatchDeleteBeatsRequest

		shouldCallBatchDeleteBeatsDAO bool
		batchDeleteBeatsDAOResponse   []*entities.Beat
		batchDeleteBeatsDAOError      error

		expectErr error
	}{
		{
			name: "OK",

			request: &services.BatchDeleteBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeleteBeatsDAO: true,
			batchDeleteBeatsDAOResponse: []*entities.Beat{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name:      "Beat 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Name:      "Beat 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "OK/NoReturn",

			request: &services.BatchDeleteBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeleteBeatsDAO: true,
		},
		{
			name: "DAO/Error",

			request: &services.BatchDeleteBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000-0000-0000-0000-000000000002",
				},
			},

			shouldCallBatchDeleteBeatsDAO: true,
			batchDeleteBeatsDAOError:      errors.New("dao error"),

			expectErr: services.ErrBatchDeleteBeats,
		},
		{
			name: "InvalidRequest/NoID",

			request: &services.BatchDeleteBeatsRequest{},

			expectErr: services.ErrInvalidBatchDeleteBeatsRequest,
		},
		{
			name: "InvalidRequest/TooManyIDs",

			request: &services.BatchDeleteBeatsRequest{
				IDs: make([]string, 129),
			},

			expectErr: services.ErrInvalidBatchDeleteBeatsRequest,
		},
		{
			name: "InvalidRequest/InvalidID",

			request: &services.BatchDeleteBeatsRequest{
				IDs: []string{
					"00000000-0000-0000-0000-000000000001",
					"00000000x0000x0000x0000x000000000002",
				},
			},

			expectErr: services.ErrInvalidBatchDeleteBeatsRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			batchDeleteBeatsDAO := daomocks.NewMockBatchDeleteBeats(t)

			if testCase.shouldCallBatchDeleteBeatsDAO {
				batchDeleteBeatsDAO.
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
					Return(testCase.batchDeleteBeatsDAOResponse, testCase.batchDeleteBeatsDAOError)
			}

			service := services.NewBatchDeleteBeats(batchDeleteBeatsDAO)
			err := service.Exec(context.Background(), testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)

			batchDeleteBeatsDAO.AssertExpectations(t)
		})
	}
}
