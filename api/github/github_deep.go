package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	git "github.com/go-git/go-git/v5"
	//gitplumbing "github.com/go-git/go-git/v5/plumbing"

	gitobject "github.com/go-git/go-git/v5/plumbing/object"
)

func (deep DeepInvestigation) GetGithubRepos() (GithubRepos, int, error) {
	err := deep.Validate()
	if err != nil {
		return nil, 0, err
	}
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", deep.Username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, 0, err
	}
	if len(deep.Tokens) >= 1 {
		if deep.Tokens[0] != "" {
			req.Header.Set("Authorization", fmt.Sprintf("token %s", deep.Tokens[0]))
		}
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

	var repos GithubRepos
	err = json.NewDecoder(resp.Body).Decode(&repos)
	if err != nil {
		return nil, remaining, err
	}

	return repos, remaining, nil
}

func (repoObj GithubRepo) ReturnEmails() (ReceivedGitHubEmails, error) {
	tmpDir, err := os.MkdirTemp("", "repo-clone-*")
	if err != nil {
		return nil, ErrCreatingTmp
	}
	defer os.RemoveAll(tmpDir) // clean up temp dir at end of function

	fmt.Printf("%s", repoObj.Url)
	repo, err := git.PlainClone(tmpDir, true, &git.CloneOptions{
		URL: repoObj.Url,
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
				CommitUrl:  fmt.Sprintf("%s/commit/%s", repoObj.Url, c.Hash.String()),
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

func (deep DeepInvestigation) GetAllEmailsFromRepos(repos GithubRepos) (ReceivedGitHubEmails, error) {
	err := deep.Validate()
	if err != nil {
		return nil, err
	}
	var wg sync.WaitGroup
	emails := make(chan ReceivedGitHubEmails)

	for _, repo := range repos {
		wg.Add(1)
		go func(repo GithubRepo) {
			defer wg.Done()
			emailMap, err := repo.ReturnEmails()
			if err == ErrCreatingTmp {
				// function should not continue running.
				log.Printf("Error creating temp dir: %#v", err)
			} else if err != nil {
				log.Printf("Error getting emails: %#v", err)
			} else {
				emails <- emailMap
			}
		}(repo)
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

	return computedEmails, nil
}

func (deep DeepInvestigation) FilterEmails(recivedEmails ReceivedGitHubEmails) ([]ReceivedGitHubEmail, error) {
	err := deep.Validate()
	if err != nil {
		return nil, err
	}
	emails := []ReceivedGitHubEmail{} // FIXME type missmatch if github recived github Emails type changes
	for _, recivedEmail := range recivedEmails {
		if strings.EqualFold(recivedEmail.User, deep.Username) && !recivedEmail.GithubMail {
			emails = append(emails, recivedEmail)
		}
	}
	return emails, nil
}

func (deep DeepInvestigation) GetEmails() ([]ReceivedGitHubEmail, int, error) { // old
	repos, rateLimitRate, err := deep.GetGithubRepos()
	if rateLimitRate == 0 || err == ErrRateLimited {
		return nil, rateLimitRate, ErrRateLimited
	} else if err != nil {
		log.Printf("Error getting repos: %v", err)
		return nil, rateLimitRate, err
	}
	recivedGitHubEmails, err := deep.GetAllEmailsFromRepos(repos)
	if err != nil {
		return nil, rateLimitRate, err
	}
	filterdEmails, err := deep.FilterEmails(recivedGitHubEmails)
	return filterdEmails, rateLimitRate, err
}
