package sources

import (
	"database/sql/driver"
	"encoding/json"
)


type Sources map[string]Source
type Source struct {
	URL string `json:"url" validate:"url" tstype:"string"`
}

func (s *Sources) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), s); err != nil {
		return err
	}
	return nil
}
func (s Sources) Value() (driver.Value, error) {
	value, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return value, nil
}
