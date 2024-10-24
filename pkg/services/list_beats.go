package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

var (
	ErrInvalidListBeatsRequest = errors.New("invalid list beats request")
	ErrListBeats               = errors.New("list beats")
)

var listBeatsValidate = validator.New(validator.WithRequiredStructEnabled())

type ListBeatsRequest struct {
	IDs []string `validate:"required,min=1,max=128,dive,required,len=36"`
}

type ListBeatsResponseBeat struct {
	ID        string
	Name      string
	Prompt    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type ListBeatsResponse struct {
	Beats []*ListBeatsResponseBeat
}

type ListBeats interface {
	Exec(ctx context.Context, data *ListBeatsRequest) (*ListBeatsResponse, error)
}

type listBeatsImpl struct {
	dao dao.ListBeats
}

func (service *listBeatsImpl) Exec(ctx context.Context, data *ListBeatsRequest) (*ListBeatsResponse, error) {
	var err error

	if err = listBeatsValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidListBeatsRequest, err)
	}

	beatIDs := make(uuid.UUIDs, len(data.IDs))
	for i, id := range data.IDs {
		beatIDs[i], err = uuid.Parse(id)
		if err != nil {
			return nil, errors.Join(ErrInvalidListBeatsRequest, fmt.Errorf("at position %v: '%s': %w", i, id, err))
		}
	}

	beats, err := service.dao.Exec(ctx, beatIDs)
	if err != nil {
		return nil, errors.Join(ErrListBeats, err)
	}

	response := &ListBeatsResponse{
		Beats: lo.Map(beats, func(item *entities.Beat, _ int) *ListBeatsResponseBeat {
			return &ListBeatsResponseBeat{
				ID:        item.ID.String(),
				Name:      item.Name,
				Prompt:    item.Prompt,
				CreatedAt: item.CreatedAt,
				UpdatedAt: item.UpdatedAt,
			}
		}),
	}

	return response, nil
}

func NewListBeats(dao dao.ListBeats) ListBeats {
	return &listBeatsImpl{dao: dao}
}
