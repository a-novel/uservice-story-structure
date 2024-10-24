package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type ListBeats interface {
	Exec(ctx context.Context, ids []uuid.UUID) ([]*entities.Beat, error)
}

type listBeatsImpl struct {
	database bun.IDB
}

func (dao *listBeatsImpl) Exec(ctx context.Context, ids []uuid.UUID) ([]*entities.Beat, error) {
	beats := make([]*entities.Beat, 0)

	err := dao.database.NewSelect().Model(&beats).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return beats, nil
}

func NewListBeats(database bun.IDB) ListBeats {
	return &listBeatsImpl{database: database}
}
