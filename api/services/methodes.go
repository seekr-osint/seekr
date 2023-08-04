package services

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
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

func (data UserServiceDataToCheck) UserExistsFunction() ServiceCheckResult {
	exists, err := data.Service.UserExistsFunc(data)
	return ServiceCheckResult{
		Error:   err,
		Result:  exists,
		Service: data.Service,
		User:    data.User,
	}

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
			Username: service.TestData.ExistingUser,
		},
	}
}

func (service Service) TestUserServiceData2() UserServiceDataToCheck {
	return UserServiceDataToCheck{
		Service: service,
		User: User{
			Username: service.TestData.NotExistingUser,
		},
	}
}

func (user User) GetServices() DataToCheck {
	services := []UserServiceDataToCheck{}
	for _, service := range DefaultServices {
		serviceWithData := UserServiceDataToCheck{
			User:    user,
			Service: service,
		}
		services = append(services, serviceWithData)
	}
	return services
}

func (results ServiceCheckResults) GetFailed() Services {
	services := Services{}
	for _, result := range results {
		if result.Error != nil {
			services = append(services, result.Service)
		}
	}
	return services
}
func (results ServiceCheckResults) GetExisting() Services {
	services := Services{}
	for _, result := range results {
		if result.Result && result.Error == nil {
			services = append(services, result.Service)
		}
	}
	return services
}

func (services Services) List() []string {
	res := []string{}
	for _, service := range services {
		res = append(res, service.Name)
	}
	return res

}

func (user User) String() string {
	return user.Username
}
func (result ServiceCheckResult) String() string {
	return fmt.Sprintf("User: %s\nExists: %t\n", result.User.Username, result.Result)
}

func (results ServiceCheckResults) String() string {
	var sb strings.Builder
	for _, result := range results {
		sb.WriteString(result.String() + "\n")

	}
	return sb.String()
}
func (user User) Scan() ServiceCheckResults {
	return user.GetServices().Scan()
}

func (services DataToCheck) Scan() ServiceCheckResults {
	results := ServiceCheckResults{}
	workers := 10
	s := make(chan UserServiceDataToCheck, workers)
	res := make(chan ServiceCheckResult, workers)
	wg := sync.WaitGroup{}
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go ServicesCheckWorker(s, res, &wg)

	}
	for _, service := range services {
		s <- service
	}
	close(s)
	wg.Wait()
	for i := 0; i < len(services); i++ {
		result := <-res
		results = append(results, result)
	}
	return results
}
