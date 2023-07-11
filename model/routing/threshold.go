package routing

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
)

func IsThresholdFailed(firstNodeId types.NodeId, secondNodeId types.NodeId, graph *types.Graph, request types.Request) bool {
	if !config.GetThresholdEnabled() {
		return false
	}
	var edgeWrapper EdgeWrapper
	edgeDataFirst := graph.GetEdge(firstNodeId, secondNodeId)
	p2pFirst := edgeDataFirst.Attrs.A2B
	edgeDataSecond := graph.GetEdge(secondNodeId, firstNodeId)
	p2pSecond := edgeDataSecond.Attrs.A2B
	edgeWrapper.Edge = *edgeDataFirst

	threshold := config.GetThreshold()
	if config.IsAdjustableThreshold() {
		threshold = edgeDataFirst.Attrs.Threshold
	}

	peerPriceChunk := utils.PeerPriceChunk(secondNodeId, request.ChunkId)

	price := p2pFirst + peerPriceChunk
	if config.GetReciprocityEnabled() {
		price = p2pFirst - p2pSecond + peerPriceChunk
	}
	//fmt.Printf("price: %d = p2pFirst: %d - p2pSecond: %d + PeerPriceChunk: %d \n", price, p2pFirst, p2pSecond, peerPriceChunk)

	if price > threshold && config.IsForgivenessEnabled() {
		newP2pFirst, forgiven := edgeWrapper.CheckForgiveness(graph, request)
		if forgiven {
			price = newP2pFirst - p2pSecond + peerPriceChunk
		}
	}

	return price > threshold
}
