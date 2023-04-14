package api

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"github.com/sundowndev/phoneinfoga/v2/lib/number"
	"log"
	"net/http"
	"strings"
)

type PhoneNumber struct {
	Number         string        `json:"number"`
	Valid          bool          `json:"valid"`
	NationalFormat string        `json:"national_format"`
	Tag            string        `json:"tag"`
	Phoneinfoga    number.Number `json:"phoneinfoga"`
}
type PhoneNumbers map[string]PhoneNumber

func (phoneNumber PhoneNumber) GetPhoneinfoga() (number.Number, error) {
	n, err := number.NewNumber(phoneNumber.Number)
	if err != nil {
		return *n, err
	}
	return *n, err
}
func (phoneNumber PhoneNumber) Markdown() string {
	var sb strings.Builder
	if phoneNumber.IsValid() {
		sb.WriteString(fmt.Sprintf("- Phone: `%s`\n", phoneNumber.Number))
	}
	return sb.String()
}

func (phoneNumber PhoneNumber) IsValid() bool {
	parsedNumber, err := phonenumbers.Parse(phoneNumber.Number, "")
	if err != nil {
		log.Printf("error parsing number: %s", err)
		return false
	}
	return phonenumbers.IsValidNumber(parsedNumber)
}

func (phoneNumber PhoneNumber) Parse() PhoneNumber {
	if !phoneNumber.IsValid() && phoneNumber.Number[0] != '+' {
		phoneNumber.Number = "+" + phoneNumber.Number
		if !phoneNumber.IsValid() {
			phoneNumber.Number = phoneNumber.Number[1:]
		}
	}
	phoneNumber.Valid = phoneNumber.IsValid()
	parsedNumber, err := phonenumbers.Parse(phoneNumber.Number, "")
	if err != nil {
		log.Printf("error parsing number: %s", err)
		return phoneNumber
	}
	phoneNumber.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
	phoneNumber.NationalFormat = phonenumbers.Format(parsedNumber, phonenumbers.NATIONAL)
	phoneNumber.Phoneinfoga, err = phoneNumber.GetPhoneinfoga() // FIXME error handeling
	return phoneNumber
}

func (numbers PhoneNumbers) Markdown() string {
	var sb strings.Builder
	for _, number := range SortMapKeys(map[string]PhoneNumber(numbers)) {
		sb.WriteString(numbers[number].Markdown())
	}
	return sb.String()
}
func (numbers PhoneNumbers) Parse() PhoneNumbers {
	newNumbers := PhoneNumbers{}
	for _, number := range SortMapKeys(map[string]PhoneNumber(numbers)) {
		// delete empty emails
		if number == "" {
			delete(numbers, number) // number == ""
			break
		}
		parsedNumber := numbers[number].Parse()
		newNumbers[parsedNumber.Number] = parsedNumber
	}
	return newNumbers
}

func (numbers PhoneNumbers) Validate() error {
	for _, number := range SortMapKeys(map[string]PhoneNumber(numbers)) {
		if number != numbers[number].Number {
			return APIError{
				Message: fmt.Sprintf("Key missmatch: Phone[%s] = %s", number, numbers[number].Number),
				Status:  http.StatusBadRequest,
			}
		}
	}
	return nil
}
