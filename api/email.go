package api

import (
	"regexp"
  "sync"
)

type MailService struct {
	Name           string         // example: "GitHub"
	UserExistsFunc MailUserExistsFunc // example: Discord()
}
type MailServices []MailService
type MailUserExistsFunc func(MailService, string) bool // (BaseUrl,email)

var DefaultMailServices = MailServices{
  MailService{
    Name:           "Discord",
    UserExistsFunc: Discord,
  },
}

func IsGmailAddress(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9._-]+@gmail.com$")
	return pattern.MatchString(email)
}

func IsValidGmailAddress(email string) bool {
	pattern := regexp.MustCompile("^[a-zA-Z0-9.]+@gmail.com$")
	return pattern.MatchString(email)
}

func IsEmailValid(email string) bool {
	// Compile the regular expression pattern
	pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z]{2,})*$")

	// Check if the email matches the pattern
	return pattern.MatchString(email)
}
func IsGitHubMail(email string) bool {
	match, _ := regexp.MatchString("@users\\.noreply\\.github\\.com$", email)
	return match
}




func MailServicesHandler(servicesToCheck MailServices, email string) []string {
	wg := &sync.WaitGroup{}

	var services []string
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i]                  // current service
			if service.UserExistsFunc(service, email) { // if service exisits
				services = append(services, service.Name) // add service to accounts
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return services
}
