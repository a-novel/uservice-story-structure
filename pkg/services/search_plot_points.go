package services

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

var (
	ErrInvalidSearchPlotPointsRequest = errors.New("invalid search plot points request")
	ErrSearchPlotPoints               = errors.New("search plot points")
)

var searchPlotPointsValidate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	searchPlotPointsValidate.RegisterCustomTypeFunc(entities.ValidateSortDirection, entities.SortDirection(""))
	searchPlotPointsValidate.RegisterCustomTypeFunc(entities.ValidateSortPlotPoint, entities.SortPlotPoint(""))
}

type SearchPlotPointsRequest struct {
	Limit         int                    `validate:"required,min=1,max=128"`
	Offset        int                    `validate:"omitempty,min=0"`
	Sort          entities.SortPlotPoint `validate:"omitempty,oneof=name created_at updated_at"`
	SortDirection entities.SortDirection `validate:"omitempty,oneof=asc desc"`
	CreatorIDs    []string               `validate:"omitempty,dive,min=1,max=128"`
}

type SearchPlotPointsResponse struct {
	IDs []string
}

type SearchPlotPoints interface {
	Exec(ctx context.Context, data *SearchPlotPointsRequest) (*SearchPlotPointsResponse, error)
}

type searchPlotPointsImpl struct {
	dao dao.SearchPlotPoints
}

func (service *searchPlotPointsImpl) Exec(
	ctx context.Context,
	data *SearchPlotPointsRequest,
) (*SearchPlotPointsResponse, error) {
	if err := searchPlotPointsValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidSearchPlotPointsRequest, err)
	}

	ids, err := service.dao.Exec(ctx, &dao.SearchPlotPointsRequest{
		Limit:         data.Limit,
		Offset:        data.Offset,
		Sort:          data.Sort,
		SortDirection: data.SortDirection,
		CreatorIDs:    data.CreatorIDs,
	})
	if err != nil {
		return nil, errors.Join(ErrSearchPlotPoints, err)
	}

	return &SearchPlotPointsResponse{IDs: ids.Strings()}, nil
}

func NewSearchPlotPoints(dao dao.SearchPlotPoints) SearchPlotPoints {
	return &searchPlotPointsImpl{dao: dao}
}
