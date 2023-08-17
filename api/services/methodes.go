package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"image"
	"image/png"
	"io"
	"time"

	_ "image/jpeg"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"

	// "github.com/seekr-osint/seekr/api/history"
)

func (data UserServiceDataToCheck) GetImagelUrl() (string, error) {
	tmpl, err := template.New("url").Parse(data.Service.UrlTemplates["image"]) // FIXME
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
func (data UserServiceDataToCheck) GetTemplate(templateString string) (string, error) {
	tmpl, err := template.New("url").Parse(templateString)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	user := Template{
		data.User,
		data.Service,
	}
	var result strings.Builder
	err = tmpl.Execute(&result, user)
	if err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	// url, err := SetProtocolURL(result.String(), data.Service.Protocol)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to set the protocol from url: %w", err)
	// }
	// log.Printf("url: %s\n", url)
	return result.String(), nil
}

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
		Errors: Errors{
			Info: err,
		},
		Exists: exists,
		InputData: InputData{
			Service: data.Service,
			User:    data.User,
		},
	}

}

func (data UserServiceDataToCheck) PatternUrlMatchUserExists(patternTemplate string) (bool, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return false, fmt.Errorf("failed to get user HTML URL: %w", err)
	}
	log.Printf("checking service %s for status code: %s\n", data.Service.Name, url)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error making request pattern matching user exsists check: %s", err)
		return false, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	pattern, err := data.GetTemplate(patternTemplate)
	if err != nil {
		return false, fmt.Errorf("failed to get pattern from pattern template: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	log.Printf("checking pattern: %s => %d\n", pattern, strings.Count(string(body), pattern))

	if strings.Count(string(body), pattern) > 0 {
		return true, nil
	}

	return false, nil
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
		if result.Errors.Info != nil {
			services = append(services, result.InputData.Service)
		}
	}
	return services
}
func (results ServiceCheckResults) GetExisting() Services {
	services := Services{}
	for _, result := range results {
		if result.Exists && result.Errors.Info == nil {
			services = append(services, result.InputData.Service)
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

func (result *ServiceCheckResult) GetInfo(data UserServiceDataToCheck) { // FIXME bad code
	if result.Exists {
		if result.Errors.Info != nil {
			result.Info, _ = EmptyInfo(data)
			return
		}
		info, err := data.Service.InfoFunc(data)
		if err != nil {
			result.Errors.Info = err
			return
		}
		result.Info = info
		result.Errors.Info = nil

	}
}
func (result ServiceCheckResult) String() string {
	return fmt.Sprintf("User: %s\nExists: %t\n", result.InputData.User.Username, result.Exists)
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
		service.Service.Parse()
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
func (img *Image) MarshalJSON() ([]byte, error) {
	if img.Img == nil {
		return json.Marshal(nil)
	}
	var buffer bytes.Buffer

	err := png.Encode(&buffer, img.Img)
	if err != nil {
		return []byte{}, err
	}

	// Encode the buffer as base64 and return the resulting string.

	return json.Marshal(base64.StdEncoding.EncodeToString(buffer.Bytes()))
}

// func DecodeImage(imgstr string) (Image, error) {
// 	reader := strings.NewReader(imgstr)
// 	decodedImg, imgType, err := image.Decode(reader)
// 	log.Printf("image type:%s", imgType)
// 	if err != nil {
// 		return Image{}, err
// 	}
// 	return Image{
// 		Img: decodedImg,
// 	}, nil
// }

func (data UserServiceDataToCheck) GetImage() (Image, error) {
	url, err := data.GetImagelUrl()
	if err != nil {
		return Image{}, fmt.Errorf("failed to get Image URL: %w", err)
	}
	return GetImage(url)
}

func GetImage(url string) (Image, error) {
	if url == "" {
		return Image{}, nil
	}
	log.Printf("image: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("error status code check: %s", err)
		return Image{}, fmt.Errorf("failed to send GET request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Image{}, nil // FIXME error handeling
	}
	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return Image{}, fmt.Errorf("failed to decode image: %w", err)
	}
	return Image{
		Img:  img,
		Url:  url,
		Date: time.Now(),
	}, nil
}
func (info *AccountInfo) GetProfilePicture(url string) error {
	pfp, err := GetImage(url)
	if err != nil {
		return err
	}
	info.ProfilePicture.AddOrUpdateLatestItem(pfp)
	return nil
}

func (service *Service) Parse() {
	if service.InfoFunc == nil {
		service.InfoFunc = EmptyInfo
	}
}

func (s1 *ServiceCheckResult) Merge(s2 ServiceCheckResult) {
	s1.Info.Bio.Merge(s2.Info.Bio)
	s1.Info.ProfilePicture.Merge(s2.Info.ProfilePicture)
  // s1Value := reflect.ValueOf(s1.Info)
  // s2Value := reflect.ValueOf(s2.Info)
	
  //   field1 := s1Value.Field(i)
  //   field2 := s2Value.Field(i)
  //   // if field1.Type() == reflect.TypeOf(history.History[any]{}) {
		// fmt.Println(field1.NumMethod())
		// fmt.Println(field1.Convert(reflect.TypeOf(field1)))

		// fmt.Println(field2.NumMethod())
		// fmt.Println(field2)

  //     parseMethod := field1.MethodByName("Merge")
  //     if parseMethod.IsValid() {
				// fmt.Printf("Merge on %s",field1.String())
  //       arguments := []reflect.Value{field2}
  //       parseMethod.Call(arguments)
  //     } else {

				// fmt.Printf("No Merge on %s",field1.String())
				// log.Println(field1.String())
			// }
  //   // }
  // }
}
