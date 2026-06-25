package Core

func NewFSM() *FSM {
	fsm := &FSM{
		States:     make(map[string]State),
		Transition: make(map[string][]FSMTransition),
		Entity:     make(map[string]FSMEntity),
	}

	return fsm
}

// Gestion Functions

func (f *FSM) AddTransition(from string, transition FSMTransition) {
	f.Transition[from] = append(f.Transition[from], transition)
}

func (f *FSM) AddState(name string, state State) {
	f.States[name] = state
}

func (f *FSM) AddEntity(name string, entity FSMEntity) {
	f.Entity[name] = entity
}

func (f *FSM) GetEntity(name string) FSMEntity {
	return f.Entity[name]
}

func (f *FSM) GetState(name string) State {
	return f.States[name]
}

func (f *FSM) GetTransitions(from string) []FSMTransition {
	return f.Transition[from]
}

func (f *FSM) GetTransition(from string, index int) FSMTransition {
	return f.Transition[from][index]
}

func (f *FSM) SetState(state string, stateData State) {
	f.States[state] = stateData
}

func (f *FSM) SetTransition(from string, index int, transition FSMTransition) {
	f.Transition[from][index] = transition
}

func (f *FSM) SetEntity(name string, entity FSMEntity) {
	f.Entity[name] = entity
}

func (f *FSM) DeleteTransition(from string, index int) {
	f.Transition[from] = append(f.Transition[from][:index], f.Transition[from][index+1:]...)
}

func (f *FSM) DeleteState(name string) {
	delete(f.States, name)
}

func (f *FSM) DeleteEntity(name string) {
	delete(f.Entity, name)
}

// Utility functions

func (f *FSM) GetCurrentState(name string) string {
	return f.Entity[name].CurrentState
}

func (f *FSM) GetNewState(name string) (FSMTransition, any) {
	worldState := NewWorldState()
	for _, transCond := range f.Transition[f.GetCurrentState(name)] {
		if ok, metaData := transCond.Condition(worldState); ok {
			return transCond, metaData
		}
	}
	ok, metaData := f.GetEntity(name).PréTrans.Condition(worldState)
	if ok {
		return f.GetEntity(name).PréTrans, metaData
	}
	return f.GetEntity(name).PréTrans, nil
}

func (f *FSM) ExecuteAction(name string, metaData any) {
	state := f.GetCurrentState(name)
	if _, exists := f.States[state]; exists {
		f.States[state].Action(name, metaData)
	}
}

func (f *FSM) UpdateEntity(name string) {
	newState, metaData := f.GetNewState(name)
	entity := FSMEntity{
		CurrentState: newState.To,
		PréTrans:     newState,
	}
	f.SetEntity(name, entity)
	f.ExecuteAction(name, metaData)
}
