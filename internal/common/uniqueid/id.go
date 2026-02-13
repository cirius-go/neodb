package uniqueid

import "github.com/google/uuid"

// New creates a new UUID string.
func New() string {
	id, err := uuid.NewV7()
	if err != nil {
		panic(err)
	}
	return id.String()
}
