// Package Used to Validate and generate enums.
package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

// Alias to any.
// Specify which types can be enums.
type ValidEnum interface {
	any
}

// Methods required for enums
type EnumT[T ValidEnum] interface {
	Values() []T
	NullValue() T
}

// Actual enum
type Enum[T EnumT[T]] struct {
	CurrentValue T
}

// checking the Valid values specified in the Values() and NullValue() methods.
func (e Enum[T]) Validate() error {
	if reflect.DeepEqual(e.CurrentValue, e.CurrentValue.NullValue()) {
		return nil
	}
	for _, v := range e.CurrentValue.Values() {
		if reflect.DeepEqual(e.CurrentValue, v) {
			return nil
		}
	}
	return fmt.Errorf("invalid enum value: %v", e.CurrentValue)
}


// Used in db
func (e *Enum[T]) Scan(value interface{}) error {
	if err := json.Unmarshal(value.([]byte), e); err != nil {
		return err
	}

	if err := e.Validate(); err != nil {
		return err
	}
	return nil
}

// Used in db
func (e Enum[T]) Value() (driver.Value, error) {
	if err := e.Validate(); err != nil {
		return "", err
	}
	value, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (e Enum[T]) MarshalJSON() ([]byte, error) {
	if err := e.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal(e.CurrentValue)
}

func (e *Enum[T]) UnmarshalJSON(data []byte) error {
	if err := json.Unmarshal(data, &e.CurrentValue); err != nil {
		return err
	}
	if err := e.Validate(); err != nil {
		return err
	}
	return nil
}

// just return the enum value as a string
func (e Enum[T]) String() string {
	return fmt.Sprintf("%v", e.CurrentValue)
}

func (e Enum[T]) TSValue() T { // error on non string enums
	return e.CurrentValue
}

func (e Enum[T]) TSAssign() string { // error on non string enums
	res := fmt.Sprintf("'%v'", e.CurrentValue.NullValue())
	for _, v := range e.CurrentValue.Values() {
		res = fmt.Sprintf("%s | '%v'", res, v)
	}
	return res
}
