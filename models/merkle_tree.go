package models

import (
	"math/rand"

	"github.com/wealdtech/go-merkletree/keccak256"
)

type Hash struct {
	Value []byte
}

type NodeData struct {
	Data string
}

func (nData *NodeData) GetBytes() []byte {
	return []byte(nData.Data)
}

type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	NodeHash Hash
	NodeId   int
	Data     NodeData
}

var TreeNodeId = 0

var KeccakHasher = keccak256.New()

const STRING_VALUE_LENGTH = 50
const KECCAK_SHA_LENGTH = 32

func BuildLeafNode() TreeNode {
	TreeNodeId += 1
	leafData := NodeData{
		Data: RandStringRunes(STRING_VALUE_LENGTH),
	}

	return TreeNode{
		Parent:   nil,
		Children: make([]*TreeNode, 0),
		NodeHash: Hash{
			Value: KeccakHasher.Hash([]byte(leafData.GetBytes())),
		},
		NodeId: TreeNodeId,
		Data:   leafData,
	}
}

func BuildBranchNode(children []*TreeNode) TreeNode {
	TreeNodeId += 1
	node := TreeNode{
		Parent: nil,
		NodeId: TreeNodeId,
		Data: NodeData{
			Data: "",
		},
	}

	node.AssignChildrenToParent(children)
	return node
}

func (tn *TreeNode) AssignChildrenToParent(children []*TreeNode) {
	for _, c := range children {
		c.Parent = tn
	}
	tn.Children = children
	tn.NodeHash = GetParentHashFromChildren(children)
}

func GetParentHashFromChildren(children []*TreeNode) Hash {
	concatenatedhashes := make([]byte, 0)
	for _, ch := range children {
		concatenatedhashes = append(concatenatedhashes, ch.NodeHash.Value...)
	}
	return Hash{
		Value: KeccakHasher.Hash(concatenatedhashes),
	}
}

func GenRandomTreeNode(r *rand.Rand) *TreeNode {
	treeNode := &TreeNode{
		Parent:   nil,
		Children: make([]*TreeNode, 0),
		NodeHash: Hash{Value: make([]byte, KECCAK_SHA_LENGTH)},
		NodeId:   int(r.Int()),
	}
	return treeNode
}

type MerkleTree struct {
	Root     *TreeNode
	RootHash Hash
}
