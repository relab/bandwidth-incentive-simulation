package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func RerouteMap(state *types.State, policyInput types.RequestResult) types.RerouteStruct {
	if constants.IsRetryWithAnotherPeer() {
		route := policyInput.Route
		originator := route[0]
		chunkId := route[len(route)-1]

		// -1 Threshold Fail, -2 Access Fail
		if !general.Contains(route, -1) && !general.Contains(route, -2) {
			reroute := state.RerouteStruct.GetRerouteMap(originator)
			if reroute != nil {
				if reroute[len(reroute)-1] == route[len(route)-1] {
					//remove rerouteMap[originator]
					state.RerouteStruct.DeleteReroute(originator)
					// If found remove from waiting queue (pending map)
					if constants.IsWaitingEnabled() {
						state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
					}
				}
			}
		} else {
			if len(route) > 3 {
				reroute := state.RerouteStruct.GetRerouteMap(originator)
				state.RerouteStruct.RerouteMutex.Lock()
				if reroute != nil {
					if !general.Contains(reroute, route[1]) {
						reroute = append([]int{route[1]}, reroute...)
						state.RerouteStruct.RerouteMap[originator] = reroute
					}
				} else {
					state.RerouteStruct.RerouteMap[originator] = []int{route[1], route[len(route)-1]}
				}
				state.RerouteStruct.RerouteMutex.Unlock()
			}
		}
		reroute := state.RerouteStruct.GetRerouteMap(originator)
		if reroute != nil {
			if len(reroute) > constants.GetBinSize() {
				state.RerouteStruct.DeleteReroute(originator)
				if constants.IsWaitingEnabled() {
					state.PendingStruct.DeleteChunkIdFromPendingQueue(originator, chunkId)
				}
			}
		}
	}
	return state.RerouteStruct
}
