package api

import (
	"log"
	//"log"
	"regexp"
	//"strings"
	//"sync"
)

func (email Email) GetExistingEmailServices(mailServices MailServices, apiConfig ApiConfig) (EmailServices, SkippedServices) {
	emailServices := make(EmailServices)
	skippedServices := make(SkippedServices)

	emailServiceChan := make(chan EmailService)
	doneChan := make(chan bool)

	for _, service := range mailServices {
		go func(service MailService) {
			emailService, err := service.UserExistsFunc(service, email.Mail, apiConfig)
			if err != nil {
				log.Printf("Error checking user existence for %s:%s\n", service.Name, err)
				skippedServices[service.Name] = true
			}
			if emailService != (EmailService{}) {
				emailServiceChan <- emailService
			}
			doneChan <- true
		}(service)
	}

	go func() {
		for range mailServices {
			<-doneChan
		}
		close(emailServiceChan)
	}()

	for emailService := range emailServiceChan {
		emailServices[emailService.Name] = emailService
	}

	return emailServices, skippedServices
}

func (email Email) CheckMail(config ApiConfig) (Email, error) {
	email.Services, email.SkippedServices = email.Parse().GetExistingEmailServices(DefaultMailServices, config)
	return email, nil
}

func (emails EmailsType) CheckMail(config ApiConfig) (EmailsType, error) { // FIXME slow
	var err error
	for _, email := range SortMapKeys(emails) {
		emails[email], err = emails[email].CheckMail(config)
	}
	return emails, err
}
func (person Person) CheckMail(config ApiConfig) (Person, error) {
	var err error
	person.Email, err = person.Email.CheckMail(config)
	return person, err
}

var RANDOM_USERNAME = "AAAANSUhEUgAAAgAAAA"
var RANDOM_PASSWORD = "64,iVBORw0KGgoAAAA$1"

func IsGitHubMail(email string) bool {
	match, _ := regexp.MatchString("@users\\.noreply\\.github\\.com$", email)
	return match
}

//func MailServicesHandler(servicesToCheck MailServices, email string, config ApiConfig) (EmailServices, SkippedServices) {
//	var mailMutex = sync.RWMutex{}
//	wg := &sync.WaitGroup{}
//
//	services := EmailServices{}
//	skippedServices := SkippedServices{}
//	// FIXME EMAIL SERVICE SCANNING
//	for i := 0; i < len(servicesToCheck); i++ { // loop over all services
//		wg.Add(1)
//		go func(i int) {
//			// Do something
//			service := servicesToCheck[i] // current service
//			err, userExsits := service.UserExistsFunc(service, email, config)
//			if err != nil {
//				log.Printf("error in service: %s,%e", service.Name, err)
//				mailMutex.Lock()
//				skippedServices[service.Name] = true
//				mailMutex.Unlock()
//			} else {
//				if userExsits { // if service exisits
//					log.Printf("User %s on %s exsits", email, service.Name)
//					mailMutex.Lock()
//					services[service.Name] = EmailService{
//						Name: service.Name,
//						Icon: service.Icon,
//						Link: strings.ReplaceAll(service.Url, "{{ email }}", email),
//					} // add service to accounts
//					mailMutex.Unlock()
//				} else {
//					log.Printf("User %s on %s does not exsit", email, service.Name)
//				}
//			}
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//	log.Println(len(services))
//	return services, skippedServices
//}
//
//func CheckMail2(newPerson Person, config ApiConfig) Person { // FIXME TODO
//	var mailMutex = sync.RWMutex{}
//	fmt.Println(newPerson)
//	if newPerson.Email == nil {
//		log.Println("nil newPerson.Email")
//		newPerson.Email = EmailsType{}
//	}
//	if len(newPerson.Email) == 0 {
//		log.Printf("newPerson.Email (ID=%s) is a empty list", newPerson.ID)
//	} else {
//		log.Println()
//		for i, mail := range newPerson.Email {
//			if mail.Mail != "" {
//				log.Printf("Checking %s", mail.Mail)
//				mail = mail.Parse()
//				// We always want to clear the skipped services
//				//if mail.SkippedServices == nil {
//				mail.SkippedServices = SkippedServices{}
//				//log.Printf("mail.SkippedServices == nil (%s)", mail.Mail)
//				//}
//				retMailServices, retSkippedMailServices := MailServicesHandler(DefaultMailServices, mail.Mail, config)
//				log.Printf("found %d services", len(retMailServices))
//				for key, value := range retMailServices {
//					log.Printf("%s = %s", key, value)
//					mailMutex.Lock()
//					log.Printf("mail.Services[%s] = %#v", key, value)
//					mail.Services[key] = value
//					mailMutex.Unlock()
//				}
//				for key, value := range retSkippedMailServices {
//					log.Printf("%s = %v", key, value)
//					mailMutex.Lock()
//					log.Printf("mail.SkippedServices[%s] = %#v", key, value)
//					mail.SkippedServices[key] = value
//					mailMutex.Unlock()
//				}
//			} else {
//				log.Println("nil mail field")
//			}
//			log.Printf("%#v", mail)
//			newPerson.Email[i] = mail
//		}
//	}
//	return newPerson
//}
