package IAMobs

import (
	"math/rand"

	"github.com/Try-si/IAM/Core"
)

var (
	fsm   Core.FSM
	brain Core.Brain
	bt    Core.BeavioursTree
)

// FSM Management

func InitFSM() {
	fsm = *Core.NewFSM()
}

func UpdateFSM() {
	for name := range fsm.Entity {
		fsm.UpdateEntity(name)
	}
}

func GetFSM() Core.FSM {
	return fsm
}

// Brain Management

func InitBrain() {
	brain = *Core.NewBrain()
}

func UpdateBrain(randSrc *rand.Rand) {
	brain.Update(randSrc)
}

func GetBrain() Core.Brain {
	return brain
}

// BeavioursTree Management

func InitBeavioursTree() {
	bt = *Core.InitBeavioursTree()
}

func ExecuteBehavioursTree(root string) {
	bt.Execute(root)
}

func GetBeavioursTree() Core.BeavioursTree {
	return bt
}
