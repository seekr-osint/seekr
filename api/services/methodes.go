package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func (data UserServiceDataToCheck) GetUserHtmlUrl() (string, error) {
	tmpl, err := template.New("url").Parse(data.Service.UserHtmlUrlTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL template: %w", err)
	}

	user := Template{
		data.User,
		data.Service,
	}
	var result strings.Builder
	err = tmpl.Execute(&result, user)
	if err != nil {
		return "", fmt.Errorf("failed to execute URL template: %w", err)
	}

	url, err := SetProtocolURL(result.String(), data.Service.Protocol)
	if err != nil {
		return "", fmt.Errorf("failed to set the protocol from url: %w", err)
	}
	log.Printf("url: %s\n", url)
	return url, nil
}
func SetProtocolURL(rawURL, protocol string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if protocol != "" {
		parsedURL.Scheme = protocol
	} else if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	} // else don't change the protocol

	return parsedURL.String(), nil
}

func (data UserServiceDataToCheck) StatusCodeUserExistsFunc() (bool, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return false, fmt.Errorf("failed to get user HTML URL: %w", err)
	}
	log.Printf("checking service %s for status code: %s\n", data.Service.Name, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error status code check: %s", err)
		return false, fmt.Errorf("failed to send GET request: %w", err)
	}
	log.Printf("status code for %s (%s): %d \n", data.Service.Name, url, resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func (service Service) TestUserServiceData() UserServiceDataToCheck {
	return UserServiceDataToCheck{
		Service: service,
		User: User{
			Username: service.TestData.ExsistingUser,
		},
	}
}
func (service Service) TestUserServiceData2() UserServiceDataToCheck {
	return UserServiceDataToCheck{
		Service: service,
		User: User{
			Username: service.TestData.NotExsistingUser,
		},
	}
}
