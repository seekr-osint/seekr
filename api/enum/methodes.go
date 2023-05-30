package enum

import (
	"fmt"
	"github.com/seekr-osint/seekr/api/tc"
	"reflect"
	"strings"
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
