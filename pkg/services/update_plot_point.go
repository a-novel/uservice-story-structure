package services

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
)

var (
	ErrInvalidUpdatePlotPointRequest = errors.New("invalid update plot point request")
	ErrUpdatePlotPoint               = errors.New("update plot point")
)

var updatePlotPointValidate = validator.New(validator.WithRequiredStructEnabled())

// UpdatePlotPointRequest is the request structure for updating a plot point.
// Note: ensure the constraints on name and prompt matches the ones defined on CreatePlotPointRequest.
type UpdatePlotPointRequest struct {
	ID     string `validate:"required,len=36"`
	Name   string `validate:"required,min=2,max=64"`
	Prompt string `validate:"required,min=2,max=1024"`
}

type UpdatePlotPointResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UpdatePlotPoint interface {
	Exec(ctx context.Context, data *UpdatePlotPointRequest) (*UpdatePlotPointResponse, error)
}

type updatePlotPointImpl struct {
	dao dao.UpdatePlotPoint
}

func (service *updatePlotPointImpl) Exec(
	ctx context.Context,
	data *UpdatePlotPointRequest,
) (*UpdatePlotPointResponse, error) {
	if err := updatePlotPointValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidUpdatePlotPointRequest, err)
	}

	plotPointID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, errors.Join(ErrInvalidUpdatePlotPointRequest, err)
	}

	request := &dao.UpdatePlotPointRequest{
		Name:   data.Name,
		Prompt: data.Prompt,
	}

	plotPoint, err := service.dao.Exec(ctx, plotPointID, time.Now(), request)
	if err != nil {
		return nil, errors.Join(ErrUpdatePlotPoint, err)
	}

	return &UpdatePlotPointResponse{
		ID:        plotPoint.ID.String(),
		Name:      plotPoint.Name,
		Prompt:    plotPoint.Prompt,
		CreatedAt: plotPoint.CreatedAt,
		UpdatedAt: plotPoint.UpdatedAt,
	}, nil
}

func NewUpdatePlotPoint(dao dao.UpdatePlotPoint) UpdatePlotPoint {
	return &updatePlotPointImpl{dao: dao}
}
