package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type DeletePlotPoint interface {
	Exec(ctx context.Context, id uuid.UUID, creatorID string) (*entities.PlotPoint, error)
}

type deletePlotPointImpl struct {
	database bun.IDB
}

func (dao *deletePlotPointImpl) Exec(
	ctx context.Context, id uuid.UUID, creatorID string,
) (*entities.PlotPoint, error) {
	plotPoint := &entities.PlotPoint{ID: id}

	query := dao.database.NewDelete().Model(plotPoint).WherePK().Returning("*")

	if creatorID != "" {
		query.Where("creator_id = ?", creatorID)
	}

	res, err := query.Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("get rows affected: %w", err)
	}
	if rows == 0 {
		return nil, ErrPlotPointNotFound
	}

	return plotPoint, nil
}

func NewDeletePlotPoint(database bun.IDB) DeletePlotPoint {
	return &deletePlotPointImpl{database: database}
}
