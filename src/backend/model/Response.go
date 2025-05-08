package model

type Response struct {
	Status      string
	NumOfRecipe int64
	Node        []string
	Edge        [][]string
}
