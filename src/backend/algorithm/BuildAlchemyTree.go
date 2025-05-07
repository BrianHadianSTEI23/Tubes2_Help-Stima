package algorithm

import "littlealchemy2/model"

func BuildAlchemyTree(rootElements []*model.AlchemyTree, listOfAllRecipes *[][]string, listOfNodes *[]*model.AlchemyTree) {
	for _, element := range rootElements {
		for _, recipe := range *listOfAllRecipes {
			if recipe[1] == element.Name {

				// search first if the nodes already made in the listOfNodes
				if SearchNodeInCreatedNode(recipe[1], *listOfNodes) != nil {

					// get that particular node to bind
					companion := SearchNodeInCreatedNode(recipe[1], *listOfNodes)

					// add the each other of each other companion
					element.Companion = append(element.Companion, companion)
					companion.Companion = append(companion.Companion, element)

					// build parent pair of current element
					parentPair := model.Pair{
						Ingridient1: element,
						Ingridient2: companion,
					}
					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], element.Children) == nil {

						// build the children of current element
						childAlchemyTree := model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						childAlchemyTree.Parent = append(childAlchemyTree.Parent, &parentPair)

						// bind the child into the children of both
						element.Children = append(element.Children, &childAlchemyTree)
						companion.Children = append(companion.Children, &childAlchemyTree)

						// add the childAlchemyTree into the rootElements and the listOfNodes
						*listOfNodes = append(*listOfNodes, &childAlchemyTree)
					}

				} else { // not made

					// build the opposite parent of current element
					companion := model.AlchemyTree{
						Name:      recipe[1],
						Parent:    nil,
						Children:  nil,
						Companion: nil,
					}

					// add the each other of each other companion
					element.Companion = append(element.Companion, &companion)
					companion.Companion = append(companion.Companion, element)

					// build parent pair of current element
					parentPair := model.Pair{
						Ingridient1: element,
						Ingridient2: &companion,
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], element.Children) == nil {

						// build the children of current element
						childAlchemyTree := model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						childAlchemyTree.Parent = append(childAlchemyTree.Parent, &parentPair)

						// bind the child into the children of both
						element.Children = append(element.Children, &childAlchemyTree)
						companion.Children = append(companion.Children, &childAlchemyTree)

						// add the childAlchemyTree into the rootElements and the listOfNodes
						*listOfNodes = append(*listOfNodes, &childAlchemyTree)
					}

					// add the childAlchemyTree into the rootElements and the listOfNodes
					(*listOfNodes) = append(*listOfNodes, &companion)
				}

			} else if recipe[2] == element.Name {
				// search first if the nodes already made in the listOfNodes
				if SearchNodeInCreatedNode(recipe[2], *listOfNodes) != nil {

					// get that particular node to bind
					companion := SearchNodeInCreatedNode(recipe[2], *listOfNodes)

					// add the each other of each other companion
					element.Companion = append(element.Companion, companion)
					companion.Companion = append(companion.Companion, element)

					// build parent pair of current element
					parentPair := model.Pair{
						Ingridient1: element,
						Ingridient2: companion,
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], element.Children) == nil {

						// build the children of current element
						childAlchemyTree := model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						childAlchemyTree.Parent = append(childAlchemyTree.Parent, &parentPair)

						// bind the child into the children of both
						element.Children = append(element.Children, &childAlchemyTree)
						companion.Children = append(companion.Children, &childAlchemyTree)

						// add the childAlchemyTree into the rootElements and the listOfNodes
						*listOfNodes = append(*listOfNodes, &childAlchemyTree)
					}

				} else { // not made
					// build the opposite parent of current element
					companion := model.AlchemyTree{
						Name:      recipe[2],
						Parent:    nil,
						Children:  nil,
						Companion: nil,
					}

					// add the each other of each other companion
					element.Companion = append(element.Companion, &companion)
					companion.Companion = append(companion.Companion, element)

					// build parent pair of current element
					parentPair := model.Pair{
						Ingridient1: element,
						Ingridient2: &companion,
					}

					// check if the node is already created in the childre
					if SearchNodeInCreatedNode(recipe[0], element.Children) == nil {

						// build the children of current element
						childAlchemyTree := model.AlchemyTree{
							Name:      recipe[0],
							Parent:    nil,
							Children:  nil,
							Companion: nil,
						}

						// bind the parent pair into child
						childAlchemyTree.Parent = append(childAlchemyTree.Parent, &parentPair)

						// bind the child into the children of both
						element.Children = append(element.Children, &childAlchemyTree)
						companion.Children = append(companion.Children, &childAlchemyTree)

						// add the childAlchemyTree into the rootElements and the listOfNodes
						*listOfNodes = append(*listOfNodes, &childAlchemyTree)
					}

					// add the childAlchemyTree into the rootElements and the listOfNodes
					(*listOfNodes) = append(*listOfNodes, &companion)

				}
			}
		}
	}

}

func SearchNodeInCreatedNode(targetName string, listOfNodes []*model.AlchemyTree) *model.AlchemyTree {
	for _, node := range listOfNodes {
		if node != nil {
			if targetName == (*node).Name {
				return node
			}
		}
	}
	return nil
}
