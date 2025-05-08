package algorithm

import (
	"littlealchemy2/model"
)

func RecipeConstructor(target *model.AlchemyTree, returnJSON *model.Response, mode int8) {
	/*
		1. do loop for all parent of target
		2. for each parent, do recursive
		3. if target is fire, earth, air, water, or time -> check the other parent
		4. else do loop again
	*/

	// main algorithm
	if target != nil {
		// change mode : search shortest first or multiple recipe
		if mode == 1 { // shortest path
			// iterate for each parent
			for _, parent := range (*target).Parent {
				// generate new response model
				newReturnJSON := &model.Response{
					Status: "",
					Node:   []string{},
					Edge:   [][]string{},
				}
				if newReturnJSON != nil {
					// setting up the status
					(*newReturnJSON).Status = "Fail"

					// turn the pair parent into array
					var edge []string
					edge = append(edge, (*parent).Ingridient1.Name)
					edge = append(edge, (*parent).Ingridient2.Name)

					// bind edge and node into newReturnJSON
					(*newReturnJSON).Edge = append((*newReturnJSON).Edge, edge)
					(*newReturnJSON).Node = append((*newReturnJSON).Node, (*target).Name)
					(*newReturnJSON).NumOfRecipe++

					// check again for doing recursive the first parent
					if (*parent).Ingridient1 == nil || (*parent).Ingridient1.Name == "Fire" || (*parent).Ingridient1.Name == "Water" || (*parent).Ingridient1.Name == "Air" || (*parent).Ingridient1.Name == "Earth" || (*parent).Ingridient1.Name == "Time" {
						// check the second parent
						if (*parent).Ingridient2 == nil || (*parent).Ingridient2.Name == "Fire" || (*parent).Ingridient2.Name == "Water" || (*parent).Ingridient2.Name == "Air" || (*parent).Ingridient2.Name == "Earth" || (*parent).Ingridient2.Name == "Time" {
							// return the json
						} else { // do recursive of recipe constructor on the second parent
							RecipeConstructor(parent.Ingridient2, newReturnJSON, mode)
						}
					} else {
						RecipeConstructor(parent.Ingridient1, newReturnJSON, mode)
					}

					// bind the newResponseJSON into the response model IF IT'S SMALLER THAN CURRENT ONE
					if len((*returnJSON).Edge) > len((*newReturnJSON).Edge) {
						returnJSON.Edge = newReturnJSON.Edge
						returnJSON.Node = newReturnJSON.Node
						(*returnJSON).Status = "Success"
					}
				}
			}
		} else if mode == 2 { // multiple recipe
			// iterate for each parent
			for _, parent := range (*target).Parent {
				// turn the pair parent into array
				var edge []string
				edge = append(edge, (*parent).Ingridient1.Name)
				edge = append(edge, (*parent).Ingridient2.Name)

				// add the edge and node into the response model
				(*returnJSON).Edge = append((*returnJSON).Edge, edge)
				(*returnJSON).Node = append((*returnJSON).Node, (*target).Name)

				// check the first parent
				if (*parent).Ingridient1 == nil || (*parent).Ingridient1.Name == "Fire" || (*parent).Ingridient1.Name == "Water" || (*parent).Ingridient1.Name == "Air" || (*parent).Ingridient1.Name == "Earth" || (*parent).Ingridient1.Name == "Time" {
					// check the second parent
					if (*parent).Ingridient2 == nil || (*parent).Ingridient2.Name == "Fire" || (*parent).Ingridient2.Name == "Water" || (*parent).Ingridient2.Name == "Air" || (*parent).Ingridient2.Name == "Earth" || (*parent).Ingridient2.Name == "Time" {
						// return the json
					} else { // do recursive of recipe constructor on the second parent
						RecipeConstructor(parent.Ingridient2, returnJSON, mode)
					}
				} else {
					RecipeConstructor(parent.Ingridient1, returnJSON, mode)
				}

				// mark the responseJSON to be success
				(*returnJSON).Status = "Success"
			}
		}
	}
}
