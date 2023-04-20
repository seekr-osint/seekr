package webserver

import (
	"errors"
	"reflect"
)

var (
	ErrUndefinedWebFS = errors.New("error: Undefined WebFS")
)

// Parsing

func (webserver Webserver) Parse() (Webserver, error) {
	return webserver, nil
}

// Validation

func (webserver Webserver) Validate() error {
	if !webserver.Disable {
		// check rather webserver.FileSystem is left default (not defined)
		if reflect.DeepEqual(webserver.FileSystem, reflect.Zero(reflect.TypeOf(webserver.FileSystem)).Interface()) {
			return ErrUndefinedWebFS
		}
	}
	return nil
}
