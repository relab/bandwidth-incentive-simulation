package routing

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"math"
)

type EdgeWrapper struct {
	types.Edge
}

func (edge *EdgeWrapper) CheckForgiveness(graph *types.Graph, request types.Request) (int, bool) {
	passedTime := request.Epoch - edge.Attrs.LastEpoch

	if passedTime <= 0 {
		return edge.Attrs.A2B, false
	}

	refreshRate := config.GetRefreshRate()
	if config.IsAdjustableThreshold() {
		refreshRate = GetAdjustedRefreshrate(edge.Attrs.Threshold, config.GetThreshold(), config.GetRefreshRate(), config.GetAdjustableThresholdExponent())
	}

	removedDeptAmount := passedTime * refreshRate
	newEdgeData := edge.Attrs
	newEdgeData.A2B -= removedDeptAmount
	if newEdgeData.A2B < 0 {
		newEdgeData.A2B = 0
	}
	newEdgeData.LastEpoch = request.Epoch
	graph.SetEdgeData(edge.FromNodeId, edge.ToNodeId, newEdgeData)

	return newEdgeData.A2B, true
}

func GetAdjustedRefreshrate(adjustedThreshold, threshold, refreshRate, power int) int {
	ratio := float64(adjustedThreshold) / float64(threshold)
	return int(math.Ceil(float64(refreshRate) * math.Pow(ratio, float64(power))))
}
