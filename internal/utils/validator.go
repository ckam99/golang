package utils

import (
	"strings"

	"github.com/ckam225/golang/fiber/internal/http/response"

	"github.com/go-playground/validator"
)

func Validator() *validator.Validate {
	return validator.New()
}

func ValidateCredentials(s interface{}) []*response.Error {
	var errors []*response.Error
	err := Validator().Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.Error
			element.Namespace = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Field = strings.ToLower(err.Field())

			errors = append(errors, &element)
		}
	}
	return errors
}
