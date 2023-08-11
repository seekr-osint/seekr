package api

import (
	"encoding/json"
	"fmt"
	"testing"

	//"encoding/json"
	"os"
	"time"

	"github.com/seekr-osint/seekr/api/civilstatus"
	"github.com/seekr-osint/seekr/api/config"
	"github.com/seekr-osint/seekr/api/enum"
	"github.com/seekr-osint/seekr/api/ethnicity"
	"github.com/seekr-osint/seekr/api/functions"
	"github.com/seekr-osint/seekr/api/gender"
	"github.com/seekr-osint/seekr/api/religion"
	"github.com/seekr-osint/seekr/api/tc"
)

func waitForFile() {
	time.Sleep(time.Second) // wait for one second before checking again
	time.Sleep(time.Second) // wait for one second before checking again
}

var requests = tc.Requests{
	"1-postPerson": { // ID 2
		RequestType: "POST",
		Name:        "Post Person",
		URL:         "http://localhost:8080/api/person",
		PostData: functions.Interface(Person{
			ID: "2",
		}),
		//ExpectedResponse: functions.ParsedConfigInterface(Person{
		//	ID: "2",
		//}),
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "custom": interface{}(nil), "gender": "", "ethnicity": "", "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "hobbies": map[string]interface{}{}, "id": "2", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"2-overwritePerson": { // Id 1
		RequestType:      "POST",
		Name:             "Overwrite Person",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"id": "1"},
		ExpectedResponse: map[string]interface{}{"message": "overwritten person"},
		StatusCode:       202,
	},
	"3-getPerson": { // ID 2
		RequestType:      "GET",
		Name:             "Get Person by ID",
		URL:              "http://localhost:8080/api/people/2",
		PostData:         nil,
		ExpectedResponse: map[string]interface{}{"accounts": map[string]interface{}{}, "custom": interface{}(nil), "gender": "", "ethnicity": "", "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "hobbies": map[string]interface{}{}, "id": "2", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       200,
	},
	"4-getPersonNotExisting": { // ID 100 NOTFOUND
		RequestType:      "GET",
		Name:             "Get Person which does not exsist",
		URL:              "http://localhost:8080/api/people/100",
		PostData:         nil,
		ExpectedResponse: nil,
		StatusCode:       404,
	},
	"5-email": { // ID 10
		RequestType: "POST",
		Name:        "Post person with included email",
		URL:         "http://localhost:8080/api/person",
		PostData: functions.Interface(Person{
			Name: "Email test",
			ID:   "10",
			Email: EmailsType{
				"fsdfadsfasdfasdf@gmail.com": Email{
					Mail: "fsdfadsfasdfasdf@gmail.com",
				},
			},
			Age: 10,
		}),
		//PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"fsdfadsfasdfasdf@gmail.com": map[string]interface{}{"mail": "fsdfadsfasdfasdf@gmail.com"}}, "id": "10", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "custom": interface{}(nil), "gender": "", "ethnicity": "", "address": "", "age": float64(10), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"fsdfadsfasdfasdf@gmail.com": map[string]interface{}{"mail": "fsdfadsfasdfasdf@gmail.com", "provider": "gmail", "services": map[string]interface{}{}, "src": "", "valid": true, "value": float64(0), "skipped_services": map[string]interface{}{}}}, "hobbies": map[string]interface{}{}, "id": "10", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"7a-emailServices": { // ID 11
		RequestType:                "POST",
		Name:                       "Post person with included email detecting only discord as a services",
		URL:                        "http://localhost:8080/api/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"has_discord_account@gmail.com": map[string]interface{}{"mail": "has_discord_account@gmail.com"}}, "id": "11", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "gender": "", "ethnicity": "", "custom": interface{}(nil), "age": float64(10), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"has_discord_account@gmail.com": map[string]interface{}{"mail": "has_discord_account@gmail.com", "provider": "fake_mail", "services": map[string]interface{}{"Discord": map[string]interface{}{"icon": "./images/mail/discord.png", "link": "", "name": "Discord", "username": ""}}, "src": "", "valid": true, "value": float64(0), "skipped_services": map[string]interface{}{}}}, "hobbies": map[string]interface{}{}, "id": "11", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"7b-allEmailServices": { // ID 11
		RequestType:                "POST",
		Name:                       "Post person with included email detecting all services",
		URL:                        "http://localhost:8080/api/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"mail": "all@gmail.com"}}, "id": "12", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "gender": "", "ethnicity": "", "age": float64(10), "bday": "", "custom": interface{}(nil), "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"mail": "all@gmail.com", "provider": "gmail", "services": map[string]interface{}{"Discord": map[string]interface{}{"icon": "./images/mail/discord.png", "link": "", "name": "Discord", "username": ""}, "Spotify": map[string]interface{}{"icon": "./images/mail/spotify.png", "link": "", "name": "Spotify", "username": ""}, "Twitter": map[string]interface{}{"icon": "./images/mail/twitter.png", "link": "", "name": "Twitter", "username": ""}, "Ubuntu GPG": map[string]interface{}{"icon": "./images/mail/ubuntu.png", "link": "https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com&op=index", "name": "Ubuntu GPG", "username": ""}, "keys.gnupg.net": map[string]interface{}{"icon": "./images/mail/gnupg.ico", "link": "https://keys.gnupg.net/pks/lookup?search=all@gmail.com&op=index", "name": "keys.gnupg.net", "username": ""}}, "skipped_services": map[string]interface{}{}, "src": "", "valid": true, "value": float64(0)}}, "hobbies": map[string]interface{}{}, "id": "12", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"7c-email-error": { // ID 13
		RequestType:                "POST",
		Name:                       "Post person with included email and discord check failing",
		URL:                        "http://localhost:8080/api/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(13), "email": map[string]interface{}{"discord_error@gmail.com": map[string]interface{}{"mail": "discord_error@gmail.com"}}, "hobbies": map[string]interface{}{}, "id": "13", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": interface{}(nil), "political": "", "prevoccupation": "", "relations": interface{}(nil), "sources": interface{}(nil), "tags": interface{}(nil)},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "gender": "", "ethnicity": "", "age": float64(13), "bday": "", "civilstatus": "", "custom": interface{}(nil), "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"discord_error@gmail.com": map[string]interface{}{"mail": "discord_error@gmail.com", "provider": "fake_mail", "services": map[string]interface{}{}, "skipped_services": map[string]interface{}{"Discord": true}, "src": "", "valid": true, "value": float64(0)}}, "hobbies": map[string]interface{}{}, "id": "13", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"7d-fakeEmail": { // ID 14
		RequestType:                "POST",
		Name:                       "Post person with included email detected as a fake email by seekr",
		URL:                        "http://localhost:8080/api/person",
		PostData:                   map[string]interface{}{"accounts": interface{}(nil), "age": float64(10), "email": map[string]interface{}{"fake_mail@gmail.com": map[string]interface{}{"mail": "fake_mail@gmail.com"}}, "id": "14", "name": "Email test"},
		ExpectedResponse:           map[string]interface{}{"accounts": map[string]interface{}{}, "address": "", "custom": interface{}(nil), "gender": "", "ethnicity": "", "age": float64(10), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"fake_mail@gmail.com": map[string]interface{}{"mail": "fake_mail@gmail.com", "provider": "fake_mail", "services": map[string]interface{}{}, "skipped_services": map[string]interface{}{}, "src": "", "valid": true, "value": float64(0)}}, "hobbies": map[string]interface{}{}, "id": "14", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Email test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:                 201,
		RequiresInternetConnection: true,
	},
	"8a-accounts": { // ID 15
		RequestType:      "GET",
		Name:             "Post Person (civil status)",
		URL:              "http://localhost:8080/api/getAccounts/snapchat-exists",
		PostData:         nil,
		ExpectedResponse: map[string]interface{}{"Snapchat-snapchat-exists": map[string]interface{}{"bio": interface{}(nil), "blog": "", "created": "", "firstname": "", "followers": float64(0), "following": float64(0), "id": "", "lastname": "", "location": "", "profilePicture": interface{}(nil), "service": "Snapchat", "updated": "", "url": "", "username": "snapchat-exists"}},
		StatusCode:       200,
	},

	"8b-config": { // No id
		RequestType:      "GET",
		Name:             "Get the current seekr config",
		URL:              "http://localhost:8080/api/config",
		PostData:         nil,
		ExpectedResponse: functions.Interface(config.DefaultConfig()),
		StatusCode:       200,
	},
	"8c-config": { // No id
		RequestType:      "POST",
		Name:             "Post a seekr config",
		URL:              "http://localhost:8080/api/config",
		PostData:         functions.Interface(config.DefaultConfig()),
		ExpectedResponse: map[string]interface{}{"message": "updated config"},
		StatusCode:       202,
	},

	"8d-info": { // No id
		RequestType:      "GET",
		Name:             "Get info about seekr",
		URL:              "http://localhost:8080/api/info",
		ExpectedResponse: map[string]interface{}{"download_url": "https://github.com/seekr-osint/seekr/releases/download/0.0.1/seekr_0.0.1_linux_arm64", "is_latest": true, "latest": "0.0.1", "version": "0.0.1"},
		StatusCode:       200,
	},
	"9b1-postPerson": enum.TcRequestInvalidEnum(civilstatus.Enum, "http://localhost:8080/api/person"),
	"9b2-postPerson": enum.TcRequestValidEnum(civilstatus.Enum, "15", "http://localhost:8080/api/person", functions.Interface(Person{}.ReplaceNil())), // 15
	"9c1-postPerson": enum.TcRequestInvalidEnum(religion.Enum, "http://localhost:8080/api/person"),
	"9c2-postPerson": enum.TcRequestValidEnum(religion.Enum, "16", "http://localhost:8080/api/person", functions.Interface(Person{}.ReplaceNil())), // 16
	"9d1-postPerson": enum.TcRequestInvalidEnum(gender.Enum, "http://localhost:8080/api/person"),
	"9d2-postPerson": enum.TcRequestValidEnum(gender.Enum, "17", "http://localhost:8080/api/person", functions.Interface(Person{}.ReplaceNil())), // 17
	"9e1-postPerson": enum.TcRequestInvalidEnum(ethnicity.Enum, "http://localhost:8080/api/person"),
	"9e2-postPerson": enum.TcRequestValidEnum(ethnicity.Enum, "18", "http://localhost:8080/api/person", functions.Interface(Person{}.ReplaceNil())), // 18
	"9f-postPerson": { // ID NONE
		RequestType:      "POST",
		Name:             "Post Person (missing id)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{},
		ExpectedResponse: map[string]interface{}{"message": "Missing ID"},
		StatusCode:       400,
	},
	"9g-postPerson": { // ID 19
		RequestType:      "POST",
		Name:             "Post Person (Email key missmatch)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"age": float64(10), "email": map[string]interface{}{"key@gmail.com": map[string]interface{}{"mail": "missmatched_value@gmail.com"}}, "id": "10", "name": "Email test"},
		ExpectedResponse: map[string]interface{}{"message": "Key missmatch: Email[key@gmail.com] = missmatched_value@gmail.com"},
		StatusCode:       400,
	},
	"9h-postPerson": { // ID 20
		RequestType:      "POST",
		Name:             "Post Person (Email key missmatch)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"6502530000": map[string]interface{}{"number": "missmatched_value"}}, "id": "20", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"message": "Key missmatch: Phone[6502530000] = missmatched_value"},
		StatusCode:       400,
	},
	"9i-postPerson": { // ID 20
		RequestType:      "POST",
		Name:             "Post Person (Email key missmatch already taken ID)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"6502530000": map[string]interface{}{"number": "missmatched_value"}}, "id": "20", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"message": "Key missmatch: Phone[6502530000] = missmatched_value"},
		StatusCode:       400,
	},

	"9j-postPerson": { // ID 21
		RequestType:      "POST",
		Name:             "Post Person (Phone number formatting)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"+1(234) 567-8901": map[string]interface{}{"number": "+1(234) 567-8901"}}, "id": "21", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"custom": interface{}(nil), "accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "gender": "", "ethnicity": "", "hobbies": map[string]interface{}{}, "id": "21", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Phone test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{"+1 234-567-8901": map[string]interface{}{"tag": "", "number": "+1 234-567-8901", "national_format": "(234) 567-8901", "phoneinfoga": map[string]interface{}{"Carrier": "", "Country": "US", "CountryCode": float64(1), "E164": "+12345678901", "International": "12345678901", "Local": "(234) 567-8901", "RawLocal": "2345678901", "Valid": true}, "valid": true}}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"9k-postPerson": { // ID 30
		RequestType:      "POST",
		Name:             "Post Person (Phone number formatting missing +)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"1-234-567-8901": map[string]interface{}{"number": "1-234-567-8901"}}, "id": "30", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"custom": interface{}(nil), "accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "gender": "", "ethnicity": "", "hobbies": map[string]interface{}{}, "id": "30", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Phone test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{"+1 234-567-8901": map[string]interface{}{"tag": "", "number": "+1 234-567-8901", "national_format": "(234) 567-8901", "phoneinfoga": map[string]interface{}{"Carrier": "", "Country": "US", "CountryCode": float64(1), "E164": "+12345678901", "International": "12345678901", "Local": "(234) 567-8901", "RawLocal": "2345678901", "Valid": true}, "valid": true}}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"9l-postPerson": { // ID 22
		RequestType:      "POST",
		Name:             "Post Person (Invalid_number)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"Invalid_number": map[string]interface{}{"number": "Invalid_number"}}, "id": "22", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"custom": interface{}(nil), "accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "gender": "", "ethnicity": "", "hobbies": map[string]interface{}{}, "id": "22", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Phone test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{"Invalid_number": map[string]interface{}{"tag": "", "national_format": "", "number": "Invalid_number", "phoneinfoga": map[string]interface{}{"Carrier": "", "Country": "", "CountryCode": float64(0), "E164": "", "International": "", "Local": "", "RawLocal": "", "Valid": false}, "valid": false}}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"9m-postPerson": { // ID 23
		RequestType:      "POST",
		Name:             "Post Person (Empty phone number)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"": map[string]interface{}{"number": ""}}, "id": "23", "name": "Phone test"},
		ExpectedResponse: map[string]interface{}{"custom": interface{}(nil), "accounts": map[string]interface{}{}, "address": "", "age": float64(0), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{}, "gender": "", "ethnicity": "", "hobbies": map[string]interface{}{}, "id": "23", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Phone test", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"9n-postPerson": { // ID 24
		RequestType:      "POST",
		Name:             "Post Person (Lot of fields)",
		URL:              "http://localhost:8080/api/person",
		PostData:         map[string]interface{}{"phone": map[string]interface{}{"+13183442908": map[string]interface{}{"number": "+13183442908"}}, "id": "24", "name": "Many fields", "age": float64(23), "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"mail": "all@gmail.com"}}},
		ExpectedResponse: map[string]interface{}{"custom": interface{}(nil), "accounts": map[string]interface{}{}, "address": "", "age": float64(23), "bday": "", "civilstatus": "", "clubs": map[string]interface{}{}, "education": "", "email": map[string]interface{}{"all@gmail.com": map[string]interface{}{"mail": "all@gmail.com", "provider": "gmail", "services": map[string]interface{}{"Discord": map[string]interface{}{"icon": "./images/mail/discord.png", "link": "", "name": "Discord", "username": ""}, "Spotify": map[string]interface{}{"icon": "./images/mail/spotify.png", "link": "", "name": "Spotify", "username": ""}, "Twitter": map[string]interface{}{"icon": "./images/mail/twitter.png", "link": "", "name": "Twitter", "username": ""}, "Ubuntu GPG": map[string]interface{}{"icon": "./images/mail/ubuntu.png", "link": "https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com&op=index", "name": "Ubuntu GPG", "username": ""}, "keys.gnupg.net": map[string]interface{}{"icon": "./images/mail/gnupg.ico", "link": "https://keys.gnupg.net/pks/lookup?search=all@gmail.com&op=index", "name": "keys.gnupg.net", "username": ""}}, "skipped_services": map[string]interface{}{}, "src": "", "valid": true, "value": float64(0)}}, "gender": "", "ethnicity": "", "hobbies": map[string]interface{}{}, "id": "24", "kids": "", "ips": map[string]interface{}{}, "legal": "", "maidenname": "", "military": "", "name": "Many fields", "notaccounts": interface{}(nil), "notes": "", "occupation": "", "pets": "", "phone": map[string]interface{}{"+1 318-344-2908": map[string]interface{}{"national_format": "(318) 344-2908", "number": "+1 318-344-2908", "phoneinfoga": map[string]interface{}{"Carrier": "", "Country": "US", "CountryCode": float64(1), "E164": "+13183442908", "International": "13183442908", "Local": "(318) 344-2908", "RawLocal": "3183442908", "Valid": true}, "tag": "", "valid": true}}, "pictures": map[string]interface{}{}, "political": "", "prevoccupation": "", "relations": map[string]interface{}{}, "religion": "", "sources": map[string]interface{}{}, "tags": []interface{}{}},
		StatusCode:       201,
	},
	"9o-postPerson": { // ID 24
		RequestType:      "GET",
		Name:             "GET Person Markdown",
		URL:              "http://localhost:8080/api/people/24/markdown",
		PostData:         nil,
		ExpectedResponse: map[string]interface{}{"markdown": "# Many fields\n- Age: `23`\n- Phone: `+1 318-344-2908`\n## Email\n### all@gmail.com\n- Mail: `all@gmail.com`\n- Provider: `gmail`\n#### Services\n##### Discord\n- Name: `Discord`\n- Icon: `./images/mail/discord.png`\n##### Spotify\n- Name: `Spotify`\n- Icon: `./images/mail/spotify.png`\n##### Twitter\n- Name: `Twitter`\n- Icon: `./images/mail/twitter.png`\n##### Ubuntu GPG\n- Name: `Ubuntu GPG`\n- Link: `https://keyserver.ubuntu.com/pks/lookup?search=all@gmail.com&op=index`\n- Icon: `./images/mail/ubuntu.png`\n##### keys.gnupg.net\n- Name: `keys.gnupg.net`\n- Link: `https://keys.gnupg.net/pks/lookup?search=all@gmail.com&op=index`\n- Icon: `./images/mail/gnupg.ico`\n\n\n"},
		StatusCode:       200,
	},

	"9p-detectLanguage": { // ID 24
		RequestType: "POST",
		Name:        "Detect language",
		URL:         "http://localhost:8080/api/detect/language",
		PostData: map[string]string{
			"text": "это русский текст и, надеюсь, он будет обнаружен. Больше текста улучшит обнаружение, поэтому я должен сказать несколько случайных вещей, чтобы получить действительно хорошее обнаружение.", // Russian text
		},
		ExpectedResponse: map[string]interface{}{"Afrikaans": float64(0), "Albanian": float64(0), "Arabic": float64(0), "Armenian": float64(0), "Azerbaijani": float64(0), "Basque": float64(0), "Belarusian": float64(0), "Bengali": float64(0), "Bokmal": float64(0), "Bosnian": float64(0), "Bulgarian": float64(0), "Catalan": float64(0), "Chinese": float64(0), "Croatian": float64(0), "Czech": float64(0), "Danish": float64(0), "Dutch": float64(0), "English": float64(0), "Esperanto": float64(0), "Estonian": float64(0), "Finnish": float64(0), "French": float64(0), "Ganda": float64(0), "Georgian": float64(0), "German": float64(0), "Greek": float64(0), "Gujarati": float64(0), "Hebrew": float64(0), "Hindi": float64(0), "Hungarian": float64(0), "Icelandic": float64(0), "Indonesian": float64(0), "Irish": float64(0), "Italian": float64(0), "Japanese": float64(0), "Kazakh": float64(0), "Korean": float64(0), "Latin": float64(0), "Latvian": float64(0), "Lithuanian": float64(0), "Macedonian": float64(0), "Malay": float64(0), "Maori": float64(0), "Marathi": float64(0), "Mongolian": float64(0), "Nynorsk": float64(0), "Persian": float64(0), "Polish": float64(0), "Portuguese": float64(0), "Punjabi": float64(0), "Romanian": float64(0), "Russian": float64(1), "Serbian": float64(0), "Shona": float64(0), "Slovak": float64(0), "Slovene": float64(0), "Somali": float64(0), "Sotho": float64(0), "Spanish": float64(0), "Swahili": float64(0), "Swedish": float64(0), "Tagalog": float64(0), "Tamil": float64(0), "Telugu": float64(0), "Thai": float64(0), "Tsonga": float64(0), "Tswana": float64(0), "Turkish": float64(0), "Ukrainian": float64(0), "Urdu": float64(0), "Vietnamese": float64(0), "Welsh": float64(0), "Xhosa": float64(0), "Yoruba": float64(0), "Zulu": float64(0)},
		StatusCode:       200,
	},

	"9q-detectLanguage": { // ID 24
		RequestType: "POST",
		Name:        "Detect language from code",
		URL:         "http://localhost:8080/api/detect/language/code",
		PostData: map[string]string{
			"code": `
package main

func main() {
  // это русский текст и, надеюсь, он будет обнаружен. Больше текста улучшит обнаружение, поэтому я должен сказать несколько случайных вещей, чтобы получить действительно хорошее обнаружение.
	// Ky është një tekst rus dhe shpresoj se do të zbulohet. Më shumë tekst do të përmirësojë zbulimin, kështu që më duhet të them disa gjëra të rastësishme për të marrë një zbulim vërtet të mirë.
	fmt.Printf("hello world")
}
			`,
			"lang": "go",
		},
		ExpectedResponse: map[string]interface{}{"result": []interface{}{map[string]interface{}{"lang": map[string]interface{}{"Afrikaans": float64(0), "Albanian": float64(0), "Arabic": float64(0), "Armenian": float64(0), "Azerbaijani": float64(0), "Basque": float64(0), "Belarusian": float64(0), "Bengali": float64(0), "Bokmal": float64(0), "Bosnian": float64(0), "Bulgarian": float64(0), "Catalan": float64(0), "Chinese": float64(0), "Croatian": float64(0), "Czech": float64(0), "Danish": float64(0), "Dutch": float64(0), "English": float64(0), "Esperanto": float64(0), "Estonian": float64(0), "Finnish": float64(0), "French": float64(0), "Ganda": float64(0), "Georgian": float64(0), "German": float64(0), "Greek": float64(0), "Gujarati": float64(0), "Hebrew": float64(0), "Hindi": float64(0), "Hungarian": float64(0), "Icelandic": float64(0), "Indonesian": float64(0), "Irish": float64(0), "Italian": float64(0), "Japanese": float64(0), "Kazakh": float64(0), "Korean": float64(0), "Latin": float64(0), "Latvian": float64(0), "Lithuanian": float64(0), "Macedonian": float64(0), "Malay": float64(0), "Maori": float64(0), "Marathi": float64(0), "Mongolian": float64(0), "Nynorsk": float64(0), "Persian": float64(0), "Polish": float64(0), "Portuguese": float64(0), "Punjabi": float64(0), "Romanian": float64(0), "Russian": float64(1), "Serbian": float64(0), "Shona": float64(0), "Slovak": float64(0), "Slovene": float64(0), "Somali": float64(0), "Sotho": float64(0), "Spanish": float64(0), "Swahili": float64(0), "Swedish": float64(0), "Tagalog": float64(0), "Tamil": float64(0), "Telugu": float64(0), "Thai": float64(0), "Tsonga": float64(0), "Tswana": float64(0), "Turkish": float64(0), "Ukrainian": float64(0), "Urdu": float64(0), "Vietnamese": float64(0), "Welsh": float64(0), "Xhosa": float64(0), "Yoruba": float64(0), "Zulu": float64(0)}, "text": "   это русский текст и, надеюсь, он будет обнаружен. Больше текста улучшит обнаружение, поэтому я должен сказать несколько случайных вещей, чтобы получить действительно хорошее обнаружение."}, map[string]interface{}{"lang": map[string]interface{}{"Afrikaans": float64(0), "Albanian": float64(1), "Arabic": float64(0), "Armenian": float64(0), "Azerbaijani": float64(0), "Basque": float64(0), "Belarusian": float64(0), "Bengali": float64(0), "Bokmal": float64(0), "Bosnian": float64(0), "Bulgarian": float64(0), "Catalan": float64(0), "Chinese": float64(0), "Croatian": float64(0), "Czech": float64(0), "Danish": float64(0), "Dutch": float64(0), "English": float64(0), "Esperanto": float64(0), "Estonian": float64(0), "Finnish": float64(0), "French": float64(0), "Ganda": float64(0), "Georgian": float64(0), "German": float64(0), "Greek": float64(0), "Gujarati": float64(0), "Hebrew": float64(0), "Hindi": float64(0), "Hungarian": float64(0), "Icelandic": float64(0), "Indonesian": float64(0), "Irish": float64(0), "Italian": float64(0), "Japanese": float64(0), "Kazakh": float64(0), "Korean": float64(0), "Latin": float64(0), "Latvian": float64(0), "Lithuanian": float64(0), "Macedonian": float64(0), "Malay": float64(0), "Maori": float64(0), "Marathi": float64(0), "Mongolian": float64(0), "Nynorsk": float64(0), "Persian": float64(0), "Polish": float64(0), "Portuguese": float64(0), "Punjabi": float64(0), "Romanian": float64(0), "Russian": float64(0), "Serbian": float64(0), "Shona": float64(0), "Slovak": float64(0), "Slovene": float64(0), "Somali": float64(0), "Sotho": float64(0), "Spanish": float64(0), "Swahili": float64(0), "Swedish": float64(0), "Tagalog": float64(0), "Tamil": float64(0), "Telugu": float64(0), "Thai": float64(0), "Tsonga": float64(0), "Tswana": float64(0), "Turkish": float64(0), "Ukrainian": float64(0), "Urdu": float64(0), "Vietnamese": float64(0), "Welsh": float64(0), "Xhosa": float64(0), "Yoruba": float64(0), "Zulu": float64(0)}, "text": "\t Ky është një tekst rus dhe shpresoj se do të zbulohet. Më shumë tekst do të përmirësojë zbulimin, kështu që më duhet të them disa gjëra të rastësishme për të marrë një zbulim vërtet të mirë."}}},
		StatusCode:       200,
	},
	"9r-ping": {
		RequestType: "GET",
		Name:        "Ping the api",
		URL:         "http://localhost:8080/api/ping",
		PostData: nil,
		ExpectedResponse: map[string]interface{}{"ping":"pong"},
		StatusCode:       200,
	},
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

	for _, key := range functions.SortMapKeys(requests) {
		value := requests[key]
		postData, _ := json.MarshalIndent(value.PostData, "", "\t")

		requestStr := fmt.Sprintf("**Curl Request:**\n\n```sh\ncurl -X %s %s", value.RequestType, value.URL)
		if value.RequestType != "GET" {
			requestStr += fmt.Sprintf(" \\\n-H 'Content-Type: application/json' \\\n-d '%s'", postData)
		}
		requestStr += "\n```\n\n"

		responseStr := fmt.Sprintf("**Response:**\n\n```json\n%s\n```\n\n", toJsonString(value.ExpectedResponse))
		statusCodeStr := fmt.Sprintf("**Status Code:** %d\n\n", value.StatusCode)
		markdownStr := fmt.Sprintf("## %s\n%s\n\n%s%s%s\n", value.Name, value.Comment, requestStr, responseStr, statusCodeStr)

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

	writeDocs()

	// tc api tests
	test := tc.ApiTest{
		Requests: requests,
	}
	test.RunApiTests(t)
}

// debug function

func Test_GetPersonID(t *testing.T) {
	var config = ApiConfig{
		DataBase: DataBase{},
	}
	personExists, _ := GetPersonByID(config, "1")
	if personExists {
		t.Fatalf("got personExists true when selecting from an empty set")
	}
}
