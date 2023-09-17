package enums

// Auto generated type definitions for enums
// DO NOT EDIT
import (
	"github.com/seekr-osint/seekr/api/enum"
)

// Type used to store the Enum in GenderEnum
type Gender string

// Enum type used by enum package.
type GenderEnum struct {
	Gender enum.Enum[Gender] `json:"gender" tstype:"'Male' | 'Female' | 'Other' | ''" example:"Male"`
}

// func (gender GenderEnum) Value() (driver.Value, error) {
// 	return gender.Gender.Vlaue()
// }

// func (gender GenderEnum) Scan(value interface{}) error {
// 	return gender.Gender.Scan(value)
// }

// returning all valid values for the enum.
// Used by the enum package.
func (gender Gender) Values() []Gender {
	return []Gender{"Male", "Female", "Other"}
}

// returning The NullValue for the enum.
// Used by the enum package.
func (gender Gender) NullValue() Gender {
	return ""
}
