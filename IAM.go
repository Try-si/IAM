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

func InitFSM(config string) {
	fsm = *Core.NewFSM(config)
}

func UpdateFSM() {
	for name := range fsm.Entity {
		fsm.UpdateEntity(name)
	}
}

func GetFSM() *Core.FSM {
	return &fsm
}

// Brain Management

func InitBrain(config string) {
	brain = *Core.NewBrain(config)
}

func UpdateBrain(randSrc *rand.Rand) {
	brain.Update(randSrc)
}

func GetBrain() *Core.Brain {
	return &brain
}

// BeavioursTree Management

func InitBeavioursTree(config string) {
	bt = *Core.InitBeavioursTree(config)
}

func ExecuteBehavioursTree(root string) {
	bt.Execute(root)
}

func GetBeavioursTree() *Core.BeavioursTree {
	return &bt
}

func UpdateBeavioursTree() {
	for name := range bt.GetRoots() {
		bt.Execute(name)
	}
}
