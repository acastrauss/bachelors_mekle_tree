package models

import (
	"math"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
)

func genTreeParamsForTest(args []reflect.Value, r *rand.Rand) {
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
	if err := quick.Check(f1, &quick.Config{MaxCount: 1000, Values: genTreeParamsForTest}); err != nil {
		t.Error(err)
	}
}
