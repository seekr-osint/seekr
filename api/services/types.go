package services

type Services []Service
type Service struct {
	Name                string
	UserExistsFunc      UserExistsFunc
	UserHtmlUrlTemplate string
	Domain              string
	Protocol            string
	TestData            TestData
}
type TestData struct {
	ExsistingUser    string
	NotExsistingUser string
}
type User struct {
	Username string
}
type Template struct {
	User
	Service
}
type UserServiceDataToCheck struct {
	User    User
	Service Service
}

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
