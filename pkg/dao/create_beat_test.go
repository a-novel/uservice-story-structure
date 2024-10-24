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

func TestCreateBeat(t *testing.T) {
	fixtures := []interface{}{
		&entities.Beat{
			ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			Name:      "Beat 1",
			Prompt:    "Prompt 1",
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: lo.ToPtr(time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC)),
		},
	}

	testCases := []struct {
		name string

		id   uuid.UUID
		now  time.Time
		data *dao.CreateBeatRequest

		expect    *entities.Beat
		expectErr error
	}{
		{
			name: "Create",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			now:  time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			data: &dao.CreateBeatRequest{
				Name:   "Beat 2",
				Prompt: "Prompt 2",
			},
			expect: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				Name:      "Beat 2",
				Prompt:    "Prompt 2",
				CreatedAt: time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: nil,
			},
		},
		{
			name: "Create/NameAndPromptNotUnique",
			id:   uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			now:  time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC),
			data: &dao.CreateBeatRequest{
				Name:   "Beat 1",
				Prompt: "Prompt 1",
			},
			expect: &entities.Beat{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				Name:      "Beat 1",
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

			createBeatDAO := dao.NewCreateBeat(transaction)

			beat, err := createBeatDAO.Exec(context.Background(), testCase.id, testCase.now, testCase.data)

			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, beat)
		})
	}
}
