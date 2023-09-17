package accounts

import (
	"fmt"
	"net/url"
	"strings"
)

// if a protocol is specified change it
// if no protocol is specified and the url already has a protocol specified don't change it
// else set it to https
func SetProtocolURL(rawURL, protocol string) (string, error) {
	if rawURL == "" {
		return "", fmt.Errorf("invalid url")
	}
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if protocol != "" {
		parsedURL.Scheme = protocol
	} else if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}

	return parsedURL.String(), nil
}

// Used to parse the string returned from the UserExistsCheck.
func ParseCheckResult(res string) (bool, bool, error) {
	switch {
	case res == "true":
		return true, false, nil
	case res == "false":
		return false, false, nil
	case strings.HasPrefix(res, "error:"):
		errorMessage := strings.TrimSpace(strings.TrimPrefix(res, "error:"))
		if errorMessage != "" {
			return false, false, fmt.Errorf("%s", errorMessage)
		} else {
			return false, false, fmt.Errorf("error message missing")
		}
	case res == "":
		return false, false, fmt.Errorf("empty result")
	case res == "rate limited":
		return false, true, fmt.Errorf("rate limited")
	default:
		return false, false, fmt.Errorf("unknown check result: %s", res)
	}
}
