package validators

import (
	"errors"
	"models"
	"reflect"
)

func IsDataValidWithinTree(data models.NodeData, tree models.MerkleTree) bool {
	leafThatHasData := GetLeafThatHasData(data, &tree)

	if leafThatHasData == nil {
		return false
	}

	return checkParentHashWithSiblings(leafThatHasData)
}

func checkParentHashWithSiblings(node *models.TreeNode) bool {
	retval := true

	if node.Parent == nil { // root
		return true
	} else {
		if siblings, err := getSiblings(node); err == nil {
			excpectedParentHash := models.GetParentHashFromChildren(siblings)
			retval = retval && reflect.DeepEqual(excpectedParentHash.Value, node.Parent.NodeHash.Value)
			retval = retval && checkParentHashWithSiblings(node.Parent)
		} else {
			panic("Shouldn't be here")
		}
	}

	return retval
}

func getSiblings(node *models.TreeNode) ([]*models.TreeNode, error) {
	if node.Parent != nil {
		return node.Parent.Children, nil
	} else {
		return nil, errors.New("node doesn't have parent")
	}
}
