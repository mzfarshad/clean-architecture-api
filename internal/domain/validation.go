package domain

import (
	"github.com/go-playground/validator/v10"
)

// var validate = new(validator.Validate)
var validate *validator.Validate = validator.New()

type validatable interface{ validate() }

type Validatable struct{}

func (Validatable) validate() {}

func Validate(v validatable) error {
	return validate.Struct(v)
}
