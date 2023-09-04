package enums

import (
	"github.com/seekr-osint/seekr/api/enum"
)

type Religion string
type ReligionEnum struct {
	Religion enum.Enum[Religion] `json:"religion" tstype:"'Christianity' | 'Islam' | 'Hinduism' | 'Buddhism' | 'Sikhism' | 'Judaism' | 'Other' | 'Atheism' | ''" example:"Christianity"`
}

// func (religion ReligionEnum) Value() (driver.Value, error) {
// 	return religion.Religion.Vlaue()
// }

// func (religion ReligionEnum) Scan(value interface{}) error {
// 	return religion.Religion.Scan(value)
// }

func (religion Religion) Values() []Religion {
	return []Religion{"Christianity", "Islam", "Hinduism", "Buddhism", "Sikhism", "Judaism", "Other", "Atheism"}
}

func (religion Religion) NullValue() Religion {
	return ""
}
