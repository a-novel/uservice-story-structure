package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
)

// Beat is a specific sub-instruction used to shape a plot point. Mixing different beats together gives the flavor
// to a particular generation.
type Beat struct {
	bun.BaseModel `bun:"table:beats,alias:beats"`

	// CreatorID is the unique identifier of the user who created the beat.
	CreatorID string `bun:"creator_id"`

	// ID is the unique identifier of the beat.
	ID uuid.UUID `bun:"id,pk,type:uuid"`
	// Name is the human-readable name of the beat.
	Name string `bun:"name"`
	// Prompt used to generate a plot point with this beat.
	Prompt string `bun:"prompt"`

	CreatedAt time.Time  `bun:"created_at"`
	UpdatedAt *time.Time `bun:"updated_at"`
}

type SortBeat string

const (
	SortBeatNone      SortBeat = ""
	SortBeatName      SortBeat = "name"
	SortBeatCreatedAt SortBeat = "created_at"
	SortBeatUpdatedAt SortBeat = "updated_at"
)

var SortBeatConverter = grpc.NewProtoConverter(
	grpc.ProtoMapper[storystructurev1.SortBeats, SortBeat]{
		storystructurev1.SortBeats_SORT_BEATS_BY_NAME:       SortBeatName,
		storystructurev1.SortBeats_SORT_BEATS_BY_CREATED_AT: SortBeatCreatedAt,
		storystructurev1.SortBeats_SORT_BEATS_BY_UPDATED_AT: SortBeatUpdatedAt,
	},
	storystructurev1.SortBeats_SORT_BEATS_UNSPECIFIED,
	SortBeatNone,
)
