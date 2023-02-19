package models

type Hash struct {
	Value []byte
}

type TreeNode struct {
	Parent   *TreeNode
	Children []*TreeNode
	IsLeaf   bool
	NodeHash Hash
}

type MerkleTree struct {
	Root     *TreeNode
	RootHash Hash
}
