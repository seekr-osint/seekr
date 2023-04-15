package api

import (
	"net/http"
)

var (
	ErrPersonNotFound = APIError{
		Message: "Person Not Found",
		Status:  http.StatusNotFound,
	}
)
