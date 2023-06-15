package services

import (
	"fmt"
	"html/template"
	"strings"
)

func (data UserServiceDataToCheck) GetUserHtmlUrl() (string, error) {
	tmpl, err := template.New("url").Parse(data.Service.UserHtmlUrlTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL template: %w", err)
	}

	user := data.User
	var result strings.Builder
	err = tmpl.Execute(&result, user)
	if err != nil {
		return "", fmt.Errorf("failed to execute URL template: %w", err)
	}

	return result.String(), nil
}
