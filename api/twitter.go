package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type twitterResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

func Twitter(mailService MailService, email string) bool {
	var endpoint string = "https://api.twitter.com/i/users/email_available.json"

	data := url.Values{}
	data.Set("email", email)

	r, err := http.Get(endpoint + "?" + data.Encode())
	if err != nil {
		log.Println(err)
		return false
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	if err != nil {
		log.Println(err)
		return false
	}
	if r.StatusCode == 200 {
		body, _ := ioutil.ReadAll(r.Body)
		var response twitterResponse
		json.Unmarshal(body, &response)
		if response.Taken {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
