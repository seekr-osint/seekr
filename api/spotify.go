
package api
// https://github.com/alpkeskin/wau/blob/main/cmd/apps/spotify.go

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type spotifyResponse struct {
	Status int `json:"status"`
	Errors struct {
		Email string `json:"email"`
	} `json:"errors"`
	Country                        string `json:"country"`
	CanAcceptLicensesInOneStep     bool   `json:"can_accept_licenses_in_one_step"`
	RequiresMarketingOptIn         bool   `json:"requires_marketing_opt_in"`
	RequiresMarketingOptInText     bool   `json:"requires_marketing_opt_in_text"`
	MinimumAge                     int    `json:"minimum_age"`
	CountryGroup                   string `json:"country_group"`
	SpecificLicenses               bool   `json:"specific_licenses"`
	TermsConditionsAcceptance      string `json:"terms_conditions_acceptance"`
	PrivacyPolicyAcceptance        string `json:"privacy_policy_acceptance"`
	SpotifyMarketingMessagesOption string `json:"spotify_marketing_messages_option"`
	PretickEula                    bool   `json:"pretick_eula"`
	ShowCollectPersonalInfo        bool   `json:"show_collect_personal_info"`
	UseAllGenders                  bool   `json:"use_all_genders"`
	UseOtherGender                 bool   `json:"use_other_gender"`
	DateEndianness                 int    `json:"date_endianness"`
	IsCountryLaunched              bool   `json:"is_country_launched"`
	AllowedCallingCodes            []struct {
		CountryCode string `json:"country_code"`
		CallingCode string `json:"calling_code"`
	} `json:"allowed_calling_codes"`
	PushNotifications bool `json:"push-notifications"`
}

func Spotify(mailService MailService, email string) bool {

	var endpoint string = "https://spclient.wg.spotify.com/signup/public/v1/account"

	data := url.Values{}
	data.Set("validate", "1")
	data.Set("email", email)

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	res, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body, _ := ioutil.ReadAll(res.Body)
		var response spotifyResponse
		json.Unmarshal(body, &response)
		if response.Status == 20 {
      return true
		} else {
      return false
		}
	} else {
      return false
	}
}
