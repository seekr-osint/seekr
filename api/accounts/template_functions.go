package accounts

import (
	"fmt"
	"net/http"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"status":       StatusCode,
		"statusSimple": StatusCodeSimple,
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
