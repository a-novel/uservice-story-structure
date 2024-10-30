package entities

import "reflect"

type SortDirection string

const (
	SortDirectionNone SortDirection = ""
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func ValidateSortDirection(field reflect.Value) interface{} {
	value, ok := field.Interface().(SortDirection)
	if !ok {
		return nil
	}

	return string(value)
}
