package validators

import (
	"strings"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	Namespace string `json:"namespace"`
	Field     string `json:"field"`
	Tag       string `json:"tag"`
	Value     string `json:"value"`
}

func Validator() *validator.Validate {
	return validator.New()
}

func ValidateSchema(s interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := Validator().Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Namespace = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			element.Field = strings.ToLower(err.Field())

			errors = append(errors, &element)
		}
	}
	return errors
}
