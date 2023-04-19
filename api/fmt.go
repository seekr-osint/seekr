package api

import (
	"fmt"
	"reflect"
	"strings"
)

func PrintTypeTreeRec(t reflect.Type, visited map[reflect.Type]bool) {
	// Avoid printing the same type twice to prevent infinite recursion
	if visited[t] {
		return
	}
	visited[t] = true

	fmt.Printf("%s\n", t.String())

	// Print all fields
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("  - %s (%s)\n", field.Name, field.Type)

		switch field.Type.Kind() {
		case reflect.Struct:
			PrintTypeTreeRec(field.Type, visited)
		case reflect.Map:
			if field.Type.Elem().Kind() == reflect.Struct {
				PrintTypeTreeRec(field.Type.Elem(), visited)
			} else {

				fmt.Printf("%s map[%s]%s\n", field.Type, field.Type.Key(), field.Type.Elem().Kind())
			}
		case reflect.Slice:
			if field.Type.Elem().Kind() == reflect.Struct {
				PrintTypeTreeRec(field.Type.Elem(), visited)
			} else {
				//fmt.Printf("%s map[%s]%s\n",field.Type,field.PkgPa.Key(), field.Type.Elem().Kind())
			}
		default:

			//fmt.Printf("case default: %s",field.Type.Kind())
		}
	}

	// Print all methods
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		cntIn := method.Type.NumIn()
		in := []string{}
		for i := 0; i < cntIn; i++ {
			in = append(in, method.Type.In(i).String())
		}
		cntOut := method.Type.NumOut()
		out := []string{}
		for i := 0; i < cntOut; i++ {
			out = append(out, method.Type.Out(i).String())
		}
		outStr := strings.Join(out, ", ")
		if cntOut > 1 {
			outStr = fmt.Sprintf("(%s)", outStr)
		}
		fmt.Printf("  + func %s(%s) %s\n", method.Name, strings.Join(in, ", "), outStr)
	}

	// Recursively print fields and methods of embedded types
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Anonymous {
			PrintTypeTreeRec(field.Type, visited)
		}
	}
}
