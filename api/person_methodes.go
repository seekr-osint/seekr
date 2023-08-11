package api

import (
	"fmt"
	"reflect"
	"strings"
)

// Markdown

func (person Person) Markdown() string {
	var sb strings.Builder
	if person.Name != "" {
		sb.WriteString(fmt.Sprintf("# %s\n", person.Name))
	} else {
		sb.WriteString(fmt.Sprintf("# ID:%s\n", person.ID))
	}

	  personValue := reflect.ValueOf(person)

    for i := 0; i < personValue.NumField(); i++ {
        field := personValue.Field(i)

        if markdownMethod := field.MethodByName("Markdown"); markdownMethod.IsValid() {
            result := markdownMethod.Call(nil)
            if len(result) > 0 {
                markdownString, ok := result[0].Interface().(string)
                if ok {
                    sb.WriteString(markdownString)
                    //sb.WriteString("\n")
                }
            }
        } else {
            //sb.WriteString(field.Interface())
					
        }
    }
	//sb.WriteString(person.Age.Markdown())
	//sb.WriteString(person.Gender.Markdown())
	//sb.WriteString(person.Ethnicity.Markdown())
	//sb.WriteString(person.Civilstatus.Markdown())
	//sb.WriteString(person.Religion.Markdown())
	//sb.WriteString(person.Phone.Markdown())
	markdown, err := person.Ips.Markdown()
	if err != nil {
		return sb.String()
	}
	sb.WriteString(markdown)
	if len(person.Email) >= 1 {
		sb.WriteString(fmt.Sprintf("## Email\n%s\n", person.Email.Markdown()))
	}
	return sb.String()
}
