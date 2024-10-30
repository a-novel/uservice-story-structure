package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type BatchDeleteBeats interface {
	Exec(ctx context.Context, ids uuid.UUIDs, creatorID string) ([]*entities.Beat, error)
}

type batchDeleteBeatsImpl struct {
	database bun.IDB
}

func (dao *batchDeleteBeatsImpl) Exec(
	ctx context.Context, ids uuid.UUIDs, creatorID string,
) ([]*entities.Beat, error) {
	beats := make([]*entities.Beat, 0)

	query := dao.database.NewDelete().
		Model(&beats).
		Where("id IN (?)", bun.In(ids)).
		Returning("*")

	if creatorID != "" {
		query = query.Where("creator_id = ?", creatorID)
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return beats, nil
}

func NewBatchDeleteBeats(database bun.IDB) BatchDeleteBeats {
	return &batchDeleteBeatsImpl{database: database}
}
