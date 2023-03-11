package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Timestep(state *types.State) int {
	curTimeStep := int(atomic.AddInt32(&state.TimeStep, 1))
	return curTimeStep

}
