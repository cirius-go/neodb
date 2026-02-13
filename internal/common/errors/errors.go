package errors

import "fmt"

// AppError represents an application error.
type AppError struct {
	Code  string
	Cause error
}

// Error implements the error interface.
func (d *AppError) Error() string {
	if d.Cause == nil {
		return d.Code
	}
	return fmt.Sprintf("%s: %s", d.Code, d.Cause.Error())
}

// WithCause sets the underlying cause of the error.
func (d *AppError) WithCause(err error) *AppError {
	return &AppError{
		Code:  d.Code,
		Cause: err,
	}
}

// Unwrap returns the underlying cause of the error.
func (d *AppError) Unwrap() error {
	return d.Cause
}

// Is checks if the target error is the same as the current error based on the
// Code.
func (e *AppError) Is(target error) bool {
	if e == nil {
		return false
	}

	t, ok := target.(*AppError)
	if !ok || t == nil {
		return false
	}

	return e.Code == t.Code
}

// NewDomain creates a new domain error with the given code.
func NewDomain(code string) *AppError {
	return &AppError{
		Code: "error:domain:" + code,
	}
}

// NewInfra creates a new infra error with the given code.
func NewInfra(code string) *AppError {
	return &AppError{
		Code: "error:infra:" + code,
	}
}
