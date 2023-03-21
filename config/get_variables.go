package config

func SetNumRoutingGoroutines(num int) int {
	//num-- // fot the outputWorker
	//if IsWriteStatesToFile() {
	//	num--
	//}
	//if IsWriteRoutesToFile() {
	//	num--
	//}
	num-- // for the requestWorker
	Variable.NumRoutingGoroutines = num
	return num
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func IsAdjustableThreshold() bool {
	return Variable.AdjustableThreshold
}

func IsForgivenessEnabled() bool {
	return Variable.ForgivenessEnabled
}

func IsForgivenessDuringRouting() bool {
	return Variable.ForgivenessDuringRouting
}

func IsCacheEnabled() bool {
	return Variable.CacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return Variable.PreferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return Variable.RetryWithAnotherPeer
}

func IsForwarderPayForceOriginatorToPay() bool {
	return Variable.ForwarderPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return Variable.PayIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return Variable.PayOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return Variable.OnlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return Variable.WaitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return Variable.MaxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return Variable.ThresholdEnabled
}

func GetPaymentEnabled() bool {
	return Variable.PaymentEnabled
}

func GetRequestsPerSecond() int {
	return Variable.RequestsPerSecond
}

func GetChunks() int {
	return Variable.Chunks
}

func GetBits() int {
	return Variable.Bits
}

func GetNetworkSize() int {
	return Variable.NetworkSize
}

func GetBinSize() int {
	return Variable.BinSize
}

func GetSimulationRuns() int {
	return 125000
}

func GetRangeAddress() int {
	return Variable.RangeAddress
}

func GetOriginators() int {
	return Variable.Originators
}

func GetRefreshRate() int {
	return Variable.RefreshRate
}

func GetThreshold() int {
	return Variable.Threshold
}

func GetRandomSeed() int64 {
	return Variable.RandomSeed
}

func GetMaxProximityOrder() int {
	return Variable.MaxProximityOrder
}

func GetPrice() int {
	return Variable.Price
}

func GetSameOriginator() bool {
	return Variable.SameOriginator
}

func GetEdgeLock() bool {
	return Variable.EdgeLock
}

func IsPrecomputeRespNodes() bool {
	return Variable.PrecomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return Variable.WriteRoutesToFile
}

func IsWriteStatesToFile() bool {
	return Variable.WriteStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return Variable.IterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return Variable.DebugPrints
}

func GetDebugInterval() int {
	return Variable.DebugInterval
}

func GetNumRoutingGoroutines() int {
	return Variable.NumRoutingGoroutines
}

func GetEpoch() int {
	return Variable.Epoch
}
