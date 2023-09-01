package services

import (
	"database/sql/driver"
	"encoding/json"
	"image"
	"time"

	"github.com/seekr-osint/seekr/api/history"
)

type Services []Service
type Service struct {
	Name                string            `json:"name"`
	UserExistsFunc      UserExistsFunc    `json:"-" tstype:"-"`
	InfoFunc            InfoFunc          `json:"-" tstype:"-"`
	UserHtmlUrlTemplate string            `json:"-" tstype:"-"`
	UrlTemplates        map[string]string `json:"-" tstype:"-"`
	Domain              string            `json:"domain"`
	Protocol            string            `json:"-" tstype:"-"`
	TestData            TestData          `json:"-" tstype:"-"`
	BlocksTor           bool              `json:"-" tstype:"-"`
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
	Exists error `json:"exists" tstype:"string"`
	Info   error `json:"info" tstype:"string"`
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
	Img  image.Image `tstype:"string"`
	Url  string
	Date time.Time `tstype:"string"`
}
type AccountInfo struct {
	Url            string                 `json:"url" tstype:"string"`
	ProfilePicture history.History[Image] `json:"profile_picture" tstype:"{ latest: { data: string } }"`
	Bio            history.History[Bio]   `json:"bio" tstype:"{ latest: { data: {bio: string} } }"`
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

func (e *MapServiceCheckResult) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), e); err != nil {
		return err
	}
	return nil
}

func (e MapServiceCheckResult) Value() (driver.Value, error) {
	value, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return value, nil
}
