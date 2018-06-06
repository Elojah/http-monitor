package monitor

import (
	uuid "github.com/satori/go.uuid"
)

// ID is an unique identifier for resources.
type ID = uuid.UUID

// NewID generates a new uniquer identifier.
func NewID() ID {
	return ID(uuid.Must(uuid.NewV4()))
}
