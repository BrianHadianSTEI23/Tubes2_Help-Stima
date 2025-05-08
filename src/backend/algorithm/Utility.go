package algorithm

import "littlealchemy2/model"

// func IsExistInSearchedRecipe(targetRecipe []string, searchedRecipe [][]string) bool {
// 	for i := 0; i < len(searchedRecipe); i++ {
// 		if searchedRecipe[i][0] == targetRecipe[0] && searchedRecipe[i][1] == targetRecipe[1] && searchedRecipe[i][2] == targetRecipe[2] {
// 			return true
// 		}
// 	}
// 	return false
// }

// func SearchAllRelatedRecipe(target string, recipes [][]string, searchedRecipe [][]string) [][]string {
// 	var listOfRecipes [][]string
// 	for _, recipe := range recipes {
// 		if recipe[0] == target && !IsExistInSearchedRecipe(recipe, searchedRecipe) {
// 			listOfRecipes = append(listOfRecipes, recipe)
// 		}
// 	}
// 	return listOfRecipes
// }

// func SearchAllParentRecipe(target string, recipes [][]string, searchedRecipe [][]string) [][]string {
// 	var AllParentRecipe [][]string
// 	for _, recipe := range recipes {
// 		if !IsExistInSearchedRecipe(recipe, searchedRecipe) && (recipe[1] == target || recipe[2] == target) {
// 			AllParentRecipe = append(AllParentRecipe, recipe)
// 		}
// 	}
// 	return AllParentRecipe
// }

// func IsTargetRecipesEmpty(targetRecipes [][]string) bool {
// 	return len(targetRecipes) == 0
// }

// func IsExistInListOfAllElements(element string, listOfAllElement []string) bool {
// 	for _, v := range listOfAllElement {
// 		if v == element {
// 			return true
// 		}
// 	}
// 	return false
// }

func SearchShortestRecipe(listOfNodes []*model.AlchemyTree, target string, Response *model.Response, mode int8) *model.Response {

	if mode == 1 { // DFS
		// if the child name is found, get the companion of the same index in the companion array
		// and then add the combination and child into the response data
		return DFSAlchemyTree(target, listOfNodes, Response, 1, 1)
	} else if mode == 2 { // BFS
		// if the child name is found, get the companion of the same index in the companion array
		// and then add the combination and child into the response data
		return BFSAlchemyTree(target, listOfNodes, Response, 1, 1)
	} else {
		return Response
	}
}

func SearchMultipleRecipe(listOfNodes []*model.AlchemyTree, target string, Response *model.Response, mode int8, askedNumOfRecipes int64) *model.Response {

	if mode == 1 { // DFS
		// if the child name is found, get the companion of the same index in the companion array
		// and then add the combination and child into the response data
		return DFSAlchemyTree(target, listOfNodes, Response, askedNumOfRecipes, 2)
	} else if mode == 2 { // BFS
		// if the child name is found, get the companion of the same index in the companion array
		// and then add the combination and child into the response data
		return BFSAlchemyTree(target, listOfNodes, Response, askedNumOfRecipes, 2)
	} else {
		return Response
	}
}
