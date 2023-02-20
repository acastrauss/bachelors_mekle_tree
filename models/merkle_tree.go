package models

import (
	"errors"
	"math"
)

type TreeParams struct {
	TreeIndex        uint
	PowerOfTreeIndex uint
}

func getNumberOfLeafNodes(treeParams TreeParams) int {
	return int(math.Pow(float64(treeParams.TreeIndex), float64(treeParams.PowerOfTreeIndex)))
}

type Hash struct {
	Value []byte
}

type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	NodeHash Hash
	NodeId   int
}

var TreeNodeId = 0

func buildMerkleNode(childrenHashes []Hash, stringValueLength int) TreeNode {
	var nodeHash Hash
	isLeaf := len(childrenHashes) == 0
	if isLeaf {
		// leaf node
		nodeHash = Hash{Value: keccakHasher.Hash([]byte(RandStringRunes(stringValueLength)))}
	} else {
		var concatenatedhashes []byte
		for _, ch := range childrenHashes {
			concatenatedhashes = append(concatenatedhashes, ch.Value...)
		}
		nodeHash = Hash{Value: keccakHasher.Hash(concatenatedhashes)}
	}

	TreeNodeId += 1

	return TreeNode{
		Parent:   nil,
		NodeHash: nodeHash,
		NodeId:   TreeNodeId,
	}
}

func assignParentToChildren(parent *TreeNode, children []*TreeNode) error {
	if parent == nil {
		return errors.New("parent is nil")
	}

	for _, c := range children {
		c.Parent = parent
	}
	parent.Children = children

	return nil
}

type MerkleTree struct {
	Root     *TreeNode
	RootHash Hash
}
