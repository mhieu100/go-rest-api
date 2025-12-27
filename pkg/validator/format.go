package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	case "min":
		return fmt.Sprintf("Should be at least %s characters", fe.Param())
	case "max":
		return fmt.Sprintf("Should be at most %s characters", fe.Param())
	}
	return "Unknown error"
}

func FormatErrors(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{
				Field:   fe.Field(),
				Message: GetErrorMsg(fe),
			}
		}
		return out
	}
	return nil
}
