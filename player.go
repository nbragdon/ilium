package ilium

import uuid "github.com/satori/go.uuid"

// Player represents a controller that is human operated
type Player interface {
	Controller
}

type player struct {
	UUID uuid.UUID
}

// NewPlayer creates a player with sensible defaults
func NewPlayer() Controller {
	newUUID := uuid.Must(uuid.NewV4())
	return player{newUUID}
}
