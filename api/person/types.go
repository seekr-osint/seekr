package person

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/seekr-osint/seekr/api/enums"
	"github.com/seekr-osint/seekr/api/services"

	// "github.com/seekr-osint/seekr/api/services"
	"github.com/seekr-osint/seekr/api/validate"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model     `json:"-" tstype:"-" skip:"yes"`
	ID             uint                           `json:"-" tstype:"-" gorm:"primaryKey"`
	Owner          string                         `json:"-" tstype:"-" validate:"alphanum" skip:"yes"`
	Name           string                         `json:"name" validate:"required" tstype:",required" example:"john"`
	Age            uint                           `json:"age" validate:"number" tstype:"number" example:"13"`
	Maidenname     string                         `json:"maidenname" tstype:"string" example:"greg"`
	Kids           string                         `json:"kids" tstype:"string" example:"no because no wife"`
	Birthday       string                         `json:"bday" tstype:"string" example:"01.01.2001"`
	Address        string                         `json:"address" tstype:"string"`
	Occupation     string                         `json:"occupation" tstype:"string"`
	Prevoccupation string                         `json:"prevoccupation" tstype:"string"`
	Education      string                         `json:"education" tstype:"string"`
	Military       string                         `json:"military" tstype:"string"`
	Pets           string                         `json:"pets" tstype:"string"`
	Legal          string                         `json:"legal" tstype:"string"`
	Political      string                         `json:"political" tstype:"string"`
	Notes          string                         `json:"notes" tstype:"string"`
	Services       services.MapServiceCheckResult `json:"accounts" grom:"embedded"`
	enums.GenderEnum
	enums.EthnicityEnum
	enums.CivilstatusEnum
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
