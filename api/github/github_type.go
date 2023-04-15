package github

type DeepInvestigation struct {
	Username  string   `json:"username"`
	Tokens    []string `json:"tokens"`
	ScanForks bool     `json:"scan_forks"`
}
type GithubRepos []GithubRepo
type GithubRepo struct {
	Name string `json:"name"`
	Fork bool   `json:"fork"`
	Url  string `json:"html_url"`
}

type ReceivedGitHubEmails map[string]ReceivedGitHubEmail
type ReceivedGitHubEmail struct {
	Author     string `json:"author"`
	Email      string `json:"email"`
	User       string `json:"user"`
	CommitHash string `json:"commit_hash"`
	CommitUrl  string `json:"commit_url"`
	GithubMail bool   `json:"github_mail"`
}
