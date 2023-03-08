package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func FailedRequestsThreshold(state *types.State, policyInput types.RequestResult) int32 {
	found := policyInput.Found
	// thresholdFailedList := policyInput.thresholdFailedList
	accessFailed := policyInput.AccessFailed
	if !found && !accessFailed {
		return atomic.AddInt32(&state.FailedRequestsThreshold, 1)
	}
	return atomic.LoadInt32(&state.FailedRequestsThreshold)
}
