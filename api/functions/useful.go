package functions

import (
	"reflect"
	"sort"
	"strings"
	"fmt"
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
	for k, v := range m {
		if k != "" {
			parsed, err := ParseRet(v)
			if err != nil {
				return newMap, err
			}
			parsedFieldValue := reflect.ValueOf(parsed).FieldByName(fieldName).String()
			newMap[parsedFieldValue] = parsed
		}
	}
	return newMap, nil
}

func ParseRet[T interface{ Parse() (T, error) }](t T) (T, error) {
	parsed, err := t.Parse()
	if err != nil {
		return parsed, err
	}
	return parsed, nil
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

func MarkdownMap[T interface{ Markdown() (string, error) }](m map[string]T) (string,error) {
  var sb strings.Builder
  for _, key:= range SortMapKeys(map[string]T(m)) {
    v := m[key]
    sb.WriteString(fmt.Sprintf("### %s\n", key))
		markdown, err := v.Markdown()
		if err != nil {
			return sb.String(),err
		}
    sb.WriteString(markdown)
    sb.WriteString("\n")
  }
  return sb.String(), nil
}
