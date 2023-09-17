package email

type Emails map[string]Email
type Email struct {
	Mail            string          `json:"mail" tstype:"string"`
	Value           int             `json:"value" tstype:"number"`
	Src             string          `json:"src"`
	Services        EmailServices   `json:"services"`
	SkippedServices SkippedServices `json:"skipped_services"`
	Valid           bool            `json:"valid" tstype:"boolean"`
	Provider        string          `json:"provider" tstype:"string"`
}

type EmailServices map[string]EmailService
type EmailService struct {
	Name     string `json:"name" tstype:"string"`
	Link     string `json:"link" tstype:"string"`
	Username string `json:"username" tstype:"string"`
	Icon     string `json:"icon" tstype:"string"`
}

type SkippedServices map[string]bool

type MailService struct {
	Name           string             // example: "Discord"
	UserExistsFunc MailUserExistsFunc `tstype:"-"` // example: Discord(MailService,string,ApiConfig) (EmailService,error)
	Icon           string             // example: "./images/mail/discord.png"
	Url            string
}
type MailServices []MailService
