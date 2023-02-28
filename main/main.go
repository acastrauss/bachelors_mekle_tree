package main

import (
	"fmt"
	"generators"
	"models"
	"validators"
)

func main() {
	var treeParams = models.TreeParams{
		TreeIndex:        3,
		PowerOfTreeIndex: 3,
	}
	tree := generators.GenerateMerkleTree(treeParams)
	fmt.Printf("Root hash:%v\n", tree.RootHash)
	fmt.Printf("Is tree valid :%v\n", validators.IsMerkleTreeValid(&tree))
	validators.InvalidateTree(&tree)
	fmt.Printf("Is tree valid :%v\n", validators.IsMerkleTreeValid(&tree))

	fmt.Printf("Are all nodes different: %v", validators.AreMerkleTreesNodesDifferent(&tree, treeParams))
}
