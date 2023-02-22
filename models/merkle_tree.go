package models

import (
	"errors"
	"math"

	"github.com/wealdtech/go-merkletree/keccak256"
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

func AreHashesEqual(h1 Hash, h2 Hash) bool {
	if (h1.Value == nil && h2.Value != nil) || (h1.Value != nil && h2.Value == nil) {
		return false
	}
	if h1.Value == nil && h2.Value == nil {
		return true
	}
	if len(h1.Value) != len(h2.Value) {
		return false
	}

	for i := 0; i < len(h1.Value) && i < len(h2.Value); i++ {
		if h1.Value[i] != h2.Value[i] {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	NodeHash Hash
	NodeId   int
}

var TreeNodeId = 0

var keccakHasher = keccak256.New()

const STRING_VALUE_LENGTH = 50
const KECCAK_SHA_LENGTH = 32

func buildMerkleNode(childrenHashes []Hash, stringValueLength int) TreeNode {
	var nodeHash Hash
	isLeaf := len(childrenHashes) == 0
	if isLeaf {
		// leaf node
		nodeHash = Hash{Value: keccakHasher.Hash([]byte(RandStringRunes(stringValueLength)))}
	} else {
		nodeHash = genParentHashFromChildrenHashes(childrenHashes)

	}
	TreeNodeId += 1

	return TreeNode{
		Parent:   nil,
		NodeHash: nodeHash,
		NodeId:   TreeNodeId,
	}
}

func genParentHashFromChildrenHashes(childrenHashes []Hash) Hash {
	concatenatedhashes := make([]byte, 0)
	for _, ch := range childrenHashes {
		concatenatedhashes = append(concatenatedhashes, ch.Value...)
	}

	return Hash{Value: keccakHasher.Hash(concatenatedhashes)}
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
