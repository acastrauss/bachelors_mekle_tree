package models

import "math"

type TreeParams struct {
	TreeIndex        uint
	PowerOfTreeIndex uint
}

func (treeParams TreeParams) GetNumberOfLeafNodes() int {
	return int(math.Pow(float64(treeParams.TreeIndex), float64(treeParams.PowerOfTreeIndex)))
}

func (treeParams TreeParams) GetTotalNumberOfNodes() int {
	return int(
		(math.Pow(float64(treeParams.TreeIndex), float64(treeParams.PowerOfTreeIndex+1)) - 1) / (float64(treeParams.TreeIndex) - 1))
}
