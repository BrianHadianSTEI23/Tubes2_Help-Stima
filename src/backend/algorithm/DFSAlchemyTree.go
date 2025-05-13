package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func DFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, mode int8, askedNumOfRecipes *int64, childNode *model.Tree, mapOfElementsTier map[string]int, currentFoundRecipe *int64) {

	/**
	algorithm
	0. make a stack that first is filled by the target
	1. search the list of created nodes till found the element you're searching for
	2. get the parent of it
	3. for each parent, search again for it in the list of created nodes
	*/

	if target == "Fire" || target == "Water" || target == "Air" || target == "Earth" || target == "Time" {
		childNode.Name = target
		return
	}

	found := false
	for _, n := range listOfCreatedNodes {
		if n == nil || n.Name != target {
			continue
		}
		for _, p := range n.Parent {
			// fmt.Println(p.Ingridient1.Name + " " + strconv.Itoa(mapOfElementsTier[p.Ingridient1.Name]))
			// fmt.Println(p.Ingridient2.Name + " " + strconv.Itoa(mapOfElementsTier[p.Ingridient2.Name]))
			// fmt.Println(n.Name + " " + strconv.Itoa(mapOfElementsTier[n.Name]))
			if (mapOfElementsTier[p.Ingridient1.Name] < mapOfElementsTier[n.Name]) || (mapOfElementsTier[p.Ingridient2.Name] < mapOfElementsTier[n.Name]) {
				ing1 := &model.Tree{Name: p.Ingridient1.Name}
				ing2 := &model.Tree{Name: p.Ingridient2.Name}

				childNode.Name = target
				childNode.Children = []*model.Tree{ing1, ing2}

				var localWg sync.WaitGroup
				localWg.Add(2)
				go func() {
					defer localWg.Done()
					DFSAlchemyTree(ing1.Name, listOfCreatedNodes, mode, askedNumOfRecipes, ing1, mapOfElementsTier, currentFoundRecipe)
				}()
				go func() {
					defer localWg.Done()
					DFSAlchemyTree(ing2.Name, listOfCreatedNodes, mode, askedNumOfRecipes, ing2, mapOfElementsTier, currentFoundRecipe)
				}()
				localWg.Wait()

				atomic.AddInt64(currentFoundRecipe, 1)

				found = true
				if mode == 1 {
					return
				}
			}
		}
	}

	if !found {
		childNode.Name = target
	}

}
