package domain

import "cirius-go/neodb/internal/common/errors"

// some standard errors
var (
	ErrUniqueViolation = errors.NewDomain("UniqueViolation")
	ErrDatabaseFailure = errors.NewDomain("DatabaseFailure")
)

// ListingParams represents parameters for listing entities.
type ListingParams[FilterType any] struct {
	Page    int64
	PerPage int64
	Sort    string
	Filter  FilterType
}
