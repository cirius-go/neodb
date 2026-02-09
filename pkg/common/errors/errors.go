package errors

import (
	"errors"
	"fmt"
	"net/http"

	"cirius-go/neodb/pkg/common"
	"cirius-go/neodb/pkg/common/slice"

	"golang.org/x/text/language"
)

// Aliases for errors package functions.
var (
	Is     = errors.Is
	Unwrap = errors.Unwrap
	As     = errors.As
)

var (
	DefaultLanguage      = language.English.String()
	UnknownErrorLocation = "unknown"
)

type (
	// ErrorDetail represents detailed information about an error.
	ErrorDetail struct {
		// Location of the error, e.g., field name or parameter.
		Location string `json:"location" yaml:"location"`
		// Value associated with the error, e.g., invalid value.
		Value any `json:"value,omitempty" yaml:"value,omitempty"`
		// Message is the human-readable error message.
		Message string `json:"message" yaml:"message"`
	}
	// StatusError is a marker interface for simple status errors.
	StatusError struct {
		Status   int            `json:"status" yaml:"status"`
		Title    string         `json:"title" yaml:"title"`
		Detail   string         `json:"detail" yaml:"detail"`
		Errors   []*ErrorDetail `json:"errors,omitempty" yaml:"errors,omitempty"`
		internal error          `json:"-" yaml:"-"`
	}
)

// Error implements error.
func (e *ErrorDetail) Error() string {
	return fmt.Sprintf("%s is having error: %s", e.Location, e.Message)
}

// Error implements common.Error.
func (e *StatusError) Error() string {
	return e.Detail
}

// Unwrap implements common.Error.
func (e *StatusError) Unwrap() error {
	return e.internal
}

// Is implements common.Error.
func (e *StatusError) Is(target error) bool {
	t, ok := target.(*StatusError)
	if !ok {
		return false
	}
	return e.Detail == t.Detail && e.Status == t.Status
}

// WithDetail implements common.Error.
func (e *StatusError) WithDetail(d *ErrorDetail) *StatusError {
	e.Errors = append(e.Errors, d)
	return e
}

// WithInternal implements common.Error.
func (e *StatusError) WithInternal(err error) *StatusError {
	e.internal = err
	return e
}

var _ common.Error[*StatusError, *ErrorDetail] = (*StatusError)(nil)

// New creates a new instance of Error.
func New(status int, msg string, errs ...error) *StatusError {
	return &StatusError{
		Status: status,
		Title:  http.StatusText(status),
		Detail: msg,
		Errors: slice.Map(errs, ConvertToErrDetail),
	}
}

// ConvertToErrDetail converts a generic error to ErrorDetail.
func ConvertToErrDetail(err error) *ErrorDetail {
	if err == nil {
		return nil
	}
	if ed, ok := err.(*ErrorDetail); ok {
		return ed
	}
	if af, ok := err.(*I18nErrorDetail); ok {
		return &ErrorDetail{
			Location: af.Location,
			Value:    af.Value,
			Message:  af.Message.Localize(DefaultLanguage),
		}
	}
	return &ErrorDetail{
		Location: UnknownErrorLocation,
		Value:    nil,
		Message:  err.Error(),
	}
}
