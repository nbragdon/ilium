package ilium

const EFFECT_MAPPING := map[string]interface{} {
	"damage": dealDamage
}

// Action has costs and associated effects as well as when the action can be used
type Action struct {
	Being Being
	Ability Ability
	Targets []interface{}
}

func (action *Action) Execute(game *Game) {
	for _, cost := range action.Ability.Costs {
		action.PayCost(cost)
	}

	for _, effect := range action.Ability.Effects {
		action.ApplyEffect(effect, game)
	}
}

func (action *Action) PayCost(cost *Cost) {
	action.Being.DecreaseStat(cost.StatName, cost.Value)
}

func (action *Action) ApplyEffect(effect *Effect, game *Game) {
	EFFECT_MAPPING[effect.]
}


func dealDamage(targets *[]Target, effect *Effect, game *Game) {
	for _, target := range targets {

	}
}
