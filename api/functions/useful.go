package functions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func SortMapKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func DeleteEmptyKey[T any](m map[string]T) map[string]T {
	newMap := make(map[string]T)
	for k, v := range m {
		if k != "" {
			newMap[k] = v
		}
	}
	return newMap
}

func FullParseMapRet[T interface{ Parse() (T, error) }](m map[string]T, fieldName string) (map[string]T, error) {
	newMap := make(map[string]T)
	m = DeleteEmptyKey(m)
	for fieldKey, fieldValue := range m {
		if fieldKey != "" {
			// Parse field
			parsed, err := Parse(fieldValue)
			if err != nil {
				return newMap, err
			}

			parsedFieldValue := reflect.ValueOf(parsed).FieldByName(fieldName).String()
			newMap[parsedFieldValue] = parsed
		}
	}
	return newMap, nil
}

func Parse[T interface{ Parse() (T, error) }](t T) (T, error) {
	parsed, err := t.Parse()
	if err != nil {
		return parsed, err
	}
	return parsed, nil
}

func Merge[T interface{}](t1 T, t2 T) (T, error) { // FIXME merge a map too
	switch reflect.TypeOf(t1).Kind() {
	case reflect.String:
		if reflect.ValueOf(t2).String() != "" {
			t1 = t2
		}
	case reflect.Int:
		if reflect.ValueOf(t2).Int() != 0 {
			t1 = t2
		}
	case reflect.Struct:

		for i := 0; i < reflect.TypeOf(t1).NumField(); i++ {
			field1 := reflect.TypeOf(t1).Field(i)
			field2 := reflect.TypeOf(t2).Field(i)
			field1Value := reflect.ValueOf(t1).Field(i)
			field2Value := reflect.ValueOf(t2).Field(i)
			if field1.Name != field2.Name {
				return t1, ErrDifferentTypes
			}
			if field1Value.Kind() != field2Value.Kind() {
				return t1, ErrDifferentTypes
			}
			switch field1.Type.Kind() {
			case reflect.String:
				merged, err := Merge(field1Value.String(), field2Value.String())
				if err != nil {
					return t1, err
				}
				reflect.ValueOf(&t1).Elem().FieldByName(field2.Name).SetString(merged)
			case reflect.Int:
				merged, err := Merge(field1Value.Int(), field2Value.Int())
				if err != nil {
					return t1, err
				}
				reflect.ValueOf(&t1).Elem().FieldByName(field2.Name).SetInt(merged)
			case reflect.Struct:
				merged, err := Merge(field1Value.Interface(), field2Value.Interface())
				if err != nil {
					return t1, err
				}
				reflect.ValueOf(&t1).Elem().FieldByName(field2.Name).Set(reflect.ValueOf(merged))
			}
		}
	}
	return t1, nil
}

func Markdown[T interface{}](t T) (string, error) {
	if reflect.TypeOf(t).Kind() != reflect.Struct {
		return "", ErrOnlyStruct
	}
	var sb strings.Builder

	typ := reflect.TypeOf(t)
	val := reflect.ValueOf(t)

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := val.Field(i)

		if field.Type.Kind() != reflect.Struct {
			sb.WriteString(fmt.Sprintf("- %s: %s\n", field.Name, fieldValue.Interface()))
		} else {
			nestedMarkdown, err := Markdown(fieldValue.Interface())
			if err != nil {
				return "", err
			}
			sb.WriteString(fmt.Sprintf("# %s\n%s", field.Name, nestedMarkdown))
		}
	}

	return sb.String(), nil
}

func MarkdownMap[T interface{ Markdown() (string, error) }](m map[string]T, header string) (string, error) {
	var sb strings.Builder
	if len(m) >= 1 {
		sb.WriteString(fmt.Sprintf("## %s\n", header))
	}
	for _, key := range SortMapKeys(map[string]T(m)) {
		v := m[key]
		sb.WriteString(fmt.Sprintf("### %s\n", key))
		markdown, err := v.Markdown()
		if err != nil {
			sb.WriteString(markdown)
			return sb.String(), err
		}
	}
	return sb.String(), nil
}
func ParsedConfigInterface[T1 interface{}, T2 interface{ Parse(T1) (T2, error) }](t T2, t1 T1) interface{} {
	t, _ = t.Parse(t1)
	return Interface(t)
}

func ParsedInterface[T interface{ Parse() (T, error) }](t T) interface{} {
	t, _ = t.Parse()
	return Interface(t)
}
func Interface[T interface{}](t T) map[string]interface{} {
	var cfg map[string]interface{}
	jsonBytes, _ := json.Marshal(t)
	_ = json.Unmarshal(jsonBytes, &cfg)
	return cfg
}

func SliceToCommaSeparatedList[T comparable](slice []T) string {
	var nonEmptyStrings []string

	for _, val := range slice {
		str := fmt.Sprintf("%v", val)
		if str != "" {
			nonEmptyStrings = append(nonEmptyStrings, str)
		}
	}

	return strings.Join(nonEmptyStrings, ",")
}
