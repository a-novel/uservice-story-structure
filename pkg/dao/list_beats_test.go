package dao_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	anoveldb "github.com/a-novel/golib/database"

	"github.com/a-novel/uservice-story-structure/migrations"
	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

func TestListBeats(t *testing.T) {
	fixtures := []interface{}{
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			CreatorID: "creator_id_1",
			Name:      "Beat 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			CreatorID: "creator_id_2",
			Name:      "Beat 2",
			Prompt:    "Prompt 2",
			CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			CreatorID: "creator_id_3",
			Name:      "Beat 3",
			Prompt:    "Prompt 3",
			CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		ids []uuid.UUID

		expect    []*entities.Beat
		expectErr error
	}{
		{
			name: "List",
			ids: []uuid.UUID{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
			expect: []*entities.Beat{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					CreatorID: "creator_id_1",
					Name:      "Beat 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					CreatorID: "creator_id_2",
					Name:      "Beat 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "IgnoreMissingIDs",
			ids: []uuid.UUID{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			},
			expect: []*entities.Beat{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					CreatorID: "creator_id_1",
					Name:      "Beat 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					CreatorID: "creator_id_2",
					Name:      "Beat 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 2, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "NoResults",
			ids: []uuid.UUID{
				uuid.MustParse("00000000-0000-0000-0000-000000000004"),
			},
			expect: []*entities.Beat{},
		},
	}

	database, closer, err := anoveldb.OpenTestDB(&migrations.SQLMigrations)
	require.NoError(t, err)
	defer closer()

	transaction := anoveldb.BeginTestTX(database, fixtures)
	defer anoveldb.RollbackTestTX(transaction)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			listBeatsDAO := dao.NewListBeats(transaction)

			beat, err := listBeatsDAO.Exec(context.Background(), testCase.ids)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, beat)
		})
	}
}