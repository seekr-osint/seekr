package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type twitterResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

func TwitterMail(mailService MailService, email string, config ApiConfig) (EmailService, error) {
	emailService := EmailService{
		Name: mailService.Name,
		Icon: mailService.Icon,
		Link: strings.ReplaceAll(mailService.Url, "{{ email }}", email),
	}
	if config.Testing {
		if email == "has_twitter_account@gmail.com" || email == "twitter@gmail.com" || email == "all@gmail.com" {
			log.Println("has_twitter_account testing case true")
			return emailService, nil
		}
		log.Println("has_twitter_account testing case false")

		return EmailService{}, nil
	}
	var endpoint = "https://api.twitter.com/i/users/email_available.json"

	data := url.Values{}
	data.Set("email", email)

	r, err := http.Get(endpoint + "?" + data.Encode())
	if err != nil {
		log.Println(err)

		return EmailService{}, err
	}
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	if err != nil {
		log.Println(err)

		return EmailService{}, err
	}
	if r.StatusCode == 200 {
		body, _ := ioutil.ReadAll(r.Body)
		var response twitterResponse
		json.Unmarshal(body, &response)
		if response.Taken {
			return emailService, nil
		} else {

			return EmailService{}, nil
		}
	} else {

		return EmailService{}, nil
	}
}
