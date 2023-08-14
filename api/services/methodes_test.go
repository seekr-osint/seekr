package services

import (
	"encoding/base64"
	"encoding/json"
	"errors"
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

func (mock *MockServer) IsServerRunning() bool {
	return mock.Server != nil
}

func (mock *MockServer) CreateMockServer() error {
	handler := http.NewServeMux()

	for _, endpoint := range mock.Endpoints {
		log.Printf("creating endpoint: %s - %d\n", endpoint.URL, endpoint.StatusCode)
		parsedURL, err := url.Parse(endpoint.URL)
		if err != nil {
			return err
		}

		path := parsedURL.Path
		url := fmt.Sprintf("/%s%s", parsedURL.Hostname(), path)
		//fmt.Printf("adding url %s: %s\n", endpoint.URL, url)

		// Create a new variable to capture the current endpoint value.
		currentEndpoint := endpoint

		handler.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
			// fmt.Printf("got request on %s with status code %d \n", currentEndpoint.URL, currentEndpoint.StatusCode)
			w.WriteHeader(currentEndpoint.StatusCode)
			fmt.Fprint(w, currentEndpoint.ResponseBody)
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

		// existing user
		url, err := service.TestUserServiceData().GetUserHtmlUrl()
		if err != nil {
			return nil, err
		}
		endpoint, err := mockHTTPServer(url)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)

		// non existing userr
		url2, err := service.TestUserServiceData2().GetUserHtmlUrl()
		if err != nil {
			return nil, err
		}
		endpoint2, err := mockHTTPServer(url2)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint2)

	}
	return endpoints, nil
}

func ServicesMockWorker(s <-chan Service, res chan<- bool, wg *sync.WaitGroup, t *testing.T) {
	defer wg.Done()
	for service := range s {
		t.Run(service.Name, func(t *testing.T) {
		})

		dataE := service.TestUserServiceData() // user exists

		dataNE := service.TestUserServiceData2() // ueser Not Exists

		statusE, err := dataE.Service.UserExistsFunc(dataE)
		if err != nil {
			t.Errorf("%v", err)
		}

		statusNE, err := dataNE.Service.UserExistsFunc(dataNE)
		if err != nil {
			t.Errorf("%v", err)
		}
		status := false
		if !statusNE && statusE {
			status = true
			//fmt.Printf("working\n")
		}
		t.Logf("Status (%s): %v\n", service.Name, status)
		res <- status
	}
}

func TestServicesMock(t *testing.T) {
	log.SetOutput(io.Discard)
	endpoints, err := DefaultServices.GetMockResp()
	if err != nil {
		fmt.Print("skipping service")
		//t.Skipf("error making http request: %v", err)

		t.Errorf("error making http request: %v", err)
	}

	mockServer := &MockServer{
		Endpoints: endpoints,
	}

	err = mockServer.CreateMockServer()
	if err != nil {
		t.Errorf("%v", err)
	}
	if mockServer.IsServerRunning() {
		log.Printf("mock running\n")
	} else {
		log.Printf("mock not running\n")
	}
	replacedSrevices := ReplaceDomains(DefaultServices, mockServer.Server.URL)
	workers := 2
	wg := sync.WaitGroup{}
	s := make(chan Service, len(replacedSrevices))
	res := make(chan bool, len(replacedSrevices))

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
		service.Domain = fmt.Sprintf("%s/%s", mockDomain, service.Domain)
		newServices = append(newServices, service)
	}
	return newServices
}

func ReadMockHttp(url string) (*MockServerEndpoint, error) {

	data, err := os.ReadFile(GetFilePath(url))
	if err != nil {
		return nil, ErrReadFile
	}

	var mockData MockServerEndpoint
	err = json.Unmarshal(data, &mockData)
	if err != nil {
		return nil, err
	}

	return &mockData, err
}

func GetFilePath(url string) string {
	encodedURL := base64.URLEncoding.EncodeToString([]byte(url))
	filePath := fmt.Sprintf("mock/%s.json", encodedURL)
	return filePath
}

var (
	ErrReadFile = errors.New("failed to read file")
)

func GetBody(url string) (*MockServerEndpoint, error) {
	fmt.Printf("request: %s\n", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &MockServerEndpoint{
		URL:          url,
		StatusCode:   resp.StatusCode,
		ResponseBody: string(body),
	}, nil

}

func mockHTTPServer(url string) (*MockServerEndpoint, error) {
	mockServer, err := ReadMockHttp(url)
	if err == ErrReadFile {
		mockServer, err := GetBody(url)
		if err != nil {
			return nil, err
		}

		err = WriteMock(*mockServer)
		if err != nil {
			return mockServer, err
		}

		return mockServer, nil
	}

	return mockServer, nil
}

//	func TestScan(t *testing.T) {
//		user := User{
//			Username: "9glenda",
//		}
//		result := user.Scan()
//		fmt.Println(result)
//		//for _, i := range result {
//		//	//fmt.Printf("User: %s\nResult %t\n\n", i.User.Username, i.Result)
//		//	fmt.Println(i)
//		//}
//		fmt.Println(result.GetExisting())
//		fmt.Println(result.GetFailed())
//	}
func WriteMock(mockServer MockServerEndpoint) error {
	url := mockServer.URL
	jsonData, err := json.MarshalIndent(mockServer, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data to JSON:", err)

		return err
	}

	// create mock dir
	err = os.MkdirAll("mock", 0755)
	if err != nil {
		fmt.Println("Error creating the 'mock' directory:", err)

		return err
	}

	err = os.WriteFile(GetFilePath(url), jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)

		return err
	}

	fmt.Println("Data successfully written to", GetFilePath(url))
	return nil

}
