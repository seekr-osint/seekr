package api

import (
	"testing"
)

func TestIsEmailValid(t *testing.T) {
	// Test cases
	testCases := []struct {
		email  string
		expect bool
	}{
		{"user@example.com", true},
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

	// Loop through test cases
	for _, tc := range testCases {
		result := IsEmailValid(tc.email)
		if result != tc.expect {
			t.Errorf("Expected %t for %s, got %t", tc.expect, tc.email, result)
		}
	}
}

func TestIsGmailAddress(t *testing.T) {
	// Test cases
	testCases := []struct {
		email  string
		expect bool
	}{
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

	// Loop through test cases
	for _, tc := range testCases {
		result := IsGmailAddress(tc.email)
		if result != tc.expect {
			t.Errorf("Expected %t for %s, got %t", tc.expect, tc.email, result)
		}
	}
}
func TestIsValidGmailAddress(t *testing.T) {
	// Test cases
	testCases := []struct {
		email  string
		expect bool
	}{
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
	// Loop through test cases
	for _, tc := range testCases {
		result := IsValidGmailAddress(tc.email)
		if result != tc.expect {
			t.Errorf("Expected %t for %s, got %t", tc.expect, tc.email, result)
		}
	}
}
