package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	//"encoding/json"
	"bytes"
	"os"
	"time"
)

func waitForFile() {
	for {
		_, err := os.Stat("/tmp/running")
		if err == nil {
			fmt.Println("File found!")
			return
		}
		time.Sleep(time.Second) // wait for one second before checking again
	}
}

type Requests = map[string]struct {
	RequestType      string
	URL              string
	PostData         interface{}
	ExpectedResponse interface{}
}

func TestAPI(t *testing.T) {
	dbData := DataBase{
		"1": Person{
			ID:   "1",
			Name: "Test",
			Age:  1,
		},
	}
	TestApi(dbData)
	waitForFile()

	requests := Requests{
		"overwritePerson": {
			RequestType:      "POST",
			URL:              "http://localhost:8080/person",
			PostData:         Person{ID: "1"},
			ExpectedResponse: map[string]interface{}{"message": "overwritten person"},
		},
		"postPerson": {
			RequestType:      "POST",
			URL:              "http://localhost:8080/person",
			PostData:         Person{ID: "2"},
			ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{}, "hobbies": "", "id": "2", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": interface{}(nil), "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": interface{}(nil)},
		},
	}
	for name, req := range requests {
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
		var respBody = make(map[string]interface{})
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			t.Fatalf("[%s] %v", name, err)
		}

		// Compare the response body to the expected value
		if !reflect.DeepEqual(req.ExpectedResponse.(map[string]interface{}), respBody) {
			t.Errorf("[%s] Unexpected response body: %#v\nExpected %#v", name, respBody, req.ExpectedResponse)
		}
	}
}

// debug function
func areMapsEqual(map1, map2 map[string]interface{}) bool {
	if len(map1) != len(map2) {
		return false
	}
	for key, val1 := range map1 {
		val2, ok := map2[key]
		if !ok || !reflect.DeepEqual(val1, val2) {
			fmt.Printf("%s = %s; %s = %s", key, val1, key, val2)
			return false
		}
	}
	return true
}

func Test_GetPersonID(t *testing.T) {
	var config = ApiConfig{
		DataBase: DataBase{},
	}
	personExists, _ := GetPersonByID(config, "1")
	if personExists {
		t.Fatalf("got personExists true when selecting from an empty set")
	}
}
