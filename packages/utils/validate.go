package utils

import (
	"github.com/go-playground/validator"
)

func ValidateStruct(b interface{}) error {
	validate := validator.New()

	return validate.Struct(b)
}
