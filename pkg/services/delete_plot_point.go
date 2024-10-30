package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
)

var (
	ErrInvalidDeletePlotPointRequest = errors.New("invalid delete plot point request")
	ErrDeletePlotPoint               = errors.New("delete plot point")
)

var deletePlotPointValidate = validator.New(validator.WithRequiredStructEnabled())

type DeletePlotPointRequest struct {
	ID        string `validate:"required,len=36"`
	CreatorID string `validate:"omitempty,min=1,max=128"`
}

type DeletePlotPoint interface {
	Exec(ctx context.Context, data *DeletePlotPointRequest) error
}

type deletePlotPointImpl struct {
	dao dao.DeletePlotPoint
}

func (service *deletePlotPointImpl) Exec(ctx context.Context, data *DeletePlotPointRequest) error {
	if err := deletePlotPointValidate.Struct(data); err != nil {
		return errors.Join(ErrInvalidDeletePlotPointRequest, err)
	}

	plotPointID, err := uuid.Parse(data.ID)
	if err != nil {
		return errors.Join(ErrInvalidDeletePlotPointRequest, fmt.Errorf("uuid value: '%s': %w", data.ID, err))
	}

	_, err = service.dao.Exec(ctx, plotPointID, data.CreatorID)
	if err != nil {
		return errors.Join(ErrDeletePlotPoint, err)
	}

	return nil
}

func NewDeletePlotPoint(dao dao.DeletePlotPoint) DeletePlotPoint {
	return &deletePlotPointImpl{dao: dao}
}
