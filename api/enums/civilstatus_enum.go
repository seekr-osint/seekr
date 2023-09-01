package enums

import (
	"github.com/seekr-osint/seekr/api/enum"
)

type Civilstatus string
type CivilstatusEnum struct {
	Civilstatus enum.Enum[Civilstatus] `json:"civilstatus" tstype:"'Single' | 'Married' | 'Windowed' | 'Divorced' | 'Seperated' | ''" example:"Single"`
}

// func (civilstatus CivilstatusEnum) Value() (driver.Value, error) {
// 	return civilstatus.Civilstatus.Vlaue()
// }

// func (civilstatus CivilstatusEnum) Scan(value interface{}) error {
// 	return civilstatus.Civilstatus.Scan(value)
// }

func (civilstatus Civilstatus) Values() []Civilstatus {
	return []Civilstatus{"Single", "Married", "Windowed", "Divorced", "Seperated"}
}

func (civilstatus Civilstatus) NullValue() Civilstatus {
	return ""
}
