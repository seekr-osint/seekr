package tc

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/seekr-osint/seekr/api/functions"
)

type TestCase[T1 comparable, T2 comparable] struct {
	Input  T1
	Expect T2
}
type TestCases[T1 comparable, T2 comparable] []TestCase[T1, T2]
type Test[T1 comparable, T2 comparable] struct {
	Cases TestCases[T1, T2]
	Func  func(T1) T2
}

// api tests
type Requests = map[string]struct {
	RequestType                string
	Name                       string
	URL                        string
	PostData                   interface{}
	ExpectedResponse           interface{}
	StatusCode                 int
	RequiresInternetConnection bool
	Comment                    string
}
type ApiTest struct {
	Requests Requests
}

func (apiTest ApiTest) RunApiTests(t *testing.T) {
	requests := apiTest.Requests
	for _, name := range functions.SortMapKeys(requests) {
		req := requests[name]
		// Convert post data to JSON if necessary
		postDataJson := []byte{}
		if req.PostData != nil {
			var err error
			postDataJson, err = json.Marshal(req.PostData)
			if err != nil {
				t.Fatalf("[%s] %v", name, err)
			}
		}

		// Send the HTTP request
		httpReq, err := http.NewRequest(req.RequestType, req.URL, bytes.NewBuffer(postDataJson))
		if err != nil {
			t.Fatalf("[%s] %v", name, err)
		}
		httpReq.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(httpReq)
		if err != nil {
			t.Fatalf("[%s] %v", name, err)
		}
		defer resp.Body.Close()

		// Decode the response body
		var respBody interface{}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			t.Fatalf("[%s] %v", name, err)
		}

		if resp.StatusCode != req.StatusCode {
			t.Errorf("[%s] Unexpected Status Code: %d\nExpected %d", name, resp.StatusCode, req.StatusCode)
		}
		// Compare the response body to the expected value
		if !reflect.DeepEqual(respBody, req.ExpectedResponse) {
			t.Errorf("[%s] Unexpected response body: %#v\nExpected %#v", name, respBody, req.ExpectedResponse)
		}
	}
}

// Tests
func NewTest[T1 comparable, T2 comparable](testCaseMap map[T1]T2, function func(T1) T2) Test[T1, T2] {
	var testCases TestCases[T1, T2]
	for input, expect := range testCaseMap {
		testCase := TestCase[T1, T2]{Input: input, Expect: expect}
		testCases = append(testCases, testCase)
	}

	return Test[T1, T2]{Cases: testCases, Func: function}
}

func NewEnumIsValidTest[T1 comparable, T2 comparable](function func(T1) T2, invalidExpect T2, invalidEnum T1, validExpect T2, validEnum ...T1) Test[T1, T2] {
	var testCases TestCases[T1, T2]

	for _, input := range validEnum {
		testCase := TestCase[T1, T2]{Input: input, Expect: validExpect}
		testCases = append(testCases, testCase)
	}

	testCase := TestCase[T1, T2]{Input: invalidEnum, Expect: invalidExpect}
	testCases = append(testCases, testCase)

	return Test[T1, T2]{Cases: testCases, Func: function}
}
