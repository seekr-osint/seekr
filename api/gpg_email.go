package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func UbuntuGPGUserExists(mailService MailService, email string, config ApiConfig) (EmailService, error) {
	emailService := EmailService{
		Name: mailService.Name,
		Icon: mailService.Icon,
		Link: strings.ReplaceAll(mailService.Url, "{{ email }}", email),
	}
	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return emailService, nil
		} else if email == "error@gmail.com" {
			return EmailService{}, errors.New("error")
		}
		log.Println("all email testing case false")

		return EmailService{}, nil
	}
	baseUrl := "https://keyserver.ubuntu.com"
	path := fmt.Sprintf("/pks/lookup?search=%s&op=index", email)
	url := baseUrl + path

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return EmailService{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return emailService, nil
	} else {
		return EmailService{}, nil
	}
}
func KeysGnuPGUserExists(mailService MailService, email string, config ApiConfig) (EmailService, error) {
	emailService := EmailService{
		Name: mailService.Name,
		Icon: mailService.Icon,
		Link: strings.ReplaceAll(mailService.Url, "{{ email }}", email),
	}
	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return emailService, nil
		} else if email == "error@gmail.com" {
			return EmailService{}, errors.New("error")
		}
		log.Println("all email testing case false")

		return EmailService{}, nil
	}
	baseUrl := "https://keys.gnupg.net"
	path := fmt.Sprintf("/pks/lookup?search=%s&op=index", email)
	url := baseUrl + path

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return EmailService{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return emailService, nil
	} else {
		return EmailService{}, nil
	}
}
func KeyserverPGPUserExists(service MailService, email string, config ApiConfig) (error, bool) {

	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return nil, true
		} else if email == "error@gmail.com" {
			return errors.New("error"), false
		}
		log.Println("all email testing case false")
		return nil, false
	}
	baseUrl := "https://keyserver.pgp.com"
	path := fmt.Sprintf("/pks/lookup?search=%s&op=index", email)
	url := baseUrl + path

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil, true
	} else {
		return nil, false
	}
}
func PgpMitUserExists(service MailService, email string, config ApiConfig) (error, bool) {

	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return nil, true
		} else if email == "error@gmail.com" {
			return errors.New("error"), false
		}
		log.Println("all email testing case false")
		return nil, false
	}
	baseUrl := "https://pgp.mit.edu"
	path := fmt.Sprintf("/pks/lookup?search=%s&op=index", email)
	url := baseUrl + path

	resp, err := http.Get(url)
	if err != nil {
		return err, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil, true
	} else {
		return nil, false
	}
}
func PoolSKSUserExists(service MailService, email string, config ApiConfig) (error, bool) {

	if config.Testing {
		if email == "all@gmail.com" {
			log.Println("all email testing case true")
			return nil, true
		} else if email == "error@gmail.com" {
			return errors.New("error"), false
		}
		log.Println("all email testing case false")
		return nil, false
	}
	baseUrl := "https://pool.sks-keyservers.net"
	path := fmt.Sprintf("/pks/lookup?search=%s&op=index", email)
	url := baseUrl + path

	resp, err := http.Get(url)
	if err != nil {
		return err, false
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return nil, true
	} else {
		return nil, false
	}
}
