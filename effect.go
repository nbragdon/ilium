package ilium

// Effect defines what should be altered or changed during an action
type Effect struct {
	Classification string
	Value          int
	TargetType     string
	Tags           []string
}
