package api

import (
	"net/http"

	"github.com/seekr-osint/seekr/api/ssn"
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
	person.SSN, err = person.SSN.Parse()
	return person, err
}

func (person Person) ReplaceNil() Person {
	if person.Email == nil {
		person.Email = EmailsType{}
	}
	if person.SSN == nil {
		person.SSN = ssn.SSNs{}
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
	err := person.SSN.Validate()
	if err != nil {
		switch err {
		case ssn.ErrInvalidSSN:
			return APIError{
				Message: "Invalid SSN",
				Status:  http.StatusBadRequest,
			}
		case ssn.ErrKeyMissmatch:
			return APIError{
				Message: "Key Missmatch",
				Status:  http.StatusBadRequest,
			}
		default:
			return err
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
	err = person.Email.Validate()
	if err != nil {
		return err
	}
	err = person.Phone.Validate()
	if err != nil {
		return err
	}
	return nil
}
