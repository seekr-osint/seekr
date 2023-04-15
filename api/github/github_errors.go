package github

import (
	"errors"
)

var (
	ErrCreatingTmp     = errors.New("failed to create temp dir")
	ErrCalcRateLimit   = errors.New("failed to calculate the rate limit")
	ErrRateLimited     = errors.New("rate limited")
	ErrMissingUsername = errors.New("Missing Username")
	//ErrEmptyEmail = err
)
