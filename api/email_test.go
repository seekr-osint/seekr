package api

import (
	"sync"
	"testing"
)

func TcTestHandler(t *testing.T, testCases []TestCase, testMethode string) { // example TcTestHandler(t,testCases,TestFunction)
	wg := &sync.WaitGroup{}

	for _, tc := range testCases {
		wg.Add(1)
		go func(tc TestCase) {
			e := Email{
        Mail: tc.Input,
      }
      var result bool
      switch testMethode {
    case "IsValidEmail":
      result = e.IsValidEmail()
    case "IsValidGmailAddress":
      result = e.IsValidGmailAddress()
    case "IsGmailAddress":
      result = e.IsGmailAddress()
      } 
			if result != tc.expect {
				t.Errorf("Expected %t for %s, got %t", tc.expect, tc.Input, result)
			}
			wg.Done()
		}(tc)
	}
	wg.Wait()
}

func TestIsEmailValid(t *testing.T) {
	// Test cases
	testCases := []TestCase{
		{"user@example.com", true},
		{"user@example.ru", true},
		{"user@fbi.gov", true},
		{"user@subdomain.fbi.gov", true},
		{"user.name@example.com", true},
		{"user_name@example.com", true},
		{"user-name@example.com", true},
		{"user+name@example.com", true},
		{"user@subdomain.example.com", true},
		{"user@123.example.com", true},
		{"user@", false},
		{"@example.com", false},
		{"user@.com", false},
		{"user@example.", false},
		{"user@example.c", false},
		{"user@example.c@m", false},
	}
	TcTestHandler(t, testCases, "IsValidEmail")
}

type TestCase struct {
	Input  string
	expect bool
}

func TestIsGmailAddress(t *testing.T) {
	// Test cases
	testCases := []TestCase{
		{"user@gmail.com", true},
		{"user.name@gmail.com", true},
		{"user_name@gmail.com", true},
		{"user-name@gmail.com", true},
		{"user@googlemail.com", false},
		{"user@example.com", false},
		{"user@gmail.co.uk", false},
		{"user@gmail.", false},
		{"@gmail.com", false},
	}

	TcTestHandler(t, testCases, "IsGmailAddress")
}

func TestIsValidGmailAddress(t *testing.T) {
	// Test cases
	testCases := []TestCase{
		{"user@gmail.com", true},
		{"user.name@gmail.com", true},
		{"user.name1@gmail.com", true},
		{"user.1name@gmail.com", true},
		{"user-name@gmail.com", false},
		{"user_name@gmail.com", false},
		{"user.name@googlemail.com", false},
		{"user.name@example.com", false},
		{"user.name@gmail.co.uk", false},
		{"user.name@gmail.", false},
		{"@gmail.com", false},
		{"user.name@gmail.c", false},
		{"user.name@gmail.co", false},
		{"user.name@gmail.c1", false},
		{"user.name@gmail.1com", false},
	}

	TcTestHandler(t, testCases, "IsValidGmailAddress")
}

func TestIsGitHubMail(t *testing.T) { // FIXME methode
	// Test cases
	testCases := []TestCase{
		{"user@gmail.com", false},
		{"user.name@gmail.com", false},
		{"user.name1@gmail.com", false},
		{"user.1name@gmail.com", false},
		{"user-name@gmail.com", false},
		{"user_name@gmail.com", false},
		{"user.name@googlemail.com", false},
		{"user.name@example.com", false},
		{"user.name@gmail.co.uk", false},
		{"user.name@gmail.", false},
		{"@gmail.com", false},
		{"user.name@gmail.c", false},
		{"user.name@gmail.co", false},
		{"user.name@gmail.c1", false},
		{"user.name@gmail.1com", false},
		{"67828948+9glenda@users.noreply.github.com", true},
	}
	// Loop through test cases
	for _, tc := range testCases {
		result := IsGitHubMail(tc.Input)
		if result != tc.expect {
			t.Errorf("Expected %t for %s, got %t", tc.expect, tc.Input, result)
		}
	}
}
