package main

import (
	"fmt"
	"models"
)

func main() {
	var treeParams = models.TreeParams{
		TreeIndex:        2,
		PowerOfTreeIndex: 3,
	}
	tree := models.GenerateMerkleTree(treeParams)
	fmt.Printf("Root hash:%v\n", tree.RootHash)
	fmt.Printf("Is tree valid :%v\n", models.ValidateMerkleTree(&tree))
	models.InvalidateTree(&tree)
	fmt.Printf("Is tree valid :%v\n", models.ValidateMerkleTree(&tree))
}
