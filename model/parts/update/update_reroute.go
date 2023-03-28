package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Reroute(state *types.State, requestResult types.RequestResult, curEpoch int) int32 {
	var retryCounter int32
	if constants.IsRetryWithAnotherPeer() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId
		originatorId := route[0]
		originator := state.Graph.GetNode(originatorId)
		reroute := originator.RerouteStruct.Reroute // reroute = rejected nodes + chunk

		if requestResult.Found {
			if reroute.RejectedNodes != nil {
				if reroute.ChunkId == chunkId { // If chunkId == chunkId --> reset reroute
					originator.RerouteStruct.ResetRerouteAndSaveToHistory(chunkId, curEpoch)
				}
			}

		} else if len(route) > 1 { // Rejection in second hop --> route have at least an originator and a lastHopNode
			lastHopNode := route[len(route)-1]
			if reroute.RejectedNodes == nil {
				reroute = originator.RerouteStruct.AddNewReroute(requestResult.AccessFailed, lastHopNode, chunkId, curEpoch)
				retryCounter = atomic.AddInt32(&state.UniqueRetryCounter, 1)
			} else {
				if !general.Contains(reroute.RejectedNodes, lastHopNode) { // if the last hop in new route have not been searched before
					originator.RerouteStruct.AddNodeToRejectedNodes(requestResult.AccessFailed, lastHopNode, curEpoch)
				}
			}
		}

		if retryCounter == 0 {
			retryCounter = atomic.LoadInt32(&state.UniqueRetryCounter)
		}

		if len(reroute.RejectedNodes) > constants.GetBinSize() {
			originator.RerouteStruct.ResetRerouteAndSaveToHistory(chunkId, curEpoch)
		}

	}
	return retryCounter
}
