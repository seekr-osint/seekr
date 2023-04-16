package github

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
)

// Parse

func (receivedGitHubEmail ReceivedGitHubEmail) Parse() ReceivedGitHubEmail {
	if !receivedGitHubEmail.GithubMail {
		receivedGitHubEmail = receivedGitHubEmail.GetUser()
	}
	return receivedGitHubEmail
}

// Checks for parsing fields

func isGitHubEmail(email string) bool {
	re := regexp.MustCompile(`^\d+\+[\w-]+@users\.noreply\.github\.com$`)
	return re.MatchString(email)
}

func extractGitHubUsername(email string) (string, error) {
	re := regexp.MustCompile(`^(\d+)\+([\w-]+)@users\.noreply\.github\.com$`)
	match := re.FindStringSubmatch(email)
	if match == nil {
		return "", fmt.Errorf("email %s is not a valid GitHub email", email)
	}
	return match[2], nil
}

// Validation

func (deep DeepInvestigation) Validate() error {
	if deep.Username == "" {
		return ErrMissingUsername
	}
	return nil
}

// Networking

func (receivedGitHubEmail ReceivedGitHubEmail) GetUser() ReceivedGitHubEmail {
	// Instantiate a new collector
	c := colly.NewCollector()

	// Find and extract the text inside the <div> tag with class "AvatarStack-body"
	c.OnHTML("div.AvatarStack-body", func(e *colly.HTMLElement) {
		username := e.Attr("aria-label")
		receivedGitHubEmail.User = username
	})

	// Visit the webpage
	c.Visit(receivedGitHubEmail.CommitUrl)
	return receivedGitHubEmail
}
