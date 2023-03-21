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
	experiments.Constant.NumRoutingGoroutines = num
	return num
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return experiments.Constant.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return experiments.Constant.ForgivenessEnabled
}

func IsForgivenessDuringRouting() bool {
	return experiments.Constant.ForgivenessDuringRouting
}

func IsCacheEnabled() bool {
	return experiments.Constant.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return experiments.Constant.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return experiments.Constant.RetryWithAnotherPeer
}

func IsForwarderPayForceOriginatorToPay() bool {
	return experiments.Constant.ForwarderPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return experiments.Constant.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return experiments.Constant.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return experiments.Constant.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return experiments.Constant.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return experiments.Constant.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return experiments.Constant.ThresholdEnabled
}

func GetPaymentEnabled() bool {
	return experiments.Constant.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return experiments.Constant.RequestsPerSecond
}

func GetChunks() int {
	return experiments.Constant.Chunks
}

func GetBits() int {
	return experiments.Constant.Bits
}

func GetNetworkSize() int {
	return experiments.Constant.NetworkSize
}

func GetBinSize() int {
	return experiments.Constant.BinSize
}

func GetSimulationRuns() int {
	return 125000
}

func GetRangeAddress() int {
	return experiments.Constant.RangeAddress
}

func GetOriginators() int {
	return experiments.Constant.Originators
}

func GetRefreshRate() int {
	return experiments.Constant.RefreshRate
}

func GetThreshold() int {
	return experiments.Constant.Threshold
}

func GetRandomSeed() int64 {
	return experiments.Constant.RandomSeed
}

func GetMaxProximityOrder() int {
	return experiments.Constant.MaxProximityOrder
}

func GetPrice() int {
	return experiments.Constant.Price
}

func GetSameOriginator() bool {
	return experiments.Constant.SameOriginator
}

func GetEdgeLock() bool {
	return experiments.Constant.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return experiments.Constant.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return experiments.Constant.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return experiments.Constant.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return experiments.Constant.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return experiments.Constant.DebugPrints
}

func GetDebugInterval() int {
	return experiments.Constant.DebugInterval
}

func GetNumRoutingGoroutines() int {
	return experiments.Constant.NumRoutingGoroutines
}

func GetEpoch() int {
	return experiments.Constant.Epoch
}
