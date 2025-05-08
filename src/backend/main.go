package main

import (
	"bufio"
	"fmt"
	"littlealchemy2/algorithm"
	"littlealchemy2/model"
	"os"
	"strings"
)

func main() {
	// variables
	var target string
	var mode int
	mode = 0
	var searchAlgorithm int
	searchAlgorithm = 0
	var listOfAllRecipes [][]string
	var listOfCreatedNodes []*model.AlchemyTree
	var rootElements []*model.AlchemyTree

	// reading file
	file, err := os.Open("./data/little_alchemy_2_elements_split.csv")
	if err != nil {
		fmt.Println("File is invalid")
		return
	}
	defer file.Close()

	// output content to the terminal
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		listOfAllRecipes = append(listOfAllRecipes, line)
	}

	/* build tree that intertwined all possible recipes:
	start with 5 basic elements, basic tree structure : name, companion, parent, and children
	*/

	// building 5 root elements
	fireAlchemyTree := model.AlchemyTree{
		Name:      "Fire",
		Parent:    nil,
		Companion: nil,
		Children:  nil,
	}
	waterAlchemyTree := model.AlchemyTree{
		Name:      "Water",
		Parent:    nil,
		Companion: nil,
		Children:  nil,
	}
	airAlchemyTree := model.AlchemyTree{
		Name:      "Air",
		Parent:    nil,
		Companion: nil,
		Children:  nil,
	}
	earthAlchemyTree := model.AlchemyTree{
		Name:      "Earth",
		Parent:    nil,
		Companion: nil,
		Children:  nil,
	}
	timeAlchemyTree := model.AlchemyTree{
		Name:      "Time",
		Parent:    nil,
		Companion: nil,
		Children:  nil,
	}

	rootElements = append(rootElements, &fireAlchemyTree)
	rootElements = append(rootElements, &earthAlchemyTree)
	rootElements = append(rootElements, &waterAlchemyTree)
	rootElements = append(rootElements, &airAlchemyTree)
	rootElements = append(rootElements, &timeAlchemyTree)

	listOfCreatedNodes = append(listOfCreatedNodes, &fireAlchemyTree)
	listOfCreatedNodes = append(listOfCreatedNodes, &earthAlchemyTree)
	listOfCreatedNodes = append(listOfCreatedNodes, &waterAlchemyTree)
	listOfCreatedNodes = append(listOfCreatedNodes, &airAlchemyTree)
	listOfCreatedNodes = append(listOfCreatedNodes, &timeAlchemyTree)

	algorithm.BuildAlchemyTree(rootElements, &listOfAllRecipes, &listOfCreatedNodes)

	// debug
	// for _, root := range rootElements {
	// 	model.DisplayAlchemyTree(root)
	// }

	// main algorithm
	fmt.Println("Give me your target : ")
	fmt.Scanln(&target)
	fmt.Println("Choose algorithm : ")
	fmt.Println("1. DFS ")
	fmt.Println("2. BFS ")
	for searchAlgorithm != 1 && searchAlgorithm != 2 {
		fmt.Println("Please enter a number...")
		fmt.Scanln(&searchAlgorithm)
	}
	fmt.Println("Choose mode : ")
	fmt.Println("1. Shortest Path ")
	fmt.Println("2. Multiple Recipe ")
	for mode != 1 && mode != 2 {
		fmt.Println("Please enter a number...")
		fmt.Scanln(&mode)
	}

	// initializing JSON response
	var response *model.Response
	(*response).Status = "Fail"

	// choosing searching algorithm : DFS or BFS
	if mode == 1 {
		// choosing mode : multiple recipe or shortest recipe
		outputJSON := algorithm.SearchShortestRecipe(listOfCreatedNodes, target, response, int8(searchAlgorithm))
		// if len(outputJSON) != 0 { // need to be implemented for new data structures
		// 	outputJSON.Status = "Success"
		// }

		// send the output JSON [NOT IMPLEMENTED]
	} else if mode == 2 {
		// get how many recipes are asked
		var askedNumOfRecipes int64
		fmt.Scanln(&askedNumOfRecipes)

		// doing search algorithm
		outputJSON := algorithm.SearchMultipleRecipe(listOfCreatedNodes, target, response, int8(searchAlgorithm), askedNumOfRecipes)
		// if len(outputJSON.Data) != 0 { // need to be implemented for new data structures
		// 	outputJSON.Status = "Success"
		// }

		// send the output JSON [NOT IMPLEMENTED]

	}
}
