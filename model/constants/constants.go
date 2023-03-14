package constants

type constant struct {
	runs                             int
	bits                             int
	networkSize                      int
	binSize                          int
	rangeAddress                     int
	originators                      int
	refreshRate                      int
	threshold                        int
	randomSeed                       int64
	maxProximityOrder                int
	price                            int
	chunks                           int
	requestsPerSecond                int
	thresholdEnabled                 bool
	forgivenessEnabled               bool
	paymentEnabled                   bool
	maxPOCheckEnabled                bool
	waitingEnabled                   bool
	onlyOriginatorPays               bool
	payOnlyForCurrentRequest         bool
	payIfOrigPays                    bool
	forwarderPayForceOriginatorToPay bool
	retryWithAnotherPeer             bool
	cacheIsEnabled                   bool
	preferredChunks                  bool
	adjustableThreshold              bool
	edgeLock                         bool
	sameOriginator                   bool
	precomputeRespNodes              bool
	writeRoutesToFile                bool
	writeStatesToFile                bool
	numGoroutines                    int
}

var Constants = constant{
	runs:                             1,
	bits:                             16,
	networkSize:                      10000,
	binSize:                          16,
	rangeAddress:                     65536, // 2 * *Bits
	originators:                      1000,  // int(0.001 * NetworkSize)
	refreshRate:                      8,
	threshold:                        16,
	randomSeed:                       123456789,
	maxProximityOrder:                16,
	price:                            1,
	chunks:                           10000,
	requestsPerSecond:                12500, // 12500
	thresholdEnabled:                 true,
	forgivenessEnabled:               false,
	paymentEnabled:                   true,
	maxPOCheckEnabled:                true,
	waitingEnabled:                   false,
	onlyOriginatorPays:               false,
	payOnlyForCurrentRequest:         false,
	payIfOrigPays:                    false,
	forwarderPayForceOriginatorToPay: false,
	retryWithAnotherPeer:             false,
	cacheIsEnabled:                   false,
	preferredChunks:                  false, // Fits well with cache
	adjustableThreshold:              false,
	edgeLock:                         true,  // Should always be true when using concurrency
	sameOriginator:                   false, // For testing the usefulness of locking the edges
	precomputeRespNodes:              true,  // Precompute the responsible nodes for every possible chunkId
	writeRoutesToFile:                true,  // Write the routes to file during run
	writeStatesToFile:                true,  // Write a subset of the states to file during the run
	numGoroutines:                    25,    // 25 seems to currently be the sweet spot
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func (c *constant) IsAdjustableThreshold() bool {
	return c.adjustableThreshold
}

func (c *constant) IsForgivenessEnabled() bool {
	return c.forgivenessEnabled
}

func (c *constant) IsCacheEnabled() bool {
	return c.cacheIsEnabled
}

func (c *constant) IsPreferredChunksEnabled() bool {
	return c.preferredChunks
}

func (c *constant) IsRetryWithAnotherPeer() bool {
	return c.retryWithAnotherPeer
}

func (c *constant) IsForwarderPayForceOriginatorToPay() bool {
	return c.forwarderPayForceOriginatorToPay
}

func (c *constant) IsPayIfOrigPays() bool {
	return c.payIfOrigPays
}

func (c *constant) IsPayOnlyForCurrentRequest() bool {
	return c.payOnlyForCurrentRequest
}

func (c *constant) IsOnlyOriginatorPays() bool {
	return c.onlyOriginatorPays
}

func (c *constant) IsWaitingEnabled() bool {
	return c.waitingEnabled
}

func (c *constant) GetMaxPOCheckEnabled() bool {
	return c.maxPOCheckEnabled
}

func (c *constant) GetThresholdEnabled() bool {
	return c.thresholdEnabled
}

func (c *constant) GetPaymentEnabled() bool {
	return c.paymentEnabled
}

func (c *constant) GetRequestsPerSecond() int {
	return c.requestsPerSecond
}

func (c *constant) GetChunks() int {
	return c.chunks
}

func (c *constant) GetBits() int {
	return c.bits
}

func (c *constant) GetNetworkSize() int {
	return c.networkSize
}

func (c *constant) GetBinSize() int {
	return c.binSize
}

func GetSimulationRuns() int {
	return 125000
}

func (c *constant) GetRangeAddress() int {
	return c.rangeAddress
}

func (c *constant) GetOriginators() int {
	return c.originators
}

func (c *constant) GetRefreshRate() int {
	return c.refreshRate
}

func (c *constant) GetThreshold() int {
	return c.threshold
}

func (c *constant) GetRandomSeed() int64 {
	return c.randomSeed
}

func (c *constant) GetMaxProximityOrder() int {
	return c.maxProximityOrder
}

func (c *constant) GetPrice() int {
	return c.price
}

func (c *constant) GetSameOriginator() bool {
	return c.sameOriginator
}

func (c *constant) GetEdgeLock() bool {
	return c.edgeLock
}

func (c *constant) IsPrecomputeRespNodes() bool {
	return c.precomputeRespNodes
}

func (c *constant) IsWriteRoutesToFile() bool {
	return c.writeRoutesToFile
}

func (c *constant) IsWriteStatesToFile() bool {
	return c.writeStatesToFile
}

func (c *constant) GetNumGoroutines() int {
	return c.numGoroutines
}
