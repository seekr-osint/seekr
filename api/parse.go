package api

import (
	"net/http"

	"github.com/seekr-osint/seekr/api/club"
	"github.com/seekr-osint/seekr/api/errortypes"
	"github.com/seekr-osint/seekr/api/hobby"
	"github.com/seekr-osint/seekr/api/ip"
	"github.com/seekr-osint/seekr/api/sources"
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
	person.Hobbies, err = person.Hobbies.Parse()
	if err != nil {
		return person, err
	}
	person.Clubs, err = person.Clubs.Parse()
	if err != nil {
		return person, err
	}
	person.Ips, err = person.Ips.Parse()
	if err != nil {
		return person, err
	}
	person.Sources, err = person.Sources.Parse()
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
	if person.Ips == nil {
		person.Ips = ip.Ips{}
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
		person.Sources = sources.Sources{}
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
	if person.ID == "" {
		return errortypes.APIError{
			Message: "Missing ID",
			Status:  http.StatusBadRequest,
		}
	}

	// enums
	err := person.Religion.Validate()
	if err != nil {
		return err
	}

	err = person.Civilstatus.Validate()
	if err != nil {
		return err
	}

	err = person.Gender.Validate()
	if err != nil {
		return err
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
