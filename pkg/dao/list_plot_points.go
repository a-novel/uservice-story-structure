package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type ListPlotPoints interface {
	Exec(ctx context.Context, ids []uuid.UUID) ([]*entities.PlotPoint, error)
}

type listPlotPointsImpl struct {
	database bun.IDB
}

func (dao *listPlotPointsImpl) Exec(ctx context.Context, ids []uuid.UUID) ([]*entities.PlotPoint, error) {
	plotPoints := make([]*entities.PlotPoint, 0)

	err := dao.database.NewSelect().Model(&plotPoints).Where("id IN (?)", bun.In(ids)).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return plotPoints, nil
}

func NewListPlotPoints(database bun.IDB) ListPlotPoints {
	return &listPlotPointsImpl{database: database}
}
