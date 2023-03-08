package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func RerouteMap(state *types.State, policyInput types.RequestResult) types.RerouteStruct {
	//rerouteStruct := state.RerouteStruct
	if constants.Constants.IsRetryWithAnotherPeer() {
		route := policyInput.Route
		originator := route[0]
		if !general.Contains(route, -1) && !general.Contains(route, -2) {
			reroute := state.RerouteStruct.GetRerouteMap(originator)
			if reroute != nil {
				if reroute[len(reroute)-1] == route[len(route)-1] {
					//remove rerouteMap[originator]
					state.RerouteStruct.DeleteReroute(originator)
				}
			}
			//if _, ok := rerouteMap[originator]; ok {
			//	val := rerouteMap[originator]
			//	if val[len(val)-1] == route[len(route)-1] {
			//		//remove rerouteMap[originator]
			//		delete(rerouteMap, originator)
			//	}
			//}
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

				//if _, ok := rerouteMap[originator]; ok {
				//	val := rerouteMap[originator]
				//	if !Contains(val, route[1]) {
				//		val = append([]int{route[1]}, val...)
				//		rerouteMap[originator] = val
				//	}
				//} else {
				//	rerouteMap[originator] = []int{route[1], route[len(route)-1]}
				//}
			}
		}
		reroute := state.RerouteStruct.GetRerouteMap(originator)
		if reroute != nil {
			if len(reroute) > constants.Constants.GetBinSize() {
				state.RerouteStruct.DeleteReroute(originator)
			}
		}
		//if _, ok := rerouteMap[originator]; ok {
		//	if len(rerouteMap[originator]) > Constants.GetBinSize() {
		//		delete(rerouteMap, originator)
		//	}
		//}
	}
	//state.RerouteStruct = rerouteStruct
	return state.RerouteStruct
}
