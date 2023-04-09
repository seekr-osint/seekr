package api

import (
	"fmt"
	"strings"
)

type CivilStatus string

const (
	Single        CivilStatus = "Single"
	Married       CivilStatus = "Married"
	Widowed       CivilStatus = "Widowed"
	Divorced      CivilStatus = "Divorced"
	Separated     CivilStatus = "Separated"
	NoCivilStatus CivilStatus = ""
)

func (cs CivilStatus) Markdown() string {
	var sb strings.Builder
	if cs.IsValid() && cs != NoCivilStatus {
		sb.WriteString(fmt.Sprintf("- Civil Status: `%s`\n", cs))
	}
	return sb.String()
}

func (cs CivilStatus) IsValid() bool {
	switch cs {
	case Single, Married, Widowed, Divorced, Separated, NoCivilStatus:
		return true
	}
	return false
}
