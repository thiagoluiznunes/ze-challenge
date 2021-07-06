package zerrors

// GeneralError represents an input validation error
type GeneralError struct {
	Field   string `json:"field_name,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

// NewValidationError returns a ValidationError instance with the provided parameters
func NewValidationError(field string, message string) *GeneralError {
	return &GeneralError{
		Field:   field,
		Message: message,
		Code:    422,
	}
}

// NewApplicationError returns a ValidationError instance with the provided parameters
func NewApplicationError(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Code:    503,
	}
}
