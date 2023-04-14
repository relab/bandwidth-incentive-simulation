package config

func GetNumRoutingGoroutines() int {
	num := Variables.NumGoroutines
	//if IsWriteStatesToFile() {
	//	num--
	//}
	//if IsWriteRoutesToFile() {
	//	num--
	//}
	num-- // for the requestWorker
	//num-- // for the outputWorker
	return num
}

func GetNumGoroutines() int {
	return Variables.NumGoroutines
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return Variables.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return Variables.ForgivenessEnabled
}

func IsForgivenessDuringRouting() bool {
	return Variables.ForgivenessDuringRouting
}

func IsCacheEnabled() bool {
	return Variables.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return Variables.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return Variables.RetryWithAnotherPeer
}

func IsForwarderPayForceOriginatorToPay() bool {
	return Variables.ForwardersPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return Variables.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return Variables.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return Variables.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return Variables.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return Variables.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return Variables.ThresholdEnabled
}

func GetPaymentEnabled() bool {
	return Variables.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return Variables.RequestsPerSecond
}

func GetBits() int {
	return Variables.Bits
}

func GetNetworkSize() int {
	return Variables.NetworkSize
}

func GetBinSize() int {
	return Variables.BinSize
}

func GetSimulationRuns() int {
	return 125000
}

func GetRangeAddress() int {
	return Variables.RangeAddress
}

func GetOriginators() int {
	return Variables.Originators
}

func GetRefreshRate() int {
	return Variables.RefreshRate
}

func GetThreshold() int {
	return Variables.Threshold
}

func GetRandomSeed() int64 {
	return Variables.RandomSeed
}

func GetMaxProximityOrder() int {
	return Variables.MaxProximityOrder
}

func GetPrice() int {
	return Variables.Price
}

func GetSameOriginator() bool {
	return Variables.SameOriginator
}

func GetEdgeLock() bool {
	return Variables.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return Variables.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return Variables.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return Variables.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return Variables.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return Variables.DebugPrints
}

func GetDebugInterval() int {
	return Variables.DebugInterval
}

func GetIterations() int {
	return Variables.Iterations
}
