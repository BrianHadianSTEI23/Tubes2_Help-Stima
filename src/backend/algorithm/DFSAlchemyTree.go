// package algorithm

// import (
// 	"littlealchemy2/model"
// 	"sync"
// 	"sync/atomic"
// )

// func DFSAlchemyTree(target string, t []*model.AlchemyTree, r *model.Response, mode int8, numOfFoundRecipe *int64) {
// 	var wg sync.WaitGroup
// 	var found int32 = 0

// 	var dfs func(t *model.AlchemyTree)
// 	dfs = func(t *model.AlchemyTree) {
// 		defer wg.Done()

// 		if t == nil || atomic.LoadInt32(&found) == 1 {
// 			return
// 		}

// 		if t.Name == target {
// 			atomic.StoreInt32(&found, 1)
// 			// if found, add to count and construct recipe
// 			tempJSON := model.Response{
// 				Status: "Fail",
// 				Data: model.Tree{
// 					NumOfRecipe: 0,
// 					Children:    []string{},
// 				},
// 			}
// 			// RecipeConstructor(t, r, mode, tempJSON)
// 			return
// 		}

// 		for _, parent := range t.Parent {
// 			if atomic.LoadInt32(&found) == 1 {
// 				return
// 			}
// 			wg.Add(1)
// 			go dfs(parent.Ingridient1)
// 			go dfs(parent.Ingridient2)
// 		}
// 	}

// 	for _, root := range t {
// 		if root != nil {
// 			wg.Add(1)
// 			go dfs(root)
// 		}
// 	}

// 	wg.Wait()
// }

package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func DFSAlchemyTree(target string, listOfCreatedNodes []*model.AlchemyTree, r *model.Response, numOfFoundRecipe *int64) {
	var wg sync.WaitGroup
	var found int32 = 0

	var dfs func(t *model.AlchemyTree)
	dfs = func(t *model.AlchemyTree) {
		defer wg.Done()

		if t == nil || atomic.LoadInt32(&found) == 1 {
			return
		}

		if (*t).Name == "Fire" || (*t).Name == "Water" || (*t).Name == "Air" || (*t).Name == "Earth" || (*t).Name == "Time" {
			// STOP : adding the tree into the response
			newTree := model.Tree{
				Name:     (*t).Name,
				Children: []model.Tree{},
			}
			r.NumOfRecipe++
			r.Data.Children = append(r.Data.Children, newTree)
			return
		}
		// adding the tree into the response
		newTree := model.Tree{
			Name:     (*t).Name,
			Children: []model.Tree{},
		}

		// traverse parent
		for _, parent := range t.Parent {
			if atomic.LoadInt32(&found) == 1 {
				return
			}
			wg.Add(1)
			dfs(parent.Ingridient1)
			wg.Add(1)
			dfs(parent.Ingridient2)
		}

		// appending the parent while in recursive
		r.Data.Children = append(r.Data.Children, newTree)
	}

	// construct the recipe
	for _, n := range listOfCreatedNodes {
		if n.Name == target && (*r).NumOfRecipe < (*numOfFoundRecipe) {
			wg.Add(1)
			go dfs(n)
		}
	}

	wg.Wait()
}
