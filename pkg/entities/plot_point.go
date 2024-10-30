package entities

import (
	"reflect"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
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

func ValidateSortPlotPoint(field reflect.Value) interface{} {
	value, ok := field.Interface().(SortPlotPoint)
	if !ok {
		return nil
	}

	return string(value)
}
