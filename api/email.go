package api

import (
	"fmt"
	"log"
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
		Name: "Discord",
		//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
		UserExistsFunc: Discord,
		Icon:           "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAgAAAAIACAYAAAD0eNT6AAAAAXNSR0IArs4c6QAAIABJ",
	},
	MailService{
		Name: "Spotify",
		//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
		UserExistsFunc: Spotify,
		Icon:           "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAB9AAAAfQCAYAAACaOMR5AAAgAElEQVR4nOzdebRsaVn",
	},
	MailService{
		Name: "Twitter",
		//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
		UserExistsFunc: Twitter,
		Icon:           "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA+gAAAPoCAYAAABNo9TkAAAAAXNSR0IArs4c6QAAIABJR",
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

func MailServicesHandler(servicesToCheck MailServices, email string) EmailServiceEnums {
	var mailMutex = sync.RWMutex{}
	wg := &sync.WaitGroup{}

	services := EmailServiceEnums{}
	// FIXME EMAIL SERVICE SCANNING
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i]               // current service
			if service.UserExistsFunc(service, email) { // if service exisits
				go func() {
					mailMutex.Lock()
					services[service.Name] = EmailServiceEnum{
						Name: service.Name,
						Icon: service.Icon,
					} // add service to accounts
					mailMutex.Unlock()
				}()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return services
}

func CheckMail(newPerson Person) Person { // FIXME TODO
	//var mailMutex = sync.RWMutex{} // FIXME Netwoking
	fmt.Println(newPerson)
	if newPerson.Email == nil {
		log.Println("nil newPerson.Email")
		newPerson.Email = EmailsType{}
	}
	log.Println("email not nil")
	if len(newPerson.Email) == 0 {
		log.Println("empty list")
	} else {
		for i, mail := range newPerson.Email {
			if mail.Mail != "" {
				log.Println("email not \"\"")
				//mail.Services = MailServices(mail.Mail)
				mail.Valid = IsEmailValid(mail.Mail)
				mail.Gmail = IsGmailAddress(mail.Mail)
				mail.ValidGmail = IsValidGmailAddress(mail.Mail)

				if mail.Services == nil {
					mail.Services = EmailServiceEnums{}
				}

				// FIXME Netwoking
				//for key, value := range MailServicesHandler(DefaultMailServices, mail.Mail) {
				//	if mail.Services == nil {
				//		mail.Services = EmailServiceEnums{}
				//	}

				//	go func(key string, value EmailServiceEnum) {
				//		mailMutex.Lock()
				//		mail.Services[key] = value
				//		mailMutex.Unlock()
				//	}(key, value)
				//}
			} else {
				log.Println("nil mail field")
			}
			newPerson.Email[i] = mail
		}
	}
	return newPerson
}
