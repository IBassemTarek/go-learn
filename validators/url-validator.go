package validators

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateURL(url validator.FieldLevel) bool {
	return !strings.Contains(url.Field().String(), "https://www.google.com")
}
