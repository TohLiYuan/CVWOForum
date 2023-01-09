package validators

import (
	"github.com/go-playground/validator/v10"
)

func ValidateCategory(fl validator.FieldLevel) bool {
	return fl.Field().String() == "general" || fl.Field().String() == "discussion" || fl.Field().String() == "advice"
}
