package accounts

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/seekr-osint/seekr/api/functions"
)

// URLTemplate

func (u URLTemplate) GetURL(user User, accountScanner AccountScanner) (string, error) {
	templateInput := accountScanner.GetURLTemplateInput(user)
	tmpl, err := template.New("url").Parse(string(u))
	if err != nil {
		return "", err
	}
	result := strings.Builder{}
	err = tmpl.Execute(&result, templateInput)
	if err != nil {
		return "", err
	}
	url, err := SetProtocolURL(result.String(), accountScanner.Protocol)
	if err != nil {
		return "", fmt.Errorf("failed to set the protocol from url: %w", err)
	}
	return url, nil
}

// AccountScanner

func (a AccountScanner) GetURL(templateName string, user User) (string, error) {
	tmpl, ok := a.URLTemplates[templateName]
	if !ok {
		return "", fmt.Errorf("no such template %s", templateName)
	}
	return tmpl.GetURL(user, a)
}

func (a AccountScanner) UserExistsCheckInput(user User) (*UserExistsCheckInput, error) {
	urls, err := a.GetURLsMap(user)
	if err != nil {
		return nil, err
	}
	userExistsCheckInput := UserExistsCheckInput{
		URLs: *urls,
	}
	return &userExistsCheckInput, nil
}

func (a AccountScanner) GetURLsMap(user User) (*URLs, error) {
	urls := URLs{}
	for _, tmpl := range functions.SortMapKeys(a.URLTemplates) {
		url, err := a.URLTemplates[tmpl].GetURL(user, a)
		if err != nil {
			return nil, err
		}
		urls[tmpl] = url
	}
	return &urls, nil
}

func (a AccountScanner) GetURLTemplateInput(user User) *URLTemplateInput {
	return &URLTemplateInput{AccountScanner: a, User: user}
}

func (a AccountScanner) RunScannerDefaultAccountResult(user User) (*ScanResult[DefaultAccount], error) {
	res := ScanResult[DefaultAccount]{
		Account: &DefaultAccount{ // FIXME add url
			Name: a.Name,
		},
	}

	userExistsCheckInput, err := a.UserExistsCheckInput(user)
	if err != nil {
		return nil, err
	}

	res.Exists, res.RateLimited, res.Errors.UserExistsCheck = userExistsCheckInput.RunUserExistsCheck(&a)
	if res.Errors.UserExistsCheck != nil {
		return &res, nil
	}
	return &res, nil
}

func (u UserExistsCheckInput) RunUserExistsCheck(accountScanner *AccountScanner) (bool, bool, error) {
	tmpl, err := template.New("url").Funcs(template.FuncMap{
		"status": StatusCode,
	}).Parse(string(accountScanner.UserExistsCheck))
	if err != nil {
		return false, false, err
	}
	result := strings.Builder{}
	err = tmpl.Execute(&result, u)
	if err != nil {
		return false, false, err
	}
	return ParseCheckResult(result.String())
}
