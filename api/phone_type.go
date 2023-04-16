package api

import (
	"github.com/sundowndev/phoneinfoga/v2/lib/number"
)

type PhoneNumber struct {
	Number         string        `json:"number"`
	Valid          bool          `json:"valid"`
	NationalFormat string        `json:"national_format"`
	Tag            string        `json:"tag"`
	Phoneinfoga    number.Number `json:"phoneinfoga"`
}
type PhoneNumbers map[string]PhoneNumber
