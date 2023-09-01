package hobbies

import (
	"database/sql/driver"
	"encoding/json"
)

type Hobbies map[string]Hobby
type Hobby struct {
	Hobby string `json:"hobby" tstype:"string"`
}
func (h *Hobbies) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), h); err != nil {
		return err
	}
	return nil
}
func (h Hobbies) Value() (driver.Value, error) {
	value, err := json.Marshal(h)
	if err != nil {
		return nil, err
	}
	return value, nil
}
