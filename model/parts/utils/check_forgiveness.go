package utils

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"math"
)

func CheckForgiveness(edgeData types.EdgeAttrs, firstNodeId int, secondNodeId int, graph *types.Graph, request types.Request) {
	passedTime := request.TimeStep - edgeData.Last/constants.GetRequestsPerSecond()
	if passedTime > 0 {
		refreshRate := constants.GetRefreshRate()
		if constants.IsAdjustableThreshold() {
			refreshRate = int(math.Ceil(float64(edgeData.Threshold / 2)))
		}
		removedDeptAmount := passedTime * refreshRate
		newEdgeData := edgeData
		newEdgeData.A2B -= removedDeptAmount
		if newEdgeData.A2B < 0 {
			newEdgeData.A2B = 0
		}
		newEdgeData.Last = request.TimeStep
		graph.SetEdgeData(firstNodeId, secondNodeId, newEdgeData)
	}

}
