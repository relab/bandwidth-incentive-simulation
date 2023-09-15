package routing

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
)

func CheckForgiveness(edgeData types.EdgeAttrs, firstNodeId types.NodeId, secondNodeId types.NodeId, graph *types.Graph, request types.Request) (int, bool) {
	passedTime := request.Epoch - edgeData.LastEpoch

	if passedTime <= 0 {
		return edgeData.A2B, false
	}

	refreshRate := edgeData.Refreshrate

	removedDeptAmount := passedTime * refreshRate
	newEdgeData := edgeData
	newEdgeData.A2B -= removedDeptAmount
	if newEdgeData.A2B < 0 {
		newEdgeData.A2B = 0
	}
	newEdgeData.LastEpoch = request.Epoch
	graph.SetEdgeData(firstNodeId, secondNodeId, newEdgeData)
	if config.IsVariableRefreshrate() {
		graph.SetEdgeDecrementThreshold(firstNodeId, secondNodeId)
	}

	return newEdgeData.A2B, true
}
