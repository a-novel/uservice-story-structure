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

func TestBatchDeletePlotPoints(t *testing.T) {
	fixtures := []interface{}{
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Name:      "Plot Point 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			Name:      "Plot Point 2",
			Prompt:    "Prompt 2",
			CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
		},
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000003"),
			Name:      "Plot Point 3",
			Prompt:    "Prompt 3",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		ids uuid.UUIDs

		expect    []*entities.PlotPoint
		expectErr error
	}{
		{
			name: "Delete",
			ids: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
			expect: []*entities.PlotPoint{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name:      "Plot Point 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Name:      "Plot Point 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "Delete/IgnoreNotFound",
			ids: uuid.UUIDs{
				uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			},
			expect: []*entities.PlotPoint{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					Name:      "Plot Point 1",
					Prompt:    "Prompt 1",
					CreatedAt: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 3, 2, 0, 0, 0, 0, time.UTC)),
				},
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Name:      "Plot Point 2",
					Prompt:    "Prompt 2",
					CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: lo.ToPtr(time.Date(2021, 4, 2, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name: "Delete/NoneDeleted",
			ids: uuid.UUIDs{
				uuid.MustParse("10000000-0000-0000-0000-000000000001"),
			},
			expect: []*entities.PlotPoint{},
		},
	}

	database, closer, err := anoveldb.OpenTestDB(&migrations.SQLMigrations)
	require.NoError(t, err)
	defer closer()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			transaction := anoveldb.BeginTestTX(database, fixtures)
			defer anoveldb.RollbackTestTX(transaction)

			deletePlotPointDAO := dao.NewBatchDeletePlotPoints(transaction)

			res, err := deletePlotPointDAO.Exec(context.Background(), testCase.ids)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)
		})
	}
}
