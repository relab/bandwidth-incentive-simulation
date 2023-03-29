package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func SuccessfulFound(state *types.State, requestResult types.RequestResult) int64 {
	if requestResult.Found {
		return atomic.AddInt64(&state.SuccessfulFound, 1)
	}
	return atomic.LoadInt64(&state.SuccessfulFound)
}
