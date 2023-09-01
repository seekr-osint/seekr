package clubs

import (
	"database/sql/driver"
	"encoding/json"
)

type Clubs map[string]Club
type Club struct {
	Club string `json:"club" tstype:"string"`
}


func (c *Clubs) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), c); err != nil {
		return err
	}
	return nil
}
func (c Clubs) Value() (driver.Value, error) {
	value, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return value, nil
}

