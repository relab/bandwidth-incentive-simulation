package utils

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func FindDistance(first types.NodeId, second types.ChunkId) int {
	return config.GetBits() - general.BitLength(first.ToInt()^second.ToInt())
}

func FindRoute(request types.Request, graph *types.Graph) ([]types.NodeId, []types.Payment, bool, bool, bool, bool) {
	chunkId := request.ChunkId
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

	if config.IsPayIfOrigPays() {
		prevNodePaid = true
	}

	depth := config.GetStorageDepth()

	if FindDistance(mainOriginatorId, chunkId) >= depth {
		found = true
	} else {
	out:
		for !(FindDistance(curNextNodeId, chunkId) >= depth) {
			nextNodeId, thresholdFailed, accessFailed, prevNodePaid, payment = getNext(request, curNextNodeId, prevNodePaid, graph)

			if !payment.IsNil() {
				paymentList = append(paymentList, payment)
			}
			if !nextNodeId.IsNil() {
				route = append(route, nextNodeId)
			}
			if !thresholdFailed && !accessFailed {
				if FindDistance(nextNodeId, chunkId) >= depth {
					found = true
					break out
				}
				if config.IsCacheEnabled() {
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

	if config.IsForwardersPayForceOriginatorToPay() {
		if !accessFailed && len(paymentList) > 0 {

			newList := make([]types.Payment, 0, len(paymentList))

			for i := 0; i < len(route)-1; i++ {
				newPayment := types.Payment{
					FirstNodeId: route[i],
					PayNextId:   route[i+1],
					ChunkId:     chunkId}
				if i == 0 {
					newPayment.IsOriginator = true
				}
				newList = append(newList, newPayment)

				oldIndex := -1
				for oi, tmp := range paymentList {
					if newPayment.FirstNodeId == tmp.FirstNodeId && newPayment.PayNextId == tmp.PayNextId {
						oldIndex = oi
						break
					}
				}

				if oldIndex > -1 {
					paymentList = append(paymentList[:oldIndex], paymentList[oldIndex+1:]...)
				}
				if len(paymentList) == 0 {
					break
				}

			}

			paymentList = newList

		} else {
			paymentList = []types.Payment{}
		}
	}

	return route, paymentList, found, accessFailed, thresholdFailed, foundByCaching
}
