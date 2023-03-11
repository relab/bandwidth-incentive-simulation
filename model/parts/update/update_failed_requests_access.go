package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func FailedRequestsAccess(state *types.State, policyInput types.RequestResult) int32 {
	accessFailed := policyInput.AccessFailed
	if accessFailed {
		return atomic.AddInt32(&state.FailedRequestsAccess, 1)
	}
	return atomic.LoadInt32(&state.FailedRequestsAccess)
}
