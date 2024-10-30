package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
)

var (
	ErrInvalidGetPlotPointRequest = errors.New("invalid get plot point request")
	ErrGetPlotPoint               = errors.New("get plot point")
)

var getPlotPointValidate = validator.New(validator.WithRequiredStructEnabled())

type GetPlotPointRequest struct {
	ID string `validate:"required,len=36"`
}

type GetPlotPointResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatorID string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type GetPlotPoint interface {
	Exec(ctx context.Context, data *GetPlotPointRequest) (*GetPlotPointResponse, error)
}

type getPlotPointImpl struct {
	dao dao.GetPlotPoint
}

func (service *getPlotPointImpl) Exec(ctx context.Context, data *GetPlotPointRequest) (*GetPlotPointResponse, error) {
	if err := getPlotPointValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidGetPlotPointRequest, err)
	}

	plotPointID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, errors.Join(ErrInvalidGetPlotPointRequest, fmt.Errorf("uuid value: '%s': %w", data.ID, err))
	}

	plotPoint, err := service.dao.Exec(ctx, plotPointID)
	if err != nil {
		return nil, errors.Join(ErrGetPlotPoint, err)
	}

	return &GetPlotPointResponse{
		ID:        plotPoint.ID.String(),
		CreatorID: plotPoint.CreatorID,
		Name:      plotPoint.Name,
		Prompt:    plotPoint.Prompt,
		CreatedAt: plotPoint.CreatedAt,
		UpdatedAt: plotPoint.UpdatedAt,
	}, nil
}

func NewGetPlotPoint(dao dao.GetPlotPoint) GetPlotPoint {
	return &getPlotPointImpl{dao: dao}
}
