package ilium

// Effect are game changes when an action is executed
type Effect interface {
	Apply(targets []interface{})
}
