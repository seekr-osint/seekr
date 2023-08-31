package db

import (
	"fmt"
	"reflect"
	"strings"
)

func GenerateCreateTableSQL(data interface{}, tableName string) string {
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Struct {
		return ""
	}

	var columns []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		extra := field.Tag.Get("extra")
		fieldType := field.Tag.Get("type")
		name := field.Tag.Get("genji")
		if name == "" {
			name = field.Name
		}

		column := fmt.Sprintf("%s %s %s", name, fieldType, extra)
		columns = append(columns, column)
	}

	createTableSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n\t%s\n)", tableName, strings.Join(columns, ",\n\t"))

	return createTableSQL
}

func GenerateFieldPointers(template interface{}) string {
	var result []string

	t := reflect.TypeOf(template)
	// v := reflect.ValueOf(template)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("type")

		if tag != "" {
			result = append(result, fmt.Sprintf("&data.%s", field.Name))
		}
	}

	return strings.Join(result, ", ")
}
