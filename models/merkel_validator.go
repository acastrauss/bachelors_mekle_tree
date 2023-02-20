package models

import (
	"fmt"
	"reflect"
)

func ValidateMerkleTree(tree *MerkleTree) bool {
	return CompareParentHashToChildrenHashes(tree.Root)
}

var stopAtNode = -1

func CompareParentHashToChildrenHashes(parent *TreeNode) bool {
	if len(parent.Children) == 0 { // leaf node
		return true
	} else {
		retval := true
		childrenHashes := make([]Hash, 0)
		for _, c := range parent.Children {
			if c.NodeId == stopAtNode {
				fmt.Printf("%d", c.NodeId)
			}
			childrenHashes = append(childrenHashes, c.NodeHash)
			retval = retval && CompareParentHashToChildrenHashes(c)
		}

		expectedHash := genParentHashFromChildrenHashes(childrenHashes)
		return retval && reflect.DeepEqual(expectedHash, parent.NodeHash)
	}
}

func InvalidateTree(tree *MerkleTree) {
	currentNode := tree.Root
	for len(currentNode.Children) > 0 {
		currentNode = currentNode.Children[0]
	}
	stopAtNode = currentNode.NodeId
	currentNode.NodeHash = Hash{Value: make([]byte, KECCAK_SHA_LENGTH)}
}
