package gender

import "github.com/seekr-osint/seekr/api/enum"

type Gender string

const (
	Male        Gender = "Male"
	Female      Gender = "Female"
	OtherGender Gender = "Other"
	NoGender    Gender = ""
)

var Enum = enum.NewEnum(Gender("invalid"), Male, Female, OtherGender, NoGender)

func (g Gender) Markdown() string {
	return enum.Markdown(Enum, g)
}

func (g Gender) IsValid() bool {
	return enum.IsValid(Enum, g)
}
