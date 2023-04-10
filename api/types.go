package api

import (
	"embed"

	"github.com/gin-gonic/gin"
)

// main data set
type Person struct {
	ID             string             `json:"id"`
	Name           string             `json:"name"`
	Gender         Gender             `json:"gender"`
	Pictures       Pictures           `json:"pictures"`
	Maidenname     string             `json:"maidenname"`
	Age            Age                `json:"age"` // has to be a float64 becuase of json Unmarshal
	Birthday       string             `json:"bday"`
	Address        string             `json:"address"`
	Phone          PhoneNumbers       `json:"phone"`
	SSN            SSN                `json:"ssn"`
	Civilstatus    CivilStatus        `json:"civilstatus"`
	Kids           string             `json:"kids"`
	Hobbies        string             `json:"hobbies"`
	Email          EmailsType         `json:"email"`
	Occupation     string             `json:"occupation"`
	Prevoccupation string             `json:"prevoccupation"`
	Education      string             `json:"education"`
	Military       string             `json:"military"`
	Religion       Religion           `json:"religion"`
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

type SaveDBFunc func(ApiConfig) error
type LoadDBFunc func(ApiConfig) (ApiConfig, error)
type ApiConfig struct {
	Ip             string        `json:"ip"`
	TempMailIp     string        `json:"temp_mail_ip"`
	WebServer      bool          `json:"web_server"`
	WebServerFS    embed.FS      `json:"web_server_fs"`
	LogFile        string        `json:"log_file"`
	DataBaseFile   string        `json:"data_base_file"`
	DataBase       DataBase      `json:"data_base"`
	SetCORSHeader  bool          `json:"set_CORS_header"`
	SaveDBFunc     SaveDBFunc    `json:"save_db_func"`
	LoadDBFunc     LoadDBFunc    `json:"load_db_func"`
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
