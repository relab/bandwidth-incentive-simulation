package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func PendingMap(state *types.State, policyInput types.RequestResult) types.PendingStruct {
	//pendingStruct := state.PendingStruct
	if constants.Constants.IsWaitingEnabled() {
		route := policyInput.Route
		originator := route[0]
		chunkId := route[len(route)-1]
		pendingNode := state.PendingStruct.GetPending(originator)

		if !general.Contains(route, -1) && !general.Contains(route, -2) {
			pendingNodesId := pendingNode.NodeIds
			if general.Contains(pendingNodesId, chunkId) {
				// delete the entire thing
				//state.PendingStruct.DeletePending(originator)
				state.PendingStruct.DeletePendingNodeId(originator, chunkId)
			}
		} else if constants.Constants.IsRetryWithAnotherPeer() {
			if general.Contains(route, -1) && general.Contains(route, -2) {
				pendingNodesId := pendingNode.NodeIds
				if !general.Contains(pendingNodesId, -1) {
					if general.Contains(pendingNodesId, chunkId) {
						if pendingNode.PendingCounter < 10 {
							// increment
							state.PendingStruct.IncrementPending(originator)
						} else {
							state.PendingStruct.AddPending(originator, chunkId)
						}

					} else {
						state.PendingStruct.AddP(originator, chunkId)
					}
				}
			}
			//} else if general.Contains(route, -1) {
			//	pendingNodesId := pendingNode.NodeIds
			//	if !general.Contains(pendingNodesId, -1) {
			//		if pendingNode.PendingCounter < 10 {
			//			// increment
			//			state.PendingStruct.IncrementPending(originator)
			//		} else {
			//			pendingNodeIdIndex := state.PendingStruct.CheckPending(originator, chunkId)
			//			if pendingNodeIdIndex != -1 {
			//				state.PendingStruct.DeletePendingNodeId(originator, pendingNodeIdIndex)
			//			}
			//		}
			//	} else {
			//		state.PendingStruct.AddPending(originator, chunkId)
			//	}
		}
	}
	//} else if general.Contains(route, -1) {
	//	pendingNodesId := pendingNode.NodeIds
	//	if pendingNodesId[0] != -1 {
	//		exsist := state.PendingStruct.CheckPending(originator, chunkId)
	//		if exsist != -1 {
	//			if pendingNode.PendingCounter < 10 {
	//				state.PendingStruct.IncrementPending(originator)
	//			} else {
	//				pendingNodeIdIndex := state.PendingStruct.CheckPending(originator, chunkId)
	//				state.PendingStruct.DeletePendingNodeId(originator, pendingNodeIdIndex)
	//			}
	//		}
	//	} else {
	//		state.PendingStruct.AddPending(originator, chunkId)
	//	}
	//}
	//}
	return state.PendingStruct
}

//func PendingMap(state *types.State, policyInput types.RequestResult) types.PendingStruct {
//	//pendingStruct := state.PendingStruct
//	if constants.Constants.IsWaitingEnabled() {
//		route := policyInput.Route
//		originator := route[0]
//		chunkId := route[len(route)-1]
//		pendingNode := state.PendingStruct.GetPending(originator)
//
//		if !general.Contains(route, -1) && !general.Contains(route, -2) {
//			pendingNodeId := pendingNode.NodeId
//			if pendingNodeId != -1 {
//				if pendingNodeId == chunkId {
//					// remove the pending request
//					state.PendingStruct.DeletePending(originator)
//				}
//			}
//		} else if constants.Constants.IsRetryWithAnotherPeer() {
//			if general.Contains(route, -1) && general.Contains(route, -2) {
//				pendingNodeId := pendingNode.NodeId
//				if pendingNodeId != -1 {
//					if pendingNode.PendingCounter < 10 {
//						state.PendingStruct.IncrementPending(originator)
//					} else {
//						// remove the pending request
//						state.PendingStruct.DeletePending(originator)
//					}
//				} else {
//					// add the pending request
//					state.PendingStruct.AddPending(originator, chunkId)
//				}
//			}
//		} else if general.Contains(route, -1) {
//			pendingNodeId := pendingNode.NodeId
//			if pendingNodeId != -1 {
//				if pendingNode.PendingCounter < 10 {
//					state.PendingStruct.IncrementPending(originator)
//				} else {
//					// remove the pending request
//					state.PendingStruct.DeletePending(originator)
//				}
//			} else {
//				// add the pending request
//				state.PendingStruct.AddPending(originator, chunkId)
//			}
//		}
//	}
//	return state.PendingStruct
//}
