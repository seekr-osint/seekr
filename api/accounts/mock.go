package accounts

import (
	// 	// "fmt"
	// 	"log"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// "testing"

	"github.com/jarcoal/httpmock"
	// "github.com/seekr-osint/seekr/api/tcmultiarg"
)

var (
	ErrReadFile = errors.New("failed to read file")
)

type MockEndpoint struct {
	URL          string `json:"url"`
	ResponseBody string `json:"resp_body"`
	StatusCode   int    `json:"status_code"`
}
type MockEndpoints []MockEndpoint

// var endpoints MockEndpoints

func (endpoints MockEndpoints) CreateMockEndoints() {
	httpmock.Activate()
	// defer httpmock.DeactivateAndReset()

	for _, e := range endpoints {
		endpoint := e
		httpmock.RegisterResponder("GET", endpoint.URL, httpmock.NewStringResponder(endpoint.StatusCode, endpoint.ResponseBody))
	}
}
func FetchURL(url string) (*MockEndpoint, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return &MockEndpoint{
		ResponseBody: string(body),
		StatusCode:   response.StatusCode,
		URL:          url,
	}, nil
}

func GetMockHTTP(url string) (*MockEndpoint, error) {
	var endpoint *MockEndpoint
	endpoint, err := ReadMockHTTP(url)
	if err == ErrReadFile {
		endpoint, err = FetchURL(url)
		if err != nil {
			return nil, err
		}
		err = endpoint.Write()
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		log.Printf("error getting mock http %s: %s", url, err)
		return nil, err
	}
	return endpoint, nil
}

func (endpoint MockEndpoint) Write() error {
	url := endpoint.URL
	jsonData, err := json.MarshalIndent(endpoint, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling data to JSON:", err)
		return err
	}

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

// read a mock http response from the file system
func ReadMockHTTP(url string) (*MockEndpoint, error) {
	data, err := os.ReadFile(GetFilePath(url))
	if err != nil {
		return nil, ErrReadFile
	}

	var mockData MockEndpoint
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

func (u URLSlice) ToMockEndpoints() (*MockEndpoints, error) {
	endpoints := MockEndpoints{}
	for _, url := range u {
		endpoint, err := GetMockHTTP(url)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, *endpoint)
	}
	return &endpoints, nil

}
