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
	ErrInvalidDeleteBeatRequest = errors.New("invalid delete beat request")
	ErrDeleteBeat               = errors.New("delete beat")
)

var deleteBeatValidate = validator.New(validator.WithRequiredStructEnabled())

type DeleteBeatRequest struct {
	ID string `validate:"required,len=36"`
}

type DeleteBeat interface {
	Exec(ctx context.Context, data *DeleteBeatRequest) error
}

type deleteBeatImpl struct {
	dao dao.DeleteBeat
}

func (service *deleteBeatImpl) Exec(ctx context.Context, data *DeleteBeatRequest) error {
	if err := deleteBeatValidate.Struct(data); err != nil {
		return errors.Join(ErrInvalidDeleteBeatRequest, err)
	}

	beatID, err := uuid.Parse(data.ID)
	if err != nil {
		return errors.Join(ErrInvalidDeleteBeatRequest, fmt.Errorf("uuid value: '%s': %w", data.ID, err))
	}

	_, err = service.dao.Exec(ctx, beatID)
	if err != nil {
		return errors.Join(ErrDeleteBeat, err)
	}

	return nil
}

func NewDeleteBeat(dao dao.DeleteBeat) DeleteBeat {
	return &deleteBeatImpl{dao: dao}
}
