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
	ErrInvalidCreatePlotPointRequest = errors.New("invalid create plot point request")
	ErrCreatePlotPoint               = errors.New("create plot point")
)

var createPlotPointValidate = validator.New(validator.WithRequiredStructEnabled())

// CreatePlotPointRequest is the request structure for updating a plot point.
// Note: ensure the constraints on name and prompt matches the ones defined on UpdatePlotPointRequest.
type CreatePlotPointRequest struct {
	Name      string `validate:"required,min=2,max=64"`
	Prompt    string `validate:"required,min=2,max=1024"`
	CreatorID string `validate:"required,min=1,max=128"`
}

type CreatePlotPointResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatorID string
	CreatedAt time.Time
}

type CreatePlotPoint interface {
	Exec(ctx context.Context, data *CreatePlotPointRequest) (*CreatePlotPointResponse, error)
}

type createPlotPointImpl struct {
	dao dao.CreatePlotPoint
}

func (service *createPlotPointImpl) Exec(
	ctx context.Context,
	data *CreatePlotPointRequest,
) (*CreatePlotPointResponse, error) {
	if err := createPlotPointValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidCreatePlotPointRequest, err)
	}

	request := &dao.CreatePlotPointRequest{
		Name:      data.Name,
		Prompt:    data.Prompt,
		CreatorID: data.CreatorID,
	}

	plotPoint, err := service.dao.Exec(ctx, uuid.New(), time.Now(), request)
	if err != nil {
		return nil, errors.Join(ErrCreatePlotPoint, err)
	}

	return &CreatePlotPointResponse{
		ID:        plotPoint.ID.String(),
		CreatorID: plotPoint.CreatorID,
		Name:      plotPoint.Name,
		Prompt:    plotPoint.Prompt,
		CreatedAt: plotPoint.CreatedAt,
	}, nil
}

func NewCreatePlotPoint(dao dao.CreatePlotPoint) CreatePlotPoint {
	return &createPlotPointImpl{dao: dao}
}
