package main

import (
	"encoding/json"
	"io"
	"littlealchemy2/algorithm"
	"littlealchemy2/model"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// initialize all variables
// variables
var target string
var mode = new(int)
var searchAlgorithm = new(int)
var listOfAllRecipes [][]string
var listOfCreatedNodes []*model.AlchemyTree
var mapOfElementsTier map[string]int
var rootElements []*model.AlchemyTree

// var listOfElementImage map[string]string

// var numOfRecipesFound = new(int64)
var getRequest = &model.GetRequest{}
var response = &model.Response{}

func init_main() {
	// reading file
	listOfAllRecipes, _, mapOfElementsTier = algorithm.Scraper()

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

	// assignment
	target = (*getRequest).Target
	(*searchAlgorithm) = (*getRequest).Algorithm
	(*mode) = (*getRequest).Mode

	// initialize the main tree
	response.Data.Name = target
	response.NumOfRecipe = 0

	// choosing searching algorithm : DFS or BFS
	if *searchAlgorithm == 1 {
		// get how many num of recipes is being asked
		var askedNumOfRecipes int64 = (int64)((*getRequest).MaxRecipes)
		var currentFoundRecipe int64 = 0

		// create initial base node
		baseNode := &model.Tree{}

		// doing search algorithm
		algorithm.DFSAlchemyTree(target, listOfCreatedNodes, (int8)(*mode), &askedNumOfRecipes, baseNode, mapOfElementsTier, &currentFoundRecipe)

		// final handling
		response.NumOfRecipe = currentFoundRecipe
		response.Data = *baseNode
	} else if *searchAlgorithm == 2 {
		// get how many num of recipes is being asked
		var askedNumOfRecipes int64 = (int64)((*getRequest).MaxRecipes)

		// create initial base node
		baseNode := &model.Tree{
			Name:     target,
			Children: []*model.Tree{},
		}
		response.Data = *baseNode

		// doing search algorithm
		algorithm.BFSAlchemyTree(target, listOfCreatedNodes, (int8)(*mode), &askedNumOfRecipes, response, mapOfElementsTier)
	}

	log.Println(*response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// main function
func main() {

	// // initialize the main graph
	init_main()

	// // // BACKEND API
	mux := http.NewServeMux()
	mux.HandleFunc("/api/post-recipe", postHandler)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	}).Handler(mux)
	log.Println("Server is running at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
