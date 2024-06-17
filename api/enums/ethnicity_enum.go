package enums

// Auto generated type definitions for enums
// DO NOT EDIT
import (
	"github.com/seekr-osint/seekr/api/enum"
)

// Type used to store the Enum in EthnicityEnum
type Ethnicity string

// Enum type used by enum package.
type EthnicityEnum struct {
	Ethnicity enum.Enum[Ethnicity] `json:"ethnicity" tstype:"'African' | 'Asian' | 'Caucasian/White' | 'Hispanic/Latino' | 'Indigenous/Native American' | 'Multiracial/Mixed' | ''" example:"African"`
}

// func (ethnicity EthnicityEnum) Value() (driver.Value, error) {
// 	return ethnicity.Ethnicity.Vlaue()
// }

// func (ethnicity EthnicityEnum) Scan(value interface{}) error {
// 	return ethnicity.Ethnicity.Scan(value)
// }

// returning all valid values for the enum.
// Used by the enum package.
func (ethnicity Ethnicity) Values() []Ethnicity {
	return []Ethnicity{"African", "Asian", "Caucasian/White", "Hispanic/Latino", "Indigenous/Native American", "Multiracial/Mixed"}
}

// returning The NullValue for the enum.
// Used by the enum package.
func (ethnicity Ethnicity) NullValue() Ethnicity {
	return ""
}
