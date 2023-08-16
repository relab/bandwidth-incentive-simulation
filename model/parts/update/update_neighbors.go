package update

import (
	"go-incentive-simulation/model/parts/types"
	"math/rand"
)

func Neighbors(globalState *types.State) bool {
	// Update neighbors with probability p
	for _, node := range(globalState.Graph.NodesMap) {
		if node.OriginatorStruct.RequestCount == 0 {
			continue
		}
		// TODO: make this a config
		if rand.Float32() < 1 {
			node.UpdateNeighbors()
		}
	}

	return true
}
