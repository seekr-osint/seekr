package services

import (
	"log"
	"net/http"

	//"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/seekr-osint/seekr/api/language"
)

func StatusCodeUserExistsFunc(data UserServiceDataToCheck) (bool, error) {
	url, _ := data.GetUserHtmlUrl()
	log.Printf("status code check:%s\n", url)
	return data.StatusCodeUserExistsFunc()
}

func EmptyInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	return AccountInfo{}, nil
}

var DefaultServices = Services{
	{
		Name:                "GitHub",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            EmptyInfo,
		Domain:              "github.com",
		UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "greg2q1412fdwkdfns",
		},
	},

	{
		Name:                "YouTube",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            YouTubeInfo,
		Domain:              "youtube.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		UrlTemplates: map[string]string{
			"bio": "{{.Domain}}/@{{.Username}}/about",
		},
		TestData: TestData{
			ExistingUser:    "mrbeast",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	{
		Name:                "TikTok",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            TikTokInfo,
		Domain:              "tiktok.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	//{
	//	Name:           "TryHackMe",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "tryhackme.com",
	//	BlocksTor: true,
	//	UserHtmlUrlTemplate: "{{.Domain}}/p/{{.Username}}",
	//	TestData: TestData{
	//		ExistingUser:    "greg",
	//		NotExistingUser: "gregdoesnotexsist",
	//	},
	//},
	{
		Name:                "Npm",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            EmptyInfo,
		Domain:              "npmjs.com",
		UserHtmlUrlTemplate: "{{.Domain}}/~{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	//{
	//	Name:           "chess.com",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "api.chess.com",
	//	UserHtmlUrlTemplate: "{{.Domain}}/pub/player/{{.Username}}",
	//  BlocksTor: true,
	//	TestData: TestData{
	//		ExistingUser:    "danielnaroditsky",
	//		NotExistingUser: "gregdoesnotexsist",
	//	},
	//},
	{
		Name:                "Asciinema",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		InfoFunc:            EmptyInfo,
		Domain:              "asciinema.org",
		UserHtmlUrlTemplate: "{{.Domain}}/~{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "gregdoesnotexsist",
		},
	},
	// blocks tor
	//{
	//	Name:           "Replit",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "replit.com",
	//	UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
	//	TestData: TestData{
	//		ExistingUser:    "greg",
	//		NotExistingUser: "gregdoesnotexsistsfdssfda",
	//	},
	//},

	//{
	//	Name:           "Lichess",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain: "lichess.org",
	//	UserHtmlUrlTemplate: "{{.Domain}}/api/user/{{.Username}}",
	//	BlocksTor: true, // ???
	//	TestData: TestData{
	//		ExistingUser:    "starwars",
	//		NotExistingUser: "gregdoesnotexsist",
	//	},
	//},

	//{
	//	Name:           "Snapchat",
	//	UserExistsFunc: StatusCodeUserExistsFunc,
	//	Domain:        "snapchat.com",
	//	UserHtmlUrlTemplate: "{{.Domain}}/add/{{.Username}}",
	//	TestData: TestData{
	//		ExistingUser:    "greg",
	//		NotExistingUser: "gregdoesnotexsistdsada",
	//	},
	//},
}

func ServicesCheckWorker(s <-chan UserServiceDataToCheck, res chan<- ServiceCheckResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for service := range s {
		status := service.UserExistsFunction()
		status.GetInfo(service)
		res <- status
	}
}
func ScrapeBioTwitterTag(response http.Response) (string, error) {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return "", err
	}

	bio := doc.Find(`meta[name="twitter:description"]`).AttrOr("content", "")
	return bio, nil
}

func YouTubeInfo(data UserServiceDataToCheck) (AccountInfo, error) { // FIXME broken
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url + "/about")
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	bio, err := ScrapeBioTwitterTag(*response)
	if err != nil {
		return AccountInfo{}, err
	}
	return AccountInfo{
		Bio: NewBio(bio),
	}, nil
}

func TikTokInfo(data UserServiceDataToCheck) (AccountInfo, error) {
	url, err := data.GetUserHtmlUrl()
	if err != nil {
		return AccountInfo{}, err
	}
	response, err := http.Get(url)
	if err != nil {
		return AccountInfo{}, err
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return AccountInfo{}, err
	}

	selector := "h2[data-e2e='user-bio']"
	userBioElement := doc.Find(selector)

	userBioText := userBioElement.Text()
	if userBioText == "No bio yet." {
		userBioText = ""
	}
	return AccountInfo{
		Bio: NewBio(userBioText),
	}, nil
}
func NewBio(bio string) Bio { // TODO discord tag regex/username regex (https://www.tiktok.com/@japan)
	if bio != "" {
		return Bio{
			Language: language.DetectLanguage(bio),
			Bio:      bio,
		}
	}
	return Bio{
		Bio: bio,
	}

}
