package newtype

import (
	//"fmt"
	"fmt"
	"reflect"
	"strings"
)

func MarkdownGen[T struct{}](t T, intendent int) (string, error) {
	var sb strings.Builder

	//	switch t.(type) {
	//	case string:
	//		sb.WriteString(fmt.Sprintf("\"%s\"", reflect.ValueOf(t).Interface()))
	//	case int:
	//		sb.WriteString(fmt.Sprintf("%d", reflect.ValueOf(t).Interface()))
	//	default:
	//		return "", ErrUnknownType
	//	}
	return sb.String(), nil
}

func Markdown[T interface{}](t T, intendent int) (string, error) {
	var sb strings.Builder

	switch reflect.TypeOf(t).Kind() {
	case reflect.String:
		sb.WriteString(fmt.Sprintf("\"%s\"", reflect.ValueOf(t).Interface()))
	case reflect.Int:
		sb.WriteString(fmt.Sprintf("%d", reflect.ValueOf(t).Interface()))
	default:
		return "", ErrUnknownType
	}
	return sb.String(), nil
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
				return t1, ErrTypeMissmatch
			}
			if field1Value.Kind() != field2Value.Kind() {
				return t1, ErrTypeMissmatch
			}
			merged, err := Merge(field1Value.Interface(), field2Value.Interface())
			if err != nil {
				return t1, err
			}
			reflect.ValueOf(&t1).Elem().FieldByName(field2.Name).Set(reflect.ValueOf(merged))
		}
	}
	return t1, nil
}
func Set[T any](v reflect.Value, t T) error {
	// .Set(reflect.ValueOf(merged))
	if v.Type() != reflect.TypeOf(t) {
		return ErrTypeMissmatch
	}
	if v.CanSet() {
		v.Set(reflect.ValueOf(t))
	}
	return nil
}
