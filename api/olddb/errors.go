package db

import "errors"

var (
	ErrEmptyDBPath    = errors.New("empty database path")
	ErrPersonNotExist = errors.New("person does not exist")
)
