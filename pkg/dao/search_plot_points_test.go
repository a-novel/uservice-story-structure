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

func TestSearchPlotPoints(t *testing.T) {
	fixtures := []interface{}{
		// Order by name: Plot Point 1, Plot Point 2, Plot Point 3
		// Order by created_at: Plot Point 3, Plot Point 2, Plot Point 1
		// Order by updated_at: Plot Point 3, Plot Point 1, Plot Point 2
		// Insertion order: Plot Point 2, Plot Point 1, Plot Point 3

		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			CreatorID: "creator_id_1",
			Name:      "Plot Point 2",
			Prompt:    "Prompt 2",
			CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			CreatorID: "creator_id_2",
			Name:      "Plot Point 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			CreatorID: "creator_id_3",
			Name:      "Plot Point 3",
			Prompt:    "Prompt 3",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		request *dao.SearchPlotPointsRequest

		expect    uuid.UUIDs
		expectErr error
	}{
		// Base.
		{
			name: "OK",
			request: &dao.SearchPlotPointsRequest{
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
			request: &dao.SearchPlotPointsRequest{
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
			request: &dao.SearchPlotPointsRequest{
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
			request: &dao.SearchPlotPointsRequest{
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
			request: &dao.SearchPlotPointsRequest{
				Limit:  3,
				Offset: 10,
			},
			expect: uuid.UUIDs{},
		},

		// Sort: name
		{
			name: "Name",
			request: &dao.SearchPlotPointsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortPlotPointName,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
		{
			name: "NameAsc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortPlotPointName,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},
		{
			name: "NameDesc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortPlotPointName,
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
			request: &dao.SearchPlotPointsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortPlotPointCreatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			name: "CreatedAtAsc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortPlotPointCreatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			name: "CreatedAtDesc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortPlotPointCreatedAt,
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
			request: &dao.SearchPlotPointsRequest{
				Limit:  3,
				Offset: 0,
				Sort:   entities.SortPlotPointUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "UpdatedAtAsc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionAsc,
				Sort:          entities.SortPlotPointUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "UpdatedAtDesc",
			request: &dao.SearchPlotPointsRequest{
				Limit:         3,
				Offset:        0,
				SortDirection: entities.SortDirectionDesc,
				Sort:          entities.SortPlotPointUpdatedAt,
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			},
		},

		// Filter: creator_ids
		{
			name: "Filter/CreatorIDs",
			request: &dao.SearchPlotPointsRequest{
				Limit:      3,
				Offset:     0,
				CreatorIDs: []string{"creator_id_1", "creator_id_2"},
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
		{
			name: "Filter/CreatorIDs/Single",
			request: &dao.SearchPlotPointsRequest{
				Limit:      3,
				Offset:     0,
				CreatorIDs: []string{"creator_id_1"},
			},
			expect: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
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
			searchPlotPointsDAO := dao.NewSearchPlotPoints(transaction)

			beat, err := searchPlotPointsDAO.Exec(context.Background(), testCase.request)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, beat)
		})
	}
}
