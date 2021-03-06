package ilium

// Being defines all the usable characters in the game
type Being interface {
	AddAbility(ability *Ability)
	AddStat(stat *Stat)
	IncreaseStat(name string, amount int)
	DecreaseStat(name string, amount int)
	SetStat(name string, value int)
}

type being struct {
	GameObject
	Abilities []Ability
	Stats     map[string]*Stat
}

// NewBeing creates a being with sensible defaults
func NewBeing() Being {
	abilities := make([]Ability, 0)
	stats := make(map[string]*Stat)
	being := &being{Abilities: abilities, Stats: stats}
	being.Init()
	return being
}

func (being *being) AddAbility(ability *Ability) {
	being.Abilities = append(being.Abilities, *ability)
}

func (being *being) AddStat(stat *Stat) {
	being.Stats[stat.Name] = stat
}

func (being *being) IncreaseStat(name string, amount int) {
	being.Stats[name].Add(amount)
}

func (being *being) DecreaseStat(name string, amount int) {
	being.Stats[name].Subtract(amount)
}

func (being *being) SetStat(name string, value int) {
	being.Stats[name].Value = value
}
