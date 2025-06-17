package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
)

type (
	ErrorResponse struct {
		Error        bool
		FailedFields string
	}

	XValidator struct {
		Validator *validator.Validate
	}

	ValidationError struct {
		Message string
	}
)

func (e ValidationError) Error() string {
	return e.Message
}

func NewXValidator() *XValidator {
	return &XValidator{Validator: validator.New()}
}

func (v XValidator) Validate(data any) error {

	if errs := v.validateData(data); len(errs) > 0 && errs[0].Error {
		errMsg := make([]string, 0)
		for _, err := range errs {
			errMsg = append(errMsg, fmt.Sprintf(
				"[%s]",
				err.FailedFields,
			))
		}
		return ValidationError{Message: "Invalid field or its absence: " + strings.Join(errMsg, " and ")}
	}
	return nil
}

func (v XValidator) validateData(data any) []ErrorResponse {
	var validationErrors []ErrorResponse

	errs := v.Validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse
			elem.FailedFields = err.Field()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}
	return validationErrors
}
