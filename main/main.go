package main

import (
	"fmt"
	"generators"
	"log"
	"models"
	"validators"

	"github.com/barkimedes/go-deepcopy"
)

func PrintLeafFound(leaf *models.TreeNode) {
	if leaf != nil {
		fmt.Printf("Leaf with id:%d has data: '%s'\n", leaf.NodeId, leaf.Data.Data)
	} else {
		fmt.Println("No leaf was found in tree for given data")
	}
}

func main() {
	var treeParams = models.TreeParams{
		TreeIndex:        2,
		PowerOfTreeIndex: 2,
	}
	tree := generators.GenerateMerkleTree(treeParams)
	fmt.Printf("Root hash:%v\n", tree.RootHash)
	fmt.Printf("Is tree valid :%v\n", validators.IsMerkleTreeValid(&tree))
	invalidTreeCopied, err := deepcopy.Anything(tree)
	if err != nil {
		log.Fatal(err)
	} else {
		invalidTree := invalidTreeCopied.(models.MerkleTree)
		validators.InvalidateTree(&invalidTree)
		fmt.Printf("Is tree valid :%v\n", validators.IsMerkleTreeValid(&invalidTree))
	}

	fmt.Printf("Are all nodes different: %v\n", validators.AreMerkleTreesNodesDifferent(&tree, treeParams))
	dataFromOneLeaf := validators.GetDataFromOneLeaf(&tree)
	fmt.Printf("Searching nodes for data:'%s'\n", dataFromOneLeaf.Data)
	leafWithGivenData := validators.GetLeafThatHasData(dataFromOneLeaf, &tree)
	PrintLeafFound(leafWithGivenData)

	fmt.Printf("Is data:'%s' valid in tree:%v\n", dataFromOneLeaf.Data, validators.IsDataValidWithinTree(dataFromOneLeaf, tree))

	dataNotInTree := models.NodeData{
		Data: "some data not in tree",
	}
	fmt.Printf("Searching nodes for data:'%s'\n", dataNotInTree.Data)
	leafWithGivenData = validators.GetLeafThatHasData(dataNotInTree, &tree)
	PrintLeafFound(leafWithGivenData)

	fmt.Printf("Is data:'%s' valid in tree:%v\n", dataNotInTree.Data, validators.IsDataValidWithinTree(dataNotInTree, tree))
}
