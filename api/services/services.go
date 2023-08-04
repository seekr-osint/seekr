package services

import (
	"log"
	"sync"
)

func StatusCodeUserExistsFunc(data UserServiceDataToCheck) (bool, error) {
	url, _ := data.GetUserHtmlUrl()
	log.Printf("status code check:%s\n", url)
	return data.StatusCodeUserExistsFunc()
}

var DefaultServices = Services{
	{
		Name:                "GitHub",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		Domain:              "github.com",
		UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
		TestData: TestData{
			ExistingUser:    "greg",
			NotExistingUser: "greg2q1412fdwkdfns",
		},
	},
	{
		Name:                "TikTok",
		UserExistsFunc:      StatusCodeUserExistsFunc,
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
		Name:           "Npm",
		UserExistsFunc: StatusCodeUserExistsFunc,
		Domain: "npmjs.com",
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
		Name:           "Asciinema",
		UserExistsFunc: StatusCodeUserExistsFunc,
		Domain: "asciinema.org",
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
		res <- status
	}
}
