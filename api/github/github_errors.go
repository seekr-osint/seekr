package github

import (
	"net/http"

	"github.com/seekr-osint/seekr/api/errortypes"
)

var (
	ErrCreatingTmp = errortypes.APIError{
		Message: "failed to create temp dir",
		Status:  http.StatusInternalServerError,
	}
	ErrCalcRateLimit = errortypes.APIError{
		Message: "failed to calculate the rate limit",
		Status:  http.StatusInternalServerError,
	}
	ErrRateLimited = errortypes.APIError{
		Message: "reate limited",
		Status:  http.StatusUnauthorized,
	}
	ErrMissingUsername = errortypes.APIError{
		Message: "missing Username",
		Status:  http.StatusBadRequest,
	}

	ErrUnmarshaling = errortypes.APIError{
		Message: "missing Unmarshaling",
		Status:  http.StatusBadRequest,
	}
)
