package services

import (
	"image"
	"time"

	"github.com/seekr-osint/seekr/api/history"
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

type Errors struct {
	Exists error `json:"exists"`
	Info   error `json:"info"`
}
type ServiceCheckResults []ServiceCheckResult
type InputData struct {
	User    User    `json:"user"`
	Service Service `json:"service"`
}

type Accounts struct {
	Existing MapServiceCheckResult `json:"existing"`
	Failed   MapServiceCheckResult `json:"failed"`
}
type MapServiceCheckResult map[string]ServiceCheckResult
type ServiceCheckResult struct {
	InputData InputData   `json:"input_data"`
	Exists    bool        `json:"exists"`
	Info      AccountInfo `json:"info"`
	Errors    Errors      `json:"errors"`
}
type Image struct {
	Img  image.Image
	Url  string
	Date time.Time
}
type AccountInfo struct {
	Url            string                 `json:"url"`
	ProfilePicture history.History[Image] `json:"profile_picture" ts_type:"{ latest: { data: string } } "`
	Bio            history.History[Bio]   `json:"bio" ts_type:"{ latest: { data: {bio: string} } }"`
}
type Bio struct {
	Bio      string             `json:"bio"`
	Language map[string]float64 `json:"language"`
}

type DataToCheck []UserServiceDataToCheck
type UserServiceDataToCheck struct {
	User    User    `json:"user"`
	Service Service `json:"service"`
	// ExistingServiceCheckResult ServiceCheckResult `json:"-"`
}

type UserExistsFunc func(UserServiceDataToCheck) (bool, error)
type InfoFunc func(UserServiceDataToCheck) (AccountInfo, error)
