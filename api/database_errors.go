package api

import (
	"net/http"

	"github.com/seekr-osint/seekr/api/errortypes"
)

var (
	ErrPersonNotFound = errortypes.APIError{
		Message: "Person Not Found",
		Status:  http.StatusNotFound,
	}
)
