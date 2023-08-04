package services

import (
	"image"
)

type Services []Service
type Service struct {
	Name                string            `json:"name"`
	UserExistsFunc      UserExistsFunc    `json:"-"`
	InfoFunc            InfoFunc          `json:"-"`
	UserHtmlUrlTemplate string            `json:"-"`
	UrlTemplates        map[string]string `json:"-"`
	Domain              string            `json:"domain"`
	Protocol            string            `json:"-"`
	TestData            TestData          `json:"-"`
	BlocksTor           bool              `json:"-"`
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
	User    User        `json:"user"`
	Service Service     `json:"service"`
	Result  bool        `json:"exists"`
	Info    AccountInfo `json:"info"`
	InfoErr error       `json:"InfoErr"`
	Error   error       `json:"error"`
}
type Image struct {
	Img image.Image
}
type AccountInfo struct {
	ProfilePicture Image `json:"profile_picture"`
	Bio            Bio   `json:"bio"`
}
type Bio struct {
	Bio      string             `json:"bio"`
	Language map[string]float64 `json:"language"`
}

type DataToCheck []UserServiceDataToCheck
type UserServiceDataToCheck struct {
	User    User    `json:"user"`
	Service Service `json:"service"`
}

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
type InfoFunc func(UserServiceDataToCheck) (AccountInfo, error)
