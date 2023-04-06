package api

import (
	"fmt"
	"github.com/nyaruka/phonenumbers"
	"log"
	"net/http"
	"strings"
)

type PhoneNumber struct {
	Number string `json:"number"`
	Valid  bool   `json:"valid"`
}
type PhoneNumbers map[string]PhoneNumber

func (number PhoneNumber) Markdown() string {
	var sb strings.Builder
	if number.IsValid() {
		sb.WriteString(fmt.Sprintf("- Phone: `%s`\n", number.Number))
	}
	return sb.String()
}

func (number PhoneNumber) IsValid() bool {
	parsedNumber, err := phonenumbers.Parse(number.Number, "")
	if err != nil {
		log.Printf("error parsing number: %s", err)
		return false
	}
	return phonenumbers.IsValidNumber(parsedNumber)
}

func (number PhoneNumber) Parse() PhoneNumber {
	number.Valid = number.IsValid()
	parsedNumber, err := phonenumbers.Parse(number.Number, "")
	if err != nil {
		log.Printf("error parsing number: %s", err)
		return number
	}
	number.Number = phonenumbers.Format(parsedNumber, phonenumbers.INTERNATIONAL)
	return number
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
