package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"littlealchemy2/algorithm"
	"littlealchemy2/model"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Text string `json:"text"`
}

// define endpoint API link
func GETHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message{Text: "Hello world!"})
}

// main function
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
	numOfRecipesFound := new(int64) // allocates memory
	*numOfRecipesFound = 0          // sets the value

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
	response := new(model.Response)
	(*response).Status = "Fail"

	// choosing searching algorithm : DFS or BFS
	if searchAlgorithm == 1 {
		algorithm.DFSAlchemyTree(target, rootElements, response, int8(mode), numOfRecipesFound)
	} else if searchAlgorithm == 2 {
		// get how many num of recipes is being asked
		var askedNumOfRecipes int64
		fmt.Println("How many recipe do you want?")
		fmt.Scanln(&askedNumOfRecipes)

		// doing search algorithm
		algorithm.BFSAlchemyTree(target, rootElements, response, int8(mode), numOfRecipesFound)
		if (*numOfRecipesFound) > askedNumOfRecipes {
			// change the number of recipes found in the response model
			(*response).NumOfRecipe = askedNumOfRecipes
		} else {
			(*response).NumOfRecipe = (*numOfRecipesFound)
		}
	}

	// debug
	model.DisplayResponse(response)

	// BACKEND API
	// http.HandleFunc("api/get-recipe", GETHandler)
	// log.Println("Go API running on https://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// send the output JSON [NOT IMPLEMENTED]
}
