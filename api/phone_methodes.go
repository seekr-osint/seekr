package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/nyaruka/phonenumbers"
	"github.com/seekr-osint/seekr/api/errortypes"
	"github.com/seekr-osint/seekr/api/functions"
	"github.com/sundowndev/phoneinfoga/v2/lib/number"
)

// Markdown
func (phoneNumber PhoneNumber) Markdown() string {
	var sb strings.Builder
	if phoneNumber.IsValid() {
		sb.WriteString(fmt.Sprintf("- Phone: `%s`\n", phoneNumber.Number))
	}
	return sb.String()
}

func (numbers PhoneNumbers) Markdown() string {
	var sb strings.Builder
	for _, number := range functions.SortMapKeys(map[string]PhoneNumber(numbers)) {
		sb.WriteString(numbers[number].Markdown())
	}
	return sb.String()
}

// Validation

func (phoneNumber PhoneNumber) IsValid() bool {
	parsedNumber, err := phonenumbers.Parse(phoneNumber.Number, "")
	if err != nil {
		log.Printf("error parsing number: %s", err)
		return false
	}
	return phonenumbers.IsValidNumber(parsedNumber)
}

func (numbers PhoneNumbers) Validate() error {
	for _, number := range functions.SortMapKeys(map[string]PhoneNumber(numbers)) {
		if number != numbers[number].Number {
			return errortypes.APIError{
				Message: fmt.Sprintf("Key missmatch: Phone[%s] = %s", number, numbers[number].Number),
				Status:  http.StatusBadRequest,
			}
		}
	}
	return nil
}

// Parsing

func (phoneNumber PhoneNumber) Parse() (PhoneNumber, error) {
	if phoneNumber.Number != "" {
		if !phoneNumber.IsValid() && phoneNumber.Number[0] != '+' {
			phoneNumber.Number = "+" + phoneNumber.Number
			if !phoneNumber.IsValid() {
				phoneNumber.Number = phoneNumber.Number[1:]
			}
		}
	}
	phoneNumber.Valid = phoneNumber.IsValid()
	if phoneNumber.Valid {
		parsedNumber, err := phonenumbers.Parse(phoneNumber.Number, "")
		if err != nil {
			log.Printf("error parsing number: %s", err)
			return phoneNumber, err
		}
		phoneNumber.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
		phoneNumber.NationalFormat = phonenumbers.Format(parsedNumber, phonenumbers.NATIONAL)
		phoneNumber.Phoneinfoga, err = phoneNumber.GetPhoneinfoga() // FIXME error handeling
		if err != nil {
			log.Printf("error getting number number: %s", err)
			return phoneNumber, err
		}
	}
	return phoneNumber, nil
}

func (numbers PhoneNumbers) Parse() (PhoneNumbers, error) {
	newNumbers, err := functions.FullParseMapRet(numbers, "Number")
	return newNumbers, err
}

// Info

func (phoneNumber PhoneNumber) GetPhoneinfoga() (number.Number, error) {
	n, err := number.NewNumber(phoneNumber.Number)
	return *n, err
}
