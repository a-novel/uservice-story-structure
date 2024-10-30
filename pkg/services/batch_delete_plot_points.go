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
	ErrInvalidBatchDeletePlotPointsRequest = errors.New("invalid batch delete plot points request")
	ErrBatchDeletePlotPoints               = errors.New("batch delete plot points")
)

var batchDeletePlotPointsValidate = validator.New(validator.WithRequiredStructEnabled())

type BatchDeletePlotPointsRequest struct {
	IDs       []string `validate:"required,min=1,max=128,dive,required,len=36"`
	CreatorID string   `validate:"omitempty,min=1,max=128"`
}

type BatchDeletePlotPoints interface {
	Exec(ctx context.Context, data *BatchDeletePlotPointsRequest) error
}

type batchDeletePlotPointsImpl struct {
	dao dao.BatchDeletePlotPoints
}

func (service *batchDeletePlotPointsImpl) Exec(ctx context.Context, data *BatchDeletePlotPointsRequest) error {
	var err error

	if err = batchDeletePlotPointsValidate.Struct(data); err != nil {
		return errors.Join(ErrInvalidBatchDeletePlotPointsRequest, err)
	}

	plotPointIDs := make(uuid.UUIDs, len(data.IDs))
	for i, id := range data.IDs {
		plotPointIDs[i], err = uuid.Parse(id)
		if err != nil {
			return errors.Join(
				ErrInvalidBatchDeletePlotPointsRequest,
				fmt.Errorf("at position %v: '%s': %w", i, id, err),
			)
		}
	}

	_, err = service.dao.Exec(ctx, plotPointIDs, data.CreatorID)
	if err != nil {
		return errors.Join(ErrBatchDeletePlotPoints, err)
	}

	return nil
}

func NewBatchDeletePlotPoints(dao dao.BatchDeletePlotPoints) BatchDeletePlotPoints {
	return &batchDeletePlotPointsImpl{dao: dao}
}
