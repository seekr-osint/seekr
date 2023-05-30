package civilstatus

import "github.com/seekr-osint/seekr/api/enum"

type CivilStatus string

const (
	Single        CivilStatus = "Single"
	Married       CivilStatus = "Married"
	Widowed       CivilStatus = "Widowed"
	Divorced      CivilStatus = "Divorced"
	Separated     CivilStatus = "Separated"
	NoCivilStatus CivilStatus = ""
)

var Enum = enum.NewEnum(CivilStatus("invalid"), Single, Married, Widowed, Divorced, Separated, NoCivilStatus)

func (c CivilStatus) Markdown() string {
	return enum.Markdown(Enum, c)
}

func (c CivilStatus) IsValid() bool {
	return enum.IsValid(Enum, c)
}

func (c CivilStatus) Validate() error {
	return enum.IsValidApiError(Enum, c)
}
