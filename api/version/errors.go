package version

import (
	"errors"
)

var (
	ErrInvalidVersionFormat         = errors.New("invalid schematic version format")
	ErrInvalidMajorVersionComponent = errors.New("invalid major version component")
	ErrInvalidMinorVersionComponent = errors.New("invalid minor version component")
	ErrInvalidPatchVersionComponent = errors.New("invalid patch version component")
)
