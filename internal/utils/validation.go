package utils

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrors(err error) map[string]string {
	var validationErrors validator.ValidationErrors

	if !errors.As(err, &validationErrors) {
		return map[string]string{
			"error": err.Error(),
		}
	}

	result := make(map[string]string)

	for _, fieldErr := range validationErrors {

		field := strings.ToLower(fieldErr.Field())

		switch fieldErr.Tag() {

		case "required":
			result[field] = fieldErr.Field() + " is required"

		case "email":
			result[field] = "Invalid email address"

		case "min":
			result[field] = fieldErr.Field() + " is too short"

		case "max":
			result[field] = fieldErr.Field() + " is too long"

		default:
			result[field] = "Invalid value"
		}
	}

	return result
}
