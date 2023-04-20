package functions

import (
	"reflect"
	"sort"
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
