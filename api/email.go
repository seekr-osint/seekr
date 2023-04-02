package api

import (
	"fmt"
	"log"
	"regexp"
	"sync"
)

var RANDOM_USERNAME = "AAAANSUhEUgAAAgAAAA"
var RANDOM_PASSWORD = "64,iVBORw0KGgoAAAA$1"


var DefaultMailServices = MailServices{
	MailService{
		Name: "Discord",
		//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
		UserExistsFunc: Discord,
		Icon:           "https://assets-global.website-files.com/6257adef93867e50d84d30e2/636e0a6a49cf127bf92de1e2_icon_clyde_blurple_RGB.png",
	},
	//MailService{
	//	Name: "Spotify",
	//	//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
	//	UserExistsFunc: Spotify,
	//	Icon:           "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAB9AAAAfQCAYAAACaOMR5AAAgAElEQVR4nOzdebRsaVn",
	//},
	//MailService{
	//	Name: "Twitter",
	//	//UserExistsFunc: func(s MailService, str string) bool { return true }, // for testing useful
	//	UserExistsFunc: Twitter,
	//	Icon:           "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAA+gAAAPoCAYAAABNo9TkAAAAAXNSR0IArs4c6QAAIABJR",
	//},
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

func MailServicesHandler(servicesToCheck MailServices, email string,config ApiConfig) EmailServiceEnums {
	var mailMutex = sync.RWMutex{}
	wg := &sync.WaitGroup{}

	services := EmailServiceEnums{}
	// FIXME EMAIL SERVICE SCANNING
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i]               // current service
      err, userExsits := service.UserExistsFunc(service, email,config)
      if err != nil {
        log.Printf("error in service: %s,%e",service.Name,err)
      } else {
			if  userExsits { // if service exisits
          log.Printf("User %s on %s exsits",email,service.Name)
					mailMutex.Lock()
					services[service.Name] = EmailServiceEnum{
						Name: service.Name,
						Icon: service.Icon,
					} // add service to accounts
					mailMutex.Unlock()
			} else {
        log.Printf("User %s on %s does not exsit",email,service.Name)
      }
    }
			wg.Done()
		}(i)
	}
	wg.Wait()
  log.Println(len(services))
	return services
}

func CheckMail(newPerson Person,config ApiConfig) Person { // FIXME TODO
	var mailMutex = sync.RWMutex{}
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
        retMailServices := MailServicesHandler(DefaultMailServices, mail.Mail,config)
        log.Printf("found %d services",len(retMailServices))
				for key, value := range  retMailServices {
					go func(key string, value EmailServiceEnum) {
            log.Printf("%s = %s",key,value)
						mailMutex.Lock()
						mail.Services[key] = value
						mailMutex.Unlock()
					}(key, value)
				}
			} else {
				log.Println("nil mail field")
			}
			newPerson.Email[i] = mail
		}
	}
	return newPerson
}
