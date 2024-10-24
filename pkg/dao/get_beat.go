package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	"github.com/a-novel/uservice-story-structure/pkg/entities"
)

type GetBeat interface {
	Exec(ctx context.Context, id uuid.UUID) (*entities.Beat, error)
}

type getBeatImpl struct {
	database bun.IDB
}

func (dao *getBeatImpl) Exec(ctx context.Context, id uuid.UUID) (*entities.Beat, error) {
	beat := new(entities.Beat)

	err := dao.database.NewSelect().Model(beat).Where("id = ?", id).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrBeatNotFound
		}

		return nil, fmt.Errorf("exec query: %w", err)
	}

	return beat, nil
}

func NewGetBeat(database bun.IDB) GetBeat {
	return &getBeatImpl{database: database}
}
