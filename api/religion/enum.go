package religion

import "github.com/seekr-osint/seekr/api/enum"

type Religion string

const (
	Christianity Religion = "Christianity"
	Islam        Religion = "Islam"
	Hinduism     Religion = "Hinduism"
	Buddhism     Religion = "Buddhism"
	Sikhism      Religion = "Sikhism"
	Judaism      Religion = "Judaism"
	Other        Religion = "Other"
	NoReligion   Religion = ""
	Atheism      Religion = "Atheism"
)

var Enum = enum.NewEnum(Religion("invalid"), Christianity, Islam, Hinduism, Buddhism, Sikhism, Judaism, Other, NoReligion, Atheism)

func (r Religion) Markdown() string {
	return enum.Markdown(Enum, r)
}

func (r Religion) IsValid() bool {
	return enum.IsValid(Enum, r)
}

func (r Religion) Validate() error {
	return enum.IsValidApiError(Enum, r)
}
