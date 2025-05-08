package algorithm

import (
	"littlealchemy2/model"
)

func RecipeConstructor(root *model.AlchemyTree, returnJSON *model.Response, mode int8, tempJSON model.Response) {
	/*
		1. do loop for all parent of root
		2. for each parent, do recursive
		3. if root is fire, earth, air, water, or time -> check the other parent
		4. else do loop again
	*/

	// main algorithm
	if root != nil {
		// bind node into newReturnJSON
		tempJSON.Node = append(tempJSON.Node, (*root).Name)
		// iterate for each parent
		for _, parent := range (*root).Parent {
			// fmt.Println((*parent).Ingridient1.Name)
			// fmt.Println((*parent).Ingridient2.Name)

			// turn the pair parent into array
			edge1 := []string{(*parent).Ingridient1.Name, (*root).Name}
			edge2 := []string{(*parent).Ingridient2.Name, (*root).Name}
			tempJSON.Edge = append(tempJSON.Edge, edge1, edge2)

			// check again for doing recursive the first parent
			if (*parent).Ingridient1 == nil || (*parent).Ingridient1.Name == "Fire" || (*parent).Ingridient1.Name == "Water" || (*parent).Ingridient1.Name == "Air" || (*parent).Ingridient1.Name == "Earth" || (*parent).Ingridient1.Name == "Time" {
				// check the second parent
				if (*parent).Ingridient2 == nil || (*parent).Ingridient2.Name == "Fire" || (*parent).Ingridient2.Name == "Water" || (*parent).Ingridient2.Name == "Air" || (*parent).Ingridient2.Name == "Earth" || (*parent).Ingridient2.Name == "Time" {
					// return the json
					(*returnJSON).NumOfRecipe++
					(*returnJSON).Status = "Success"
					// change mode : search shortest first or multiple recipe
					if mode == 1 { // shortest path
						/*
							1. check if the current returnJSON.Edge has length 0 or not, if yes, put right into it
							2. if the length > 0, check if it's less
						*/
						if (len((*returnJSON).Edge) == 0) || (len((*returnJSON).Edge) > len(tempJSON.Edge)) {
							(*returnJSON).Edge = tempJSON.Edge
							(*returnJSON).Node = tempJSON.Node
						}
					} else if mode == 2 { // multiple recipe
						/* just keep adding into the returnJSON */
						for _, e := range tempJSON.Edge {
							// debug
							// fmt.Println(e[0])
							// fmt.Println(e[1])
							if !IsPairInArray(e, (*returnJSON).Edge) {
								(*returnJSON).Edge = append((*returnJSON).Edge, e)
							}
						}
						for _, n := range tempJSON.Node {
							// debug
							// fmt.Println(n)
							if !IsElementInArray(n, (*returnJSON).Node) {
								(*returnJSON).Node = append((*returnJSON).Node, n)
							}
						}
					}
				} else { // do recursive of recipe constructor on the second parent
					// newTempJSON := model.DeepCopyResponse(tempJSON)
					go RecipeConstructor(parent.Ingridient2, returnJSON, mode, tempJSON)
				}
			} else {
				// newTempJSON := model.DeepCopyResponse(tempJSON)
				go RecipeConstructor(parent.Ingridient1, returnJSON, mode, tempJSON)
			}
		}
	}
}

func IsPairInArray(pair []string, array [][]string) bool {
	for _, a := range array {
		if a[0] == pair[0] && a[1] == pair[1] {
			return true
		}
	}
	return false
}

func IsElementInArray(el string, array []string) bool {
	for _, a := range array {
		if el == a {
			return true
		}
	}
	return false
}
