package api

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Services []Service
type Service struct {
	Name           string         // example: "github"
	UserExistsFunc UserExistsFunc // example: SimpleUserExistsCheck()
	GetInfoFunc    GetInfoFunc    // example: EmptyAccountInfo()
	BaseUrl        string         // example: "https://github.com"
}

type Accounts map[string]Account
type Account struct {
	Service  string   `json:"service"`  // example: GitHub
	Id       string   `json:"id"`       // example: 1224234
	Username string   `json:"username"` // example: 9glenda
	Url      string   `json:"url"`      // example: https://github.com/9glenda
	Pricture []string `json:"profilePicture"`
	Bio      []string `json:"bio"` // example: pro hacka
}

type GetInfoFunc func(string, Service) Account // (username)
type UserExistsFunc func(string, string) bool  // (BaseUrl,username)

var DefaultServices = Services{
	Service{
		Name:           "github",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    GithubInfo,
		BaseUrl:        "https://github.com/",
	},
	Service{
		Name:           "slideshare",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    EmptyAccountInfo,
		BaseUrl:        "https://slideshare.net/",
	},
}

func SimpleUserExistsCheck(BaseUrl, username string) bool {
	return GetStatusCode(BaseUrl+username) == 200
}

func EmptyAccountInfo(username string, service Service) Account {
	return Account{
		Service:  service.Name,
		Username: username,
		Bio:      nil,
	}
}

// maybe remove
func DefaultServicesHandler(username string) Accounts {
	return ServicesHandler(DefaultServices, username)
}

func HttpRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	return string(body)
}

func ServicesHandler(servicesToCheck Services, username string) Accounts {
	var accounts = make(Accounts)
	for i := 0; i < len(servicesToCheck); i++ {
		service := servicesToCheck[i]
		if service.UserExistsFunc(service.BaseUrl, username) {
			accounts[service.Name] = service.GetInfoFunc(username, service)
		}
	}
	return accounts
}

func EncodeBase64(url string) string {
	return base64.StdEncoding.EncodeToString([]byte(HttpRequest(url)))
}

func GithubInfo(username string, service Service) Account {
	var data struct {
		Id         int    `json:"id"`
		Bio        string `json:"bio"`
		Avatar_url string `json:"avatar_url"`
		Url        string `json:"html_url"`
	}
	jsonData := HttpRequest("https://api.github.com/users/" + username)
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
	}
	account := Account{
		Service:  "GitHub",
		Username: username,
		Url:      data.Url,
		Id:       strconv.Itoa(data.Id),
		Bio:      []string{data.Bio},
		Pricture: []string{EncodeBase64(data.Avatar_url)},
	}
	return account
}
