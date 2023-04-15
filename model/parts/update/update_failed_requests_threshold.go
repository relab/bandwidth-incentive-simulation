package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func FailedRequestsThreshold(state *types.State, requestResult types.RequestResult) int64 {
	if requestResult.ThresholdFailed {
		return atomic.AddInt64(&state.FailedRequestsThreshold, 1)
	}
	return atomic.LoadInt64(&state.FailedRequestsThreshold)
}
