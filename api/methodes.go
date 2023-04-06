package api

import (
	"fmt"
	"strings"
)

func (person Person) Markdown() string {
	var sb strings.Builder
	if person.Name != "" {
		sb.WriteString(fmt.Sprintf("# %s\n", person.Name))
	}

	sb.WriteString(person.Gender.Markdown())
	sb.WriteString(person.Age.Markdown())
	sb.WriteString(person.Civilstatus.Markdown())
	sb.WriteString(person.Phone.Markdown())
	sb.WriteString(fmt.Sprintf("## Email\n%s\n", person.Email.Markdown()))
	return sb.String()
}
