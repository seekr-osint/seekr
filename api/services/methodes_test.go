package services

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/seekr-osint/seekr/api/tc"
)

type StringAndErr struct {
	Error  error
	String string
}

func GetUserHtmlUrlHandler(url string) string {
	data := UserServiceDataToCheck{
		User: User{
			Username: "greg",
		},
		Service: Service{
			Name:                "TestService",
			UserHtmlUrlTemplate: url,
		},
	}
	b, err := data.GetUserHtmlUrl()
	if err != nil {
		return ""
	}
	return b
}

func TestNewTest(t *testing.T) {
	test := tc.Test[string, string]{
		Cases: tc.TestCases[string, string]{
			{
				Input:  "github.com/{{.Username}}",
				Expect: "github.com/greg",
				Title:  "Simple username in Url",
			},
		},
		Func: GetUserHtmlUrlHandler,
	}

	test.TcTestHandler(t)
	err := ioutil.WriteFile("doc.md", []byte(fmt.Sprintf("# Service url templating\n\n%s", test.TcTestToMarkdown())), 0644)
	if err != nil {
		panic("Error writing to file:")
		return
	}
}
