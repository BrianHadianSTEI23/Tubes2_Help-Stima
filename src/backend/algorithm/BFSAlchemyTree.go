package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func BFSAlchemyTree(target string, t []*model.AlchemyTree, r *model.Response, askedNumOfRecipes int64, mode int8) *model.Response {
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
			// construct the recipe and add into the response

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
	return r

}
