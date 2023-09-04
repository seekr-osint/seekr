package accounts

type User struct {
	Username string `json:"username" tstype:"string"`
}

type TestData struct {
	UserExistsCheck map[User]struct {
		Exists bool
		Error  error
	}
}

type Services []AccountScanner
type AccountScanner struct {
	Name            string       `json:"name" tstype:"string"`
	Domain          string       `json:"domain" tstype:"string"`
	Protocol        string       `json:"protocol" tstype:"string"`
	URLTemplates    URLTemplates `json:"url_templates"`
	UserExistsCheck URLTemplate  `json:"user_exists_check" tstype:"string"`
}

type URLTemplates map[string]URLTemplate
type URLTemplate string

type URLs map[string]string

// type URL string // TODO protocol as method

type UserExistsCheckInput struct {
	URLs URLs `json:"urls"`
}
type URLTemplateInput struct {
	User
	AccountScanner
}

//	type ScanResult[T Account] struct {
//		Account T      `json:"account"`
//		Errors  Errors `json:"errors"`
//		RateLimited bool `json:"rate_limited" tstype:"bolean"`
//	}
type ScanResult[T Account] struct {
	Exists      bool   `json:"exists"`
	Account     *T     `json:"account"`
	Errors      Errors `json:"errors"`
	RateLimited bool   `json:"rate_limited" tstype:"bolean"`
}
type Errors struct {
	UserExistsCheck error `json:"user_exists_check" tstype:"string"`
}

type Account interface {
	DefaultAccount | GithubAccount
}

type DefaultAccount struct {
	Name string `json:"name" tstype:"string"`
	URL  string `json:"url" tstype:"string"`
}
type GithubAccount struct{}
