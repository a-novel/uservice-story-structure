package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type UpdatePlotPointRequest struct {
	Name   string
	Prompt string
}

type UpdatePlotPoint interface {
	Exec(ctx context.Context, id uuid.UUID, now time.Time, data *UpdatePlotPointRequest) (*entities.PlotPoint, error)
}

type updatePlotPointImpl struct {
	database bun.IDB
}

func (dao *updatePlotPointImpl) Exec(
	ctx context.Context,
	id uuid.UUID,
	now time.Time,
	data *UpdatePlotPointRequest,
) (*entities.PlotPoint, error) {
	model := &entities.PlotPoint{
		ID:        id,
		Name:      data.Name,
		Prompt:    data.Prompt,
		UpdatedAt: &now,
	}

	res, err := dao.database.
		NewUpdate().
		Model(model).
		WherePK().
		Column("name", "prompt", "updated_at").
		Returning("*").
		Exec(ctx)
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

	return model, nil
}

func NewUpdatePlotPoint(database bun.IDB) UpdatePlotPoint {
	return &updatePlotPointImpl{database: database}
}
