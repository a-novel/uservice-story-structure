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
	ErrInvalidBatchDeleteBeatsRequest = errors.New("invalid batch delete beats request")
	ErrBatchDeleteBeats               = errors.New("batch delete beats")
)

var batchDeleteBeatsValidate = validator.New(validator.WithRequiredStructEnabled())

type BatchDeleteBeatsRequest struct {
	IDs       []string `validate:"required,min=1,max=128,dive,required,len=36"`
	CreatorID string   `validate:"omitempty,min=1,max=128"`
}

type BatchDeleteBeats interface {
	Exec(ctx context.Context, data *BatchDeleteBeatsRequest) error
}

type batchDeleteBeatsImpl struct {
	dao dao.BatchDeleteBeats
}

func (service *batchDeleteBeatsImpl) Exec(ctx context.Context, data *BatchDeleteBeatsRequest) error {
	var err error

	if err = batchDeleteBeatsValidate.Struct(data); err != nil {
		return errors.Join(ErrInvalidBatchDeleteBeatsRequest, err)
	}

	beatIDs := make(uuid.UUIDs, len(data.IDs))
	for i, id := range data.IDs {
		beatIDs[i], err = uuid.Parse(id)
		if err != nil {
			return errors.Join(ErrInvalidBatchDeleteBeatsRequest, fmt.Errorf("at position %v: '%s': %w", i, id, err))
		}
	}

	_, err = service.dao.Exec(ctx, beatIDs, data.CreatorID)
	if err != nil {
		return errors.Join(ErrBatchDeleteBeats, err)
	}

	return nil
}

func NewBatchDeleteBeats(dao dao.BatchDeleteBeats) BatchDeleteBeats {
	return &batchDeleteBeatsImpl{dao: dao}
}
