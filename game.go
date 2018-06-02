package ilium

import (
	"github.com/satori/go.uuid"
)

// Game object to store all game information
type Game interface {
	addController(*Controller)
	addPlayer()
	addComputer()
	availableActions(*Controller) []Action
}

type game struct {
	UUID        uuid.UUID
	controllers []Controller
}

// NewGame creates new game object with sensible defaults
func NewGame() Game {
	newUUID := uuid.Must(uuid.NewV4())
	controllers := make([]Controller, 0)
	return &game{newUUID, controllers}
}

func (game *game) addController(controller *Controller) {
	game.controllers = append(game.controllers, controller)
}

func (game *game) addPlayer() {
	game.controllers = append(game.controllers, NewPlayer())
}

func (game *game) addComputer() {
	game.controllers = append(game.controllers, NewComputer())
}

func (game *game) availableActions(controller *Controller) []Action {
	return make([]Action, 0)
}
