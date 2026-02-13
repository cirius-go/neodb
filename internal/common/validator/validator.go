package validator

import (
	"sync/atomic"

	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
)

var defaultValidator atomic.Pointer[Validator]

// Validator represents a data validator.
type Validator struct {
	v *validator.Validate
}

// New creates a new Validator instance.
func New() *Validator {
	v := validator.New()

	en.RegisterDefaultTranslations(v, nil)
	return &Validator{
		v: v,
	}
}

// Struct validates a struct.
func (v *Validator) Struct(s any) error {
	return v.v.Struct(s)
}

// Var validates a variable with a tag.
func (v *Validator) Var(s any, tag string) error {
	return v.v.Var(s, tag)
}

// Get returns the default Validator instance.
func Get() *Validator {
	if nv := defaultValidator.Load(); nv != nil {
		return nv
	}
	nv := New()
	defaultValidator.CompareAndSwap(nil, nv)
	return defaultValidator.Load()
}

// Set sets the default Validator instance.
func Set(v *Validator) {
	defaultValidator.Store(v)
}

// Struct validates a struct using the default Validator instance.
func Struct(v any) error {
	validator := Get()
	return validator.Struct(v)
}

// Var validates a variable with a tag using the default Validator instance.
func Var(s any, tag string) error {
	validator := Get()
	return validator.Var(s, tag)
}
