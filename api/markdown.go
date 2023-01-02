package api

import (
	"fmt"
	"reflect"
)

func GenMD(input person) string {
	md := "## Description of the `person` struct\n\n" +
		"The `person` struct has the following fields:\n\n" +
		"Field | Type\n" +
		"----- | ----\n"

	// Iterate over the fields of the struct and add them to the Markdown string
	v := reflect.ValueOf(input)
	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		fieldType := v.Type().Field(i).Type.String()

		// Only include fields which are not nil or an empty string
		if !v.Field(i).IsNil() && v.Field(i).String() != "" {
			md += fmt.Sprintf("%s | %s\n", fieldName, fieldType)
		}
	}
	return v.String()
}
