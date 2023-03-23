package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
)

func Pending(state *types.State, requestResult types.RequestResult, curEpoch int) int32 {
	var pendingCounter int32
	if constants.IsWaitingEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId
		originatorId := route[0]
		originator := state.Graph.GetNode(originatorId)

		if constants.IsRetryWithAnotherPeer() {
			if requestResult.ThresholdFailed || requestResult.AccessFailed {
				pendingCounter = originator.PendingStruct.AddPendingChunkId(state, chunkId, curEpoch)
			} else {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		} else {
			if requestResult.ThresholdFailed {
				pendingCounter = originator.PendingStruct.AddPendingChunkId(state, chunkId, curEpoch)
			} else {
				originator.PendingStruct.DeletePendingChunkId(chunkId)
			}
		}

		//// -1 Threshold Fail, -2 Access Fail
		//if constants.IsRetryWithAnotherPeer() {
		//	if requestResult.ThresholdFailed || requestResult.AccessFailed {
		//		state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
		//	} else {
		//		state.PendingStruct.DeletePendingChunkId(originator, chunkId)
		//	}
		//} else {
		//	if requestResult.ThresholdFailed {
		//		state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
		//	} else {
		//		state.PendingStruct.DeletePendingChunkId(originator, chunkId)
		//	}
		//}
	}
	return pendingCounter
}
