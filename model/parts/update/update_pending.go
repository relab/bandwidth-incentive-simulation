package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
)

func Pending(state *types.State, requestResult types.RequestResult, curEpoch int) {
	if !config.IsWaitingEnabled() {
		return
	}

	route := requestResult.Route
	chunkId := requestResult.ChunkId
	originatorId := route[0]
	originator := state.Graph.GetNode(originatorId)

	if config.IsRetryWithAnotherPeer() {
		if requestResult.ThresholdFailed || requestResult.AccessFailed {
			originator.PendingStruct.AddPendingChunkId(chunkId, curEpoch)
		} else if requestResult.Found {
			if len(originator.PendingStruct.PendingQueue) > 0 {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		}
	} else {
		if requestResult.ThresholdFailed {
			originator.PendingStruct.AddPendingChunkId(chunkId, curEpoch)
		} else if requestResult.Found || requestResult.AccessFailed {
			if len(originator.PendingStruct.PendingQueue) > 0 {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		}
	}
}
