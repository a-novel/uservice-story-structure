package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type BatchDeletePlotPoints interface {
	Exec(ctx context.Context, ids uuid.UUIDs) ([]*entities.PlotPoint, error)
}

type batchDeletePlotPointsImpl struct {
	database bun.IDB
}

func (dao *batchDeletePlotPointsImpl) Exec(ctx context.Context, ids uuid.UUIDs) ([]*entities.PlotPoint, error) {
	beats := make([]*entities.PlotPoint, 0)

	_, err := dao.database.NewDelete().Model(&beats).Where("id IN (?)", bun.In(ids)).Returning("*").Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return beats, nil
}

func NewBatchDeletePlotPoints(database bun.IDB) BatchDeletePlotPoints {
	return &batchDeletePlotPointsImpl{database: database}
}
