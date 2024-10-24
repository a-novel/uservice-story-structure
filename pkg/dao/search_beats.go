package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type SearchBeatsRequest struct {
	Limit         int
	Offset        int
	Sort          entities.SortBeat
	SortDirection entities.SortDirection
}

type SearchBeats interface {
	Exec(ctx context.Context, request *SearchBeatsRequest) (uuid.UUIDs, error)
}

type searchBeatsImpl struct {
	database bun.IDB
}

func (dao *searchBeatsImpl) Exec(ctx context.Context, request *SearchBeatsRequest) (uuid.UUIDs, error) {
	beats := make([]*entities.Beat, 0)

	query := dao.database.
		NewSelect().
		Model(&beats).
		Column("id").
		Limit(request.Limit).
		Offset(request.Offset)

	// Only apply sorting direction if a sort value is present. Otherwise, ignore it and use default sorting.
	if request.Sort != entities.SortBeatNone {
		direction := lo.Switch[entities.SortDirection, string](request.SortDirection).
			Case(entities.SortDirectionAsc, "ASC").
			Case(entities.SortDirectionDesc, "DESC").
			Default("ASC")

		sort := lo.Switch[entities.SortBeat, string](request.Sort).
			Case(entities.SortBeatName, "beats.name").
			Case(entities.SortBeatCreatedAt, "beats.created_at").
			Case(entities.SortBeatUpdatedAt, "beats.updated_at").
			Default("beats.name")

		query = query.Order(sort + " " + direction)
	} else {
		query = query.Order("beats.name ASC")
	}

	err := query.Scan(ctx, &beats)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	ids := lo.Map(beats, func(item *entities.Beat, _ int) uuid.UUID {
		return item.ID
	})

	return ids, nil
}

func NewSearchBeats(database bun.IDB) SearchBeats {
	return &searchBeatsImpl{database: database}
}
