package api

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"sync"
)

var RANDOM_USERNAME = "AAAANSUhEUgAAAgAAAA"
var RANDOM_PASSWORD = "64,iVBORw0KGgoAAAA$1"

var DefaultMailServices = MailServices{
	MailService{
		Name:           "Discord",
		UserExistsFunc: Discord,
		Icon:           "./images/mail/discord.png",
	},
	MailService{
		Name:           "Spotify",
		UserExistsFunc: Spotify,
		Icon:           "./images/mail/spotify.png",
	},
	MailService{
		Name:           "Twitter",
		UserExistsFunc: Twitter,
		Icon:           "./images/mail/twitter.png",
	},
	MailService{
		Name:           "Ubuntu GPG",
		UserExistsFunc: UbuntuGPGUserExists,
		Icon:           "https://ubuntu.com/favicon.ico",
		Url:            "https://keyserver.ubuntu.com/pks/lookup?search={{ email }}&op=index",
	},
	MailService{
		Name:           "keys.gnupg.net",
		UserExistsFunc: KeysGnuPGUserExists,
		Icon:           "https://www.gnupg.org/favicon.ico",
		Url:            "https://keys.gnupg.net/pks/lookup?search={{ email }}&op=index",
	},

	MailService{
		Name:           "keyserver.pgp.com",
		UserExistsFunc: KeyserverPGPUserExists,
		Icon:           "https://pgp.com/favicon.ico",
		Url:            "https://keyserver.pgp.com/pks/lookup?search={{ email }}&op=index",
	},
	//MailService{ // FIXME
	//    Name: "pgp.mit.edu",
	//    UserExistsFunc: PgpMitUserExists,
	//    Icon: "https://pgp.mit.edu/favicon.ico",
	//},

	// MailService{ // FIXME
	//
	//	   Name: "pool.sks-keyservers.net",
	//	   UserExistsFunc: PoolSKSUserExists,
	//	   Icon: "https://sks-keyservers.net/favicon.ico",
	//	},
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

func MailServicesHandler(servicesToCheck MailServices, email string, config ApiConfig) (EmailServiceEnums, SkippedServicesEnum) {
	var mailMutex = sync.RWMutex{}
	wg := &sync.WaitGroup{}

	services := EmailServiceEnums{}
	skippedServices := SkippedServicesEnum{}
	// FIXME EMAIL SERVICE SCANNING
	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
		wg.Add(1)
		go func(i int) {
			// Do something
			service := servicesToCheck[i] // current service
			err, userExsits := service.UserExistsFunc(service, email, config)
			if err != nil {
				log.Printf("error in service: %s,%e", service.Name, err)
				mailMutex.Lock()
				skippedServices[service.Name] = true
				mailMutex.Unlock()
			} else {
				if userExsits { // if service exisits
					log.Printf("User %s on %s exsits", email, service.Name)
					mailMutex.Lock()
					services[service.Name] = EmailServiceEnum{
						Name: service.Name,
						Icon: service.Icon,
						Link: strings.ReplaceAll(service.Url, "{{ email }}", email),
					} // add service to accounts
					mailMutex.Unlock()
				} else {
					log.Printf("User %s on %s does not exsit", email, service.Name)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println(len(services))
	return services, skippedServices
}

func CheckMail(newPerson Person, config ApiConfig) Person { // FIXME TODO
	var mailMutex = sync.RWMutex{}
	fmt.Println(newPerson)
	if newPerson.Email == nil {
		log.Println("nil newPerson.Email")
		newPerson.Email = EmailsType{}
	}
	if len(newPerson.Email) == 0 {
		log.Printf("newPerson.Email (ID=%s) is a empty list", newPerson.ID)
	} else {
		log.Println()
		for i, mail := range newPerson.Email {
			if mail.Mail != "" {
				log.Printf("Checking %s", mail.Mail)
				//mail.Services = MailServices(mail.Mail)
				mail.Valid = IsEmailValid(mail.Mail)
				mail.Gmail = IsGmailAddress(mail.Mail)
				mail.ValidGmail = IsValidGmailAddress(mail.Mail)

				if mail.Services == nil {
					mail.Services = EmailServiceEnums{}
					log.Printf("mail.Services == nil (%s)", mail.Mail)
				}
				// We always want to clear the skipped services
				//if mail.SkippedServices == nil {
				mail.SkippedServices = SkippedServicesEnum{}
				//log.Printf("mail.SkippedServices == nil (%s)", mail.Mail)
				//}
				retMailServices, retSkippedMailServices := MailServicesHandler(DefaultMailServices, mail.Mail, config)
				log.Printf("found %d services", len(retMailServices))
				for key, value := range retMailServices {
					log.Printf("%s = %s", key, value)
					mailMutex.Lock()
					log.Printf("mail.Services[%s] = %#v", key, value)
					mail.Services[key] = value
					mailMutex.Unlock()
				}
				for key, value := range retSkippedMailServices {
					log.Printf("%s = %v", key, value)
					mailMutex.Lock()
					log.Printf("mail.SkippedServices[%s] = %#v", key, value)
					mail.SkippedServices[key] = value
					mailMutex.Unlock()
				}
			} else {
				log.Println("nil mail field")
			}
			log.Printf("%#v", mail)
			newPerson.Email[i] = mail
		}
	}
	return newPerson
}
