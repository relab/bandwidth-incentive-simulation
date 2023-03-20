package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func FailedRequestsThreshold(state *types.State, requestResult types.RequestResult) int32 {
	if requestResult.ThresholdFailed {
		return atomic.AddInt32(&state.FailedRequestsThreshold, 1)
	}
	return atomic.LoadInt32(&state.FailedRequestsThreshold)
}
