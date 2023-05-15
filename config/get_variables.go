package config

import "fmt"

func GetNumRoutingGoroutines() int {
	num := Variables.confOptions.NumGoroutines
	//if IsWriteStatesToFile() {
	//	num--
	//}
	//if IsWriteRoutesToFile() {
	//	num--
	//}
	num-- // for the requestWorker
	if IsOutputEnabled() {
		num-- // for the outputWorker
	}
	return num
}

func GetNumGoroutines() int {
	return Variables.confOptions.NumGoroutines
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return Variables.experimentOptions.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return Variables.experimentOptions.ForgivenessEnabled
}

func IsCacheEnabled() bool {
	return Variables.experimentOptions.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return Variables.experimentOptions.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return Variables.experimentOptions.RetryWithAnotherPeer
}

func IsForwardersPayForceOriginatorToPay() bool {
	return Variables.experimentOptions.ForwardersPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return Variables.experimentOptions.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return Variables.experimentOptions.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return Variables.experimentOptions.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return Variables.experimentOptions.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return Variables.experimentOptions.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return Variables.experimentOptions.ThresholdEnabled
}

func GetPaymentEnabled() bool {
	return Variables.experimentOptions.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return Variables.confOptions.RequestsPerSecond
}

func GetIterations() int {
	return Variables.confOptions.Iterations
}

func GetBits() int {
	return Variables.confOptions.Bits
}

func GetNetworkSize() int {
	return Variables.confOptions.NetworkSize
}

func GetBinSize() int {
	return Variables.confOptions.BinSize
}

func GetRangeAddress() int {
	return Variables.confOptions.RangeAddress
}

func GetOriginators() int {
	return Variables.confOptions.Originators
}

func GetRefreshRate() int {
	return Variables.confOptions.RefreshRate
}

func GetThreshold() int {
	return Variables.confOptions.Threshold
}

func GetRandomSeed() int64 {
	return Variables.confOptions.RandomSeed
}

func GetMaxProximityOrder() int {
	return Variables.confOptions.MaxProximityOrder
}

func GetPrice() int {
	return Variables.confOptions.Price
}

func GetSameOriginator() bool {
	return Variables.confOptions.SameOriginator
}

func IsEdgeLock() bool {
	return Variables.confOptions.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return Variables.confOptions.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return Variables.confOptions.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return Variables.confOptions.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return Variables.confOptions.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return Variables.confOptions.DebugPrints
}

func GetDebugInterval() int {
	return Variables.confOptions.DebugInterval
}

func TimeForDebugPrints(timeStep int) bool {
	if IsDebugPrints() {
		return timeStep%GetDebugInterval() == 0
	}
	return false
}

func TimeForNewEpoch(timeStep int) bool {
	return timeStep%GetRequestsPerSecond() == 0
}

func IsOutputEnabled() bool {
	return Variables.confOptions.OutputEnabled
}

func GetMeanRewardPerForward() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.MeanRewardPerForward
	}
	return false
}

func GetAverageNumberOfHops() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.AverageNumberOfHops
	}
	return false
}

func GetAverageFractionOfTotalRewardsK8() bool {
	return Variables.confOptions.OutputOptions.AverageFractionOfTotalRewardsK8
}

func GetAverageFractionOfTotalRewardsK16() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		Variables.confOptions.BinSize == 16 &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.AverageFractionOfTotalRewardsK16
	}
	return false

}

func GetRewardFairnessForForwardingAction() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.RewardFairnessForForwardingAction
	}
	return false
}

func GetRewardFairnessForStoringAction() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.RewardFairnessForStoringAction
	}
	return false
}

func GetRewardFairnessForAllActions() bool {
	if Variables.experimentOptions.MaxPOCheckEnabled &&
		!Variables.experimentOptions.ThresholdEnabled &&
		!Variables.experimentOptions.ForgivenessEnabled &&
		!Variables.experimentOptions.PaymentEnabled &&
		!Variables.experimentOptions.WaitingEnabled &&
		!Variables.experimentOptions.RetryWithAnotherPeer {
		return Variables.confOptions.OutputOptions.RewardFairnessForAllActions
	}
	return false
}

// GetNegativeIncome TODO: Kan ver merr det må ver mer checks her
func GetNegativeIncome() bool {
	if Variables.experimentOptions.PaymentEnabled &&
		Variables.experimentOptions.ForgivenessEnabled {
		return Variables.confOptions.OutputOptions.NegativeIncome
	}
	return false
}

func GetComputeWorkFairness() bool {
	return Variables.confOptions.OutputOptions.ComputeWorkFairness
}

func GetExperimentString() (exp string) {
	exp = fmt.Sprintf("O%dT%dsW%dTh%dFg%dW%d",
		GetOriginators()*100/GetNetworkSize(),
		GetIterations()/GetRequestsPerSecond(),
		GetBinSize(),
		GetThreshold(),
		GetRefreshRate(),
		GetMaxProximityOrder())
	if GetPaymentEnabled() {
		exp += "Pay"
	}
	if IsCacheEnabled() {
		exp += "Cache"
	}
	return exp
}
