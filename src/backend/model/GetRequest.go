package model

type GetRequest struct {
	Target     string
	Algorithm  int
	Mode       int
	MaxRecipes int
}
