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
	for _, root := range rootElements {
		model.DisplayAlchemyTree(root)
	}

	// main algorithm
	fmt.Println("Give me your target : ")
	fmt.Scanln(&target)
	fmt.Println("Choose mode : ")
	fmt.Println("1. DFS ")
	fmt.Println("2. BFS ")
	for mode != 1 && mode != 2 {
		fmt.Println("Please enter a number...")
		fmt.Scanln(&mode)
	}

	// choosing mode : multiple recipe or shortest recipe

	if mode == 1 {
		if algorithm.DFSAlchemyTree(target, rootElements) {
			fmt.Println("found!")
			fmt.Println("Getting recipes...")
			return
		}
		fmt.Print("Not found!")
	} else if mode == 2 {
		if algorithm.BFSAlchemyTree(target, rootElements) {
			fmt.Println("found!")
			fmt.Println("Getting recipes...")
			return
		}
		fmt.Print("Not found!")

	}
}
