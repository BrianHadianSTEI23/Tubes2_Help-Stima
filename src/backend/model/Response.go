package model

import "fmt"

type Response struct {
	Status      string
	NumOfRecipe int64
	Node        []string
	Edge        [][]string
}

func DisplayResponse(r *Response) {
	fmt.Println((*r).Status)
	fmt.Println((*r).NumOfRecipe)
	for _, n := range (*r).Node {
		fmt.Println(n)
	}
	for _, e := range (*r).Edge {
		fmt.Println(e)
	}
}
