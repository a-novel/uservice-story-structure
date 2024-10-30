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
	ErrInvalidCreateBeatRequest = errors.New("invalid create beat request")
	ErrCreateBeat               = errors.New("create beat")
)

var createBeatValidate = validator.New(validator.WithRequiredStructEnabled())

// CreateBeatRequest is the request structure for updating a beat.
// Note: ensure the constraints on name and prompt matches the ones defined on UpdateBeatRequest.
type CreateBeatRequest struct {
	Name      string `validate:"required,min=2,max=64"`
	Prompt    string `validate:"required,min=2,max=1024"`
	CreatorID string `validate:"required,min=1,max=128"`
}

type CreateBeatResponse struct {
	ID        string
	Name      string
	Prompt    string
	CreatorID string
	CreatedAt time.Time
}

type CreateBeat interface {
	Exec(ctx context.Context, data *CreateBeatRequest) (*CreateBeatResponse, error)
}

type createBeatImpl struct {
	dao dao.CreateBeat
}

func (service *createBeatImpl) Exec(ctx context.Context, data *CreateBeatRequest) (*CreateBeatResponse, error) {
	if err := createBeatValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidCreateBeatRequest, err)
	}

	request := &dao.CreateBeatRequest{
		Name:      data.Name,
		Prompt:    data.Prompt,
		CreatorID: data.CreatorID,
	}

	beat, err := service.dao.Exec(ctx, uuid.New(), time.Now(), request)
	if err != nil {
		return nil, errors.Join(ErrCreateBeat, err)
	}

	return &CreateBeatResponse{
		ID:        beat.ID.String(),
		CreatorID: data.CreatorID,
		Name:      beat.Name,
		Prompt:    beat.Prompt,
		CreatedAt: beat.CreatedAt,
	}, nil
}

func NewCreateBeat(dao dao.CreateBeat) CreateBeat {
	return &createBeatImpl{dao: dao}
}
