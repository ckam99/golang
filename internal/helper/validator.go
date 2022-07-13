package helper

import (
	"project-struct/internal/http/response"
	"strings"

	"github.com/go-playground/validator"
)

func Validator() *validator.Validate {
	return validator.New()
}

func ValidateSchema(s interface{}) []*response.ErrorResponse {
	var errors []*response.ErrorResponse
	err := Validator().Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element response.ErrorResponse
			element.Namespace = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Field = strings.ToLower(err.Field())

			errors = append(errors, &element)
		}
	}
	return errors
}
