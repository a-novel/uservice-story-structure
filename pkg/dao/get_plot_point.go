package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type GetPlotPoint interface {
	Exec(ctx context.Context, id uuid.UUID) (*entities.PlotPoint, error)
}

type getPlotPointImpl struct {
	database bun.IDB
}

func (dao *getPlotPointImpl) Exec(ctx context.Context, id uuid.UUID) (*entities.PlotPoint, error) {
	plotPoint := new(entities.PlotPoint)

	err := dao.database.NewSelect().Model(plotPoint).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPlotPointNotFound
		}

		return nil, fmt.Errorf("exec query: %w", err)
	}

	return plotPoint, nil
}

func NewGetPlotPoint(database bun.IDB) GetPlotPoint {
	return &getPlotPointImpl{database: database}
}
