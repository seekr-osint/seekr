package api

import (
	"github.com/seekr-osint/seekr/api/civilstatus"
	"github.com/seekr-osint/seekr/api/club"
	"github.com/seekr-osint/seekr/api/ethnicity"
	"github.com/seekr-osint/seekr/api/gender"
	"github.com/seekr-osint/seekr/api/hobby"
	"github.com/seekr-osint/seekr/api/ip"
	"github.com/seekr-osint/seekr/api/religion"
	"github.com/seekr-osint/seekr/api/services"
	"github.com/seekr-osint/seekr/api/sources"
)

// "github.com/seekr-osint/seekr/api/ssn"

// main data set
type Person struct {
	ID             string                         `json:"id" ts_transform:"__VALUE__ || ''"`
	Name           string                         `json:"name" ts_transform:"__VALUE__ || ''"`
	Gender         gender.Gender                  `json:"gender" ts_transform:"__VALUE__ || ''"`
	Ethnicity      ethnicity.Ethnicity            `json:"ethnicity" ts_transform:"__VALUE__ || ''"`
	Pictures       Pictures                       `json:"pictures"`
	Maidenname     string                         `json:"maidenname" ts_transform:"__VALUE__ || ''"`
	Age            Age                            `json:"age" ts_transform:"__VALUE__ || 0` // has to be a float64 becuase of json Unmarshal
	Birthday       string                         `json:"bday" ts_transform:"__VALUE__ || ''"`
	Address        string                         `json:"address" ts_transform:"__VALUE__ || ''"`
	Phone          PhoneNumbers                   `json:"phone"`
	Ips            ip.Ips                         `json:"ips"`
	Civilstatus    civilstatus.CivilStatus        `json:"civilstatus" ts_transform:"__VALUE__ || ''"`
	Kids           string                         `json:"kids" ts_transform:"__VALUE__ || ''"`
	Hobbies        hobby.Hobbies                  `json:"hobbies"`
	Email          EmailsType                     `json:"email"`
	Occupation     string                         `json:"occupation" ts_transform:"__VALUE__ || ''"`
	Prevoccupation string                         `json:"prevoccupation" ts_transform:"__VALUE__ || ''"`
	Education      string                         `json:"education" ts_transform:"__VALUE__ || ''"`
	Military       string                         `json:"military" ts_transform:"__VALUE__ || ''"`
	Religion       religion.Religion              `json:"religion" ts_transform:"__VALUE__ || ''"`
	Pets           string                         `json:"pets" ts_transform:"__VALUE__ || ''"`
	Clubs          club.Clubs                     `json:"clubs"`
	Legal          string                         `json:"legal" ts_transform:"__VALUE__ || ''"`
	Political      string                         `json:"political" ts_transform:"__VALUE__ || ''"`
	Notes          string                         `json:"notes" ts_transform:"__VALUE__ || ''"`
	Relations      Relation                       `json:"relations"` // FIXME
	Sources        sources.Sources                `json:"sources"`
	Accounts       Accounts                       `json:"accounts"`
	NewAccounts    services.MapServiceCheckResult `json:"-"`
	Tags           Tags                           `json:"tags"`
	NotAccounts    map[string]Account             `json:"notaccounts"`
	Custom         interface{}                    `json:"custom"`
}

type Relation map[string][]string

type Tags []Tag
type Tag struct {
	Name string `json:"name"`
}
type Pictures map[string]Picture
type Picture struct {
	Img     string `json:"img"`
	ImgHash uint64 `json:"img_hash"`
}
type Bios map[string]Bio
type Bio struct {
	Bio string `json:"bio"`
}

// type Accounts map[string]Account
type Accounts map[string]Account
type Account struct {
	Service   string   `json:"service"`  // example: GitHub
	Id        string   `json:"id"`       // example: 1224234
	Username  string   `json:"username"` // example: 9glenda
	Url       string   `json:"url"`      // example: https://github.com/9glenda
	Picture   Pictures `json:"profilePicture"`
	Bio       Bios     `json:"bio"`       // example: pro hacka
	Firstname string   `json:"firstname"` // example: Glenda
	Lastname  string   `json:"lastname"`  // example: Belov
	Location  string   `json:"location"`  // example: Moscow
	Created   string   `json:"created"`   // example: 2020-07-31T13:04:48Z
	Updated   string   `json:"updated"`
	Blog      string   `json:"blog"`
	Followers int      `json:"followers"`
	Following int      `json:"following"`
}
type Services []Service
type Service struct {
	Name              string         // example: "GitHub"
	UserExistsFunc    UserExistsFunc // example: SimpleUserExistsCheck()
	GetInfoFunc       GetInfoFunc    // example: EmptyAccountInfo()
	ImageFunc         ImageFunc
	ExternalImageFunc bool
	ScrapeImage       bool
	Scrape            ScrapeStruct
	BaseUrl           string // example: "https://github.com"
	AvatarUrl         string
	Check             string // example: "status_code"
	HtmlUrl           string
	Pattern           string
	BlockedPattern    string
}
type ScrapeStruct struct {
	FindElement string
	Attr        string
}
type GetInfoFunc func(string, Service, ApiConfig) (error, Account) // (username)
type ImageFunc func(string, Service) string                        // (username)
type UserExistsFunc func(Service, string, ApiConfig) (error, bool) // (service,username)
