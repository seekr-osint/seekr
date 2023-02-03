package api

import (
	"encoding/json"
	"fmt"
	//"fmt"
	"log"
	"os"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	//"github.com/go-git/go-git/v5/plumbing"
)

func GithubInfoDeep(username string, fork bool) EmailsType {
	log.Println("github")
	var data []struct {
		//Id     string `json:"id"`
		//NodeId string `json:"node_id"`
		Name string `json:"name"`

		FullName   string `json:"full_name"`
		Fork       bool   `json:"fork"`
		Url        string `json:"url"`
		GitUrl     string `json:"git_url"`
		SshUrl     string `json:"ssh_url"`
		CloneUrl   string `json:"clone_url"`
		OpenIssues int    `json:"open_issues"`
		Forks      int    `json:"forks"`
		Homepage   string `json:"homepage"`
		Created_at string `json:"created_at"`
		Updated_at string `json:"updated_at"`
		Pushed_at  string `json:"pushed_at"`
	}

	jsonData := HttpRequest("https://api.github.com/users/" + username + "/repos")

	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Println(err)
	}

	contributors := make(map[string]bool)
	foundEmail := make(map[string]bool)
	for _, repo := range data {
		//if repo.Fork == fork || repo.Fork {
		log.Println(repo.Name)

		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL: repo.CloneUrl,
		})
		Check(err)
		//head, err := r.Head()
		//Check(err)
		//commitIter, err := r.Log(&git.LogOptions{From: head.Hash()})

		commitIter, err := r.Log(&git.LogOptions{})
		Check(err)
		err = commitIter.ForEach(func(c *object.Commit) error {
			if !contributors[c.Author.Email] && !IsGitHubMail(c.Author.Email) {
				type Author struct {
					Name  string `json:"name"`
					Email string `json:"email"`
				}
				var commitInfo struct {
					Author Author `json:"author"`
				}

				jsonData := HttpRequest(fmt.Sprintf("https://api.github.com/repos/%s/git/commits/%s", repo.FullName, c.Hash.String()))

				err := json.Unmarshal([]byte(jsonData), &commitInfo)
				if err != nil {
					log.Println(err)
				}
				if commitInfo.Author.Name == username {
					log.Println("found:")
					log.Println(c.Author.Email)
				}
				foundEmail[c.Author.Email] = true
			}
			contributors[c.Author.Email] = true
			//log.Println(c.Hash.String())
			return nil
		})
		Check(err)

	}
	foundEmailArray := EmailsType{}
	for c := range foundEmail {
		foundEmailArray[c] = Email{
			Mail: c,
			Src:  "github",
			Services: EmailServiceEnums{
				"github": {
					Name: "GitHub",
					Link: fmt.Sprintf("https://github.com/%s", username),
					Icon: "https://github.githubassets.com/images/modules/logos_page/GitHub-Mark.png",
				},
			},
		}
	}
	//}
	return foundEmailArray
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Remove(path string) {
	if Exists(path) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Println(err)
		}
	}
}

func Mkdir(path string) {
	if !Exists(path) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
}
