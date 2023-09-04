package accounts

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"status":                 StatusCode,
		"statusSimple":           StatusCodeSimple,
		"rawStatus":              RawStatusCode,
		"rawBody":                RawBody,
		"bodyPatternMatch":       BodyPatternMatch,
		"bodyPatternMatchSimple": BodyPatternMatchSimple,
	}
}

func StatusCodeSimple(url string, expectedStatusCode int) string {
	return StatusCode(url, expectedStatusCode, -1)
}

func StatusCode(url string, expectedStatusCode, rateLimitStatusCode int) string {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}
	defer response.Body.Close()

	if response.StatusCode == expectedStatusCode {
		return "true"
	} else if rateLimitStatusCode != -1 && response.StatusCode == rateLimitStatusCode { // -1 check unecessary
		return "false"
	} else {
		return "false"
	}
}
func RawStatusCode(url string) string {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}
	defer response.Body.Close()

	return fmt.Sprintf("%d", response.StatusCode)
}

// TODO error handeling
// err => ""
func RawBody(url string) string {
	response, err := http.Get(url)
	if err != nil {
		// return fmt.Sprintf("error: %s", err)
		return ""
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		// return fmt.Sprintf("error: %s", err)
		return ""
	}
	return string(body)
	// return fmt.Sprintf("%d", response.StatusCode)
}

// err => "error: %s"
func BodyPatternMatch(url, pattern string, cnt int) string {
	if cnt <= 0 {
		return fmt.Sprintf("error: cnt too small: %d", cnt)
	}
	response, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Sprintf("error: %s", err)
	}

	if strings.Count(string(body), pattern) >= cnt {
		return "true"
	} else {
		return "false"
	}
}

func BodyPatternMatchSimple(url, pattern string) string {
	return BodyPatternMatch(url, pattern, 1)
}
