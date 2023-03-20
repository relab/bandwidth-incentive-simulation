package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func SuccessfulFound(state *types.State, requestResult types.RequestResult) int32 {
	if requestResult.Found {
		return atomic.AddInt32(&state.SuccessfulFound, 1)
	}
	return atomic.LoadInt32(&state.SuccessfulFound)
}
