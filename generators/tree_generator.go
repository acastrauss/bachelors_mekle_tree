package generators

import (
	"fmt"
	"math/rand"
	"models"
)

func GenerateMerkleTree(treeParams models.TreeParams) models.MerkleTree {
	nofLeafNodes := models.GetNumberOfLeafNodes(treeParams)
	currentLevelNofNodes := nofLeafNodes

	var prevLevelNodes []models.TreeNode

	for currentLevelNofNodes > 0 {
		// generate nodes for current level
		if len(prevLevelNodes) == 0 {
			// leafs
			prevLevelNodes = generateLeafNodes(currentLevelNofNodes)
		} else {
			var currentLevelNodes []models.TreeNode

			for i := 0; i < int(currentLevelNofNodes); i++ {
				oneParent := generateParentOfChildren(treeParams, &prevLevelNodes)
				currentLevelNodes = append(currentLevelNodes, oneParent)
			}

			prevLevelNodes = currentLevelNodes
		}
		currentLevelNofNodes = int(currentLevelNofNodes / int(treeParams.TreeIndex))
	}

	return models.MerkleTree{
		Root:     &prevLevelNodes[0],
		RootHash: prevLevelNodes[0].NodeHash,
	}
}

func generateLeafNodes(nOfLeafNodes int) []models.TreeNode {
	var emptyHashes []models.Hash
	var leafNodes []models.TreeNode
	for i := 0; i < int(nOfLeafNodes); i++ {
		leafNodes = append(leafNodes, models.BuildMerkleNode(emptyHashes, models.STRING_VALUE_LENGTH))
	}
	return leafNodes
}

func generateParentOfChildren(treeParams models.TreeParams, availableChildren *[]models.TreeNode) models.TreeNode {
	var childrenOfCurrentNode []*models.TreeNode
	var childrenHashes []models.Hash
	for j := 0; j < (int(treeParams.TreeIndex)); j++ {
		oneChild := getChildFromAvailableChildren(availableChildren)
		childrenOfCurrentNode = append(childrenOfCurrentNode, oneChild)
		childrenHashes = append(childrenHashes, oneChild.NodeHash)
	}
	parent := models.BuildMerkleNode(childrenHashes, models.STRING_VALUE_LENGTH)

	models.AssignParentToChildren(&parent, childrenOfCurrentNode)
	return parent
}

func getChildFromAvailableChildren(availableChildren *[]models.TreeNode) *models.TreeNode {
	if len(*availableChildren) == 0 {
		fmt.Println()
	}

	oneChild := (*availableChildren)[0]
	*availableChildren = removeFromNodeLevelAtIndex(*availableChildren, 0)
	return &oneChild
}

func getRandomNodeFromLevel(nodesAtLevel []models.TreeNode) (*models.TreeNode, int) {
	indx := rand.Intn(len(nodesAtLevel))
	return &nodesAtLevel[indx], indx
}

func removeFromNodeLevelAtIndex(level []models.TreeNode, indx int) []models.TreeNode {
	return append(level[:indx], level[indx+1:]...)
}
