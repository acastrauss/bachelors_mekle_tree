package models

import (
	"math/rand"

	"github.com/wealdtech/go-merkletree/keccak256"
)

var keccakHasher = keccak256.New()

const STRING_VALUE_LENGTH = 50

func GenerateBinaryTree(treeParams TreeParams) MerkleTree {
	nofLeafNodes := getNumberOfLeafNodes(treeParams)
	currentLevelNofNodes := nofLeafNodes

	var prevLevelNodes []TreeNode

	for currentLevelNofNodes > 0 {
		// generate nodes for current level
		if len(prevLevelNodes) == 0 {
			// leafs
			prevLevelNodes = generateLeafNodes(currentLevelNofNodes, STRING_VALUE_LENGTH)
		} else {
			var currentLevelNodes []TreeNode

			for i := 0; i < int(currentLevelNofNodes); i++ {
				oneParent := generateParentOfChildren(treeParams, &prevLevelNodes)
				currentLevelNodes = append(currentLevelNodes, oneParent)
			}

			prevLevelNodes = currentLevelNodes
		}
		currentLevelNofNodes = int(currentLevelNofNodes / 2)
	}

	return MerkleTree{
		Root:     &prevLevelNodes[0],
		RootHash: prevLevelNodes[0].NodeHash,
	}
}

func generateLeafNodes(nOfLeafNodes int, stringValueLength int) []TreeNode {
	var emptyHashes []Hash
	var leafNodes []TreeNode
	for i := 0; i < int(nOfLeafNodes); i++ {
		leafNodes = append(leafNodes, buildMerkleNode(emptyHashes, stringValueLength))
	}
	return leafNodes
}

func generateParentOfChildren(treeParams TreeParams, availableChildren *[]TreeNode) TreeNode {
	var childrenOfCurrentNode []*TreeNode
	var childrenHashes []Hash
	for j := 0; j < (int(treeParams.TreeIndex)); j++ {
		oneChild := getChildFromAvailableChildren(availableChildren)
		childrenOfCurrentNode = append(childrenOfCurrentNode, oneChild)
		childrenHashes = append(childrenHashes, oneChild.NodeHash)
	}
	parent := buildMerkleNode(childrenHashes, STRING_VALUE_LENGTH)
	assignParentToChildren(&parent, childrenOfCurrentNode)
	return parent
}

func getChildFromAvailableChildren(availableChildren *[]TreeNode) *TreeNode {
	oneChild, indx := getRandomNodeFromLevel(*availableChildren)
	*availableChildren = removeFromNodeLevelAtIndex(*availableChildren, indx)
	return oneChild
}

func getRandomNodeFromLevel(nodesAtLevel []TreeNode) (*TreeNode, int) {
	indx := rand.Intn(len(nodesAtLevel))
	return &nodesAtLevel[indx], indx
}

func removeFromNodeLevelAtIndex(level []TreeNode, indx int) []TreeNode {
	return append(level[:indx], level[indx+1:]...)
}
