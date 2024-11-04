package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"

	storystructurev1 "buf.build/gen/go/a-novel/proto/protocolbuffers/go/storystructure/v1"

	"github.com/a-novel/golib/grpc"
)

// PlotPoint is the main instruction used to generate a plot point. A plot point is shaped by multiple beats.
type PlotPoint struct {
	bun.BaseModel `bun:"table:plot_points,alias:plot_points"`

	// CreatorID is the unique identifier of the user who created the plot point.
	CreatorID string `bun:"creator_id"`

	// ID is the unique identifier of the plot point.
	ID uuid.UUID `bun:"id,pk,type:uuid"`
	// Name is the human-readable name of the plot point.
	Name string `bun:"name"`
	// Prompt is the global instruction, used to generate the plot point.
	Prompt string `bun:"prompt"`

	CreatedAt time.Time  `bun:"created_at"`
	UpdatedAt *time.Time `bun:"updated_at"`
}

type SortPlotPoint string

const (
	SortPlotPointNone      SortPlotPoint = ""
	SortPlotPointName      SortPlotPoint = "name"
	SortPlotPointCreatedAt SortPlotPoint = "created_at"
	SortPlotPointUpdatedAt SortPlotPoint = "updated_at"
)

var SortPlotPointConverter = grpc.NewProtoConverter(
	grpc.ProtoMapper[storystructurev1.SortPlotPoints, SortPlotPoint]{
		storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_NAME:       SortPlotPointName,
		storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_CREATED_AT: SortPlotPointCreatedAt,
		storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_BY_UPDATED_AT: SortPlotPointUpdatedAt,
	},
	storystructurev1.SortPlotPoints_SORT_PLOT_POINTS_UNSPECIFIED,
	SortPlotPointNone,
)
