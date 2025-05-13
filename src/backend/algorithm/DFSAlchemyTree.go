package algorithm

import (
	"littlealchemy2/model"
)

func DFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, mode int8, askedNumOfRecipes *int64, childNode *model.Tree, mapOfElementsTier map[string]int) {
	/**
	algorithm
	0. make a stack that first is filled by the target
	1. search the list of created nodes till found the element you're searching for
	2. get the parent of it
	3. for each parent, search again for it in the list of created nodes
	*/

	// stop if the target is either fire, water, air, earth, or time
	if target == "Fire" || target == "Water" || target == "Air" || target == "Earth" || target == "Time" {
		// STOP, bind, POP and return
		childNode.Name = target
		return
	}

	// searching the stack for target
	for _, n := range listOfCreatedNodes {
		if n != nil && n.Name == target {
			for _, p := range n.Parent {
				if (mapOfElementsTier[p.Ingridient1.Name] > mapOfElementsTier[n.Name]) && (mapOfElementsTier[p.Ingridient2.Name] > mapOfElementsTier[n.Name]) {
					ingridient1 := model.Tree{
						Name:     p.Ingridient1.Name,
						Children: nil,
					}

					ingridient2 := model.Tree{
						Name:     p.Ingridient2.Name,
						Children: nil,
					}

					DFSAlchemyTree(ingridient1.Name, listOfCreatedNodes, mode, askedNumOfRecipes, &ingridient1, mapOfElementsTier)
					DFSAlchemyTree(ingridient2.Name, listOfCreatedNodes, mode, askedNumOfRecipes, &ingridient2, mapOfElementsTier)

					childNode.Children = append(childNode.Children, &ingridient1, &ingridient2)

					if mode == 1 { // first found
						return
					}

				}
			}

		}
	}
}
