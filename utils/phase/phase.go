package phase

var initialPhase = &phase{
	ID:   10,
	Name: "Initial",
}

var commandPhase = &phase{
	ID:   20,
	Name: "Command",
}

// OnInitialPhase will set current phase to initial
func OnInitialPhase() {
	setPhase(initialPhase)
}

// OnInitialPhase will set current phase to initial
func OnCommandPhase() {
	setPhase(commandPhase)
}

func OnNewPhase(id int, name string) {
	setPhase(&phase{
		ID:   id,
		Name: name,
	})
}
