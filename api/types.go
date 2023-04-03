package api

import (
	"github.com/gin-gonic/gin"
)

// main data set
type person Person // legacy stuff
type Person struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Pictures       Pictures           `json:"pictures"`
	Maidenname     string             `json:"maidenname"`
	Age            float64            `json:"age"` // has to be a float64 becuase of json Unmarshal
	Birthday       string             `json:"bday"`
	Address        string             `json:"address"`
	Phone          string             `json:"phone"`
	SSN            string             `json:"ssn"`
	Civilstatus    string             `json:"civilstatus"`
	Kids           string             `json:"kids"`
	Hobbies        string             `json:"hobbies"`
	Email          EmailsType         `json:"email"`
	Occupation     string             `json:"occupation"`
	Prevoccupation string             `json:"prevoccupation"`
	Education      string             `json:"education"`
	Military       string             `json:"military"`
	Religion       string             `json:"religion"`
	Pets           string             `json:"pets"`
	Club           string             `json:"club"`
	Legal          string             `json:"legal"`
	Political      string             `json:"political"`
	Notes          string             `json:"notes"`
	Relations      Relation           `json:"relations"` // FIXME
	Sources        Sources            `json:"sources"`
	Accounts       Accounts           `json:"accounts"`
	Tags           Tags               `json:"tags"`
	NotAccounts    map[string]Account `json:"notaccounts"`
}

type DataBase map[string]Person
type Relation map[string][]string
type Sources map[string]Source
type Source struct {
	Url string `json:"url"`
}
type Tags []Tag
type Tag struct {
	Name string `json:"name"`
}
type EmailServiceEnum struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}
type Pictures map[string]Picture
type Picture struct {
	Img     string `json:"img"`
	ImgHash uint64 `json:"img_hash"`
}
type EmailsType map[string]Email
type EmailServiceEnums map[string]EmailServiceEnum
type Bios map[string]Bio
type Bio struct {
	Bio string `json:"bio"`
}
type Email struct {
	Mail       string            `json:"mail"`
	Value      int               `json:"value"`
	Src        string            `json:"src"`
	Services   EmailServiceEnums `json:"services"`
  SkippedServices          SkippedServicesEnum     `json:"skipped_services"`
	Valid      bool              `json:"valid"`
	Gmail      bool              `json:"gmail"`
	ValidGmail bool              `json:"validGmail"`
	Provider   string            `json:"provider"`
}
type SkippedServicesEnum map[string]bool

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

type SaveJsonFunc func(ApiConfig)
type ApiConfig struct {
	Ip             string        `json:"ip"`
	LogFile        string        `json:"log_file"`
	DataBaseFile   string        `json:"data_base_file"`
	DataBase       DataBase      `json:"data_base"`
	SetCORSHeader  bool          `json:"set_CORS_header"`
	SaveJsonFunc   SaveJsonFunc  `json:"save_json_func"`
	GinRouter      *gin.Engine   `json:"gin_router"`
	ApiKeysComplex ApiKeys       `json:"api_keys_complex"`
	ApiKeysSimple  ApiKeysSimple `json:"api_keys"`
	Testing        bool          `json:"testing"`
}
type ApiKeysSimple map[string][]string // map["serviceName"]["key1","key2"]
type ApiKeys struct {
	Github ApiKeyEnum `json:"github"`
}
type ApiKeyEnum map[string]ApiKey
type ApiKey struct {
}
type MailService struct {
	Name           string             // example: "GitHub"
	UserExistsFunc MailUserExistsFunc // example: Dis10cord()
	Icon           string             // example: https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6a49cf127bf92de1e2_icon_clyde_blurple_RGB.png
}
type MailServices []MailService
type MailUserExistsFunc func(MailService, string,ApiConfig) (error,bool) // (BaseUrl,email)

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
type GetInfoFunc func(string, Service, ApiConfig) (error,Account) // (username)
type ImageFunc func(string, Service) string               // (username)
type UserExistsFunc func(Service, string, ApiConfig) (error,bool) // (service,username)
