package api

import (
	"regexp"
)

func IsGmailAddress(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9._-]+@gmail.com$")
	return pattern.MatchString(email)
}

func IsValidGmailAddress(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.]+@gmail.com$")
	return pattern.MatchString(email)
}

func IsEmailValid(email string) bool {
	// Compile the regular expression pattern
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z]{2,})*$")

	// Check if the email matches the pattern
	return pattern.MatchString(email)
}
