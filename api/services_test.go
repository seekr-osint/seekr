package api

import (
	"testing"
)

func TestUsernames(t *testing.T) {
	return
	username := RandomString(16)
	for i := 0; i < len(DefaultServices); i++ { // loop over all services
		service := DefaultServices[i]                  // current service
		if service.UserExistsFunc(service, username) { // if service exisits
			t.Errorf("This no work %s", service.Name)
		}
	}
}
