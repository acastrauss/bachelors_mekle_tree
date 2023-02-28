package validators

import (
	"models"
)

func IsMerkleTreeValid(tree *models.MerkleTree) bool {
	return compareParentHashToChildrenHashes(tree.Root)
}

func compareParentHashToChildrenHashes(parent *models.TreeNode) bool {
	if len(parent.Children) == 0 { // leaf node
		return true
	} else {
		retval := true
		childrenHashes := make([]models.Hash, 0)
		for _, c := range parent.Children {
			childrenHashes = append(childrenHashes, c.NodeHash)
			retval = retval && compareParentHashToChildrenHashes(c)
		}

		expectedHash := models.GenParentHashFromChildrenHashes(childrenHashes)
		retval = retval && models.AreHashesEqual(expectedHash, parent.NodeHash)
		return retval
	}
}

func InvalidateTree(tree *models.MerkleTree) {
	currentNode := tree.Root
	for len(currentNode.Children) > 0 {
		currentNode = currentNode.Children[0]
	}
	currentNode.NodeHash = models.Hash{Value: make([]byte, models.KECCAK_SHA_LENGTH)}
}

func AreMerkleTreesNodesDifferent(tree *models.MerkleTree, treeParams models.TreeParams) bool {
	expectedNofNodes := models.GetTotalNumberOfNodes(treeParams)
	idsInTree := getSubtreeIds(tree.Root)
	idsInTree = removeDuplicates(idsInTree)

	return expectedNofNodes == len(idsInTree)
}

func getSubtreeIds(root *models.TreeNode) []int {
	retval := make([]int, 0)

	if len(root.Children) == 0 {
		return []int{root.NodeId}
	} else {
		retval = append(retval, root.NodeId)
		for _, c := range root.Children {
			retval = append(retval, getSubtreeIds(c)...)
		}
	}

	return retval
}

func removeDuplicates(slice []int) []int {
	allKeys := make(map[int]bool)
	newSlice := make([]int, 0)
	for _, item := range slice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}
