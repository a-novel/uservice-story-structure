package dao

import (
	"errors"
)

var ErrBeatNotFound = errors.New("beat not found")

var ErrPlotPointNotFound = errors.New("plot point not found")
