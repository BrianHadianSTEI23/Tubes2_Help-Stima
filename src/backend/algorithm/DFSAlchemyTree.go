package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func DFSAlchemyTree(target string, t []*model.AlchemyTree, r *model.Response, askedNumOfRecipes int64, mode int8) *model.Response {
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
			// creating the child-parent and add the parent
			// for _, parentPair := range (*t).Parent {
			// 	var recipe []string
			// 	recipe = append(recipe, parentPair.Ingridient1.Name)
			// 	recipe = append(recipe, parentPair.Ingridient2.Name)
			// 	recipe = append(recipe, (*t).Name)
			// 	(*r).Data = append((*r).Data, recipe)
			// }
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
	return r
}
