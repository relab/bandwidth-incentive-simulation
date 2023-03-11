package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func SuccessfulFound(state *types.State, policyInput types.RequestResult) int32 {
	if policyInput.Found {
		return atomic.AddInt32(&state.SuccessfulFound, 1)
	}
	return atomic.LoadInt32(&state.SuccessfulFound)
}
