package domain

import (
	"time"

	"github.com/google/uuid"
)

// Model represents a generic model in the domain.
type Model struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// OnInit initializes the model by generating a unique ID if it is not already
// set.
func (m *Model) OnInit() error {
	if m.ID == "" {
		id, err := uuid.NewV7()
		if err != nil {
			return err
		}
		m.ID = id.String()
	}
	return nil
}
