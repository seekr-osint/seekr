package accounts

import (
	"log"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestRunCheck(t *testing.T) {
	services := DefaultServices()
	urls := URLSlice{}
	for _, service := range services {
		for _, mockData := range service.TestData.MockData {
			u, err := service.GetURLsMap(mockData.User)
			if err != nil {
				log.Panicf("error: %v", err)
			}
			urls = append(urls, u.ToSlice()...)
		}
	}
	endpoints := MockEndpoints{}
	for _, url := range urls {
		endpoint, err := GetMockHTTP(url)
		if err != nil {
			log.Fatalf("Error %s", err)
		}
		endpoints = append(endpoints, *endpoint)
	}
	endpoints.CreateMockEndoints()
	t.Cleanup(httpmock.Deactivate)
	for _, service := range services {
		for _, mockData := range service.TestData.MockData {
			t.Run(service.Name, func(t *testing.T) {
				res, err := service.RunScannerDefaultAccountResult(mockData.User)
				if err != nil {
					log.Panicf("error: %v", err)
				}
				if !reflect.DeepEqual(*res.Account, *mockData.Result.Account) {
					t.Errorf("expected %+v got %+v", res.Account, mockData.Result.Account)
				}
				if res.Exists != mockData.Result.Exists {
					t.Errorf("expected Exists = %v got Exists = %v", res.Exists, mockData.Result.Exists)
				}
				if res.Errors != mockData.Result.Errors {
					t.Errorf("expected %+v got %+v", res.Errors, mockData.Result.Errors)
				}
				if res.RateLimited != mockData.Result.RateLimited {
					t.Errorf("expected RateLimited = %v got RateLimited = %v", res.RateLimited, mockData.Result.RateLimited)
				}
			})
		}
	}
}
