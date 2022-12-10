package api

type Services []Service
type Service struct {
	Name           string         // example: "github"
	UserExistsFunc UserExistsFunc // example: SimpleUserExistsCheck()
	GetInfoFunc    GetInfoFunc    // example: EmptyAccountInfo()
	BaseUrl        string         // example: "https://github.com"
}

type Accounts map[string]Account
type Account struct {
	Service  string   `json:"service"`  // example: GitHub
	Id       string   `json:"id"`       // example: 1224234
	Username string   `json:"username"` // example: 9glenda
	Url      string   `json:"url"`      // example: https://github.com/9glenda
	Pricture []string `json:"profilePicture"`
	Bio      []string `json:"bio"` // example: pro hacka
}

type GetInfoFunc func(string, Service) Account // (username)
type UserExistsFunc func(string, string) bool  // (BaseUrl,username)

var DefaultServices = Services{
	Service{
		Name:           "github",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    EmptyAccountInfo,
		BaseUrl:        "https://github.com/",
	},
	Service{
		Name:           "slideshare",
		UserExistsFunc: SimpleUserExistsCheck,
		GetInfoFunc:    EmptyAccountInfo,
		BaseUrl:        "https://slideshare.net/",
	},
}

func SimpleUserExistsCheck(BaseUrl, username string) bool {
	return GetStatusCode(BaseUrl+username) == 200
}

func EmptyAccountInfo(username string, service Service) Account {
	return Account{
		Service:  service.Name,
		Username: username,
	}
}

// maybe remove
func DefaultServicesHandler(username string) Accounts {
	return ServicesHandler(DefaultServices, username)
}

func ServicesHandler(servicesToCheck Services, username string) Accounts {
  var accounts = make(Accounts)
	for i := 0; i < len(servicesToCheck); i++ {
		service := servicesToCheck[i]
		if service.UserExistsFunc(service.BaseUrl, username) {
			//accounts = append(accounts, service.GetInfoFunc(username, service))
      accounts[service.Name] = service.GetInfoFunc(username, service)
		}
	}
	return accounts
}

func CheckUsername(username string) []string {
	services := []string{
		"https://github.com/" + username,
		"https://www.shutterstock.com/fi/g/" + username,
		"https://www.myfitnesspal.com/user/idanshina/profile/" + username,
		"https://nitter.net/" + username,
		"https://slideshare.net/" + username,
	}

	valid := []string{}
	for _, service := range services {
		if GetStatusCode(service) == 200 {
			valid = append(valid, service)
		}
	}
	return valid
}
