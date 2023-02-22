package models

func IsMerkleTreeValid(tree *MerkleTree) bool {
	return CompareParentHashToChildrenHashes(tree.Root)
}

func CompareParentHashToChildrenHashes(parent *TreeNode) bool {
	if len(parent.Children) == 0 { // leaf node
		return true
	} else {
		retval := true
		childrenHashes := make([]Hash, 0)
		for _, c := range parent.Children {
			childrenHashes = append(childrenHashes, c.NodeHash)
			retval = retval && CompareParentHashToChildrenHashes(c)
		}

		expectedHash := genParentHashFromChildrenHashes(childrenHashes)
		retval = retval && AreHashesEqual(expectedHash, parent.NodeHash)
		return retval
	}
}

func InvalidateTree(tree *MerkleTree) {
	currentNode := tree.Root
	for len(currentNode.Children) > 0 {
		currentNode = currentNode.Children[0]
	}
	currentNode.NodeHash = Hash{Value: make([]byte, KECCAK_SHA_LENGTH)}
}
