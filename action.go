package ilium

// Action has costs and associated effects as well as when the action can be used
type Action struct {
	Being Being
	Ability Ability
	Targets []interface{}
}

func (action *Action) Execute(game *Game) {
	for _, cost := range action.Ability.Costs {
		action.PayCost(cost, game)
	}

	for _, effect := range action.Effects {
		action.ApplyEffect(effect, game)
	}
}

func (action *Action) PayCost(being *Being, cost *Cost) {
	being.DecreaseStat(cost.StatName, cost.Value)
}

func (action *Action) ApplyEffect(effect *Effect, game *Game) {

}

const EFFECT_MAPPING := map[string]interface{} {
	"damage": dealDamage
}

func dealDamage(targets *[]Target, effect *Effect, game *Game) {

}
