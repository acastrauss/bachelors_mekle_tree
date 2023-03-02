package generators

import (
	"models"
)

func GenerateMerkleTree(treeParams models.TreeParams) models.MerkleTree {
	nofLeafNodes := treeParams.GetNumberOfLeafNodes()
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
	var leafNodes []models.TreeNode
	for i := 0; i < int(nOfLeafNodes); i++ {
		leafNodes = append(leafNodes, models.BuildLeafNode())
	}
	return leafNodes
}

func generateParentOfChildren(treeParams models.TreeParams, availableChildren *[]models.TreeNode) models.TreeNode {
	var childrenOfCurrentNode []*models.TreeNode
	for j := 0; j < (int(treeParams.TreeIndex)); j++ {
		oneChild := getChildFromAvailableChildren(availableChildren)
		childrenOfCurrentNode = append(childrenOfCurrentNode, oneChild)
	}
	parent := models.BuildBranchNode(childrenOfCurrentNode)
	return parent
}

func getChildFromAvailableChildren(availableChildren *[]models.TreeNode) *models.TreeNode {
	oneChild := (*availableChildren)[0]
	*availableChildren = removeFromNodeLevelAtIndex(*availableChildren, 0)
	return &oneChild
}

func removeFromNodeLevelAtIndex(level []models.TreeNode, indx int) []models.TreeNode {
	return append(level[:indx], level[indx+1:]...)
}
