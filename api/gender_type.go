package api

import (
	"fmt"
	"strings"
)

type Gender string

const (
	Male        Gender = "Male"
	Female      Gender = "Female"
	OtherGender        = "Other"
	NoGender    Gender = ""
)

func (g Gender) Markdown() string {
	var sb strings.Builder
	if g.IsValid() && g != NoGender {
		sb.WriteString(fmt.Sprintf("- Gender: `%s`\n", g))
	}
	return sb.String()
}

func (cs Gender) IsValid() bool {
	switch cs {
	case Male, Female, OtherGender, NoGender:
		return true
	}
	return false
}
