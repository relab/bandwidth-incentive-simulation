package update

import (
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func FailedRequestsAccess(state *types.State, requestResult types.RequestResult) int64 {
	accessFailed := requestResult.AccessFailed
	if accessFailed {
		return atomic.AddInt64(&state.FailedRequestsAccess, 1)
	}
	return atomic.LoadInt64(&state.FailedRequestsAccess)
}
