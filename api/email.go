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

func (emails EmailsType) CheckMail(config ApiConfig) (EmailsType, error) {
	var (
		err      error
		emailsCh = make(chan Email, len(emails))
	)

	// Start a goroutine for each email to check the mail concurrently.
	for _, email := range SortMapKeys(emails) {
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
//		for _, email := range SortMapKeys(emails) {
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
