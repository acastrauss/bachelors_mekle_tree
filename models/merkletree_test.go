package models

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func genTreeParamsTestData(args []reflect.Value, r *rand.Rand) {
	treeParams := &TreeParams{
		TreeIndex:        uint(r.Uint32()),
		PowerOfTreeIndex: uint(r.Uint32()),
	}
	args[0] = reflect.ValueOf(treeParams)
}

func TestGetNumberOfLeafNodes(t *testing.T) {
	f1 := func(treeParams *TreeParams) bool {
		actual := getNumberOfLeafNodes(*treeParams)
		excpeted := int(math.Pow(float64(treeParams.TreeIndex), float64(treeParams.PowerOfTreeIndex)))
		return actual == excpeted
	}
	if err := quick.Check(f1, &quick.Config{MaxCount: 1000, Values: genTreeParamsTestData}); err != nil {
		t.Error(err)
	}
}

func genRandomTreeNode(r *rand.Rand) *TreeNode {
	treeNode := &TreeNode{
		Parent:   nil,
		Children: make([]*TreeNode, 0),
		NodeHash: Hash{Value: make([]byte, KECCAK_SHA_LENGTH)},
		NodeId:   int(r.Int()),
	}
	return treeNode
}

func genTreeNodeForAssignTestData(args []reflect.Value, r *rand.Rand) {
	args[0] = reflect.ValueOf(genRandomTreeNode(r))
	children := make([]*TreeNode, 0)
	nofChildren := r.Intn(10)

	for i := 0; i < nofChildren; i++ {
		children = append(children, genRandomTreeNode(r))
	}

	args[1] = reflect.ValueOf(children)
}

func TestAssignParentToChildren(t *testing.T) {
	f1 := func(parent *TreeNode, children []*TreeNode) bool {
		err := assignParentToChildren(parent, children)
		if err != nil {
			return true
		}
		retval := true
		for _, c := range children {
			retval = retval && c.Parent == parent
		}
		retval = retval && (reflect.DeepEqual(parent.Children, children))
		return retval
	}
	if err := quick.Check(f1, &quick.Config{MaxCount: 1000, Values: genTreeNodeForAssignTestData}); err != nil {
		t.Error(err)
	}
}

func genBuildMerkleNodeTestData(args []reflect.Value, r *rand.Rand) {
	nofChildrenHashes := r.Intn(10)
	childrenHashes := make([]Hash, 0)
	for i := 0; i < nofChildrenHashes; i++ {
		childrenHashes = append(childrenHashes, Hash{Value: make([]byte, KECCAK_SHA_LENGTH)})
	}
	args[0] = reflect.ValueOf(childrenHashes)
	args[1] = reflect.ValueOf(STRING_VALUE_LENGTH)
}

func TestBuildMerkleNode(t *testing.T) {
	f1 := func(childrenHashes []Hash, stringValueLength int) bool {
		if len(childrenHashes) == 0 {
			return true
		}

		actualNode := buildMerkleNode(childrenHashes, stringValueLength)
		var concatenatedhashes []byte
		for _, ch := range childrenHashes {
			concatenatedhashes = append(concatenatedhashes, ch.Value...)
		}
		excpetedHash := Hash{Value: keccakHasher.Hash(concatenatedhashes)}

		return reflect.DeepEqual(actualNode.NodeHash.Value, excpetedHash.Value)
	}
	if err := quick.Check(f1, &quick.Config{MaxCount: 1000, Values: genBuildMerkleNodeTestData}); err != nil {
		t.Error(err)
	}
}
