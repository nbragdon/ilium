package ilium

const hp = "hp"

// Damage reduces targets hp by the specified amount
type Damage struct {
	Amount int
}

// Apply reduces targets hp but the specified amount
func (damage *Damage) Apply(targets []interface{}) {
	for _, target := range targets {
		being, ok := target.(Being)
		if ok {
			being.DecreaseStat(hp, damage.Amount)
		}
	}
}
