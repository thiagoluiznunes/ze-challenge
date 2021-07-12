package zerrors

import (
	"fmt"
)

const (
	PartnerNotFoundError         string = "partner not found"
	PartnerAlreadyExistError     string = "partner already exists"
	GetClosestPartnerByAreaError string = "fail to get closest partner by area"
)

// GeneralError represents an input validation error
type GeneralError struct {
	Field   string `json:"field_name,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (e *GeneralError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%v: %v", e.Field, e.Message)
	}
	return fmt.Sprintf("%v", e.Message)
}

// NewValidationError returns a ValidationError instance with the provided parameters
func NewValidationError(field string, message string) *GeneralError {
	return &GeneralError{
		Field:   field,
		Message: message,
		Code:    422,
	}
}

// NewNotFoundError returns a ValidationError instance with the provided parameters
func NewNotFoundError(err error) *GeneralError {
	return &GeneralError{
		Message: err.Error(),
		Code:    404,
	}
}

// NewApplicationError returns a ValidationError instance with the provided parameters
func NewApplicationError(err error) *GeneralError {
	return &GeneralError{
		Message: err.Error(),
		Code:    503,
	}
}

// NewConflictError returns a ValidationError instance with the provided parameters
func NewConflictError(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Code:    409,
	}
}
