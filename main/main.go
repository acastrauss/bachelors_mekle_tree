package main

import (
	"fmt"
	"models"
)

func main() {
	var treeParams = models.TreeParams{
		TreeIndex:        3,
		PowerOfTreeIndex: 4,
	}
	tree := models.GenerateMerkleTree(treeParams)
	fmt.Printf("Root hash:%v\n", tree.RootHash)
	fmt.Printf("Is tree valid :%v\n", models.IsMerkleTreeValid(&tree))
	models.InvalidateTree(&tree)
	fmt.Printf("Is tree valid :%v\n", models.IsMerkleTreeValid(&tree))
}
