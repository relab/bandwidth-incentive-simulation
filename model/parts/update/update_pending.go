package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
)

func PendingMap(state *types.State, requestResult types.RequestResult, curEpoch int) types.PendingStruct {
	if constants.IsWaitingEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId
		originator := route[0]

		// -1 Threshold Fail, -2 Access Fail
		if constants.IsRetryWithAnotherPeer() {
			if requestResult.ThresholdFailed || requestResult.AccessFailed {
				state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
			} else {
				state.PendingStruct.DeletePendingChunkId(originator, chunkId)
			}
		} else {
			if requestResult.ThresholdFailed {
				state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
			} else {
				state.PendingStruct.DeletePendingChunkId(originator, chunkId)
			}
		}
	}
	return state.PendingStruct
}
