package services

import (
	"log"
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
			ExsistingUser:    "greg",
			NotExsistingUser: "greg2q1412fdwkdfns",
		},
	},
	{
		Name:                "TikTok",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		Domain:              "tiktok.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		TestData: TestData{
			ExsistingUser:    "greg",
			NotExsistingUser: "gregdoesnotexsist",
		},
	},
}

func CheckUser(user User) {

}
