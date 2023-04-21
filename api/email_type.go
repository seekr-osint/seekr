package api

// Email

// Types
type EmailsType map[string]Email
type Email struct {
	Mail            string          `json:"mail"`
	Value           int             `json:"value"`
	Src             string          `json:"src"`
	Services        EmailServices   `json:"services"`
	SkippedServices SkippedServices `json:"skipped_services"`
	Valid           bool            `json:"valid"`
	Provider        string          `json:"provider"`
}

type EmailServices map[string]EmailService
type EmailService struct {
	Name     string `json:"name"`
	Link     string `json:"link"`
	Username string `json:"username"`
	Icon     string `json:"icon"`
}

type SkippedServices map[string]bool

type MailUserExistsFunc func(MailService, string, ApiConfig) (EmailService, error)
type MailService struct {
	Name           string             // example: "Discord"
	UserExistsFunc MailUserExistsFunc // example: Discord(MailService,string,ApiConfig) (EmailService,error)
	Icon           string             // example: "./images/mail/discord.png"
	Url            string
}
type MailServices []MailService
