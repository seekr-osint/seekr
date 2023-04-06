package api

import (
	"fmt"
	"strings"
)

type Age float64

func (age Age) Markdown() string {
	var sb strings.Builder
	if age.IsValid() && age != 0 {
		sb.WriteString(fmt.Sprintf("- age: `%d`\n", int(age)))
	}
	return sb.String()
}

func (age Age) IsValid() bool {
	return age >= 0 && age <= 128
}
