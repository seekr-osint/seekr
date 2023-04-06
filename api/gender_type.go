package api

import (
	"fmt"
	"strings"
)

type Gender string

const (
	Male        Gender = "Male"
	Fmale       Gender = "Fmale"
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

func (g Gender) IsValid() bool {
	switch g {
	case Male, Fmale, OtherGender, NoGender:
		return true
	}
	return false
}
