package services

func StatusCodeUserExistsFunc(data UserServiceDataToCheck) (bool, error) {
	return data.StatusCodeUserExistsFunc()
}

var DefaultServices = Services{
	{
		Name:                "GitHub",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		Domain:              "github.com",
		UserHtmlUrlTemplate: "{{.Domain}}/{{.Username}}",
		TestData: TestData{
			ExsistingUser: "greg",
		},
	},
	{
		Name:                "TikTok",
		UserExistsFunc:      StatusCodeUserExistsFunc,
		Domain:              "tiktok.com",
		UserHtmlUrlTemplate: "{{.Domain}}/@{{.Username}}",
		TestData: TestData{
			ExsistingUser: "greg",
		},
	},
}

func CheckUser(user User) {

}
