package api

import (
	"log"
	//"log"
	"regexp"

	"github.com/seekr-osint/seekr/api/functions"
	//"strings"
	//"sync"
)

func (ess1 EmailServices) Merge(ess2 EmailServices) (EmailServices, error) { // FIXME don't merge just return ess2
	var err error
	// Merge fields one by one
	//merged := EmailServices{}
	//for name, es1 := range ess1 {
	//	merged[name],err = es1.Merge(ess2[name])
	//	//merged[name], err = functions.Merge(es1, ess2[name])
	//	if err != nil {
	//		return merged, err
	//	}
	//}
	return ess2, err
}

func (emailService1 EmailService) Merge(emailService2 EmailService) (EmailService, error) {

	return functions.Merge(emailService1, emailService2)
}

func (email Email) GetExistingEmailServices(mailServices MailServices, apiConfig ApiConfig) (EmailServices, SkippedServices, error) {
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
	var err error
	for emailService := range emailServiceChan {
		if oldEmailService, ok := email.Services[emailService.Name]; ok {
			emailServices[emailService.Name], err = oldEmailService.Merge(emailService) // Merging and prfering the new service
			if err != nil {
				return emailServices, skippedServices, err
			}
		} else {
			emailServices[emailService.Name] = emailService
		}
	}

	return emailServices, skippedServices, nil
}

func (email Email) CheckMail(config ApiConfig) (Email, error) {
	var err error
	emailServices, emailSkippedServices, err := email.Parse().GetExistingEmailServices(DefaultMailServices, config)
	if err != nil {
		return email, err
	}
	email.SkippedServices = emailSkippedServices
	email.Services, err = email.Services.Merge(emailServices)
	if err != nil {
		return email, err
	}
	return email, nil
}

func (emails EmailsType) CheckMail(config ApiConfig) (EmailsType, error) {
	var (
		err      error
		emailsCh = make(chan Email, len(emails))
	)

	// Start a goroutine for each email to check the mail concurrently.
	for _, email := range functions.SortMapKeys(emails) {
		go func(email Email, config ApiConfig) {
			email, err := email.CheckMail(config)
			if err != nil {
				// No error handeling here
				log.Println(err)
			} else {
				emailsCh <- email
			}
		}(emails[email], config)
	}

	// Collect the results from the goroutines and update the emails map.
	for range emails {
		email := <-emailsCh

		emails[email.Mail] = email

	}

	close(emailsCh)

	return emails, err
}

// func (emails EmailsType) CheckMail(config ApiConfig) (EmailsType, error) { // FIXME slow
//
//		var err error
//		for _, email := range functions.SortMapKeys(emails) {
//			emails[email], err = emails[email].CheckMail(config)
//		}
//		return emails, err
//	}
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
