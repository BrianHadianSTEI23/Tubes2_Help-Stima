package algorithm

import (
	"fmt"
	"littlealchemy2/model"
	"strconv"
	"sync"
	"sync/atomic"
)

func BFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, mode int8, askedNumOfRecipes *int64, response *model.Response, mapOfElementsTier map[string]int, totalVisitedNode *int64) {

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
		// do concurrency here
		var localWg sync.WaitGroup
		nextLevelChan := make(chan QueueItem, 100)
		stopFlag := int32(0)
		(*totalVisitedNode)++

		for _, item := range BFSQueue {
			localWg.Add(1)
			go func(item QueueItem) {
				defer localWg.Done()

				// stop condition
				if item.Name == "Fire" || item.Name == "Water" || item.Name == "Air" || item.Name == "Earth" || item.Name == "Time" {
					return
				}

				// found the element in created nodes
				for _, node := range listOfCreatedNodes {
					if node == nil || item.Name != node.Name {
						continue
					}
					// search the parent of the head
					for _, p := range node.Parent {
						fmt.Println(p.Ingridient1.Name + " " + strconv.Itoa(mapOfElementsTier[p.Ingridient1.Name]))
						fmt.Println(p.Ingridient2.Name + " " + strconv.Itoa(mapOfElementsTier[p.Ingridient2.Name]))
						fmt.Println(node.Name + " " + strconv.Itoa(mapOfElementsTier[node.Name]))

						if (mapOfElementsTier[p.Ingridient1.Name] <= mapOfElementsTier[item.Name]) || (mapOfElementsTier[p.Ingridient2.Name] <= mapOfElementsTier[item.Name]) {

							if mode == 1 && atomic.LoadInt32(&stopFlag) == 1 {
								return
							}

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
							item.Tree.Children = append(item.Tree.Children, ing1, ing2)

							// add those parent into queue
							nextLevelChan <- QueueItem{Name: ing1.Name, Tree: ing1}
							nextLevelChan <- QueueItem{Name: ing2.Name, Tree: ing2}

							atomic.AddInt64(&response.NumOfRecipe, 1)

							if mode == 1 && response.NumOfRecipe >= *askedNumOfRecipes { // first found
								atomic.StoreInt32(&stopFlag, 1)
								return
							}

						}
					}
				}
			}(item)
		}
		localWg.Wait()
		close(nextLevelChan)

		// Build next level from the channel
		var nextLevel []QueueItem
		for item := range nextLevelChan {
			nextLevel = append(nextLevel, item)
		}

		// Stop if early exit condition met
		if mode == 1 && atomic.LoadInt32(&stopFlag) == 1 {
			break
		}

		BFSQueue = nextLevel
	}
}
