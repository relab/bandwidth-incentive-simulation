package config

import "fmt"

func GetNumRoutingGoroutines() int {
	num := variables.confOptions.NumGoroutines
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
	if num < 1 {
		if IsOutputEnabled() {
			panic("You need at least 3 goroutines for the requestWorker, routingWorker and outputWorker")
		}
		panic("You need at least 2 goroutines for the requestWorker and routingWorker")
	}
	return num
}

func GetNumGoroutines() int {
	return variables.confOptions.NumGoroutines
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return variables.experimentOptions.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return variables.experimentOptions.ForgivenessEnabled
}

func IsCacheEnabled() bool {
	return variables.experimentOptions.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return variables.experimentOptions.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return variables.experimentOptions.RetryWithAnotherPeer
}

func IsForwardersPayForceOriginatorToPay() bool {
	return variables.experimentOptions.ForwardersPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return variables.experimentOptions.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return variables.experimentOptions.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return variables.experimentOptions.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return variables.experimentOptions.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return variables.experimentOptions.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return variables.experimentOptions.ThresholdEnabled
}

func GetReciprocityEnabled() bool {
	return variables.experimentOptions.ReciprocityEnabled
}

func GetPaymentEnabled() bool {
	return variables.experimentOptions.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return variables.confOptions.RequestsPerSecond
}

func GetIterations() int {
	return variables.confOptions.Iterations
}

func GetBits() int {
	return variables.confOptions.Bits
}

func GetNetworkSize() int {
	return variables.confOptions.NetworkSize
}

func GetBinSize() int {
	return variables.confOptions.BinSize
}

func GetRangeAddress() int {
	return variables.confOptions.RangeAddress
}

func GetOriginators() int {
	return variables.confOptions.Originators
}

func GetRefreshRate() int {
	return variables.confOptions.RefreshRate
}

func GetThreshold() int {
	return variables.confOptions.Threshold
}

func GetRandomSeed() int64 {
	return variables.confOptions.RandomSeed
}

func GetMaxProximityOrder() int {
	return variables.confOptions.MaxProximityOrder
}

func GetPrice() int {
	return variables.confOptions.Price
}

func GetSameOriginator() bool {
	return variables.confOptions.SameOriginator
}

func IsEdgeLock() bool {
	return variables.confOptions.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return variables.confOptions.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return variables.confOptions.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return variables.confOptions.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return variables.confOptions.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return variables.confOptions.DebugPrints
}

func GetDebugInterval() int {
	return variables.confOptions.DebugInterval
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
	return variables.confOptions.OutputEnabled
}

func JustPrintOutPut() bool {
	if !variables.confOptions.OutputOptions.MeanRewardPerForward &&
		!variables.confOptions.OutputOptions.AverageNumberOfHops &&
		!variables.confOptions.OutputOptions.AverageFractionOfTotalRewardsK16 &&
		!variables.confOptions.OutputOptions.RewardFairnessForForwardingAction &&
		!variables.confOptions.OutputOptions.RewardFairnessForStoringAction &&
		!variables.confOptions.OutputOptions.RewardFairnessForAllActions &&
		!variables.confOptions.OutputOptions.NegativeIncome {
		return true
	}
	return false
}

func GetMeanRewardPerForward() bool {
	if variables.confOptions.OutputEnabled && variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.MeanRewardPerForward
	}
	return false
}

func GetAverageNumberOfHops() bool {
	if variables.confOptions.OutputEnabled && variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.AverageNumberOfHops
	}
	return false
}

func GetAverageFractionOfTotalRewardsK8() bool {
	return false
}

func GetAverageFractionOfTotalRewardsK16() bool {
	if variables.confOptions.OutputEnabled &&
		variables.confOptions.BinSize == 16 &&
		variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.AverageFractionOfTotalRewardsK16
	}
	return false

}

func GetRewardFairnessForForwardingAction() bool {
	if variables.confOptions.OutputEnabled && variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.RewardFairnessForForwardingAction
	}
	return false
}

func GetRewardFairnessForStoringAction() bool {
	if variables.confOptions.OutputEnabled && variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.RewardFairnessForStoringAction
	}
	return false
}

func GetRewardFairnessForAllActions() bool {
	if variables.confOptions.OutputEnabled && variables.experimentOptions.MaxPOCheckEnabled {
		return variables.confOptions.OutputOptions.RewardFairnessForAllActions
	}
	return false
}

func GetNegativeIncome() bool {
	if variables.experimentOptions.PaymentEnabled {
		return variables.confOptions.OutputOptions.NegativeIncome
	}
	return false
}

func GetComputeWorkFairness() bool {
	return variables.confOptions.OutputOptions.ComputeWorkFairness
}

func GetBucketInfo() bool {
	return variables.confOptions.OutputOptions.BucketInfo
}

func GetLinkInfo() bool {
	return variables.confOptions.OutputOptions.LinkInfo
}

func GetExpeimentId() string {
	return variables.confOptions.OutputOptions.ExperimentId
}

func DoReset() bool {
	return variables.confOptions.OutputOptions.Reset
}

func GetEvaluateInterval() (i int) {
	i = variables.confOptions.OutputOptions.EvaluateInterval
	if i <= 0 {
		return GetIterations()
	}
	return i
}

func GetExperimentString() (exp string) {
	exp = fmt.Sprintf("O%dT%dsS%dk%dTh%dFg%dW%d",
		GetOriginators()*100/GetNetworkSize(),
		GetIterations()/GetRequestsPerSecond(),
		GetIterations(),
		GetBinSize(),
		GetThreshold(),
		GetRefreshRate(),
		GetMaxProximityOrder(),
	)
	if GetPaymentEnabled() {
		exp += "Pay"
	}
	if !GetReciprocityEnabled() {
		exp += "NoRec"
	}
	if IsCacheEnabled() {
		exp += "Cache"
	}
	if IsPreferredChunksEnabled() {
		exp += "Skew"
	}
	if IsAdjustableThreshold() {
		exp += "FgAdj"
	}

	exp += "-" + GetExpeimentId()
	return exp
}
