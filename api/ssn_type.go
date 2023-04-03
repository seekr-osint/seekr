package api

import (
	"regexp"
)

type SSN string

func (ssn SSN) IsValid() bool {
	if ssn == "" {
		return true
	}
	// regular expression for SSN pattern
	pattern := `^\d{3}-?\d{2}-?\d{4}$`
	match, _ := regexp.MatchString(pattern, string(ssn))
	return match
}
