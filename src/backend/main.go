package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"littlealchemy2/algorithm"
	"littlealchemy2/model"
	"net/http"
	"os"
	"strings"
)

// initialize all variables
// variables
var target string
var mode = new(int)
var searchAlgorithm = new(int)
var listOfAllRecipes [][]string
var listOfCreatedNodes []*model.AlchemyTree
var rootElements []*model.AlchemyTree
var numOfRecipesFound = new(int64)
var getRequest = &model.GetRequest{}
var response = &model.Response{Status: "Fail"}

func init_main() {
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

}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Can't read body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, getRequest); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// debug
	// fmt.Fprintf(w, "target: %s\n", getRequest.Target)
	// fmt.Fprintf(w, "algorithm: %d\n", getRequest.Algorithm)
	// fmt.Fprintf(w, "mode: %d\n", getRequest.Mode)
	// fmt.Fprintf(w, "maxRecipes: %d\n", getRequest.MaxRecipes)

	// main algorithm
	// assignment
	target = (*getRequest).Target
	(*searchAlgorithm) = (*getRequest).Algorithm
	(*mode) = (*getRequest).Mode

	// choosing searching algorithm : DFS or BFS
	if *searchAlgorithm == 1 {
		// get how many num of recipes is being asked
		var askedNumOfRecipes int64 = (int64)((*getRequest).MaxRecipes)
		// if (*mode) == 2 { // multiple recipe
		// 	fmt.Println("How many recipe do you want?")
		// 	fmt.Scanln(&askedNumOfRecipes)
		// }

		// doing search algorithm
		algorithm.DFSAlchemyTree(target, rootElements, response, int8(*mode), numOfRecipesFound)

		// final handling
		if (*mode) == 2 {
			if (*response).NumOfRecipe > askedNumOfRecipes {
				// change the number of recipes found in the response model
				(*response).NumOfRecipe = askedNumOfRecipes
			} else {
				fmt.Printf("Found only : %d recipes\n", ((*response).NumOfRecipe))
			}
		}
	} else if *searchAlgorithm == 2 {
		// get how many num of recipes is being asked
		var askedNumOfRecipes int64 = (int64)((*getRequest).MaxRecipes)
		// if (*mode) == 2 { // multiple recipe
		// 	fmt.Println("How many recipe do you want?")
		// 	fmt.Scanln(&askedNumOfRecipes)
		// }

		// doing search algorithm
		algorithm.BFSAlchemyTree(target, rootElements, response, int8(*mode), numOfRecipesFound)

		// final handling
		if (*mode) == 2 {
			if (*response).NumOfRecipe > askedNumOfRecipes {
				// change the number of recipes found in the response model
				(*response).NumOfRecipe = askedNumOfRecipes
			} else {
				fmt.Printf("Found only : %d recipes\n", ((*response).NumOfRecipe))
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// main function
func main() {

	// initialize the main graph
	init_main()

	// // BACKEND API
	http.HandleFunc("/api/post-recipe", postHandler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

	// debug
	// model.DisplayResponse(response)
}
