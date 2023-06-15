package services

type Services []Service
type Service struct {
	Name                string         // example: "GitHub"
	UserExistsFunc      UserExistsFunc // example: SimpleUserExistsCheck()
	UserHtmlUrlTemplate string         // example: "https://github.com"
}
type User struct {
	Username string
}
type UserServiceDataToCheck struct {
	User    User
	Service Service
}

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
