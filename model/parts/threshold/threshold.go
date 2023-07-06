package threshold

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/forgiveness"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
)

func IsThresholdFailed(firstNodeId types.NodeId, secondNodeId types.NodeId, graph *types.Graph, request types.Request) bool {
	if config.GetThresholdEnabled() {
		edgeDataFirst := graph.GetEdgeData(firstNodeId, secondNodeId)
		p2pFirst := edgeDataFirst.A2B
		edgeDataSecond := graph.GetEdgeData(secondNodeId, firstNodeId)
		p2pSecond := edgeDataSecond.A2B

		threshold := config.GetThreshold()
		if config.IsAdjustableThreshold() {
			threshold = edgeDataFirst.Threshold
		}

		peerPriceChunk := utils.PeerPriceChunk(secondNodeId, request.ChunkId)

		price := p2pFirst + peerPriceChunk
		if config.GetReciprocityEnabled() {
			price = p2pFirst - p2pSecond + peerPriceChunk
		}
		//fmt.Printf("price: %d = p2pFirst: %d - p2pSecond: %d + PeerPriceChunk: %d \n", price, p2pFirst, p2pSecond, peerPriceChunk)

		if price > threshold {
			if config.IsForgivenessEnabled() {
				newP2pFirst, forgiven := forgiveness.CheckForgiveness(edgeDataFirst, firstNodeId, secondNodeId, graph, request)
				if forgiven {
					price = newP2pFirst - p2pSecond + peerPriceChunk
				}
			}
		}
		return price > threshold
	}
	return false
}