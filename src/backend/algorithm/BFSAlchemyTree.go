package algorithm

import "littlealchemy2/model"

func BFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, mode int8, askedNumOfRecipes *int64, response *model.Response, mapOfElementsTier map[string]int) {

	/*
		algorithm
		1. input is a queue that has the first element
		2. partition the queue to be first and the rest
		2.5 for the first element, search the parent of that element inside list of created nodes
		3. add each parent into the childNode.parent
		4.
	*/

	// define queue
	type QueueItem struct {
		Name string
		Tree *model.Tree
	}

	// initialize bfs queue
	BFSQueue := []QueueItem{
		{
			Name: target,
			Tree: &response.Data,
		},
	}

	// searching
	for len(BFSQueue) > 0 {
		// partition the queue into two parts : first one and then the rest
		head := BFSQueue[0]
		rest := BFSQueue[1:]

		BFSQueue = rest

		// stop condition
		if head.Name == "Fire" || head.Name == "Water" || head.Name == "Air" || head.Name == "Earth" || head.Name == "Time" {
			continue
		}

		// found the element in created nodes
		for _, node := range listOfCreatedNodes {
			if (node != nil) && (head.Name == node.Name) {
				// search the parent of the head
				for _, p := range node.Parent {
					// if (mapOfElementsTier[p.Ingridient1.Name] > mapOfElementsTier[head.Name]) && (mapOfElementsTier[p.Ingridient2.Name] > mapOfElementsTier[head.Name]) {
					// creating tree out of parentNode
					ing1 := &model.Tree{
						Name:     p.Ingridient1.Name,
						Children: []*model.Tree{},
					}
					ing2 := &model.Tree{
						Name:     p.Ingridient2.Name,
						Children: []*model.Tree{},
					}

					// bind those parent with the head
					head.Tree.Children = append(head.Tree.Children, ing1, ing2)

					// add those parent into queue
					BFSQueue = append(BFSQueue, QueueItem{Name: ing1.Name, Tree: ing1})
					BFSQueue = append(BFSQueue, QueueItem{Name: ing2.Name, Tree: ing2})

					(*response).NumOfRecipe++

					if mode == 1 && response.NumOfRecipe >= *askedNumOfRecipes { // first found
						return
					}

					// }
				}
			}
		}
	}
}
