package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func BFSAlchemyTree(target string, t []*model.AlchemyTree, r *model.Response, mode int8, numOfFoundRecipe *int64) {
	var wg sync.WaitGroup
	var found int32 = 0
	var visited sync.Map

	var bfs func(t *model.AlchemyTree)
	bfs = func(t *model.AlchemyTree) {
		defer wg.Done()

		if t == nil || atomic.LoadInt32(&found) == 1 {
			return
		}

		_, loaded := visited.LoadOrStore(t.Name, true)
		if loaded {
			return
		}

		if t.Name == target {
			atomic.StoreInt32(&found, 1)
			// if found, add to count and construct recipe
			tempJSON := model.Response{
				Status:      "Fail",
				NumOfRecipe: 0,
				Node:        []string{},
				Edge:        [][]string{},
			}
			RecipeConstructor(t, r, mode, tempJSON)
			return
		}

		// traverse companion
		for _, comp := range t.Companion {
			if atomic.LoadInt32(&found) == 1 {
				return
			}
			wg.Add(1)
			go bfs(comp)
		}

		// traverse children
		for _, child := range t.Children {
			if atomic.LoadInt32(&found) == 1 {
				return
			}
			wg.Add(1)
			go bfs(child)
		}
	}

	for _, root := range t {
		if root != nil {
			wg.Add(1)
			go bfs(root)
		}
	}

	wg.Wait()
}
