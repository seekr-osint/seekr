package enums

// Auto generated type definitions for enums
// DO NOT EDIT
import (
	"github.com/seekr-osint/seekr/api/enum"
)

// Type used to store the Enum in CivilstatusEnum
type Civilstatus string

// Enum type used by enum package.
type CivilstatusEnum struct {
	Civilstatus enum.Enum[Civilstatus] `json:"civilstatus" tstype:"'Single' | 'Married' | 'Windowed' | 'Divorced' | 'Seperated' | ''" example:"Single"`
}

// func (civilstatus CivilstatusEnum) Value() (driver.Value, error) {
// 	return civilstatus.Civilstatus.Vlaue()
// }

// func (civilstatus CivilstatusEnum) Scan(value interface{}) error {
// 	return civilstatus.Civilstatus.Scan(value)
// }

// returning all valid values for the enum.
// Used by the enum package.
func (civilstatus Civilstatus) Values() []Civilstatus {
	return []Civilstatus{"Single", "Married", "Windowed", "Divorced", "Seperated"}
}

// returning The NillValue for the enum.
// Used by the enum package.
func (civilstatus Civilstatus) NullValue() Civilstatus {
	return ""
}
