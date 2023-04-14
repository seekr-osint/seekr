package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func searchGithubUsersByEmail(email, token string) ([]GithubUser, error) {
	url := fmt.Sprintf("https://api.github.com/search/users?q=%s", email)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var searchResult GithubSearchResult
	err = json.NewDecoder(resp.Body).Decode(&searchResult)
	if err != nil {
		return nil, err
	}

	return searchResult.Items, nil
}

type GithubUser struct {
	Login     string `json:"login"`
	ID        int    `json:"id"`
	AvatarURL string `json:"avatar_url"`
	URL       string `json:"html_url"`
}

type GithubSearchResult struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []GithubUser `json:"items"`
}

func GitHubEmail(mailService MailService, email string, config ApiConfig) (EmailService, error) { // FIXME multiple results
	emailService := EmailService{
		Name: mailService.Name,
		Icon: mailService.Icon,
	}
	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return emailService, nil
		} else if email == "error@gmail.com" {
			return EmailService{}, errors.New("error")
		}
		log.Println("all email testing case false")

		return EmailService{}, nil
	}
	githubUser, err := searchGithubUsersByEmail(email, "")
	if len(githubUser) >= 1 {
		emailService.Username = githubUser[0].Login
		emailService.Link = githubUser[0].URL
	}

	return emailService, err
}
