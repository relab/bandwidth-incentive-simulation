package constants

import (
	"go-incentive-simulation/model/experiments"
)

func SetNumRoutingGoroutines(num int) int {
	//num-- // fot the outputWorker
	//if IsWriteStatesToFile() {
	//	num--
	//}
	//if IsWriteRoutesToFile() {
	//	num--
	//}
	num-- // for the requestWorker
	experiments.Constants.NumRoutingGoroutines = num
	return num
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return experiments.Constants.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return experiments.Constants.ForgivenessEnabled
}

func IsForgivenessDuringRouting() bool {
	return experiments.Constants.ForgivenessDuringRouting
}

func IsCacheEnabled() bool {
	return experiments.Constants.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return experiments.Constants.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return experiments.Constants.RetryWithAnotherPeer
}

func IsForwarderPayForceOriginatorToPay() bool {
	return experiments.Constants.ForwarderPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return experiments.Constants.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return experiments.Constants.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return experiments.Constants.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return experiments.Constants.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return experiments.Constants.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return experiments.Constants.ThresholdEnabled
}

func GetPaymentEnabled() bool {
	return experiments.Constants.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return experiments.Constants.RequestsPerSecond
}

func GetChunks() int {
	return experiments.Constants.Chunks
}

func GetBits() int {
	return experiments.Constants.Bits
}

func GetNetworkSize() int {
	return experiments.Constants.NetworkSize
}

func GetBinSize() int {
	return experiments.Constants.BinSize
}

func GetSimulationRuns() int {
	return 125000
}

func GetRangeAddress() int {
	return experiments.Constants.RangeAddress
}

func GetOriginators() int {
	return experiments.Constants.Originators
}

func GetRefreshRate() int {
	return experiments.Constants.RefreshRate
}

func GetThreshold() int {
	return experiments.Constants.Threshold
}

func GetRandomSeed() int64 {
	return experiments.Constants.RandomSeed
}

func GetMaxProximityOrder() int {
	return experiments.Constants.MaxProximityOrder
}

func GetPrice() int {
	return experiments.Constants.Price
}

func GetSameOriginator() bool {
	return experiments.Constants.SameOriginator
}

func GetEdgeLock() bool {
	return experiments.Constants.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return experiments.Constants.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return experiments.Constants.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return experiments.Constants.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return experiments.Constants.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return experiments.Constants.DebugPrints
}

func GetDebugInterval() int {
	return experiments.Constants.DebugInterval
}

func GetNumRoutingGoroutines() int {
	return experiments.Constants.NumRoutingGoroutines
}

func GetEpoch() int {
	return experiments.Constants.Epoch
}
