package services

type Services []Service
type Service struct {
	Name                string `json:"name"`
	UserExistsFunc      UserExistsFunc `json:"-"`
	UserHtmlUrlTemplate string `json:"-"`
	Domain              string `json:"domain"`
	Protocol            string `json:"-"`
	TestData            TestData `json:"-"`
	BlocksTor 					bool `json:"-"`
}
type TestData struct {
	ExistingUser    string
	NotExistingUser string
}
type User struct {
	Username string
}
type Template struct {
	User
	Service
}

type ServiceCheckResults []ServiceCheckResult
type ServiceCheckResult struct {
	User    User `json:"user"`
	Service Service `json:"service"`
	Result  bool `json:"exists"`
	Error   error `json:"error"`
}

type DataToCheck []UserServiceDataToCheck
type UserServiceDataToCheck struct {
	User    User `json:"user"`
	Service Service `json:"service"`
}

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
