package models

import (
	"errors"
	"math"
	"math/rand"

	"github.com/wealdtech/go-merkletree/keccak256"
)

const TREE_INDEX = 2
const STRING_LENGTH_FOR_HASH = 50

var keccakHasher = keccak256.New()

func GenerateBinaryTree() {
	maxPowerOfTwo := 3
	nofLeafNodes := int(math.Pow(2, float64(maxPowerOfTwo)))
	currentLevelNofNodes := nofLeafNodes

	var prevLevelNodes []TreeNode
	var emptyHashes []Hash

	for currentLevelNofNodes > 0 {
		// generate nodes for current level
		if len(prevLevelNodes) == 0 {
			// leafs
			for i := 0; i < int(currentLevelNofNodes); i++ {
				prevLevelNodes = append(prevLevelNodes, BuildMerkleNode(emptyHashes))
			}
		} else {
			var tempCurrentNodes []TreeNode

			for i := 0; i < int(currentLevelNofNodes); i++ {
				var childrenOfCurrentNode []*TreeNode
				var childrenHashes []Hash
				for j := 0; j < TREE_INDEX; j++ {
					oneChild, indx, _ := GetRandomNodeFromLevel(prevLevelNodes)
					prevLevelNodes, _ = RemoveFromNodeLevelAtIndex(prevLevelNodes, indx)
					childrenOfCurrentNode = append(childrenOfCurrentNode, oneChild)
					childrenHashes = append(childrenHashes, oneChild.NodeHash)
				}
				currentNode := BuildMerkleNode(childrenHashes)
				for _, c := range childrenOfCurrentNode {
					c.Parent = &currentNode
				}
				currentNode.Children = childrenOfCurrentNode
				tempCurrentNodes = append(tempCurrentNodes, currentNode)
			}

			prevLevelNodes = tempCurrentNodes
		}

		currentLevelNofNodes = int(currentLevelNofNodes / 2)
	}

}

func BuildMerkleNode(childrenHashes []Hash) TreeNode {
	var nodeHash Hash
	isLeaf := len(childrenHashes) == 0
	if isLeaf {
		// leaf node
		nodeHash = Hash{Value: keccakHasher.Hash([]byte(RandStringRunes(STRING_LENGTH_FOR_HASH)))}
	} else {
		var concatenatedhashes []byte
		for _, ch := range childrenHashes {
			concatenatedhashes = append(concatenatedhashes, ch.Value...)
		}
		nodeHash = Hash{Value: keccakHasher.Hash(concatenatedhashes)}
	}

	return TreeNode{
		Parent:   nil,
		IsLeaf:   isLeaf,
		NodeHash: nodeHash,
	}
}

func GetRandomNodeFromLevel(nodesAtLevel []TreeNode) (*TreeNode, int, error) {
	indx := rand.Intn(len(nodesAtLevel))
	if indx < 0 || indx > len(nodesAtLevel)-1 {
		return nil, indx, errors.New("invalid index")
	}
	return &nodesAtLevel[indx], indx, nil
}

func RemoveFromNodeLevelAtIndex(level []TreeNode, indx int) ([]TreeNode, error) {
	if indx < 0 || indx > len(level)-1 {
		return nil, errors.New("invalid index")
	}
	return append(level[:indx], level[indx+1:]...), nil
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
