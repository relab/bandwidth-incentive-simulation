package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func PendingMap(state *types.State, requestResult types.RequestResult, curEpoch int) types.PendingStruct {
	if config.IsWaitingEnabled() {
		route := requestResult.Route
		originator := route[0]
		chunkId := route[len(route)-1]

		// -1 Threshold Fail, -2 Access Fail
		if config.IsRetryWithAnotherPeer() {
			if general.Contains(route, -1) || general.Contains(route, -2) {
				state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
			} else {
				state.PendingStruct.DeletePendingChunkId(originator, chunkId)
			}
		} else {
			if general.Contains(route, -1) {
				state.PendingStruct.AddPendingChunkId(originator, chunkId, curEpoch)
			} else {
				state.PendingStruct.DeletePendingChunkId(originator, chunkId)
			}
		}
	}
	return state.PendingStruct
}
