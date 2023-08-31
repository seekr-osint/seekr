package person

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/seekr-osint/seekr/api/enum"
	"github.com/seekr-osint/seekr/api/enums"
	"github.com/seekr-osint/seekr/api/validate"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model `json:"-" tstype:"-" skip:"yes"`
	ID         uint                       `json:"-" tstype:"-" gorm:"primaryKey"`
	Owner      string                     `json:"-" tstype:"-" validate:"alphanum" skip:"yes"`
	Name       string                     `json:"name" validate:"required" tstype:",required" example:"john"`
	Age        uint                       `json:"age" validate:"number" tstype:"number" example:"13"`
	Gender     enum.Enum[enums.Gender]    `json:"gender" tstype:"'male' | 'female' | 'other' | ''" example:"male"`
	Ethnicity  enum.Enum[enums.Ethnicity] `json:"ethnicity" tstype:"'African' | 'Asian' | 'Caucasian/White' | 'Hispanic/Latino' | 'Indigenous/Native American' | 'Multiracial/Mixed'" example:"Asian"`
	Maidenname string                     `json:"maidenname" tstype:"string" example:"greg"`
	Kids       string                     `json:"kids" tstype:"string" example:"no because no wife"`
}

func (p Person) Validate(personValidator *validate.XValidator) error {
	if personValidator == nil {
		personValidator = NewValidator()
	}
	if errs := personValidator.Validate(p); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMsgs, " and "),
		}
	}
	return nil
}
func NewValidator() *validate.XValidator {
	v := &validate.XValidator{
		Validator: validator.New(),
	}
	// v.Validator.RegisterValidation("enum",ValidateValuer,false)
	return v
}

func ValidateValuer(field validator.FieldLevel) bool {
	if valuer, ok := field.Field().Interface().(driver.Valuer); ok {

		_, err := valuer.Value()
		return err == nil
	}

	return false
}
