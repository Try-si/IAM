# Description

IAM (IAMobs) est une librairie Go qui permet de créer facilement des systèmes d'IA pour les jeux vidéo.

## Utilisation

Exemple incomplet pour un zombie, un villageois et un garde :

```go
package main

import (
	"math/rand"

	IAMobs "github.com/Try-si/IAM"
	"github.com/Try-si/IAM/Core"
)

var (
	randSrc *rand.Rand
)

func main() {
	InitIA()
	randSrc = rand.New(rand.NewSource(0))

	GameLoop()
}

func GameLoop() {
	for {
		IAMobs.UpdateBrain(randSrc)
		IAMobs.UpdateFSM()
		IAMobs.UpdateBeavioursTree()
	}
}

func InitIA() {
	initFSM()
	initBrain()
	initBeavioursTree()
}

func initFSM() {
	IAMobs.InitFSM()

	fsm := IAMobs.GetFSM()

	fsm.AddState("zombie_idle", Core.State{
		Action: func(state any) {
			println("Zombie is idle")
		},
	})
	fsm.AddState("zombie_walk", Core.State{
		Action: func(state any) {
			println("Zombie is walking")
		},
	})

	fsm.AddTransition("zombie_idle", Core.FSMTransition{
		To: "zombie_walk",
		Condition: func(state Core.WorldState) (bool, any) {
			return false, nil // si il voit un villagois
		},
	})

	fsm.AddTransition("zombie_walk", Core.FSMTransition{
		To: "zombie_idle",
		Condition: func(state Core.WorldState) (bool, any) {
			return false, nil // si il perd le villagois de vue ou le vois pas
		},
	})

	fsm.AddEntity("zombie", Core.FSMEntity{
		CurrentState: "zombie_idle",
	})
}

func initBrain() {
	IAMobs.InitBrain()

	brain := IAMobs.GetBrain()

	brain.AddState("villagois_idle", Core.State{
		Action: func(state any) {
			println("Villagois is idle")
		},
	})
	brain.AddState("villagois_leak", Core.State{
		Action: func(state any) {
			println("Villagois is leaking")
		},
	})

	brain.AddTransition("villagois_idle", Core.BrainTransition{
		Weight: 1,
		To:     "villagois_leak",
		Condition: func(ws Core.WorldState) (bool, any) {
			return false, nil // si il voit un zombie
		},
	})
	brain.AddTransition("villagois_leak", Core.BrainTransition{
		Weight: 1,
		To:     "villagois_idle",
		Condition: func(ws Core.WorldState) (bool, any) {
			return false, nil // si il ne voit plus de zombie ou le voit pas
		},
	})

	brain.AddEntity("villagois", Core.BrainEntity{
		CurrentState: "villagois_idle",
	})
}

func initBeavioursTree() {
	IAMobs.InitBeavioursTree()

	Bt := IAMobs.GetBeavioursTree()

	Bt.AddRoot("guard", &Core.BehaviourNode{
		Condition: func(ws Core.WorldState) (bool, any) {
			return false, nil // si il voit un zombie
		},
		TrueNode: &Core.BehaviourNode{
			Condition: func(ws Core.WorldState) (bool, any) {
				return false, nil // si il voit un zombies
			},
			Action: func(a any) {
				println("Guarding")
			},
		},
		FalseNode: &Core.BehaviourNode{
			Condition: func(ws Core.WorldState) (bool, any) {
				return false, nil // si il ne voit pas de zombie
			},
			Action: func(a any) {
				println("Attacking")
			},
		},
	})
}

```