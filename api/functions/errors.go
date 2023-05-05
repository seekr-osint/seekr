package functions

import "errors"

var (
	ErrOnlyStruct = errors.New("only works with structs")
	ErrOnlyMap    = errors.New("only works with maps")

	ErrDifferentTypes = errors.New("error: different types")
)
