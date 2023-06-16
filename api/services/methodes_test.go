package services

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sync"
	"testing"

	"github.com/seekr-osint/seekr/api/tc"
)

type StringAndErr struct {
	Error  error
	String string
}

type MockServerEndpoint struct {
	URL          string `json:"url"`
	StatusCode   int    `json:"status_code"`
	ResponseBody string `json:"response_body"`
}

type MockServer struct {
	Endpoints []*MockServerEndpoint
	Server    *httptest.Server
}

func (mock *MockServer) CreateMockServer() error {
	handler := http.NewServeMux()

	for _, endpoint := range mock.Endpoints {
		parsedURL, err := url.Parse(endpoint.URL)
		if err != nil {
			return err
		}

		path := parsedURL.Path

		fmt.Printf("adding url %s: %s", endpoint.URL, path)
		handler.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(endpoint.StatusCode)
			fmt.Fprint(w, endpoint.ResponseBody)
		})
	}

	mock.Server = httptest.NewServer(handler)
	return nil
}

func GetUserHtmlUrlHandler(url string) string {
	data := UserServiceDataToCheck{
		User: User{
			Username: "greg",
		},
		Service: Service{
			Name:                "TestService",
			UserHtmlUrlTemplate: url,
			Domain:              "mockurl.com",
		},
	}
	b, err := data.GetUserHtmlUrl()
	if err != nil {
		return ""
	}
	return b
}

func TestGetUserHtmUrl(t *testing.T) {
	log.SetOutput(io.Discard)
	test := tc.Test[string, string]{
		Cases: tc.TestCases[string, string]{
			{
				Input:  "{{.Domain}}/{{.Username}}",
				Expect: "https://mockurl.com/greg",
				Title:  "Simple username in Url",
			},
		},
		Func: GetUserHtmlUrlHandler,
	}

	test.TcTestHandler(t)
	err := os.WriteFile("doc.md", []byte(fmt.Sprintf("# Service url templating\n\n%s", test.TcTestToMarkdown())), 0644)
	if err != nil {
		panic("Error writing to file:")
	}
}

func (services Services) GetMockResp() ([]*MockServerEndpoint, error) {
	endpoints := []*MockServerEndpoint{}
	for _, service := range services {
		url, err := service.TestUserServiceData().GetUserHtmlUrl()
		if err != nil {
			return nil, err
		}
		endpoint, err := mockHTTPServer(url)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)

	}
	return endpoints, nil
}

func ServicesMockWorker(s <-chan Service, res chan<- bool, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	for service := range s {
		t.Run(service.Name, func(t *testing.T) {
		})
		data := service.TestUserServiceData()
		t.Logf("Checking service %s: %s\n", data.Service.Name, data.Service.Domain)
		status, err := data.StatusCodeUserExistsFunc()
		if err != nil {
			t.Errorf("%v", err)
		}
		t.Logf("Status: %v\n", status)
		res <- status
	}
}

func TestServicesMock(t *testing.T) {
	log.SetOutput(io.Discard)
	endpoints, err := DefaultServices.GetMockResp()
	if err != nil {
		t.Skipf("error making http request: %v", err)
	}

	mockServer := &MockServer{
		Endpoints: endpoints,
	}

	err = mockServer.CreateMockServer()
	if err != nil {
		t.Errorf("%v", err)
	}
	replacedSrevices := ReplaceDomains(DefaultServices, mockServer.Server.URL)
	workers := 2
	wg := sync.WaitGroup{}
	s := make(chan Service, workers)
	res := make(chan bool, workers)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go ServicesMockWorker(s, res, &wg, t)
		t.Logf("Added worker: %d", i)
	}
	for _, service := range replacedSrevices {
		s <- service
	}
	close(s)
	wg.Wait()
	for i := 0; i < len(DefaultServices); i++ {
		status := <-res
		t.Logf("Recived: %v\n", status)
		if !status {
			t.Errorf("err status")
		}
	}

}

func ReplaceDomains(services Services, mockDomain string) Services {
	newServices := []Service{}
	for _, service := range services {
		service.Domain = mockDomain
		newServices = append(newServices, service)
	}
	return newServices
}

func mockHTTPServer(url string) (*MockServerEndpoint, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	mockServer := &MockServerEndpoint{
		URL:          url,
		StatusCode:   resp.StatusCode,
		ResponseBody: string(body),
	}

	return mockServer, nil
}
