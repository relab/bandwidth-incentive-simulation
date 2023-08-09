package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func Reroute(state *types.State, requestResult types.RequestResult, curEpoch int) {
	if !config.IsRetryWithAnotherPeer() {
		return
	}

	route := requestResult.Route
	chunkId := requestResult.ChunkId
	originatorId := route[0]
	originator := state.Graph.GetNode(originatorId)
	reroute := originator.RerouteStruct.Reroute // reroute = rejected nodes + chunk

	// If the request was successful and the chunk is in the current reroute, reset the reroute
	if requestResult.Found && reroute.RejectedNodes != nil && reroute.ChunkId == chunkId {
		originator.RerouteStruct.ResetRerouteAndSaveToHistory(chunkId, curEpoch)
	} else if len(route) > 1 { // Rejection in second hop -> route has at least an originator and a lastHopNode
		lastHopNode := route[len(route)-1]
		if reroute.RejectedNodes == nil {
			// Create a new reroute if it doesn't exist
			reroute = originator.RerouteStruct.AddNewReroute(requestResult.AccessFailed, lastHopNode, chunkId, curEpoch)
		} else if !general.Contains(reroute.RejectedNodes, lastHopNode) {
			// Add the last hop node to rejected nodes if it hasn't been searched before
			originator.RerouteStruct.AddNodeToRejectedNodes(requestResult.AccessFailed, lastHopNode, curEpoch)
		}
	}
	// If the reroute exceeds the bin size, reset the reroute
	if len(reroute.RejectedNodes) > config.GetBinSize() {
		originator.RerouteStruct.ResetRerouteAndSaveToHistory(chunkId, curEpoch)
	}
}
