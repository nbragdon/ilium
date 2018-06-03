package ilium

import uuid "github.com/satori/go.uuid"

// GameObject provides common functionality for all objects within a game
type GameObject struct {
	UUID uuid.UUID
}

func (gameObject *GameObject) Init() {
	gameObject.UUID = uuid.Must(uuid.NewV4())
}
