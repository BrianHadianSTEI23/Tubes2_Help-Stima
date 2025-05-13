package util

import (
	"fmt"
	"littlealchemy2/model"
)

func Debug(listOfCreatedNodes []*model.AlchemyTree, target string, mode *int, searchAlgorithm *int) {
	// debug
	for _, node := range listOfCreatedNodes {
		fmt.Println((*node).Name)
		fmt.Println("PARENTS")
		for _, p := range node.Parent {
			if (*p).Ingridient1 != nil {
				fmt.Print((*p).Ingridient1.Name + " | ")
			} else {
				fmt.Print("nil")
			}
			if (*p).Ingridient2 != nil {
				fmt.Print((*p).Ingridient2.Name + " | ")
			} else {
				fmt.Print("nil")
			}
		}
		fmt.Println()
		fmt.Println("CHILDREN")
		for _, c := range node.Children {
			if c != nil {
				fmt.Print((*c).Name + " | ")
			} else {
				fmt.Print("nil")
			}
			if c != nil {
				fmt.Print((*c).Name + " | ")
			} else {
				fmt.Print("nil")
			}
		}
		fmt.Println()
		fmt.Println("COMPANION")
		for _, co := range node.Companion {
			if co != nil {
				fmt.Print((*co).Name + " | ")
			} else {
				fmt.Print("nil")
			}
			if co != nil {
				fmt.Print((*co).Name + " | ")
			} else {
				fmt.Print("nil")
			}
		}
		fmt.Println()

		model.DisplayAlchemyTree(node)
	}

	// main algorithm
	fmt.Println("Give me your target : ")
	fmt.Scanln(&target)
	fmt.Println("Choose algorithm : ")
	fmt.Println("1. DFS ")
	fmt.Println("2. BFS ")
	for *searchAlgorithm != 1 && *searchAlgorithm != 2 {
		fmt.Println("Please enter a number...")
		fmt.Scanln(searchAlgorithm)
	}
	fmt.Println("Choose mode : ")
	fmt.Println("1. Shortest Path ")
	fmt.Println("2. Multiple Recipe ")
	for *mode != 1 && *mode != 2 {
		fmt.Println("Please enter a number...")
		fmt.Scanln(mode)
	}
}
