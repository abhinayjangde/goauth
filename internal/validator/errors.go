package validator

import (
	"errors"

	playground "github.com/go-playground/validator/v10"
)

func FormatErrors(err error) map[string]string {

	var validationErrors playground.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return map[string]string{
			"error": err.Error(),
		}
	}

	result := make(map[string]string)

	for _, e := range validationErrors {

		switch e.Tag() {

		case "required":
			result[e.Field()] = e.Field() + " is required"

		case "email":
			result[e.Field()] = "Invalid email address"

		case "min":
			result[e.Field()] =
				e.Field() + " must be at least " + e.Param() + " characters"

		case "max":
			result[e.Field()] =
				e.Field() + " must be at most " + e.Param() + " characters"

		default:
			result[e.Field()] = "Invalid value"
		}
	}

	return result
}
