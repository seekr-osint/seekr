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
	"github.com/seekr-osint/seekr/api/errortypes"
	"github.com/seekr-osint/seekr/api/reqcache"

	//gitplumbing "github.com/go-git/go-git/v5/plumbing"

	gitobject "github.com/go-git/go-git/v5/plumbing/object"
)

func (deep DeepInvestigation) GetGithubRepos() (GithubRepos, int, error) {
	err := deep.Validate()
	if err != nil {
		log.Printf("error validating: %s\n",err)
		return nil, 0, errortypes.APIError{
			Message: fmt.Sprintf("%s",err),
			Status: http.StatusInternalServerError,
		}
	}
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", deep.Username)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("error creating request: %s\n",err)
		return nil, 0, err
	}
	if len(deep.Tokens) >= 1 {
		if deep.Tokens[0] != "" {
			req.Header.Set("Authorization", fmt.Sprintf("token %s", deep.Tokens[0]))
		}
	}

	reqresp,err := reqcache.Reqcache(*req)
	if err != nil {
		log.Printf("err cachreq: %s\n",err)
		return nil, 0, errortypes.APIError{
			Message: fmt.Sprintf("%s",err),
			Status: http.StatusInternalServerError,
		}
	}

	fmt.Printf("%d\n",len(reqresp.Header))
	remaining, err := strconv.Atoi(reqresp.Header.Get("X-Ratelimit-Remaining"))
	if err != nil {
		log.Printf("error calculating rate limit: %s\n",err)
		return nil, 0, ErrCalcRateLimit
	}
	if remaining == 0 {
		return nil, remaining, ErrRateLimited
	}

	var repos GithubRepos
	err = json.Unmarshal(reqresp.Body,&repos)
	if err != nil {
		log.Printf("error Unmarshaling: %s\n",err)
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
		log.Printf("foundemail: %s", c.Author.Email)

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
			log.Printf("email: %s", email)
			computedEmails[email] = user
		}
	}

	return computedEmails, nil
}

func (deep DeepInvestigation) FilterEmails(recievedEmails ReceivedGitHubEmails) ([]ReceivedGitHubEmail, error) {
	err := deep.Validate()
	if err != nil {
		return nil, errortypes.APIError{
			Message: fmt.Sprintf("%s",err),
			Status: http.StatusInternalServerError,
		}
	}
	emails := []ReceivedGitHubEmail{} // FIXME type missmatch if github recieved github Emails type changes
	for _, recievedEmail := range recievedEmails {
		if strings.EqualFold(recievedEmail.User, deep.Username) && !recievedEmail.GithubMail {
			log.Printf("verified email: %s\n", recievedEmail.Email)
			emails = append(emails, recievedEmail)
		}
	}
	return emails, nil
}

func (deep DeepInvestigation) GetEmails() ([]ReceivedGitHubEmail, int, error) { // old
	repos, rateLimitRate, err := deep.GetGithubRepos()
	if rateLimitRate == 0 || err == ErrRateLimited {
		return nil, rateLimitRate, ErrRateLimited
	} else if err != nil {
		log.Printf("Error getting repos: %s\n", err)
		return nil, rateLimitRate, errortypes.APIError {
			Message: fmt.Sprintf("%s",err),
			Status: http.StatusInternalServerError,
		}
	}
	recievedGitHubEmails, err := deep.GetAllEmailsFromRepos(repos)
	if err != nil {
		return nil, rateLimitRate, err
	}
	filterdEmails, err := deep.FilterEmails(recievedGitHubEmails)
	return filterdEmails, rateLimitRate, err
}
