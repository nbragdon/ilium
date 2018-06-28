package ilium

// Action has costs and associated effects as well as when the action can be used
type Action struct {
	Being   Being
	Ability Ability
	Targets []interface{}
}

// Execute pays the cost of an ability and then activates all the effects
func (action *Action) Execute() {
	for _, cost := range action.Ability.Costs {
		action.PayCost(&cost)
	}

	for _, effect := range action.Ability.Effects {
		effect.Apply(action.Targets)
	}
}

// PayCost reduces stats but the amount specified
func (action *Action) PayCost(cost *Cost) {
	action.Being.DecreaseStat(cost.StatName, cost.Value)
}
