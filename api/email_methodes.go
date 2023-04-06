package api

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

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
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z]{2,})*$")
	return pattern.MatchString(e.Mail)
}

// FIXME unnecessary code
//
//	func (skippedServices SkippedServices) Parse() SkippedServices {
//		for _, skippedService := range SortMapKeys(map[string]bool(skippedServices)) {
//			if skippedServices[skippedService] == false {
//				//skippedServices = delete(map[string]bool(skippedServices), skippedService)
//			}
//		}
//		return skippedServices
//	}

func (emailAdresses EmailsType) Validate() error {
	for _, emailAdress := range SortMapKeys(map[string]Email(emailAdresses)) {
		if emailAdress != emailAdresses[emailAdress].Mail {
			return APIError{
				Message: fmt.Sprintf("Key missmatch: Email[%s] = %s", emailAdress, emailAdresses[emailAdress].Mail),
				Status:  http.StatusBadRequest,
			}
		}
	}
	return nil
}

func (emailAdresses EmailsType) Parse() EmailsType {
	for _, emailAdress := range SortMapKeys(map[string]Email(emailAdresses)) {
		emailAdresses[emailAdress] = emailAdresses[emailAdress].Parse()
	}
	return emailAdresses
}

func (e Email) Parse() Email {
	if e.Services == nil {
		log.Printf("mail.Services == nil (%s)", e.Mail)
		e.Services = EmailServices{}
	}
	if e.SkippedServices == nil {
		log.Printf("mail.SkippedServices == nil (%s)", e.Mail)
		e.SkippedServices = SkippedServices{}
	}

	Gmail := e.IsGmailAddress()
	ValidGmail := e.IsValidGmailAddress()
	if Gmail && ValidGmail {
		e.Provider = "gmail"
	}
	if Gmail && !ValidGmail {
		e.Provider = "fake_mail"
	}

	e.Valid = e.IsValidEmail()
	if !e.Valid {
		e.Provider = "invalid_email"
	}

	return e
}

func (et EmailsType) Markdown() string {
	var sb strings.Builder

	for _, emailAdress := range SortMapKeys(map[string]Email(et)) {
		v := et[emailAdress]
		sb.WriteString(fmt.Sprintf("### %s\n", emailAdress))
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
