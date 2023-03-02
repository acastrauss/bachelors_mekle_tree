package validators

import (
	"models"
	"reflect"
)

func IsMerkleTreeValid(tree *models.MerkleTree) bool {
	return compareParentHashToChildrenHashes(tree.Root)
}

func compareParentHashToChildrenHashes(parent *models.TreeNode) bool {
	if len(parent.Children) == 0 { // leaf node
		return true
	} else {
		retval := true
		for _, c := range parent.Children {
			retval = retval && compareParentHashToChildrenHashes(c)
		}
		excpectedHash := models.GetParentHashFromChildren(parent.Children)
		retval = retval && reflect.DeepEqual(excpectedHash.Value, parent.NodeHash.Value)
		return retval
	}
}

func InvalidateTree(tree *models.MerkleTree, invalidData models.NodeData) {
	currentNode := tree.Root
	for len(currentNode.Children) > 0 {
		currentNode = currentNode.Children[0]
	}
	currentNode.Data = invalidData
	currentNode.NodeHash = models.Hash{Value: models.KeccakHasher.Hash(invalidData.GetBytes())}
}

func AreMerkleTreesNodesDifferent(tree *models.MerkleTree, treeParams models.TreeParams) bool {
	expectedNofNodes := treeParams.GetTotalNumberOfNodes()
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
