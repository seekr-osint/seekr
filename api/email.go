package api

import (
	"regexp"
	"sync"
)

type MailService struct {
	Name           string             // example: "GitHub"
	UserExistsFunc MailUserExistsFunc // example: Discord()
	Icon           string             // example: https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6a49cf127bf92de1e2_icon_clyde_blurple_RGB.png
}
type MailServices []MailService
type MailUserExistsFunc func(MailService, string) bool // (BaseUrl,email)

var DefaultMailServices = MailServices{
	MailService{
		Name:           "Discord",
		//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
    UserExistsFunc: Discord,
		Icon:           "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6cc3c481a15a141738_icon_clyde_white_RGB.png",
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

func MailServicesHandler(servicesToCheck MailServices, email string) []EmailServiceEnum {
	wg := &sync.WaitGroup{}

	var services []EmailServiceEnum
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i]               // current service
			if service.UserExistsFunc(service, email) { // if service exisits
				services = append(services, EmailServiceEnum{
					Name: service.Name,
					Icon: service.Icon,
				}) // add service to accounts
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return services
}
