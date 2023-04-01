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

var requests = Requests{
	"overwritePerson": {
		RequestType:      "POST",
		Name:             "Overwrite Person",
		URL:              "http://localhost:8080/person",
		PostData:         map[string]interface{}{"id": "1"},
		ExpectedResponse: map[string]interface{}{"message": "overwritten person"},
		StatusCode:       202,
	},
	"postPerson": {
		RequestType:      "POST",
		Name:             "Post Person",
		URL:              "http://localhost:8080/person",
		PostData:         map[string]interface{}{"id": "2"},
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{}, "hobbies": "", "id": "2", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": interface{}(nil), "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": interface{}(nil)},
		StatusCode:       201,
	},
	"getPerson": {
		RequestType:      "GET",
		Name:             "Get Person by ID",
		URL:              "http://localhost:8080/people/1",
		PostData:         nil,
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{}, "hobbies": "", "id": "1", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": interface{}(nil), "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": interface{}(nil)},
		StatusCode:       200,
	},
	"getPersonNotExisting": {
		RequestType:      "GET",
		Name:             "Get Person which does not exsist",
		URL:              "http://localhost:8080/people/100",
		PostData:         nil,
		ExpectedResponse: nil,
		StatusCode:       404,
	},
}

type Requests = map[string]struct {
	RequestType      string
	Name             string
	URL              string
	PostData         interface{}
	ExpectedResponse interface{}
	StatusCode       int
}

func toJsonString(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "    ")
	return string(jsonBytes)
}

func writeDocs() {
	file, err := os.Create("doc.md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, value := range requests {
		postData, _ := json.Marshal(value.PostData)

		requestStr := fmt.Sprintf("**Curl Request:**\n\n```sh\ncurl -X %s %s", value.RequestType, value.URL)
		if value.RequestType != "GET" {
			requestStr += fmt.Sprintf(" \\\n-H 'Content-Type: application/json' \\\n-d '%s'", postData)
		}
		requestStr += "\n```\n\n"

		responseStr := fmt.Sprintf("**Response:**\n\n```json\n%s\n```\n\n", toJsonString(value.ExpectedResponse))
		statusCodeStr := fmt.Sprintf("**Status Code:** %d\n\n", value.StatusCode)
		markdownStr := fmt.Sprintf("## %s\n\n%s%s%s\n", value.Name, requestStr, responseStr, statusCodeStr)

		// write the markdown strings to the file
		_, err = file.WriteString(markdownStr)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

	}
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

	// WRITE DOCS
	writeDocs()
	// WRITE DOCS END

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
		var respBody interface{}
		err = json.NewDecoder(resp.Body).Decode(&respBody)
		if err != nil {
			t.Fatalf("[%s] %v", name, err)
		}

		if resp.StatusCode != req.StatusCode {
			t.Errorf("[%s] Unexpected Status Code: %d\nExpected %d", name, resp.StatusCode, req.StatusCode)
		}
		// Compare the response body to the expected value
		if !reflect.DeepEqual(req.ExpectedResponse, respBody) {
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
