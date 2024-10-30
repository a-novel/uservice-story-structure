package dao

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type UpdateBeatRequest struct {
	Name      string
	Prompt    string
	CreatorID string
}

type UpdateBeat interface {
	Exec(ctx context.Context, id uuid.UUID, now time.Time, data *UpdateBeatRequest) (*entities.Beat, error)
}

type updateBeatImpl struct {
	database bun.IDB
}

func (dao *updateBeatImpl) Exec(
	ctx context.Context,
	id uuid.UUID,
	now time.Time,
	data *UpdateBeatRequest,
) (*entities.Beat, error) {
	model := &entities.Beat{
		ID:        id,
		Name:      data.Name,
		Prompt:    data.Prompt,
		UpdatedAt: &now,
	}

	res, err := dao.database.
		NewUpdate().
		Model(model).
		WherePK().
		Where("creator_id = ?", data.CreatorID).
		Column("name", "prompt", "updated_at").
		Returning("*").
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("get rows affected: %w", err)
	}

	if rows == 0 {
		return nil, ErrBeatNotFound
	}

	return model, nil
}

func NewUpdateBeat(database bun.IDB) UpdateBeat {
	return &updateBeatImpl{database: database}
}
