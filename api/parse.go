package api

import (
	//"errors"
	"net/http"
)

func (person Person) Parse(config ApiConfig) (Person, error) { // TODO error handeling and Validate person
	person = person.ReplaceNil()
	person.Email = person.Email.Parse()
	person, err := person.CheckMail(config)
	return person, err
}

func (person Person) ReplaceNil() Person {
	if person.Email == nil {
		person.Email = EmailsType{}
	}
	if person.Pictures == nil {
		person.Pictures = Pictures{}
	}
	if person.Accounts == nil {
		person.Accounts = Accounts{}
	}
	if person.Tags == nil {
		person.Tags = Tags{}
	}
	if person.Sources == nil {
		person.Sources = Sources{}
	}
	if person.Relations == nil {
		person.Relations = Relation{}
	}
	return person
}

func (person Person) Validate() error {
	if !person.Civilstatus.IsValid() {
		return APIError{
			Message: "Invalid civil status",
			Status:  http.StatusBadRequest,
		}
	}
	if !person.Religion.IsValid() {
		return APIError{
			Message: "Invalid religion",
			Status:  http.StatusBadRequest,
		}
	}
	if !person.SSN.IsValid() {
		return APIError{
			Message: "Invalid SSN",
			Status:  http.StatusBadRequest,
		}
	}
	if person.ID == "" {
		return APIError{
			Message: "Missing ID",
			Status:  http.StatusBadRequest,
		}
	}
	if !person.Gender.IsValid() {
		return APIError{
			Message: "Invalid gender",
			Status:  http.StatusBadRequest,
		}
	}
	err := person.Email.Validate()
	if err != nil {
		return person.Email.Validate()
	}
	return nil
}
