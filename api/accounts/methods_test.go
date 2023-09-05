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
	for _, service := range *services {
		for _, mockData := range service.TestData.MockData {
			u, err := service.GetURLsMap(mockData.User)
			if err != nil {
				log.Panicf("error: %v", err)
			}
			urls = append(urls, u.ToSlice()...)
		}
	}
	endpoints, err := urls.ToMockEndpoints()
	if err != nil {
		t.Fatalf("error converting urls to mock endpoints: %s", err)
	}
	endpoints.CreateMockEndoints()
	t.Cleanup(httpmock.Deactivate)
	for _, service := range *services {
		for _, mockData := range service.TestData.MockData {
			t.Run(service.Name, func(t *testing.T) {
				res, err := service.RunScannerDefaultAccountResult(mockData.User)
				if err != nil {
					log.Panicf("error: %v", err)
				}
				if !reflect.DeepEqual(res.Account, mockData.Result.Account) {
					t.Errorf("expected %+v got %+v", mockData.Result.Account, res.Account)
				}
				if res.Exists != mockData.Result.Exists {
					t.Errorf("expected Exists = %v got Exists = %v", mockData.Result.Exists, res.Exists)
				}
				if res.Errors != mockData.Result.Errors {
					t.Errorf("expected %+v got %+v", mockData.Result.Errors, res.Errors)
				}
				if res.RateLimited != mockData.Result.RateLimited {
					t.Errorf("expected RateLimited = %v got RateLimited = %v", mockData.Result.RateLimited, res.RateLimited)
				}
			})
		}
	}
}
