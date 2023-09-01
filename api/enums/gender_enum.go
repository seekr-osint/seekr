package enums

import (
	"github.com/seekr-osint/seekr/api/enum"
)
type Gender string
type GenderEnum struct{
	Gender enum.Enum[Gender] `json:"gender" tstype:"'Male' | 'Female' | 'Other' | ''" example:"Male"`
}

// func (gender GenderEnum) Value() (driver.Value, error) {
// 	return gender.Gender.Vlaue()
// }

// func (gender GenderEnum) Scan(value interface{}) error {
// 	return gender.Gender.Scan(value)
// }


func (gender Gender) Values() []Gender {
	return []Gender{ "Male", "Female", "Other" }
}

func (gender Gender) NullValue() Gender {
	return ""
}


