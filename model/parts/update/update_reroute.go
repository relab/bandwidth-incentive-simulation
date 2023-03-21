package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func RerouteMap(state *types.State, requestResult types.RequestResult, curEpoch int) types.RerouteStruct {
	if constants.IsRetryWithAnotherPeer() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId
		originator := route[0]

		if requestResult.Found {
			routeStruct := state.RerouteStruct.GetRerouteMap(originator) // reroute = rejected nodes + chunk
			if routeStruct.Reroute != nil {
				if routeStruct.ChunkId == chunkId { // If chunkId == chunkId
					state.RerouteStruct.DeleteReroute(originator)
				}
			}

		} else if len(route) > 1 { // Rejection in second hop --> route have at least an originator and a firstHopeNode
			routeStruct := state.RerouteStruct.GetRerouteMap(originator)
			firstHopNode := route[1]
			if routeStruct.Reroute != nil {
				if !general.Contains(routeStruct.Reroute, firstHopNode) { // if the first hop in new route have not been searched before
					state.RerouteStruct.AddNodeToReroute(originator, firstHopNode)
				}
			} else {
				state.RerouteStruct.AddNewReroute(originator, firstHopNode, chunkId, curEpoch)
			}
		}

		routeStruct := state.RerouteStruct.GetRerouteMap(originator)
		if routeStruct.Reroute != nil {
			if len(routeStruct.Reroute) > constants.GetBinSize() {
				state.RerouteStruct.DeleteReroute(originator)
			}
		}
	}
	return state.RerouteStruct
}
