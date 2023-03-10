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
		//pendingNode := state.PendingStruct.GetPending(originator)
		// -1 Threshold Fail, -2 Access Fail

		if constants.IsRetryWithAnotherPeer() {
			if general.Contains(route, -1) || general.Contains(route, -2) {
				state.PendingStruct.AddPendingChunkIdToQueue(originator, chunkId)
				//if state.PendingStruct.IsEmpty(originator) {
				//	// if the pendingNode is empty then add the chunkId to the pendingNode
				//	state.PendingStruct.AddPending(originator, chunkId)
				//} else {
				//	// if the pendingNode is not empty then add the chunkId to the pendingNodeIds
				//	state.PendingStruct.AddToPendingQueue(originator, chunkId)
				//}
			} else if !state.PendingStruct.IsEmpty(originator) {
				pendingNodeIndex := state.PendingStruct.GetPendingIndex(originator, chunkId)
				if pendingNodeIndex != -1 {
					// if found then delete the chunkId from the pendingNode
					state.PendingStruct.DeletePendingNodeId(originator, pendingNodeIndex)
				}
			}

		} else {
			if general.Contains(route, -1) {
				state.PendingStruct.AddPendingChunkIdToQueue(originator, chunkId)
				//if state.PendingStruct.IsEmpty(originator) {
				//	// if the pendingNode is empty then add the chunkId to the pendingNode
				//	state.PendingStruct.AddPending(originator, chunkId)
				//} else {
				//	// if the pendingNode is not empty then add the chunkId to the pendingNodeIds
				//	state.PendingStruct.AddToPendingQueue(originator, chunkId)
				//}

			} else if !state.PendingStruct.IsEmpty(originator) {
				pendingNodeIndex := state.PendingStruct.GetPendingIndex(originator, chunkId)
				if pendingNodeIndex != -1 {
					// if found then delete the chunkId from the pendingNode
					state.PendingStruct.DeletePendingNodeId(originator, pendingNodeIndex)
				}
			}
		}
	}

	return state.PendingStruct
}

//func PendingMap(state *types.State, policyInput types.RequestResult) types.PendingStruct {
//	//pendingStruct := state.PendingStruct
//	if constants.constants.IsWaitingEnabled() {
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
//		} else if constants.constants.IsRetryWithAnotherPeer() {
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
