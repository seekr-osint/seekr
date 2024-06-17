package enums

// Auto generated type definitions for enums
// DO NOT EDIT
import (
	"github.com/seekr-osint/seekr/api/enum"
)

// Type used to store the Enum in ReligionEnum
type Religion string

// Enum type used by enum package.
type ReligionEnum struct {
	Religion enum.Enum[Religion] `json:"religion" tstype:"'Christianity' | 'Islam' | 'Hinduism' | 'Buddhism' | 'Sikhism' | 'Judaism' | 'Other' | 'Atheism' | ''" example:"Christianity"`
}

// func (religion ReligionEnum) Value() (driver.Value, error) {
// 	return religion.Religion.Vlaue()
// }

// func (religion ReligionEnum) Scan(value interface{}) error {
// 	return religion.Religion.Scan(value)
// }

// returning all valid values for the enum.
// Used by the enum package.
func (religion Religion) Values() []Religion {
	return []Religion{"Christianity", "Islam", "Hinduism", "Buddhism", "Sikhism", "Judaism", "Other", "Atheism"}
}

// returning The NullValue for the enum.
// Used by the enum package.
func (religion Religion) NullValue() Religion {
	return ""
}
