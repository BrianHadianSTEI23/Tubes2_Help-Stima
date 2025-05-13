package model

type Tree struct {
	Name     string
	Children []*Tree
}

type Response struct {
	NumOfRecipe      int64
	TotalVisitedNode int64
	ExecutionTime    int64
	Data             Tree
}

// func DisplayResponse(r *Response) {
// 	fmt.Println((*r).Status)
// 	fmt.Println((*r).NumOfRecipe)
// 	fmt.Printf("Num of node : %d\n", len((*r).Node))
// 	for _, n := range (*r).Node {
// 		fmt.Println(n)
// 	}
// 	fmt.Printf("Num of edge : %d\n", len((*r).Edge))
// 	for _, e := range (*r).Edge {
// 		fmt.Println(e)
// 	}
// }

// func DeepCopyResponse(src Response) Response {
// 	// Copy Node slice
// 	nodeCopy := make([]string, len(src.Node))
// 	copy(nodeCopy, src.Node)

// 	// Copy Edge slice
// 	edgeCopy := make([][]string, len(src.Edge))
// 	for i, pair := range src.Edge {
// 		pairCopy := make([]string, len(pair))
// 		copy(pairCopy, pair)
// 		edgeCopy[i] = pairCopy
// 	}

// 	// Return new deep copied Response
// 	return Response{
// 		Node:        nodeCopy,
// 		Edge:        edgeCopy,
// 		NumOfRecipe: src.NumOfRecipe,
// 		Status:      src.Status,
// 	}
// }
