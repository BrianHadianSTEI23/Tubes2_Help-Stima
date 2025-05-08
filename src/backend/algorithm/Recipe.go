package algorithm

import (
	"littlealchemy2/model"
)

func RecipeConstructor(target *model.AlchemyTree, returnJSON *model.Response) bool {
	/*
		1. do loop for all parent of target
		2. for each parent, do recursive
		3. if target is fire, earth, air, water, or time -> check the other parent
		4. else do loop again
	*/

	// main algorithm
	if target != nil {
		// iterate for each parent
		for _, parent := range (*target).Parent {
			// turn the pair parent into array
			// var graphData [3]string
			// graphData[0] = (*parent).Ingridient1.Name
			// graphData[1] = (*parent).Ingridient2.Name
			// graphData[2] = (*target).Name
			// // add the graphdata into the response model
			// (*returnJSON).Data = append((*returnJSON).Data, graphData[:])
			// check the first parent
			if (*parent).Ingridient1 == nil || (*parent).Ingridient1.Name == "Fire" || (*parent).Ingridient1.Name == "Water" || (*parent).Ingridient1.Name == "Air" || (*parent).Ingridient1.Name == "Earth" || (*parent).Ingridient1.Name == "Time" {
				// check the second parent
				if (*parent).Ingridient2 == nil || (*parent).Ingridient2.Name == "Fire" || (*parent).Ingridient2.Name == "Water" || (*parent).Ingridient2.Name == "Air" || (*parent).Ingridient2.Name == "Earth" || (*parent).Ingridient2.Name == "Time" {
					// return the json
					return true
				} else { // do recursive of recipe constructor on the second parent
					RecipeConstructor(parent.Ingridient2, returnJSON)
				}
			} else {
				RecipeConstructor(parent.Ingridient1, returnJSON)
			}
		}
	}
	return false

}
