package Core

import (
	"math/rand"
	"sort"
)

func NewBrain() *Brain {
	brain := &Brain{
		States:     make(map[string]State),
		Transition: make(map[string][]BrainTransition),
		Entity:     make(map[string]BrainEntity),
	}

	return brain
}

// Gestion Functions

func (f *Brain) AddTransition(from string, transition BrainTransition) {
	f.Transition[from] = append(f.Transition[from], transition)
}

func (f *Brain) AddState(name string, state State) {
	f.States[name] = state
}

func (f *Brain) AddEntity(name string, entity BrainEntity) {
	f.Entity[name] = entity
}

func (f *Brain) GetEntity(name string) BrainEntity {
	return f.Entity[name]
}

func (f *Brain) GetState(name string) State {
	return f.States[name]
}

func (f *Brain) GetTransitions(from string) []BrainTransition {
	return f.Transition[from]
}

func (f *Brain) GetTransition(from string, index int) BrainTransition {
	return f.Transition[from][index]
}

func (f *Brain) SetState(state string, stateData State) {
	f.States[state] = stateData
}

func (f *Brain) SetTransition(from string, index int, transition BrainTransition) {
	f.Transition[from][index] = transition
}

func (f *Brain) SetEntity(name string, entity BrainEntity) {
	f.Entity[name] = entity
}

func (f *Brain) DeleteTransition(from string, index int) {
	f.Transition[from] = append(f.Transition[from][:index], f.Transition[from][index+1:]...)
}

func (f *Brain) DeleteTransitions(from string) {
	delete(f.Transition, from)
}

func (f *Brain) DeleteState(name string) {
	delete(f.States, name)
}

func (f *Brain) DeleteEntity(name string) {
	delete(f.Entity, name)
}

// Utility functions

func (f *Brain) GetCurrentState(name string) string {
	return f.Entity[name].CurrentState
}

func (f *Brain) GetNewState(name string, randSrc *rand.Rand) (BrainTransition, any) {
	worldState := NewWorldState()
	all := []struct {
		Transition BrainTransition
		MetaData   any
	}{}
	for _, transCond := range f.Transition[f.GetCurrentState(name)] {
		if ok, metaData := transCond.Condition(worldState); ok {
			all = append(all, struct {
				Transition BrainTransition
				MetaData   any
			}{transCond, metaData})
		}
	}

	WeightTotal := 0.0
	for _, trans := range all {
		WeightTotal += float64(trans.Transition.Weight)
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].Transition.Weight > all[j].Transition.Weight
	})

	for _, trans := range all {
		if randSrc.Float64() >= float64(trans.Transition.Weight)/WeightTotal { // if Weight < random [0, WeightTotal]
			return trans.Transition, trans.MetaData
		}
		WeightTotal -= float64(trans.Transition.Weight)
	}

	ok, metaData := f.GetEntity(name).PréTrans.Condition(worldState)
	if ok {
		return f.GetEntity(name).PréTrans, metaData
	}
	return f.GetEntity(name).PréTrans, nil
}

func (f *Brain) ExecuteAction(name string, metaData any) {
	state := f.GetCurrentState(name)
	if _, exists := f.States[state]; exists {
		f.States[state].Action(name, metaData)
	}
}

func (f *Brain) UpdateEntity(name string, randSrc *rand.Rand) {
	newState, metaData := f.GetNewState(name, randSrc)
	entity := BrainEntity{
		CurrentState: newState.To,
		PréTrans:     newState,
	}
	f.SetEntity(name, entity)
	f.ExecuteAction(name, metaData)
}

func (f *Brain) Update(randSrc *rand.Rand) {
	for name := range f.Entity {
		f.UpdateEntity(name, randSrc)
	}
}
