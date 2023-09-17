package ips

import (
	"database/sql/driver"
	"encoding/json"
)

type IPs map[string]IP
type IP struct {
	IP string `json:"ip" validate:"ip" tstype:"string"`
}

func (i *IPs) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), i); err != nil {
		return err
	}
	return nil
}
func (i IPs) Value() (driver.Value, error) {
	value, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	return value, nil
}
