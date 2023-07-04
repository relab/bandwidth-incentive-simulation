package config

import (
	"fmt"
)

func GetNumRoutingGoroutines() int {
	num := theconfig.BaseOptions.NumGoroutines
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
	return theconfig.BaseOptions.NumGoroutines
}

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return theconfig.ExperimentOptions.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return theconfig.ExperimentOptions.ForgivenessEnabled
}

func IsCacheEnabled() bool {
	return theconfig.ExperimentOptions.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return theconfig.ExperimentOptions.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return theconfig.ExperimentOptions.RetryWithAnotherPeer
}

func IsForwardersPayForceOriginatorToPay() bool {
	return theconfig.ExperimentOptions.ForwardersPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return theconfig.ExperimentOptions.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return theconfig.ExperimentOptions.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return theconfig.ExperimentOptions.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return theconfig.ExperimentOptions.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return theconfig.ExperimentOptions.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return theconfig.ExperimentOptions.ThresholdEnabled
}

func GetReciprocityEnabled() bool {
	return theconfig.ExperimentOptions.ReciprocityEnabled
}

func GetPaymentEnabled() bool {
	return theconfig.ExperimentOptions.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return theconfig.BaseOptions.RequestsPerSecond
}

func GetIterations() int {
	return theconfig.BaseOptions.Iterations
}

func GetBits() int {
	return theconfig.BaseOptions.Bits
}

func GetNetworkSize() int {
	return theconfig.BaseOptions.NetworkSize
}

func GetBinSize() int {
	return theconfig.BaseOptions.BinSize
}

func GetAddressRange() int {
	return theconfig.BaseOptions.AddressRange
}

func GetStorageDepth() int {
	return theconfig.BaseOptions.StorageDepth
}

func GetOriginators() int {
	return theconfig.BaseOptions.Originators
}

func GetRefreshRate() int {
	return theconfig.BaseOptions.RefreshRate
}

func GetThreshold() int {
	return theconfig.BaseOptions.Threshold
}

func GetRandomSeed() int64 {
	return theconfig.BaseOptions.RandomSeed
}

func GetMaxProximityOrder() int {
	return theconfig.BaseOptions.MaxProximityOrder
}

func GetPrice() int {
	return theconfig.BaseOptions.Price
}

func GetSameOriginator() bool {
	return theconfig.BaseOptions.SameOriginator
}

func IsEdgeLock() bool {
	return theconfig.BaseOptions.EdgeLock
}

func IsWriteStatesToFile() bool {
	return theconfig.BaseOptions.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return theconfig.BaseOptions.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return theconfig.BaseOptions.DebugPrints
}

func GetDebugInterval() int {
	return theconfig.BaseOptions.DebugInterval
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

func GetReplicationFactor() int {
	return theconfig.BaseOptions.ReplicationFactor
}

func IsOutputEnabled() bool {
	return theconfig.BaseOptions.OutputEnabled
}

func JustPrintOutPut() bool {
	if !theconfig.BaseOptions.OutputOptions.MeanRewardPerForward &&
		!theconfig.BaseOptions.OutputOptions.AverageNumberOfHops &&
		!theconfig.BaseOptions.OutputOptions.AverageFractionOfTotalRewardsK16 &&
		!theconfig.BaseOptions.OutputOptions.RewardFairnessForForwardingAction &&
		!theconfig.BaseOptions.OutputOptions.RewardFairnessForStoringAction &&
		!theconfig.BaseOptions.OutputOptions.RewardFairnessForAllActions &&
		!theconfig.BaseOptions.OutputOptions.NegativeIncome {
		return true
	}
	return false
}

func GetMeanRewardPerForward() bool {
	if theconfig.BaseOptions.OutputEnabled && theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.MeanRewardPerForward
	}
	return false
}

func GetAverageNumberOfHops() bool {
	if theconfig.BaseOptions.OutputEnabled && theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.AverageNumberOfHops
	}
	return false
}

func GetAverageFractionOfTotalRewardsK8() bool {
	return false
}

func GetAverageFractionOfTotalRewardsK16() bool {
	if theconfig.BaseOptions.OutputEnabled &&
		theconfig.BaseOptions.BinSize == 16 &&
		theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.AverageFractionOfTotalRewardsK16
	}
	return false

}

func GetRewardFairnessForForwardingAction() bool {
	if theconfig.BaseOptions.OutputEnabled && theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.RewardFairnessForForwardingAction
	}
	return false
}

func GetRewardFairnessForStoringAction() bool {
	if theconfig.BaseOptions.OutputEnabled && theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.RewardFairnessForStoringAction
	}
	return false
}

func GetRewardFairnessForAllActions() bool {
	if theconfig.BaseOptions.OutputEnabled && theconfig.ExperimentOptions.MaxPOCheckEnabled {
		return theconfig.BaseOptions.OutputOptions.RewardFairnessForAllActions
	}
	return false
}

func GetNegativeIncome() bool {
	if theconfig.ExperimentOptions.PaymentEnabled {
		return theconfig.BaseOptions.OutputOptions.NegativeIncome
	}
	return false
}

func GetComputeWorkFairness() bool {
	return theconfig.BaseOptions.OutputOptions.ComputeWorkFairness
}

func GetBucketInfo() bool {
	return theconfig.BaseOptions.OutputOptions.BucketInfo
}

func GetLinkInfo() bool {
	return theconfig.BaseOptions.OutputOptions.LinkInfo
}

func GetExpeimentId() string {
	return theconfig.BaseOptions.OutputOptions.ExperimentId
}

func DoReset() bool {
	return theconfig.BaseOptions.OutputOptions.Reset
}

func GetEvaluateInterval() (i int) {
	return theconfig.BaseOptions.OutputOptions.EvaluateInterval
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
