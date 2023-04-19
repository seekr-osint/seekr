package ssn

import (
	"fmt"
	"testing"
)

func FuzzParse(f *testing.F) {
	f.Fuzz(func(t *testing.T, ssnStr string) {
		ssn := SSN{SSN: ssnStr}
		if !ssn.IsValid() {
			t.Skipf("Invalid SSN format: %s", ssnStr)
		}

		err := ssn.Parse()
		if err != nil {
			t.Errorf("Error parsing SSN: %v,%#v", err, ssn)
		}

		if !ssn.IsValid() {
			t.Errorf("Invalid SSN format after parsing: %v", ssn)
		}
	})
}

func TestSSNParse(t *testing.T) {
	validSSNs := []string{}
	for i := 1; i <= 999; i++ {
		for i2 := 1; i2 <= 99; i2 += i + i2 - (i2 / i) + 18 - i { // this gives me a goo pseudo random slice of SSNs
			validSSNs = append(validSSNs, fmt.Sprintf("%03d-%02d-6%02d9", i, i2, i2))
		}
	}
	t.Logf("Testing %d valid SSNs", len(validSSNs))

	// run subtests in parallel
	t.Run("parse and validate", func(t *testing.T) {
		t.Parallel()
		ssns := SSNs{}
		for _, ssnStr := range validSSNs {
			ssn := SSN{SSN: ssnStr}
			ssns[ssn.SSN] = ssn

			// subtest for each SSN
			t.Run(ssnStr, func(t *testing.T) {
				if !ssn.IsValid() {
					t.Fatalf("Invalid SSN format: %s", ssnStr)
				}

				if err := ssn.Parse(); err != nil {
					t.Fatalf("Error parsing SSN: %v, %#v", err, ssn)
				}

				if err := ssn.formatSSN(); err != nil {
					t.Fatalf("Error parsing SSN: %v, %#v", err, ssn)
				}

				if err := ssn.parseSSN(); err != nil {
					t.Fatalf("Error parsing SSN: %v, %#v", err, ssn)
				}

				if !ssn.IsValid() {
					t.Errorf("Invalid SSN format after parsing: %v", ssn)
				}
			})
		}
		invalidSSNs := SSNs{
			"fdsfs": SSN{
				SSN: "991-58-6589",
			},
		}
		t.Run("ErrKeyMissmatch SSNs", func(t *testing.T) {
			err := invalidSSNs.Validate()
			if err != ErrKeyMissmatch {
				t.Errorf("Expected ErrKeyMissmatch got: %v", err)
			}
			_,err = invalidSSNs.Parse()
			if err != ErrKeyMissmatch {
				t.Errorf("Expected ErrKeyMissmatch got: %v", err)
			}
		})

		invalidSSNs = SSNs{
			"fdsfs": SSN{
				SSN: "fdsfs",
			},
		}
		t.Run("ErrInvalidSSN SSNs", func(t *testing.T) {
			err := invalidSSNs.Validate()
			if err != ErrInvalidSSN {
				t.Errorf("Expected ErrInvalidSSN got: %v", err)
			}
			_,err = invalidSSNs.Parse()
			if err != ErrInvalidSSN {
				t.Errorf("Expected ErrInvalidSSN got: %v", err)
			}
		})

		t.Run("SSNs", func(t *testing.T) {
			err := ssns.Validate()
			if err != nil {
				t.Errorf("Invalid SSN format after validating: %v", err)
			}

			ssns, err = ssns.Parse()
			if err != nil {
				t.Errorf("Invalid SSN format after parsing: %v", err)
			}
		})

	})
}
