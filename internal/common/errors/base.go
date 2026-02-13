package errors

import (
	"errors"
)

// Aliases for errors package functions.
var (
	Is     = errors.Is
	Unwrap = errors.Unwrap
	As     = errors.As
	New    = errors.New
)
