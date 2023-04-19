package ssn

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrKeyMissmatch         = errors.New("key missmatch")
	ErrInvalidSSN           = errors.New("invalid ssn")
	ErrPrivateFiledsMissing = errors.New("private fileds missing")
	ErrStateNotFound        = errors.New("state not found")
	ErrInvalidAreaNumber    = errors.New("invalid areaNumber")
)

// Validation
func (ssns SSNs) Validate() error {
	for ssnStr, ssn := range ssns {
		if ssnStr != ssn.SSN {
			return ErrKeyMissmatch
		}
		if ssn.SSN != "" {
			if !ssn.IsValid() {
				return ErrInvalidSSN
			}
		} else {
			delete(ssns, ssnStr)
		}
	}
	return nil
}

func (ssn SSN) IsValid() bool {
	// regular expression for SSN pattern
	ssnPattern := `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	match, _ := regexp.MatchString(ssnPattern, ssn.SSN)
	if !match {
		return false
	}

	// valid area numbers

	area, err := strconv.Atoi(ssn.SSN[0:3])
	if err != nil {
		return false
	}
	if area <= 0 || area >= 1000 {
		return false
	}

	return match
}

// Parsing

func (ssns SSNs) Parse() (SSNs, error) {
	err := ssns.Validate()
	if err != nil {
		return ssns, err
	}
	for ssnStr, ssn := range ssns {
		err := ssn.parseSSN()
		if err != nil {
			return ssns, err
		}
		ssns[ssnStr] = ssn
	}
	return ssns, nil
}

func (ssn *SSN) Parse() error {
	if !ssn.IsValid() {
		return ErrInvalidSSN
	}
	err := ssn.parseSSN()
	if err != nil {
		return err
	}
	err = ssn.formatSSN()
	if err != nil {
		return err
	}
	return nil
}

// Private

func (ssn *SSN) parseSSN() error {
	if !ssn.IsValid() {
		return ErrInvalidSSN
	}
	// Remove any non-digit characters from SSN string
	ssnDigits := regexp.MustCompile(`\D`).ReplaceAllString(ssn.SSN, "")
	if len(ssnDigits) != 9 {
		return ErrInvalidSSN
	}

	// Parse area, group, and serial numbers from SSN string
	ssnArea, err := strconv.Atoi(ssnDigits[0:3])
	if err != nil {
		return err
	}
	ssnGroup, err := strconv.Atoi(ssnDigits[3:5])
	if err != nil {
		return err
	}
	ssnSerial, err := strconv.Atoi(ssnDigits[5:9])
	if err != nil {
		return err
	}

	// Set areaNumber, groupNumber, and serialNumber fields
	ssn.areaNumber = ssnArea
	ssn.groupNumber = ssnGroup
	ssn.serialNumber = ssnSerial
	return nil
}

// Format

func (ssn *SSN) formatSSN() error {
	err := ssn.parseSSN()
	if err != nil {
		return err
	}
	if ssn.areaNumber+ssn.groupNumber+ssn.serialNumber == 0 {
		return ErrPrivateFiledsMissing
	}

	if ssn.areaNumber == 0 {
		return ErrInvalidAreaNumber
	}
	//If ssn.areaNumber >= 900 && ssn.areaNumber <= 999 {
	//	return ErrInvalidAreaNumber
	//}
	if ssn.areaNumber > 999 {
		return ErrInvalidAreaNumber
	}

	areaStr := fmt.Sprintf("%03d", ssn.areaNumber)
	groupStr := fmt.Sprintf("%02d", ssn.groupNumber)
	serialStr := fmt.Sprintf("%04d", ssn.serialNumber)

	if len(areaStr) != 3 || len(groupStr) != 2 || len(serialStr) != 4 {
		return ErrInvalidSSN
	}

	ssn.SSN = fmt.Sprintf("%s-%s-%s", areaStr, groupStr, serialStr)
	return nil
}
