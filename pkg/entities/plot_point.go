package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

// PlotPoint is the main instruction used to generate a plot point. A plot point is shaped by multiple beats.
type PlotPoint struct {
	bun.BaseModel `bun:"table:plot_points,alias:plot_points"`

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
