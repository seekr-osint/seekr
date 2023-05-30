package ethnicity

import "github.com/seekr-osint/seekr/api/enum"

type Ethnicity string

const (
	African                  Ethnicity = "African"
	Asian                    Ethnicity = "Asian"
	CaucasianWhite           Ethnicity = "Caucasian/White"
	HispanicLatino           Ethnicity = "Hispanic/Latino"
	IndigenousNativeAmerican Ethnicity = "Indigenous/Native American"
	MultiracialMixed         Ethnicity = "Multiracial/Mixed"
	NoEthnicity              Ethnicity = ""
)

var Enum = enum.NewEnum(Ethnicity("invalid"), African, Asian, CaucasianWhite, HispanicLatino, IndigenousNativeAmerican, MultiracialMixed, NoEthnicity)

func (e Ethnicity) Markdown() string {
	return enum.Markdown(Enum, e)
}

func (e Ethnicity) IsValid() bool {
	return enum.IsValid(Enum, e)
}
