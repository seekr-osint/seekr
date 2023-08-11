package client

import (
	"github.com/seekr-osint/seekr/api/tc"
	"testing"
)

var client = &Client{
	Ip:   "localhost",
	Port: 8080,
}

func ApiCall(endpoint string) string {
	return client.ApiCall(endpoint)
}

func TestApiCall(t *testing.T) {
	testCases := map[string]string{
		"ping":        "http://localhost:8080/api/ping",
		"/ping":       "http://localhost:8080/api/ping",
		"ping/":       "http://localhost:8080/api/ping",
		"/ping/":      "http://localhost:8080/api/ping",
		"/person/id/": "http://localhost:8080/api/person/id",
	}

	test := tc.NewTest(testCases, ApiCall)

	test.TcTestHandler(t)
}
