package ilium

// Stat represents a value on a being
type Stat struct {
	Name  string
	Value int
}

// Add specified amount to value
func (stat *Stat) Add(amount int) {
	stat.Value += amount
}

// Subtract specified amount from value
func (stat *Stat) Subtract(amount int) {
	stat.Value -= amount
}
