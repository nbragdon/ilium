package ilium

import uuid "github.com/satori/go.uuid"

// Computer represents a controller that is non-player operated
type Computer interface {
	Controller
}

type computer struct {
	UUID uuid.UUID
}

// NewComputer creates a computer with sensible defaults
func NewComputer() Controller {
	newUUID := uuid.Must(uuid.NewV4())
	return player{newUUID}
}
