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
	"1-postPerson": { // ID 2
		RequestType:      "POST",
		Name:             "Post Person",
		URL:              "http://localhost:8080/person",
		PostData:         map[string]interface{}{"id": "2"},
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{}, "hobbies": "", "id": "2", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:       201,
	},
	"2-overwritePerson": {
		RequestType:      "POST",
		Name:             "Overwrite Person",
		URL:              "http://localhost:8080/person",
		PostData:         map[string]interface{}{"id": "1"},
		ExpectedResponse: map[string]interface{}{"message": "overwritten person"},
		StatusCode:       202,
	},
	"3-getPerson": { // ID 1
		RequestType:      "GET",
		Name:             "Get Person by ID",
		URL:              "http://localhost:8080/people/2",
		PostData:         nil,
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{}, "hobbies": "", "id": "2", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:       200,
	},
	"4-getPersonNotExisting": { // ID 100 NOTFOUND
		RequestType:      "GET",
		Name:             "Get Person which does not exsist",
		URL:              "http://localhost:8080/people/100",
		PostData:         nil,
		ExpectedResponse: nil,
		StatusCode:       404,
	},
	"5-email": { // ID 10
		RequestType:                "POST",
		Name:                       "Post person with included email",
		URL:                        "http://localhost:8080/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "address": "", "age": float64(10), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"fsdfadsfasdfasdf@gmail.com": map[string]interface{}{"mail": "fsdfadsfasdfasdf@gmail.com"}}, "hobbies": "", "id": "10", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": interface{}(nil), "political": "", "prevoccupation": "", "relations": interface{}(nil), "religion": "", "sources": interface{}(nil), "ssn": "", "tags": interface{}(nil)},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(10), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"fsdfadsfasdfasdf@gmail.com": map[string]interface{}{"gmail": true, "mail": "fsdfadsfasdfasdf@gmail.com", "provider": "", "services": map[string]interface{}{}, "src": "", "valid": true, "validGmail": true, "value": float64(0), "skipped_services": map[string]interface{}{}}}, "hobbies": "", "id": "10", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"6-emailServices": { // ID 11
		RequestType:                "POST",
		Name:                       "Post person with included email detecting only discord as a services",
		URL:                        "http://localhost:8080/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"has_discord_account@gmail.com": map[string]interface{}{"mail": "has_discord_account@gmail.com"}}, "id": "11", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(10), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"has_discord_account@gmail.com": map[string]interface{}{"gmail": true, "mail": "has_discord_account@gmail.com", "provider": "", "services": map[string]interface{}{"Discord": map[string]interface{}{"icon": "./images/mail/discord.png", "link": "", "name": "Discord", "username": ""}}, "src": "", "valid": true, "validGmail": false, "value": float64(0), "skipped_services": map[string]interface{}{}}}, "hobbies": "", "id": "11", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"7-allEmailServices": { // ID 11
		RequestType:                "POST",
		Name:                       "Post person with included email detecting all services",
		URL:                        "http://localhost:8080/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"mail": "all@gmail.com"}}, "id": "12", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(10), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"gmail": true, "mail": "all@gmail.com", "provider": "", "services": map[string]interface{}{"Discord": map[string]interface{}{"icon": "./images/mail/discord.png", "link": "", "name": "Discord", "username": ""}, "Spotify": map[string]interface{}{"icon": "./images/mail/spotify.png", "link": "", "name": "Spotify", "username": ""}, "Twitter": map[string]interface{}{"icon": "./images/mail/twitter.png", "link": "", "name": "Twitter", "username": ""}}, "src": "", "valid": true, "validGmail": true, "value": float64(0), "skipped_services": map[string]interface{}{}}}, "hobbies": "", "id": "12", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"8-email-error": { // ID 13
		RequestType:                "POST",
		Name:                       "Post person with included email and discord check failing",
		URL:                        "http://localhost:8080/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "address": "", "age": float64(13), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"discord_error@gmail.com": map[string]interface{}{"mail": "discord_error@gmail.com"}}, "hobbies": "", "id": "13", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": interface{}(nil), "political": "", "prevoccupation": "", "relations": interface{}(nil), "sources": interface{}(nil), "tags": interface{}(nil)},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "age": float64(13), "bday": "", "civilstatus": "", "club": "", "education": "", "email": map[string]interface{}{"discord_error@gmail.com": map[string]interface{}{"gmail": true, "mail": "discord_error@gmail.com", "provider": "", "services": map[string]interface{}{}, "skipped_services": map[string]interface{}{"Discord": true}, "src": "", "valid": true, "validGmail": false, "value": float64(0)}}, "hobbies": "", "id": "13", "kids": "", "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": "", "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "ssn": "", "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
}

type Requests = map[string]struct {
	RequestType                string
	Name                       string
	URL                        string
	PostData                   interface{}
	ExpectedResponse           interface{}
	StatusCode                 int
	RequiresInternetConnection bool
}

func toJsonString(data interface{}) string {
	jsonBytes, _ := json.MarshalIndent(data, "", "\t")
	return string(jsonBytes)
}

func writeDocs() {
	file, err := os.Create("doc.md")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, key := range SortMapKeys(requests) {
		value := requests[key]
		postData, _ := json.MarshalIndent(value.PostData, "", "\t")

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
			fmt.Printf("Error when writing to file: %e\n", err)
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

	for _, name := range SortMapKeys(requests) {
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
