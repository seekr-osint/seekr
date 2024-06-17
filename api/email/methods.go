package email

import (
	"database/sql/driver"
	"encoding/json"
)

// db

func (e *Emails) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), e); err != nil {
		return err
	}

	return nil
}
func (e Emails) Value() (driver.Value, error) {
	value, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return value, nil
}
