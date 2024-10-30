package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type CreateBeatRequest struct {
	Name      string
	Prompt    string
	CreatorID string
}

type CreateBeat interface {
	Exec(ctx context.Context, id uuid.UUID, now time.Time, data *CreateBeatRequest) (*entities.Beat, error)
}

type createBeatImpl struct {
	database bun.IDB
}

func (dao *createBeatImpl) Exec(
	ctx context.Context,
	id uuid.UUID,
	now time.Time,
	data *CreateBeatRequest,
) (*entities.Beat, error) {
	model := &entities.Beat{
		ID:        id,
		CreatorID: data.CreatorID,
		Name:      data.Name,
		Prompt:    data.Prompt,
		CreatedAt: now,
	}

	_, err := dao.database.NewInsert().Model(model).Returning("*").Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	return model, nil
}

func NewCreateBeat(database bun.IDB) CreateBeat {
	return &createBeatImpl{database: database}
}
