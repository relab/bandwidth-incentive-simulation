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
				isNewChunk = originator.PendingStruct.AddPendingChunkId(chunkId, curEpoch)
			} else if requestResult.Found {
				if len(originator.PendingStruct.PendingQueue) > 0 {
					originator.PendingStruct.DeletePendingChunkId(chunkId)
				}
			}

		} else {
			if requestResult.ThresholdFailed {
				isNewChunk = originator.PendingStruct.AddPendingChunkId(chunkId, curEpoch)
			} else if requestResult.Found || requestResult.AccessFailed {
				if len(originator.PendingStruct.PendingQueue) > 0 {
					originator.PendingStruct.DeletePendingChunkId(chunkId)
				}
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
