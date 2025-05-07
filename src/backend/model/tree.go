package model

import (
	"fmt"
	"sync"
)

type Tree struct {
	Name        string
	Parent      *Tree
	Ingridient1 *Tree
	Ingridient2 *Tree
}

func DisplayTree(target *Tree, depth *int64, wg *sync.WaitGroup) {
	defer wg.Done()

	if target == nil {
		// stop
	} else {
		(*depth)++
		fmt.Printf("Name : %s | Depth : %d\n", target.Name, *depth)
		wg.Add(2)
		go DisplayTree(target.Ingridient1, depth, wg)
		go DisplayTree(target.Ingridient2, depth, wg)
	}
}

func IsIngridientInTree(ingrident string, tree *Tree) bool {
	if tree != nil {
		if tree.Name == ingrident {
			return true
		} else {
			return IsIngridientInTree(ingrident, tree.Ingridient1) || IsIngridientInTree(ingrident, tree.Ingridient2)
		}
	}
	return false
}

func SearchIngridientInTree(ingrident string, tree *Tree) *Tree {
	if tree == nil {
		return nil
	}

	if tree.Name == ingrident {
		return tree
	}

	if result := SearchIngridientInTree(ingrident, tree.Ingridient1); result != nil {
		return tree
	}

	return SearchIngridientInTree(ingrident, tree.Ingridient2)
}
