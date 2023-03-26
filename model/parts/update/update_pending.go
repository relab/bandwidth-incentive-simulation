package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Pending(state *types.State, requestResult types.RequestResult, curEpoch int) int32 {
	var waitingCounter int32
	if constants.IsWaitingEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId
		originatorId := route[0]
		originator := state.Graph.GetNode(originatorId)
		isNewChunk := false

		if constants.IsRetryWithAnotherPeer() {
			if requestResult.ThresholdFailed || requestResult.AccessFailed {
				isNewChunk = originator.PendingStruct.AddPendingChunkId(state, chunkId, curEpoch)
			} else if requestResult.Found {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		} else {
			if requestResult.ThresholdFailed {
				isNewChunk = originator.PendingStruct.AddPendingChunkId(state, chunkId, curEpoch)
			} else if requestResult.Found || requestResult.AccessFailed {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		}
		if isNewChunk {
			waitingCounter = atomic.AddInt32(&state.UniqueWaitingCounter, 1)
		} else {
			waitingCounter = atomic.LoadInt32(&state.UniqueWaitingCounter)
		}
	}

	return waitingCounter
}
