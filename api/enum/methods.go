package enum

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/seekr-osint/seekr/api/errortypes"
	"github.com/seekr-osint/seekr/api/functions"
	"github.com/seekr-osint/seekr/api/tc"
)

func TcIsValidTest[T1 comparable](enum Enum[T1]) tc.Test[T1, bool] {
	test := tc.NewEnumIsValidTest(func(v T1) bool { return IsValid(enum, v) }, false, enum.Invalid, true, enum.Values...)
	return test
}

func Markdown[T1 comparable](enum Enum[T1], value T1) string {
	var sb strings.Builder
	if IsValid(enum, value) && fmt.Sprintf("%v", value) != "" {
		sb.WriteString(fmt.Sprintf("- %s: `%v`\n", reflect.TypeOf(value).Name(), value))
	}
	return sb.String()
}

func IsValid[T1 comparable](enum Enum[T1], input T1) bool {
	for _, value := range enum.Values {
		if input == value {
			return true
		}
	}
	return false
}

func IsValidApiError[T1 comparable](enum Enum[T1], input T1) error {
	if !IsValid(enum, input) {
		return errortypes.APIError{
			Message: fmt.Sprintf("Invalid %s", reflect.TypeOf(input).Name()),
			Status:  http.StatusBadRequest,
		}

	}
	return nil
}
func TcRequestValidEnum[T1 comparable](enum Enum[T1], id string, url string, responsePerson map[string]interface{}) tc.Request {
	person := map[string]interface{}{
		"id": id,
	}
	person[strings.ToLower(reflect.TypeOf(enum.Invalid).Name())] = reflect.ValueOf(enum.Values[0]).String()
	responsePerson[strings.ToLower(reflect.TypeOf(enum.Invalid).Name())] = reflect.ValueOf(enum.Values[0]).String()
	responsePerson["id"] = id
	request := tc.Request{
		RequestType:      "POST",
		URL:              url,
		Name:             fmt.Sprintf("Post person (valid %s)", reflect.TypeOf(enum.Invalid).Name()),
		Comment:          fmt.Sprintf("Possible values are: %s", functions.SliceToCommaSeparatedList(enum.Values)),
		PostData:         person,
		ExpectedResponse: responsePerson,
		StatusCode:       201,
	}

	return request
}

func TcRequestInvalidEnum[T1 comparable](enum Enum[T1], url string) tc.Request {
	person := map[string]interface{}{
		"id": "1",
	}
	person[strings.ToLower(reflect.TypeOf(enum.Invalid).Name())] = enum.Invalid
	request := tc.Request{
		RequestType:      "POST",
		URL:              url,
		Name:             fmt.Sprintf("Post person (invalid %s)", reflect.TypeOf(enum.Invalid).Name()),
		Comment:          fmt.Sprintf("Possible values are: %s", functions.SliceToCommaSeparatedList(enum.Values)),
		PostData:         person,
		ExpectedResponse: map[string]interface{}{"message": fmt.Sprintf("Invalid %s", reflect.TypeOf(enum.Invalid).Name())},
		StatusCode:       http.StatusBadRequest,
	}

	return request
}
