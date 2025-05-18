package algorithm

import "littlealchemy2/model"

func BuildAlchemyTree(rootElements []*model.AlchemyTree, listOfAllRecipes *[][]string, listOfNodes *[]*model.AlchemyTree) {
	/*
		algorithm
		1. iterate for all the recipes and create the node and put it into list of nodes with parents, children, and companion is nil
		2. iterate for each node in list of nodes
	*/
	for i := 0; i < len(*listOfNodes); i++ {
		for _, recipe := range *listOfAllRecipes {
			if recipe[1] == (*listOfNodes)[i].Name {
				tempCompanion := SearchNodeInCreatedNode(recipe[2], *listOfNodes)
				// search first if the nodes already made in the listOfNodes
				if tempCompanion != nil {

					// get that particular node to bind
					companion := tempCompanion

					// add the each other of each other companion
					if SearchNodeInCreatedNode(companion.Name, (*listOfNodes)[i].Companion) == nil {
						(*listOfNodes)[i].Companion = append((*listOfNodes)[i].Companion, companion)
					}
					if SearchNodeInCreatedNode((*listOfNodes)[i].Name, companion.Companion) == nil {
						companion.Companion = append(companion.Companion, (*listOfNodes)[i])
					}

					// build parent pair of current element
					parentPair := &model.Pair{
						Ingridient1: (*listOfNodes)[i],
						Ingridient2: companion,
					}
					// check if the child node is already created
					if SearchNodeInCreatedNode(recipe[0], *listOfNodes) == nil {

						// build the children of current element
						childAlchemyTree := &model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						if SearchPairInCreatedPair(parentPair.Ingridient1.Name, parentPair.Ingridient2.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}

						// bind the child into the children of both
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}

						// add the childAlchemyTree into the rootElements and the listOfNodes
						if SearchNodeInCreatedNode(childAlchemyTree.Name, *listOfNodes) == nil {
							*listOfNodes = append(*listOfNodes, childAlchemyTree)
						}
					} else {
						childAlchemyTree := SearchNodeInCreatedNode(recipe[0], *listOfNodes)
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}
						// bind the parent pair into child
						if SearchPairInCreatedPair((*listOfNodes)[i].Name, companion.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
					}

				} else { // not made

					// build the opposite parent of current element
					companion := &model.AlchemyTree{
						Name:      recipe[2],
						Parent:    nil,
						Children:  nil,
						Companion: nil,
					}

					// add the each other of each other companion
					if SearchNodeInCreatedNode(companion.Name, (*listOfNodes)[i].Companion) == nil {
						(*listOfNodes)[i].Companion = append((*listOfNodes)[i].Companion, companion)
					}
					if SearchNodeInCreatedNode((*listOfNodes)[i].Name, companion.Companion) == nil {
						companion.Companion = append(companion.Companion, (*listOfNodes)[i])
					}

					// build parent pair of current element
					parentPair := &model.Pair{
						Ingridient1: (*listOfNodes)[i],
						Ingridient2: companion,
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], *listOfNodes) == nil {

						// build the children of current element
						childAlchemyTree := &model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						if SearchPairInCreatedPair(parentPair.Ingridient1.Name, parentPair.Ingridient2.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
						// bind the child into the children of both
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(childAlchemyTree.Name, *listOfNodes) == nil {
							*listOfNodes = append(*listOfNodes, childAlchemyTree)
						}

					} else {
						childAlchemyTree := SearchNodeInCreatedNode(recipe[0], *listOfNodes)
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}
						// bind the parent pair into child
						if SearchPairInCreatedPair((*listOfNodes)[i].Name, companion.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
					}

					// add the childAlchemyTree into the rootElements and the listOfNodes
					if SearchNodeInCreatedNode(companion.Name, *listOfNodes) == nil {
						*listOfNodes = append(*listOfNodes, companion)
					}
				}

			} else if recipe[2] == (*listOfNodes)[i].Name {
				tempCompanion := SearchNodeInCreatedNode(recipe[1], *listOfNodes)
				// search first if the nodes already made in the listOfNodes
				if tempCompanion != nil {

					// get that particular node to bind
					companion := tempCompanion

					// add the each other of each other companion
					if SearchNodeInCreatedNode(companion.Name, (*listOfNodes)[i].Companion) == nil {
						(*listOfNodes)[i].Companion = append((*listOfNodes)[i].Companion, companion)
					}
					if SearchNodeInCreatedNode((*listOfNodes)[i].Name, companion.Companion) == nil {
						companion.Companion = append(companion.Companion, (*listOfNodes)[i])
					}

					// build parent pair of current element
					parentPair := &model.Pair{
						Ingridient1: companion,
						Ingridient2: (*listOfNodes)[i],
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], *listOfNodes) == nil {

						// build the children of current element
						childAlchemyTree := &model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						if SearchPairInCreatedPair(parentPair.Ingridient1.Name, parentPair.Ingridient2.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
						// bind the child into the children of both
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}

						// add the childAlchemyTree into the rootElements and the listOfNodes
						if SearchNodeInCreatedNode(childAlchemyTree.Name, *listOfNodes) == nil {
							*listOfNodes = append(*listOfNodes, childAlchemyTree)
						}
					} else {
						childAlchemyTree := SearchNodeInCreatedNode(recipe[0], *listOfNodes)
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}
						// bind the parent pair into child
						if SearchPairInCreatedPair(companion.Name, (*listOfNodes)[i].Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
					}

				} else { // not made
					// build the opposite parent of current element
					companion := &model.AlchemyTree{
						Name:      recipe[1],
						Parent:    nil,
						Children:  nil,
						Companion: nil,
					}

					// add the each other of each other companion
					if SearchNodeInCreatedNode(companion.Name, (*listOfNodes)[i].Companion) == nil {
						(*listOfNodes)[i].Companion = append((*listOfNodes)[i].Companion, companion)
					}
					if SearchNodeInCreatedNode((*listOfNodes)[i].Name, companion.Companion) == nil {
						companion.Companion = append(companion.Companion, (*listOfNodes)[i])
					}

					// build parent pair of current element
					parentPair := &model.Pair{
						Ingridient1: companion,
						Ingridient2: (*listOfNodes)[i],
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], *listOfNodes) == nil {

						// build the children of current element
						childAlchemyTree := &model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						if SearchPairInCreatedPair(parentPair.Ingridient1.Name, parentPair.Ingridient2.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
						// bind the child into the children of both
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}

						// add the childAlchemyTree into the rootElements and the listOfNodes
						if SearchNodeInCreatedNode(childAlchemyTree.Name, *listOfNodes) == nil {
							*listOfNodes = append(*listOfNodes, childAlchemyTree)
						}
					} else {
						childAlchemyTree := SearchNodeInCreatedNode(recipe[0], *listOfNodes)
						if SearchNodeInCreatedNode(recipe[0], (*listOfNodes)[i].Children) == nil {
							(*listOfNodes)[i].Children = append((*listOfNodes)[i].Children, childAlchemyTree)
						}
						if SearchNodeInCreatedNode(recipe[0], companion.Children) == nil {
							companion.Children = append(companion.Children, childAlchemyTree)
						}
						// bind the parent pair into child
						if SearchPairInCreatedPair((*listOfNodes)[i].Name, companion.Name, childAlchemyTree.Parent) == nil {
							childAlchemyTree.Parent = append(childAlchemyTree.Parent, parentPair)
						}
					}

					// add the childAlchemyTree into the rootElements and the listOfNodes
					if SearchNodeInCreatedNode(companion.Name, *listOfNodes) == nil {
						*listOfNodes = append(*listOfNodes, companion)
					}
				}
			}
		}
	}

}

func SearchNodeInCreatedNode(targetName string, listOfNodes []*model.AlchemyTree) *model.AlchemyTree {
	for _, node := range listOfNodes {
		if node != nil && targetName == (*node).Name {
			return node
		}
	}
	return nil
}

func SearchPairInCreatedPair(a, b string, listOfPairs []*model.Pair) *model.Pair {
	for _, pair := range listOfPairs {
		if pair != nil {
			n1 := pair.Ingridient1.Name
			n2 := pair.Ingridient2.Name
			if (a == n1 && b == n2) || (a == n2 && b == n1) {
				return pair
			}
		}
	}
	return nil
}
