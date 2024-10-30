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

func TestCreatePlotPoint(t *testing.T) {
	fixtures := []interface{}{
		&entities.PlotPoint{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Name:      "Plot Point 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		id   uuid.UUID
		now  time.Time
		data *dao.CreatePlotPointRequest

		expect    *entities.PlotPoint
		expectErr error
	}{
		{
			name: "Create",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			now:  time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			data: &dao.CreatePlotPointRequest{
				Name:      "Plot Point 2",
				Prompt:    "Prompt 2",
				CreatorID: "creator_id",
			},
			expect: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				CreatorID: "creator_id",
				Name:      "Plot Point 2",
				Prompt:    "Prompt 2",
				CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: nil,
			},
		},
		{
			name: "Create/NameAndPromptNotUnique",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			now:  time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			data: &dao.CreatePlotPointRequest{
				Name:      "Plot Point 1",
				Prompt:    "Prompt 1",
				CreatorID: "creator_id",
			},
			expect: &entities.PlotPoint{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				CreatorID: "creator_id",
				Name:      "Plot Point 1",
				Prompt:    "Prompt 1",
				CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: nil,
			},
		},
	}

	database, closer, err := anoveldb.OpenTestDB(&migrations.SQLMigrations)
	require.NoError(t, err)
	defer closer()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			transaction := anoveldb.BeginTestTX(database, fixtures)
			defer anoveldb.RollbackTestTX(transaction)

			createPlotPointDAO := dao.NewCreatePlotPoint(transaction)

			result, err := createPlotPointDAO.Exec(context.Background(), testCase.id, testCase.now, testCase.data)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, result)
		})
	}
}
