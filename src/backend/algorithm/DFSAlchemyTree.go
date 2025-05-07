package algorithm

import (
	"littlealchemy2/model"
	"sync"
	"sync/atomic"
)

func DFSAlchemyTree(target string, t []*model.AlchemyTree) bool {
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
	return atomic.LoadInt32(&found) == 1
}
