package api

import (
	"log"
	//"log"
	"regexp"

	"github.com/seekr-osint/seekr/api/functions"
	//"strings"
	//"sync"
)

func (es1 EmailServices) Merge(es2 EmailServices) EmailServices {
	// Merge fields one by one
	merged := EmailServices{}
	for name, obj := range es1 {
		merged[name] = obj
	}
	// overwriting the Services
	for name, obj := range es2 {
		merged[name] = obj
	}
	return merged
}

func (emailService1 EmailService) Merge(emailService2 EmailService) EmailService {
	// Merge fields one by one
	merged := EmailService{
		Name:     emailService1.Name,
		Link:     emailService1.Link,
		Username: emailService1.Username,
		Icon:     emailService1.Icon,
	}
	if emailService2.Name != "" {
		merged.Name = emailService2.Name
	}
	if emailService2.Link != "" {
		merged.Link = emailService2.Link
	}
	if emailService2.Username != "" {
		merged.Username = emailService2.Username
	}
	if emailService2.Icon != "" {
		merged.Icon = emailService2.Icon
	}
	return merged
}

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
		if oldEmailService, ok := email.Services[emailService.Name]; ok {
			emailServices[emailService.Name] = emailService.Merge(oldEmailService) // Merging and prfering the new service
		} else {
			emailServices[emailService.Name] = emailService
		}
	}

	return emailServices, skippedServices
}

func (email Email) CheckMail(config ApiConfig) (Email, error) {
	emailServices, emailSkippedServices := email.Parse().GetExistingEmailServices(DefaultMailServices, config)
	email.SkippedServices = emailSkippedServices
	email.Services = email.Services.Merge(emailServices)
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
