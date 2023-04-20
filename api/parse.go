package api

import (
	"net/http"

	"github.com/seekr-osint/seekr/api/club"
	"github.com/seekr-osint/seekr/api/hobby"
	"github.com/seekr-osint/seekr/api/errortypes"
)

func (person Person) Parse(config ApiConfig) (Person, error) { // TODO error handeling and Validate person
	var err error
	person = person.ReplaceNil()
	person.Phone, err = person.Phone.Parse()
	if err != nil {
		return person, err
	}
	person.Email = person.Email.Parse()
	person, err = person.CheckMail(config)
	if err != nil {
		return person, err
	}
	return person, err
}

func (person Person) ReplaceNil() Person {
	if person.Email == nil {
		person.Email = EmailsType{}
	}
	if person.Hobbies == nil {
		person.Hobbies = hobby.Hobbies{}
	}
	if person.Clubs == nil {
		person.Clubs = club.Clubs{}
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
	if person.Phone == nil {
		person.Phone = PhoneNumbers{}
	}
	return person
}

func (person Person) Validate() error {
	if !person.Civilstatus.IsValid() {
		return errortypes.APIError{
			Message: "Invalid civil status",
			Status:  http.StatusBadRequest,
		}
	}
	if !person.Religion.IsValid() {
		return errortypes.APIError{
			Message: "Invalid religion",
			Status:  http.StatusBadRequest,
		}
	}

	if person.ID == "" {
		return errortypes.APIError{
			Message: "Missing ID",
			Status:  http.StatusBadRequest,
		}
	}
	if !person.Gender.IsValid() {
		return errortypes.APIError{
			Message: "Invalid gender",
			Status:  http.StatusBadRequest,
		}
	}
	err := person.Email.Validate()
	if err != nil {
		return err
	}
	err = person.Phone.Validate()
	if err != nil {
		return err
	}
	return nil
}
