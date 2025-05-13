package model

import "fmt"

type AlchemyTree struct {
	Name      string
	Parent    []*Pair
	Children  []*AlchemyTree
	Companion []*AlchemyTree
}

func DisplayAlchemyTree(t *AlchemyTree) {
	visited := make(map[*AlchemyTree]bool)
	var count int = 0
	displayHelper(t, visited, &count)
	// return count
}

func displayHelper(t *AlchemyTree, visited map[*AlchemyTree]bool, count *int) {
	if t == nil || visited[t] {
		return
	}

	visited[t] = true
	(*count)++
	fmt.Println(t.Name)

	for _, comp := range t.Companion {
		displayHelper(comp, visited, count)
	}
	for _, child := range t.Children {
		displayHelper(child, visited, count)
	}
	for _, parent := range t.Parent {
		if parent.Ingridient1 != nil {
			displayHelper(parent.Ingridient1, visited, count)
		}
		if parent.Ingridient2 != nil {
			displayHelper(parent.Ingridient2, visited, count)
		}
	}
}
