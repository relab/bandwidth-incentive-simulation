package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func PendingMap(state *types.State, policyInput types.RequestResult) types.PendingStruct {
	if constants.IsWaitingEnabled() {
		route := policyInput.Route
		originator := route[0]
		chunkId := route[len(route)-1]

		// -1 Threshold Fail, -2 Access Fail
		if constants.IsRetryWithAnotherPeer() {
			if general.Contains(route, -1) || general.Contains(route, -2) {
				state.PendingStruct.AddPendingChunkIdToQueue(originator, chunkId)
			} else {
				state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
			}
		} else {
			if general.Contains(route, -1) {
				state.PendingStruct.AddPendingChunkIdToQueue(originator, chunkId)
			} else {
				state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
			}
		}
	}
	return state.PendingStruct
}
