package services

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"

	"github.com/a-novel/golib/database"

	"github.com/a-novel/uservice-story-structure/pkg/dao"
	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

var (
	ErrInvalidSearchBeatsRequest = errors.New("invalid search beats request")
	ErrSearchBeats               = errors.New("search beats")
)

var searchBeatsValidate = validator.New(validator.WithRequiredStructEnabled())

func init() {
	database.RegisterSortDirection(searchBeatsValidate)
	searchBeatsValidate.RegisterCustomTypeFunc(entities.RegisterSortBeat, entities.SortBeat(""))
}

type SearchBeatsRequest struct {
	Limit         int                    `validate:"required,min=1,max=128"`
	Offset        int                    `validate:"omitempty,min=0"`
	Sort          entities.SortBeat      `validate:"omitempty,oneof=name created_at updated_at"`
	SortDirection database.SortDirection `validate:"omitempty,oneof=asc desc"`
	CreatorIDs    []string               `validate:"omitempty,dive,min=1,max=128"`
}

type SearchBeatsResponse struct {
	IDs []string
}

type SearchBeats interface {
	Exec(ctx context.Context, data *SearchBeatsRequest) (*SearchBeatsResponse, error)
}

type searchBeatsImpl struct {
	dao dao.SearchBeats
}

func (service *searchBeatsImpl) Exec(ctx context.Context, data *SearchBeatsRequest) (*SearchBeatsResponse, error) {
	if err := searchBeatsValidate.Struct(data); err != nil {
		return nil, errors.Join(ErrInvalidSearchBeatsRequest, err)
	}

	ids, err := service.dao.Exec(ctx, &dao.SearchBeatsRequest{
		Limit:         data.Limit,
		Offset:        data.Offset,
		Sort:          data.Sort,
		SortDirection: data.SortDirection,
		CreatorIDs:    data.CreatorIDs,
	})
	if err != nil {
		return nil, errors.Join(ErrSearchBeats, err)
	}

	return &SearchBeatsResponse{IDs: ids.Strings()}, nil
}

func NewSearchBeats(dao dao.SearchBeats) SearchBeats {
	return &searchBeatsImpl{dao: dao}
}
