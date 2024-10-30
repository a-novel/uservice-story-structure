package dao

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type DeleteBeat interface {
	Exec(ctx context.Context, id uuid.UUID, creatorID string) (*entities.Beat, error)
}

type deleteBeatImpl struct {
	database bun.IDB
}

func (dao *deleteBeatImpl) Exec(
	ctx context.Context, id uuid.UUID, creatorID string,
) (*entities.Beat, error) {
	beat := &entities.Beat{ID: id}

	query := dao.database.NewDelete().Model(beat).WherePK().Returning("*")

	if creatorID != "" {
		query.Where("creator_id = ?", creatorID)
	}

	res, err := query.Exec(ctx)
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

	return beat, nil
}

func NewDeleteBeat(database bun.IDB) DeleteBeat {
	return &deleteBeatImpl{database: database}
}
