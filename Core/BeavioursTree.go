package Core

import (
	"encoding/json"
	"os"
)

// Root management

func InitBeavioursTree(config string) *BeavioursTree {
	bt := &BeavioursTree{
		Roots: make(map[string]*BehaviourNode),
	}

	if config != "" && len(config) > 4 {
		if config[len(config)-4:] == ".json" {
			bt = LoadTFromFile[*BeavioursTree](config)
		} else {
			bt = LoadTFromString[*BeavioursTree](config)
		}
	}

	return bt
}

func (bt *BeavioursTree) AddRoot(name string, node *BehaviourNode) {
	bt.Roots[name] = node
}

func (bt *BeavioursTree) RemoveRoot(name string) {
	delete(bt.Roots, name)
}

func (bt *BeavioursTree) GetRoot(name string) *BehaviourNode {
	return bt.Roots[name]
}

func (bt *BeavioursTree) GetRoots() map[string]*BehaviourNode {
	return bt.Roots
}

// node management

// tree execution

func (bt *BeavioursTree) Execute(root string) {
	WorldState := NewWorldState()
	node := bt.Roots[root]
	for i := 0; i < 100; i++ {
		if node == nil {
			break
		}
		if node.Condition == nil {
			break
		}
		ok, metaData := node.Condition(WorldState)
		if node.Verify() {
			if ok {
				node = node.TrueNode
			} else {
				node = node.FalseNode
			}
		} else {
			node.Execute(metaData)
			break
		}
	}
}

func (node *BehaviourNode) Execute(Any any) {
	node.Action(Any)
}

// si bool = false alors sa veut dire que c'est une feuille sinon c'est un noeud
func (node *BehaviourNode) Verify() bool {
	if node.FalseNode == nil || node.TrueNode == nil {
		return false
	}
	return true
}

/////

func LoadTFromString[T any](config string) T {
	var result T
	json.Unmarshal([]byte(config), &result)
	return result
}

func LoadTFromFile[T any](path string) T {
	var result T
	fileRead, err := os.ReadFile(path)
	if err != nil {
		return result
	}
	result = LoadTFromString[T](string(fileRead))
	return result
}
