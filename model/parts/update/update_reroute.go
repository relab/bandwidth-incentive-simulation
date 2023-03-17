package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func RerouteMap(state *types.State, requestResult types.RequestResult) types.RerouteStruct {
	if constants.IsRetryWithAnotherPeer() {
		route := requestResult.Route
		originator := route[0]
		firstHopNode := route[1]
		chunkId := route[len(route)-1]

		// -1 = Threshold Fail, -2 = Access Fail
		if !general.Contains(route, -1) && !general.Contains(route, -2) {
			reroute := state.RerouteStruct.GetRerouteMap(originator) // reroute = rejected nodes + chunk
			if reroute != nil {
				if reroute[len(reroute)-1] == chunkId { // If chunkId == chunkId
					state.RerouteStruct.DeleteReroute(originator)
					// If found remove from waiting queue (pending map)
					//if constants.IsWaitingEnabled() {
					//	state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
					//}
				}
			}

		} else if len(route) > 3 { // Rejection in second hop (?), route contains at least originator, -1/-2, chunkId
			reroute := state.RerouteStruct.GetRerouteMap(originator)
			state.RerouteStruct.RerouteMutex.Lock()
			if reroute != nil {
				if !general.Contains(reroute, firstHopNode) { // if the first hop in new route have not been searched before
					reroute = append([]int{firstHopNode}, reroute[:]...)
					state.RerouteStruct.RerouteMap[originator] = reroute
				}
			} else {
				state.RerouteStruct.RerouteMap[originator] = []int{firstHopNode, chunkId}
			}
			state.RerouteStruct.RerouteMutex.Unlock()
		}

		reroute := state.RerouteStruct.GetRerouteMap(originator)
		if reroute != nil {
			if len(reroute) > constants.GetBinSize() {
				state.RerouteStruct.DeleteReroute(originator)
				//if constants.IsWaitingEnabled() {
				//	state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
				//}
			}
		}
	}
	return state.RerouteStruct
}
