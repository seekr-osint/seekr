package newtype

import "errors"

var (
	ErrTypeMissmatch = errors.New("type missmatch")
	ErrUnknownType = errors.New("unknown type")
)
