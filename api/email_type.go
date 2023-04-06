package api

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Email
// Comperabel
// Methodes
func (e Email) IsGmailAddress() bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9._-]+@gmail.com$")
	return pattern.MatchString(e.Mail)
}

func (e Email) IsValidGmailAddress() bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.]+@gmail.com$")
	return pattern.MatchString(e.Mail)
}

func (e Email) IsValidEmail() bool {
	// Compile the regular expression pattern
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z]{2,})*$")

	// Check if the email matches the pattern
	return pattern.MatchString(e.Mail)
}

func (e Email) Parse() Email {

	Gmail := e.IsGmailAddress()
	ValidGmail := e.IsValidGmailAddress()
	if e.Services == nil {
		log.Printf("mail.Services == nil (%s)", e.Mail)
		e.Services = EmailServices{}
	}
	if Gmail && ValidGmail {
		e.Provider = "gmail"
	}
	if Gmail && !ValidGmail {
		e.Provider = "fake_mail"
	}
	e.Valid = e.IsValidEmail()
	if e.SkippedServices == nil {
		e.SkippedServices = SkippedServices{}
	}
	return e
}

func (et EmailsType) Markdown() string {
	var sb strings.Builder

	for _, k := range SortMapKeys(map[string]Email(et)) {
		v := et[k]
		sb.WriteString(fmt.Sprintf("### %s\n", k))
		sb.WriteString(v.Markdown())
		sb.WriteString("\n")
	}

	return sb.String()
}

func (e Email) Markdown() string {
	var sb strings.Builder

	// Write email header
	sb.WriteString(fmt.Sprintf("- Mail: `%s`\n", e.Mail))
	if e.Value != 0 {
		sb.WriteString(fmt.Sprintf("- Value: `%d`\n", e.Value))
	}
	if e.Src != "" {
		sb.WriteString(fmt.Sprintf("- Src: `%s`\n", e.Src))
	}
	if e.Provider != "" {
		sb.WriteString(fmt.Sprintf("- Provider: `%s`\n", e.Provider))
	}
	if !e.Valid {
		sb.WriteString(fmt.Sprintf("- Valid: `%t`\n", e.Valid))
	}

	// Write email services
	sb.WriteString(e.Services.Markdown())

	// Write skipped services
	sb.WriteString(e.SkippedServices.Markdown())
	return sb.String()
}
func (es EmailServices) Markdown() string {
	var sb strings.Builder
	if len(es) > 0 {
		sb.WriteString("#### Services\n")
		for _, k := range SortMapKeys(es) {
			v := es[k]
			sb.WriteString(v.Markdown(k))
		}
	}
	return sb.String()
}
func (ss SkippedServices) Markdown() string {
	var sb strings.Builder
	if len(ss) > 0 {
		sb.WriteString("#### Skipped Services\n")
		for _, k := range SortMapKeys(ss) {
			v := ss[k]
			sb.WriteString(fmt.Sprintf("- %s: `%t`\n", k, v))
		}
	}

	return sb.String()
}

func (s EmailService) Markdown(name string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("##### %s\n", name))
	sb.WriteString(fmt.Sprintf("Name: `%s`\n", s.Name))
	if s.Link != "" {
		sb.WriteString(fmt.Sprintf("Link: `%s`\n", s.Link))
	}
	if s.Username != "" {
		sb.WriteString(fmt.Sprintf("Username: `%s`\n", s.Username))
	}
	if s.Icon != "" {
		sb.WriteString(fmt.Sprintf("Icon: `%s`\n", s.Icon))
	}

	return sb.String()
}

// types
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
