package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Timestep(prevState *types.State) int {
	curTimeStep := int(atomic.AddInt32(&prevState.TimeStep, 1))
	return curTimeStep

}
