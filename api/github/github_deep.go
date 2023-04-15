package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	git "github.com/go-git/go-git/v5"
	//gitplumbing "github.com/go-git/go-git/v5/plumbing"
	"regexp"

	gitobject "github.com/go-git/go-git/v5/plumbing/object"
	"github.com/gocolly/colly"
)

type GithubRepo struct {
	Name string `json:"name"`
	//Fork string `json:"fork"`
}

var (
	ErrCreatingTmp   = errors.New("failed to create temp dir")
	ErrCalcRateLimit = errors.New("failed to calculate the rate limit")
	ErrRateLimited   = errors.New("rate limited")
)

type ReceivedGitHubEmails map[string]ReceivedGitHubEmail
type ReceivedGitHubEmail struct {
	Author     string `json:"author"`
	Email      string `json:"email"`
	User       string `json:"user"`
	CommitHash string `json:"commit_hash"`
	CommitUrl  string `json:"commit_url"`
	GithubMail bool   `json:"github_mail"`
}

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

func (receivedGitHubEmail ReceivedGitHubEmail) Parse() ReceivedGitHubEmail {
	if !receivedGitHubEmail.GithubMail {
		receivedGitHubEmail = receivedGitHubEmail.GetUser()
	}
	return receivedGitHubEmail
}

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

func GetGithubRepos(username, token string) ([]GithubRepo, int, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}

	if token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", token))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	remaining, err := strconv.Atoi(resp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		return nil, 0, ErrCalcRateLimit
	}
	if remaining == 0 {
		return nil, remaining, ErrRateLimited
	}

	var repos []GithubRepo
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, remaining, err
	}

	return repos, remaining, nil
}

func returnEmails(repoUrl string) (ReceivedGitHubEmails, error) {
	tmpDir, err := os.MkdirTemp("", "repo-clone-*")
	if err != nil {
		return nil, ErrCreatingTmp
	}
	defer os.RemoveAll(tmpDir) // clean up temp dir at end of function

	repo, err := git.PlainClone(tmpDir, true, &git.CloneOptions{
		URL: repoUrl,
	})
	if err != nil {
		return nil, err
	}

	// Traverse the commit history to obtain contributors' email addresses
	emailMap := make(ReceivedGitHubEmails) // use a map to avoid duplicates
	iter, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get commit history: %v", err)
	}
	err = iter.ForEach(func(c *gitobject.Commit) error {
		if _, ok := emailMap[c.Author.Email]; !ok {
			emailMap[c.Author.Email] = ReceivedGitHubEmail{
				Email:      c.Author.Email,
				Author:     c.Author.Name,
				CommitHash: c.Hash.String(),
				CommitUrl:  fmt.Sprintf("%s/commit/%s", repoUrl, c.Hash.String()),
				GithubMail: isGitHubEmail(c.Author.Email),
			}.Parse()
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to traverse commit history: %v", err)
	}

	return emailMap, nil
}

func GetAllEmails(username, token string) (ReceivedGitHubEmails, int, error) {
	repos, rateLimitRate, err := GetGithubRepos(username, token)
	if rateLimitRate == 0 || err == ErrRateLimited {
		return nil, rateLimitRate, ErrRateLimited
	} else if err != nil {
		log.Printf("Error getting repos: %v", err)
		return nil, rateLimitRate, err
	}

	var wg sync.WaitGroup
	emails := make(chan ReceivedGitHubEmails)

	for _, repo := range repos {
		wg.Add(1)
		go func(repoName string) {
			defer wg.Done()
			url := fmt.Sprintf("https://github.com/%s/%s", username, repoName)
			emailMap, err := returnEmails(url)
			if err == ErrCreatingTmp {
				// function should not continue running.
				log.Printf("Error creating temp dir: %#v", err)
			} else if err != nil {
				log.Printf("Error getting emails: %#v", err)
			} else {
				emails <- emailMap
			}
		}(repo.Name)
	}

	go func() {
		wg.Wait()
		close(emails)
	}()

	computedEmails := make(ReceivedGitHubEmails)
	for emailMap := range emails {
		for email, user := range emailMap {
			computedEmails[email] = user
		}
	}

	return computedEmails, rateLimitRate, nil
}
func GetEmailsOfUser(username, token string) ([]ReceivedGitHubEmail, int, error) {
	emails := []ReceivedGitHubEmail{}
	recivedEmails, rateLimitRate, err := GetAllEmails(username, token)
	if rateLimitRate == 0 || err == ErrRateLimited {
		return nil, rateLimitRate, ErrRateLimited
	} else if err != nil {
		return nil, rateLimitRate, err
	}
	for _, recivedEmail := range recivedEmails {
		if strings.EqualFold(recivedEmail.User, username) && !recivedEmail.GithubMail {
			emails = append(emails, recivedEmail)
		}
	}
	return emails, rateLimitRate, nil
}

//func main() {
//	emails, rateLimitRate, err := GetEmailsOfUser("niteletsplay", "")
//	fmt.Printf("RateLimitRate: %d\n", rateLimitRate)
//	if err != nil {
//		panic(fmt.Sprintf("Error: %e", err))
//	}
//	for _, i := range emails {
//		fmt.Printf("%s\n", i.Email)
//	}
//}
