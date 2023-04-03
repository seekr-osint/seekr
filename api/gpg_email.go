package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func UbuntuGPGUserExists(service MailService, email string, config ApiConfig) (error, bool) {
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
	baseUrl := "https://keyserver.ubuntu.com"
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
func KeysGnuPGUserExists(service MailService, email string, config ApiConfig) (error, bool) {

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
	baseUrl := "https://keys.gnupg.net"
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
