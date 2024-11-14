package validate

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// ValidationError represents validation errors with friendly messages.
type ValidationError struct {
	Errors map[string][]string
}

// Error implements the error interface.
func (ve *ValidationError) Error() string {
	// Concatenate all error messages into a single string.
	msg := "Validation failed:\n"
	for field, messages := range ve.Errors {
		for _, message := range messages {
			msg += fmt.Sprintf("- %s: %s\n", field, message)
		}
	}
	return msg
}

// Validator wraps the go-playground/validator.
type Validator struct {
	validator *validator.Validate
}

// New creates a new instance of Validator.
func New() *Validator {
	return &Validator{validator: validator.New()}
}

// Validate returns a ValidationError if any.
func (cv *Validator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return err
	}

	ve := &ValidationError{
		Errors: make(map[string][]string),
	}

	for _, fieldErr := range validationErrors {
		fieldName := fieldErr.Field()
		tag := fieldErr.Tag()
		param := fieldErr.Param()
		message := errorMessage(tag, param)

		ve.Errors[fieldName] = append(ve.Errors[fieldName], message)
	}

	return ve
}

// errorMessage creates a friendly error message based on the validation tag.
func errorMessage(tag, param string) string {
	switch tag {
	case "required":
		return "This field is required."
	case "min":
		return fmt.Sprintf("This field must be at least %s characters long.", param)
	case "email":
		return "This field must be a valid email address."
	default:
		return fmt.Sprintf("This field failed validation on the '%s' rule.", tag)
	}
}
