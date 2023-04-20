package typetree

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	colorBlack   = "\033[30m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorGray    = "\033[90m"
	colorReset   = "\033[0m"
)

func ColorPrint(color, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, colorReset)
}

func ColorType(t reflect.Kind) string {

	switch t {
	case reflect.String:
		return ColorPrint(colorYellow, t.String())
	case reflect.Int:
		return ColorPrint(colorGreen, t.String())
	case reflect.Interface:
		return ColorPrint(colorCyan, t.String())
	case reflect.Ptr:
		return ColorPrint(colorGray, t.String())
	case reflect.Bool:
		return ColorPrint(colorRed, t.String())
	case reflect.Func:
		return ColorPrint(colorMagenta, t.String())
	case reflect.Map:
		return ColorPrint(colorGreen, t.String())
	case reflect.Slice:
		return ColorPrint(colorYellow, t.String())
	}
	return t.String()
}
func PrintTypeTreeRec(t reflect.Type, visited map[reflect.Type]bool, indentLevel int, initialIndentLevel int, recursion bool) string {
	var sb strings.Builder
	methodes_printed := false

	// Print all fields with indentation
	switch t.Kind() {
	case reflect.Struct:
		// Avoid printing the same type twice to prevent infinite recursion
		if visited[t] {
			return fmt.Sprintf("%s (skipping)\n", t.Name())
		}
		visited[t] = recursion

		// Print type name with indentation
		sb.WriteString(fmt.Sprintf("%s%s {\n", strings.Repeat(" ", initialIndentLevel), ColorPrint(colorCyan, t.String())))
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)

			switch field.Type.Kind() {
			case reflect.Struct:
				sb.WriteString(fmt.Sprintf("%s  %s\t%s", strings.Repeat(" ", indentLevel+2), field.Name, PrintTypeTreeRec(field.Type, visited, indentLevel+4, initialIndentLevel, recursion)))
			case reflect.Map:
				sb.WriteString(fmt.Sprintf("%s%s\t%s[%s]%s", strings.Repeat(" ", indentLevel+4), field.Name, ColorType(field.Type.Kind()), ColorType(field.Type.Key().Kind()), PrintTypeTreeRec(field.Type.Elem(), visited, indentLevel+4, 0, recursion)))
			case reflect.Slice:
				sb.WriteString(fmt.Sprintf("%s%s\t[]%s", strings.Repeat(" ", indentLevel+4), field.Name, PrintTypeTreeRec(field.Type.Elem(), visited, indentLevel+4, 0, recursion)))
			default:
				sb.WriteString(fmt.Sprintf("%s  %s\t\t%s", strings.Repeat(" ", indentLevel+2), field.Name, PrintTypeTreeRec(field.Type, visited, indentLevel+4, 0, recursion)))
				//sb.WriteString(fmt.Sprintf("case default: %s",field.Type.Kind())
			}
		}
		//sb.WriteString(Methodes(t,indentLevel))
		//methodes_printed = true

		// Recursively print fields and methods of embedded types with increased indentation
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			if field.Anonymous {
				sb.WriteString(PrintTypeTreeRec(field.Type, visited, indentLevel+4, 0, recursion))
			}
		}

		// Close the struct with indentation
		sb.WriteString(fmt.Sprintf("%s}\n", strings.Repeat(" ", indentLevel)))
	case reflect.Map:
		sb.WriteString(fmt.Sprintf("%s%s[%s]%s\n", strings.Repeat(" ", initialIndentLevel), ColorType(t.Kind()), ColorType(t.Key().Kind()), PrintTypeTreeRec(t.Elem(), visited, indentLevel, 0, recursion)))
	case reflect.Slice:
		sb.WriteString(fmt.Sprintf("%s[]%s\n", strings.Repeat(" ", initialIndentLevel), PrintTypeTreeRec(t.Elem(), visited, indentLevel, 0, recursion)))
	default:
		sb.WriteString(fmt.Sprintf("%s\n", ColorType(t.Kind())))

	}
	if !methodes_printed {
		sb.WriteString(Methodes(t, indentLevel))
	}

	//sb.WriteString(Methodes(t))
	return sb.String()
}

func Methodes(t reflect.Type, indentLevel int) string {
	var sb strings.Builder
	if t.NumMethod() > 0 {
		sb.WriteString(fmt.Sprintf("%s(%s) {\n", strings.Repeat(" ", indentLevel), t.String()))
	}
	// Print all methods with indentation
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		cntIn := method.Type.NumIn()
		in := []string{}
		for i := 0; i < cntIn; i++ {
			in = append(in, ColorType(method.Type.In(i).Kind()))
		}
		cntOut := method.Type.NumOut()
		out := []string{}
		for i := 0; i < cntOut; i++ {
			out = append(out, ColorType(method.Type.Out(i).Kind()))
		}
		outStr := strings.Join(out, ", ")
		if cntOut > 1 {
			outStr = fmt.Sprintf("(%s)", outStr)
		}

		sb.WriteString(fmt.Sprintf("%s  %s (%s) %s(%s) %s\n", strings.Repeat(" ", indentLevel+2), ColorPrint(colorMagenta, "func"), t.String(), method.Name, strings.Join(in, ", "), outStr))
	}

	if t.NumMethod() > 0 {
		sb.WriteString(fmt.Sprintf("%s}\n\n", strings.Repeat(" ", indentLevel)))
	}
	return sb.String()
}
