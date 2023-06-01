package utils

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"math"
)

func CheckForgiveness(edgeData types.EdgeAttrs, firstNodeId types.NodeId, secondNodeId types.NodeId, graph *types.Graph, request types.Request) (int, bool) {
	//passedTime := (request.TimeStep - edgeData.Last) / constants.GetRequestsPerSecond() // Old method without Epoch
	passedTime := request.Epoch - edgeData.LastEpoch
	if passedTime > 0 {
		refreshRate := config.GetRefreshRate()
		if config.IsAdjustableThreshold() {
			// TODO: The exponent (6) here should be a configuration parameter.
			refreshRate = GetAdjustedRefreshrate(edgeData.Threshold, config.GetThreshold(), config.GetRefreshRate(), 3)
			//refreshRate = int(math.Ceil(float64(edgeData.Threshold / 2)))
		}
		removedDeptAmount := passedTime * refreshRate
		newEdgeData := edgeData
		newEdgeData.A2B -= removedDeptAmount
		if newEdgeData.A2B < 0 {
			newEdgeData.A2B = 0
		}
		newEdgeData.LastEpoch = request.Epoch
		graph.SetEdgeData(firstNodeId, secondNodeId, newEdgeData)
		return newEdgeData.A2B, true
	}
	return edgeData.A2B, false
}

func GetAdjustedRefreshrate(adjustedThreshold, threshold, refreshRate, power int) int {
	ratio := float64(adjustedThreshold) / float64(threshold)
	return int(math.Ceil(float64(refreshRate) * math.Pow(ratio, float64(power))))
}
