package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func BFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, r *model.Response, askedNumOfRecipes *int64) {
	var wg sync.WaitGroup
	var found int32 = 0

	var bfs func(t *model.AlchemyTree)
	bfs = func(t *model.AlchemyTree) {
		defer wg.Done()

		if t == nil || atomic.LoadInt32(&found) == 1 {
			return
		}

		if (*t).Name == "Fire" || (*t).Name == "Water" || (*t).Name == "Air" || (*t).Name == "Earth" || (*t).Name == "Time" {
			// STOP : adding the tree into the response
			newTree := &model.Tree{
				Name:     (*t).Name,
				Children: []*model.Tree{},
			}
			r.NumOfRecipe++
			r.Data.Children = append(r.Data.Children, newTree)
			return
		}

		// adding the tree into the response
		newTree := &model.Tree{
			Name:     (*t).Name,
			Children: []*model.Tree{},
		}

		// traverse parent
		for _, parent := range t.Parent {
			if atomic.LoadInt32(&found) == 1 {
				return
			}
			wg.Add(1)
			go bfs(parent.Ingridient1)
			wg.Add(1)
			go bfs(parent.Ingridient2)
		}

		// appending the parent while in recursive
		r.Data.Children = append(r.Data.Children, newTree)
	}

	// construct the recipe
	for _, n := range listOfCreatedNodes {
		if n.Name == target && (*r).NumOfRecipe < (*askedNumOfRecipes) {
			wg.Add(1)
			go bfs(n)
		}
	}

	wg.Wait()
}
