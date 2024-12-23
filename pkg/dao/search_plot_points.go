package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/uptrace/bun"

	"github.com/a-novel/golib/database"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type SearchPlotPointsRequest struct {
	Limit         int
	Offset        int
	Sort          entities.SortPlotPoint
	SortDirection database.SortDirection
	CreatorIDs    []string
}

type SearchPlotPoints interface {
	Exec(ctx context.Context, request *SearchPlotPointsRequest) (uuid.UUIDs, error)
}

type searchPlotPointsImpl struct {
	database bun.IDB
}

func (dao *searchPlotPointsImpl) Exec(ctx context.Context, request *SearchPlotPointsRequest) (uuid.UUIDs, error) {
	plotPoints := make([]*entities.PlotPoint, 0)

	query := dao.database.
		NewSelect().
		Model(&plotPoints).
		Column("id").
		Limit(request.Limit).
		Offset(request.Offset)

	// Only apply sorting direction if a sort value is present. Otherwise, ignore it and use default sorting.
	if request.Sort != entities.SortPlotPointNone {
		direction := lo.Switch[database.SortDirection, string](request.SortDirection).
			Case(database.SortDirectionAsc, "ASC").
			Case(database.SortDirectionDesc, "DESC").
			Default("ASC")

		sort := lo.Switch[entities.SortPlotPoint, string](request.Sort).
			Case(entities.SortPlotPointName, "plot_points.name").
			Case(entities.SortPlotPointCreatedAt, "plot_points.created_at").
			Case(entities.SortPlotPointUpdatedAt, "plot_points.updated_at").
			Default("plot_points.name")

		query = query.Order(sort + " " + direction)
	} else {
		query = query.Order("plot_points.name ASC")
	}

	if len(request.CreatorIDs) > 1 {
		query = query.Where("creator_id IN (?)", bun.In(request.CreatorIDs))
	} else if len(request.CreatorIDs) == 1 {
		query = query.Where("creator_id = ?", request.CreatorIDs[0])
	}

	err := query.Scan(ctx, &plotPoints)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	ids := lo.Map(plotPoints, func(item *entities.PlotPoint, _ int) uuid.UUID {
		return item.ID
	})

	return ids, nil
}

func NewSearchPlotPoints(database bun.IDB) SearchPlotPoints {
	return &searchPlotPointsImpl{database: database}
}
