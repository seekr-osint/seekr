package enum_test

import (
	"fmt"

	"github.com/seekr-osint/seekr/api/enum"
)

// Type used to store the Enum in GenderEnum
type Gender string

// Enum type used by enum package.
type GenderEnum struct {
	Gender enum.Enum[Gender] `json:"gender" tstype:"'Male' | 'Female' | 'Other' | ''" example:"Male"`
}

// returning all valid values for the enum.
// Used by the enum package.
func (gender Gender) Values() []Gender {
	return []Gender{"Male", "Female", "Other"}
}

// returning The NullValue for the enum.
// Used by the enum package.
func (gender Gender) NullValue() Gender {
	return ""
}

func ExampleEnum_Validate() {
	g := GenderEnum{
		Gender: enum.Enum[Gender]{"Male"},
	}
	err := g.Gender.Validate()
	fmt.Printf("%v", err)
	// Output:
	// 
	// <nil>

}
