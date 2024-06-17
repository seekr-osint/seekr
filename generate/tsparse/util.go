package tsparse

import (
	"fmt"
	"reflect"
)

type FieldInfo struct {
	FieldName  string
	JSName     string
	Value      any
	Assign     string
	PrettyName string
	Pattern    string
	Type       string
}

func GetFieldNames(v interface{}) []FieldInfo {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	var fields []FieldInfo
	for i := 0; i < t.NumField(); i++ {
		val := value.Field(i)
		field := t.Field(i)
		skip := field.Tag.Get("skip")
		if skip == "" {
			val1 := val.Interface()

			jsName := field.Tag.Get("json")
			if jsName == "" {
				jsName = field.Name
			}

			prettyName := field.Tag.Get("prettyName")
			if prettyName == "" {
				prettyName = field.Name
			}

			pattern := field.Tag.Get("pattern")
			if pattern == "" {
				pattern = ".*"
			}

			t := "text"
			assign := jsName
			fmt.Println(field.Type.Kind())
			if field.Type.Kind() == reflect.Int || field.Type.Kind() == reflect.Uint {
				assign = fmt.Sprintf("parseInt(%s,10)", jsName)
				t = "number"
			} else if field.Type.Kind() == reflect.Struct {
				tsassign := val.MethodByName("TSAssign")
				if tsassign.IsValid() {
					valuesResult := tsassign.Call(nil)
					assign = fmt.Sprintf("%s as %v", jsName, valuesResult[0])
				}

				tsvalue := val.MethodByName("TSValue")
				if tsvalue.IsValid() {
					valuesResult := tsvalue.Call(nil)
					val1 = valuesResult[0]
				}
			}
			fields = append(fields, FieldInfo{
				FieldName:  field.Name,
				JSName:     jsName,
				Assign:     assign,
				Value:      val1,
				PrettyName: prettyName,
				Pattern:    pattern,
				Type:       t,
			})
		}
	}
	return fields
}
