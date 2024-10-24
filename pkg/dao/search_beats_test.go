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

func TestSearchBeats(t *testing.T) {
	fixtures := []interface{}{
		// Order by name: Beat 1, Beat 2, Beat 3
		// Order by created_at: Beat 3, Beat 2, Beat 1
		// Order by updated_at: Beat 3, Beat 1, Beat 2
		// Insertion order: Beat 2, Beat 1, Beat 3

		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Name:      "Beat 2",
			Prompt:    "Prompt 2",
			CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Name:      "Beat 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			Name:      "Beat 3",
			Prompt:    "Prompt 3",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		request *dao.SearchBeatsRequest

		expect    uuid.UUIDs
		expectErr error
	}{
		// Base.
		{
			name: "OK",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 0,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},

		// Limit.
		{
			name: "LimitTooLow",
			request: &dao.SearchBeatsRequest{
				Limit:  2,
				Offset: 0,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "LimitTooHigh",
			request: &dao.SearchBeatsRequest{
				Limit:  10,
				Offset: 0,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},

		// Offset.
		{
			name: "Offset",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 1,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
		{
			name: "OffsetTooHigh",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 10,
			},
			expect: uuid.UUIDs{},
		},

		// Sort: name
		{
			name: "Name",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortBeatName,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
		{
			name: "NameAsc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortBeatName,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
		{
			name: "NameDesc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortBeatName,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},

		// Sort: created_at
		{
			name: "CreatedAt",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortBeatCreatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			name: "CreatedAtAsc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortBeatCreatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			name: "CreatedAtDesc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortBeatCreatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},

		// Sort: updated_at
		{
			name: "UpdatedAt",
			request: &dao.SearchBeatsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortBeatUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "UpdatedAtAsc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortBeatUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "UpdatedAtDesc",
			request: &dao.SearchBeatsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortBeatUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
	}

	database, closer, err := anoveldb.OpenTestDB(&migrations.SQLMigrations)
	require.NoError(t, err)
	defer closer()

	transaction := anoveldb.BeginTestTX(database, fixtures)
	defer anoveldb.RollbackTestTX(transaction)

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			searchBeatsDAO := dao.NewSearchBeats(transaction)

			beat, err := searchBeatsDAO.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, beat)
		})
	}
}
