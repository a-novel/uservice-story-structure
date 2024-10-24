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
	ErrInvalidGetBeatRequest = errors.New("invalid get beat request")
	ErrGetBeat               = errors.New("get beat")
)

var getBeatValidate = validator.New(validator.WithRequiredStructEnabled())

type GetBeatRequest struct {
	ID string `validate:"required,len=36"`
}

type GetBeatResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type GetBeat interface {
	Exec(ctx context.Context, data *GetBeatRequest) (*GetBeatResponse, error)
}

type getBeatImpl struct {
	dao dao.GetBeat
}

func (service *getBeatImpl) Exec(ctx context.Context, data *GetBeatRequest) (*GetBeatResponse, error) {
	if err := getBeatValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidGetBeatRequest, err)
	}

	beatID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, errors.Join(ErrInvalidGetBeatRequest, fmt.Errorf("uuid value: '%s': %w", data.ID, err))
	}

	beat, err := service.dao.Exec(ctx, beatID)
	if err != nil {
		return nil, errors.Join(ErrGetBeat, err)
	}

	return &GetBeatResponse{
		ID:        beat.ID.String(),
		Name:      beat.Name,
		Prompt:    beat.Prompt,
		CreatedAt: beat.CreatedAt,
		UpdatedAt: beat.UpdatedAt,
	}, nil
}

func NewGetBeat(dao dao.GetBeat) GetBeat {
	return &getBeatImpl{dao: dao}
}
