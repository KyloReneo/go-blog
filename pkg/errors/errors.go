package errors

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/kyloReneo/go-blog/internal/providers/validation"

)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

// Accepts error variables
func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), fieldError.Tag())
		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMessage(tag string) string {
	return validation.ErrorMessages()[tag]
}

func Get() map[string]string {
	return errorsList
}
