package api

import (
	"testing"
)

func TestParseEmail(t *testing.T) {
	validEmail := Email{Mail: "glenda@gmail.com"}
	parsedEmail := validEmail.Parse()

	if parsedEmail.Provider != "gmail" {
		t.Errorf("Expected provider to be 'gmail', but got '%s'", parsedEmail.Provider)
	}

}
