package Core

// FSM

type FSM struct {
	Transition map[string][]FSMTransition
	States     map[string]State
	Entity     map[string]FSMEntity
}

type FSMEntity struct {
	CurrentState string
	PréTrans     FSMTransition
}

type FSMTransition struct {
	To        string
	Condition func(WorldState) (bool, any)
}

type WorldState struct {
}

// Brain

type Brain struct {
	Transition map[string][]BrainTransition
	States     map[string]State
	Entity     map[string]BrainEntity
}

type BrainTransition struct {
	Weight    float32
	To        string
	Condition func(WorldState) (bool, any)
}

type BrainEntity struct {
	CurrentState string
	PréTrans     BrainTransition
}

// BeavioursTree

// Other

type State struct {
	Action func(any)
}
