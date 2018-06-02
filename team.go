package ilium

import (
	"github.com/satori/go.uuid"
)

// Team represents all the characters that a partipant controls
type Team interface{}

type team struct {
	UUID uuid.UUID
}

// NewTeam creates an instance of Team with sensible defaults
func NewTeam() Team {
	newUUID := uuid.Must(uuid.NewV4())
	return team{newUUID}
}
