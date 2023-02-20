package models

import (
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func genDataForRemoveNodeLevelAtIndex(args []reflect.Value, r *rand.Rand) {
	nodesAtLevel := make([]TreeNode, 0)
	nofNodes := r.Intn(10) + 1

	for i := 0; i < nofNodes; i++ {
		nodesAtLevel = append(nodesAtLevel, *genRandomTreeNode(r))
	}
	args[0] = reflect.ValueOf(nodesAtLevel)
	args[1] = reflect.ValueOf(r.Intn(len(nodesAtLevel)))
}

func TestRemoveFromNodeLevelAtIndex(t *testing.T) {
	f1 := func(level []TreeNode, indx int) bool {
		nodeAtIndx := level[indx]
		lenBeforeRemoving := len(level)
		level = removeFromNodeLevelAtIndex(level, indx)
		retval := true
		if indx < len(level) {
			retval = retval && !reflect.DeepEqual(nodeAtIndx, level[indx])
		}

		retval = retval && (lenBeforeRemoving-1 == len(level))
		return retval
	}
	if err := quick.Check(f1, &quick.Config{MaxCount: 1000, Values: genDataForRemoveNodeLevelAtIndex}); err != nil {
		t.Error(err)
	}
}
