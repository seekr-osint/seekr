package accounts

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewTest(t *testing.T) {
	services := DefaultServices()
	user := User{Username: "greg"}

	// create Mock Responses
	urls, err := services.ToURLSlice(user)
	if err != nil {
		t.Fatalf("error converting services to urls: %s", err)
	}
	endpoints, err := urls.ToMockEndpoints()
	if err != nil {
		t.Fatalf("error converting urls to mock endpoints: %s", err)
	}
	endpoints.CreateMockEndoints()
	t.Cleanup(httpmock.Deactivate)

	// run scan
	jobs := services.ToJobs(user)
	scanResults, err := jobs.StartWorkers(10)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	// print results
	for _, scanResult := range *scanResults {
		if !scanResult.Exists {
			fmt.Printf("Service %s has no accnount %s", scanResult.Account.Name, user.Username)
		}
	}
	// fmt.Println(scanResults)
}
