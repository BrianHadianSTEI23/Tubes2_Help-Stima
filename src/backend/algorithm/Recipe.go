package algorithm

import (
	"littlealchemy2/model"
)

type Response struct {
	Status string
	Data   []string
}

func RecipeConstructor(target *model.AlchemyTree) JSON {
	if target != nil {
		// iterate for each parent
		for _, parent := range (*target).Parent {
			// check the first parent
			if parent[0] != nil || parent[0].Name {

			}
		}
	}

	/*
		1. do loop for all parent of target
		2. for each parent, do recursive
		3. if target is fire, earth, air, water, or time -> check the other parent
		4. else do loop again
	*/

}
