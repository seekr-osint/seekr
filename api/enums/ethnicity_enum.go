package enums

import (
	"github.com/seekr-osint/seekr/api/enum"
)

type Ethnicity string
type EthnicityEnum struct {
	Ethnicity enum.Enum[Ethnicity] `json:"ethnicity" tstype:"'African' | 'Asian' | 'Caucasian/White' | 'Hispanic/Latino' | 'Indigenous/Native American' | 'Multiracial/Mixed' | ''" example:"African"`
}

// func (ethnicity EthnicityEnum) Value() (driver.Value, error) {
// 	return ethnicity.Ethnicity.Vlaue()
// }

// func (ethnicity EthnicityEnum) Scan(value interface{}) error {
// 	return ethnicity.Ethnicity.Scan(value)
// }

func (ethnicity Ethnicity) Values() []Ethnicity {
	return []Ethnicity{"African", "Asian", "Caucasian/White", "Hispanic/Latino", "Indigenous/Native American", "Multiracial/Mixed"}
}

func (ethnicity Ethnicity) NullValue() Ethnicity {
	return ""
}
