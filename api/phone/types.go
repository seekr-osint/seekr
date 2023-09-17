package phone

import (
	"github.com/sundowndev/phoneinfoga/v2/lib/number"
)

type PhoneNumber struct {
	Number         string        `json:"number" tstype:"string"`
	Valid          bool          `json:"valid" tstype:"boolean"`
	NationalFormat string        `json:"national_format" tstype:"string"`
	Tag            string        `json:"tag" tstype:"string"`
	Phoneinfoga    number.Number `json:"phoneinfoga" tstype:"{ Country: string }"`
}
type PhoneNumbers map[string]PhoneNumber
