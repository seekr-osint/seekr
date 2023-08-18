package services

import (
	"io"
	"net/http"
	"time"
)

func httpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}
func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, body)
}
