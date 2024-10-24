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
	ErrInvalidUpdateBeatRequest = errors.New("invalid update beat request")
	ErrUpdateBeat               = errors.New("update beat")
)

var updateBeatValidate = validator.New(validator.WithRequiredStructEnabled())

// UpdateBeatRequest is the request structure for updating a beat.
// Note: ensure the constraints on name and prompt matches the ones defined on CreateBeatRequest.
type UpdateBeatRequest struct {
	ID     string `validate:"required,len=36"`
	Name   string `validate:"required,min=2,max=64"`
	Prompt string `validate:"required,min=2,max=1024"`
}

type UpdateBeatResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UpdateBeat interface {
	Exec(ctx context.Context, data *UpdateBeatRequest) (*UpdateBeatResponse, error)
}

type updateBeatImpl struct {
	dao dao.UpdateBeat
}

func (service *updateBeatImpl) Exec(ctx context.Context, data *UpdateBeatRequest) (*UpdateBeatResponse, error) {
	if err := updateBeatValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidUpdateBeatRequest, err)
	}

	beatID, err := uuid.Parse(data.ID)
	if err != nil {
		return nil, errors.Join(ErrInvalidUpdateBeatRequest, err)
	}

	request := &dao.UpdateBeatRequest{
		Name:   data.Name,
		Prompt: data.Prompt,
	}

	beat, err := service.dao.Exec(ctx, beatID, time.Now(), request)
	if err != nil {
		return nil, errors.Join(ErrUpdateBeat, err)
	}

	return &UpdateBeatResponse{
		ID:        beat.ID.String(),
		Name:      beat.Name,
		Prompt:    beat.Prompt,
		CreatedAt: beat.CreatedAt,
		UpdatedAt: beat.UpdatedAt,
	}, nil
}

func NewUpdateBeat(dao dao.UpdateBeat) UpdateBeat {
	return &updateBeatImpl{dao: dao}
}
