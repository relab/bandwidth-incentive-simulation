package update

import "go-incentive-simulation/model/parts/types"

func Epoch(state *types.State) int {
	state.Epoch++
	return state.Epoch
}
