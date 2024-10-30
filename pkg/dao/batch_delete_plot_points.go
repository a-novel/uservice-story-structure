package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type BatchDeletePlotPoints interface {
	Exec(ctx context.Context, ids uuid.UUIDs, creatorID string) ([]*entities.PlotPoint, error)
}

type batchDeletePlotPointsImpl struct {
	database bun.IDB
}

func (dao *batchDeletePlotPointsImpl) Exec(
	ctx context.Context, ids uuid.UUIDs, creatorID string,
) ([]*entities.PlotPoint, error) {
	plotPoints := make([]*entities.PlotPoint, 0)

	query := dao.database.NewDelete().
		Model(&plotPoints).
		Where("id IN (?)", bun.In(ids)).
		Returning("*")

	if creatorID != "" {
		query = query.Where("creator_id = ?", creatorID)
	}

	_, err := query.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return plotPoints, nil
}

func NewBatchDeletePlotPoints(database bun.IDB) BatchDeletePlotPoints {
	return &batchDeletePlotPointsImpl{database: database}
}
