package utils

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func FindRoute(request types.Request, graph *types.Graph) ([]types.NodeId, []types.Payment, bool, bool, bool, bool) {
	chunkId := request.ChunkId
	respNodes := request.RespNodes
	mainOriginatorId := request.OriginatorId
	curNextNodeId := request.OriginatorId
	route := []types.NodeId{
		mainOriginatorId,
	}
	found := false
	accessFailed := false
	thresholdFailed := false
	foundByCaching := false
	prevNodePaid := false
	var paymentList []types.Payment
	var payment types.Payment
	var nextNodeId types.NodeId

	if constants.IsPayIfOrigPays() {
		prevNodePaid = true
	}
	if general.ArrContains(respNodes, mainOriginatorId) {
		// originator has the chunk --> chunk is found
		found = true
	} else {
	out:
		for !general.ArrContains(respNodes, curNextNodeId) {
			//fmt.Printf("\n orig: %d, chunk_id: %d", mainOriginatorId, chunkId)

			nextNodeId, thresholdFailed, accessFailed, prevNodePaid, payment = getNext(request, curNextNodeId, prevNodePaid, graph)

			if !payment.IsNil() {
				paymentList = append(paymentList, payment)
			}
			if !nextNodeId.IsNil() {
				route = append(route, nextNodeId)
			}
			if !thresholdFailed && !accessFailed {
				if general.ArrContains(respNodes, nextNodeId) {
					//fmt.Println("is not in cache")
					found = true
					break out
				}
				if constants.IsCacheEnabled() {
					node := graph.GetNode(nextNodeId)
					if node.CacheStruct.Contains(chunkId) {
						foundByCaching = true
						found = true
						break out
					}
				}
				// NOTE !
				curNextNodeId = nextNodeId
			} else {
				break out
			}
		}
	}

	if constants.IsForwardersPayForceOriginatorToPay() {
		if !accessFailed && len(paymentList) > 0 {

			for i := 0; i < len(route)-1; i++ {
				newPayment := types.Payment{
					FirstNodeId: route[i],
					PayNextId:   route[i+1],
					ChunkId:     chunkId}

				paymentAlreadyExists := false
				for _, tmp := range paymentList {
					if newPayment.FirstNodeId == tmp.FirstNodeId && newPayment.PayNextId == tmp.PayNextId {
						paymentAlreadyExists = true
						break
					}
				}
				if !paymentAlreadyExists {
					if i == 0 {
						newPayment.IsOriginator = true
					}
					if i+1 > len(paymentList) {
						paymentList = append(paymentList, newPayment)
					} else {
						paymentList = append(paymentList[:i+1], paymentList[i:]...)
						paymentList[i] = newPayment
					}
				}
			}
		} else {
			paymentList = []types.Payment{}
		}
	}

	return route, paymentList, found, accessFailed, thresholdFailed, foundByCaching
}
