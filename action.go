package ilium

import (
	"github.com/satori/go.uuid"
)

// Action represents a thing to do during a controller's turn
type Action interface{}

type action struct {
	UUID uuid.UUID
}

// NewAction creates an action with sensible defaults
func NewAction() Action {
	newUUID := uuid.Must(uuid.NewV4())
	return &action{newUUID}
}
