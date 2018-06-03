package ilium

import (
	"github.com/satori/go.uuid"
)

// Controller represents the entity that is managing and choosing actions for a team
type Controller interface{}

type controller struct {
	UUID uuid.UUID
}

// NewController creates a controller with sensible defaults
func NewController() Controller {
	newUUID := uuid.Must(uuid.NewV4())
	return &controller{newUUID}
}
