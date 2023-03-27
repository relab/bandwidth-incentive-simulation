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
		rerouteStruct := originator.RerouteStruct // reroute = rejected nodes + chunk

		if requestResult.Found {
			if rerouteStruct.Reroute.CheckedNodes != nil {
				if rerouteStruct.Reroute.ChunkId == chunkId { // If chunkId == chunkId --> reset reroute
					originator.RerouteStruct.Reroute = types.Reroute{}
				}
			}

		} else if len(route) > 1 { // Rejection in second hop --> route have at least an originator and a firstHopeNode
			lastNode := route[len(route)-1]
			if rerouteStruct.Reroute.CheckedNodes == nil {
				//firstHopNode := route[1]
				rerouteStruct = originator.RerouteStruct.AddNewReroute(originator, lastNode, chunkId, curEpoch)
				retryCounter = atomic.AddInt32(&state.UniqueRetryCounter, 1)
			} else {
				if !general.Contains(rerouteStruct.Reroute.CheckedNodes, lastNode) { // if the first hop in new route have not been searched before
					originator.RerouteStruct.AddNodeToCheckedNodes(originator, lastNode)
				}
			}
			//for _, node := range route[1:] { // skipping the originatorId
			//	if !general.Contains(rerouteStruct.Reroute.CheckedNodes, node) { // if the first hop in new route have not been searched before
			//		originator.RerouteStruct.AddNodeToCheckedNodes(originator, node)
			//	}
			//}

		} else {
			retryCounter = atomic.LoadInt32(&state.UniqueRetryCounter)
		}

		if len(rerouteStruct.Reroute.CheckedNodes) > constants.GetBinSize() {
			originator.RerouteStruct.Reroute = types.Reroute{}
		}

	}
	return retryCounter
}
