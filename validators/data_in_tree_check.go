package validators

import (
	"fmt"
	"models"
	"reflect"
)

func GetLeafThatHasData(data models.NodeData, tree *models.MerkleTree) *models.TreeNode {
	hashedData := models.Hash{
		Value: models.KeccakHasher.Hash(data.GetBytes()),
	}

	return getLeafNodeWithData(hashedData, tree.Root)
}

func getLeafNodeWithData(hashedData models.Hash, subtreeRoot *models.TreeNode) *models.TreeNode {
	if len(subtreeRoot.Children) == 0 { // is leaf with data
		if reflect.DeepEqual(subtreeRoot.NodeHash.Value, hashedData.Value) {
			return subtreeRoot
		} else {
			return nil
		}
	} else {
		for _, c := range subtreeRoot.Children {

			if reflect.DeepEqual(make([]byte, models.KECCAK_SHA_LENGTH), c.NodeHash.Value) {
				fmt.Print("asas")
			}

			childThatHasData := getLeafNodeWithData(hashedData, c)
			if childThatHasData != nil {
				return childThatHasData
			}
		}

		return nil
	}
}

func GetDataFromOneLeaf(tree *models.MerkleTree) models.NodeData {
	currentNode := tree.Root
	for len(currentNode.Children) > 0 {
		currentNode = currentNode.Children[0]
	}

	return currentNode.Data
}
