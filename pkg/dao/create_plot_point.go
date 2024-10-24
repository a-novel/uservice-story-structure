package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type CreatePlotPointRequest struct {
	Name   string
	Prompt string
}

type CreatePlotPoint interface {
	Exec(ctx context.Context, id uuid.UUID, now time.Time, data *CreatePlotPointRequest) (*entities.PlotPoint, error)
}

type createPlotPointImpl struct {
	database bun.IDB
}

func (dao *createPlotPointImpl) Exec(
	ctx context.Context,
	id uuid.UUID,
	now time.Time,
	data *CreatePlotPointRequest,
) (*entities.PlotPoint, error) {
	model := &entities.PlotPoint{
		ID:        id,
		Name:      data.Name,
		Prompt:    data.Prompt,
		CreatedAt: now,
	}

	_, err := dao.database.NewInsert().Model(model).Returning("*").Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return model, nil
}

func NewCreatePlotPoint(database bun.IDB) CreatePlotPoint {
	return &createPlotPointImpl{database: database}
}
