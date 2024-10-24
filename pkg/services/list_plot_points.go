package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

var (
	ErrInvalidListPlotPointsRequest = errors.New("invalid list plot points request")
	ErrListPlotPoints               = errors.New("list plot points")
)

var listPlotPointsValidate = validator.New(validator.WithRequiredStructEnabled())

type ListPlotPointsRequest struct {
	IDs []string `validate:"required,min=1,max=128,dive,required,len=36"`
}

type ListPlotPointsResponsePlotPoint struct {
	ID        string
	Name      string
	Prompt    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type ListPlotPointsResponse struct {
	PlotPoints []*ListPlotPointsResponsePlotPoint
}

type ListPlotPoints interface {
	Exec(ctx context.Context, data *ListPlotPointsRequest) (*ListPlotPointsResponse, error)
}

type listPlotPointsImpl struct {
	dao dao.ListPlotPoints
}

func (service *listPlotPointsImpl) Exec(
	ctx context.Context,
	data *ListPlotPointsRequest,
) (*ListPlotPointsResponse, error) {
	var err error

	if err = listPlotPointsValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidListPlotPointsRequest, err)
	}

	plotPointIDs := make(uuid.UUIDs, len(data.IDs))
	for i, id := range data.IDs {
		plotPointIDs[i], err = uuid.Parse(id)
		if err != nil {
			return nil, errors.Join(ErrInvalidListPlotPointsRequest, fmt.Errorf("at position %v: '%s': %w", i, id, err))
		}
	}

	plotPoints, err := service.dao.Exec(ctx, plotPointIDs)
	if err != nil {
		return nil, errors.Join(ErrListPlotPoints, err)
	}

	response := &ListPlotPointsResponse{
		PlotPoints: lo.Map(plotPoints, func(item *entities.PlotPoint, _ int) *ListPlotPointsResponsePlotPoint {
			return &ListPlotPointsResponsePlotPoint{
				ID:        item.ID.String(),
				Name:      item.Name,
				Prompt:    item.Prompt,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			}
		}),
	}

	return response, nil
}

func NewListPlotPoints(dao dao.ListPlotPoints) ListPlotPoints {
	return &listPlotPointsImpl{dao: dao}
}
