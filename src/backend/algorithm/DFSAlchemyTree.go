package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func DFSAlchemyTree(target string, t []*model.AlchemyTree, r *model.Response, mode int8, numOfFoundRecipe *int64) {
	var wg sync.WaitGroup
	var found int32 = 0

	var dfs func(t *model.AlchemyTree)
	dfs = func(t *model.AlchemyTree) {
		defer wg.Done()

		if t == nil || atomic.LoadInt32(&found) == 1 {
			return
		}

		if t.Name == target {
			atomic.StoreInt32(&found, 1)
			// if found, add to count and construct recipe
			if mode == 2 || (mode == 1 && (*numOfFoundRecipe) < 1) { // not shortest path mode
				(*numOfFoundRecipe)++
				RecipeConstructor(t, r, mode)
			}
			return
		}

		for _, child := range t.Children {
			if atomic.LoadInt32(&found) == 1 {
				return
			}
			wg.Add(1)
			go dfs(child)
		}
	}

	for _, root := range t {
		if root != nil {
			wg.Add(1)
			go dfs(root)
		}
	}

	wg.Wait()
}
