package api

// https://github.com/alpkeskin/wau/blob/main/cmd/apps/discord.go

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type discordResponse struct {
	Errors struct {
		Email struct {
			Errors []struct {
				Code    string `json:"code"`
				Message string `json:"message"`
			} `json:"_errors"`
		} `json:"email"`
	} `json:"errors"`
}

func Discord(mailService MailService, email string) bool {
	var endpoint = "https://discord.com/api/v9/auth/register"

	var jsonStr = []byte(`{"email":"` + email + `","username":"` + RandomString(10) + `","password":"` + RandomString(10) + `","invite":null,"consent":true,"date_of_birth":"1973-05-09","gift_code_sku_id":null,"captcha_key":null,"promotional_email_opt_in":false}`)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr)) // URL-encoded payload
	if err != nil {
    log.Println(err)
    return false
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	r.Header.Add("X-Debug-Options", "bugReporterEnabled")

	res, err := client.Do(r)
	if err != nil {
    log.Println(err)
    return false
	}
	defer res.Body.Close()
	if res.StatusCode == 400 {
		body, _ := ioutil.ReadAll(res.Body)
		var response discordResponse
		json.Unmarshal(body, &response)
		if len(response.Errors.Email.Errors) > 0 {
			if response.Errors.Email.Errors[0].Code == "EMAIL_ALREADY_REGISTERED" {
				return true
			}

		}
	} else if res.StatusCode == 429 {
		//("Too many requests to Discord!")
    log.Println("to many requests")
    return false
	}
	return false
}
