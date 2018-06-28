package ilium

import (
	"github.com/satori/go.uuid"
)

// Game object to store all game information
type Game interface {
	AddController(*Controller)
	AddPlayer()
	AddComputer()
	AvailableActions(*Controller) []Action
	NextRound()
	GetActiveController() Controller
	AddAction(*Action)
}

type game struct {
	UUID                  uuid.UUID
	controllers           []Controller
	round                 uint8
	activeControllerIndex uint8
	actions               []Action
}

// NewGame creates new game object with sensible defaults
func NewGame() Game {
	newUUID := uuid.Must(uuid.NewV4())
	controllers := make([]Controller, 0)
	actions := make([]Action, 0)
	return &game{newUUID, controllers, 1, 0, actions}
}

func (game *game) AddController(controller *Controller) {
	game.controllers = append(game.controllers, controller)
}

func (game *game) AddPlayer() {
	game.controllers = append(game.controllers, NewPlayer())
}

func (game *game) AddComputer() {
	game.controllers = append(game.controllers, NewComputer())
}

func (game *game) AvailableActions(controller *Controller) []Action {
	return make([]Action, 0)
}

func (game *game) NextRound() {
	game.activeControllerIndex = (game.activeControllerIndex + 1) % uint8(len(game.controllers))
	game.round++
}

func (game *game) GetActiveController() Controller {
	return game.controllers[game.activeControllerIndex]
}

func (game *game) AddAction(action *Action) {
	game.actions = append(game.actions, *action)
}

func (game *game) ExecuteActions() {
	for i := len(game.actions) - 1; i >= 0; i-- {
		game.actions[i].Execute()
	}
}
