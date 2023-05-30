package gender

import (
	"fmt"
	"strings"
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
	case Male, Female, OtherGender, NoGender:
		return true
	}
	return false
}
