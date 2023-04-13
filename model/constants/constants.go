package constants

type constant struct {
	runs                              int
	bits                              int
	networkSize                       int
	binSize                           int
	rangeAddress                      int
	originators                       int
	refreshRate                       int
	threshold                         int
	randomSeed                        int64
	maxProximityOrder                 int
	price                             int
	chunks                            int
	requestsPerSecond                 int
	thresholdEnabled                  bool
	forgivenessEnabled                bool
	forgivenessDuringRouting          bool
	paymentEnabled                    bool
	maxPOCheckEnabled                 bool
	waitingEnabled                    bool
	onlyOriginatorPays                bool
	payOnlyForCurrentRequest          bool
	payIfOrigPays                     bool
	forwarderPayForceOriginatorToPay  bool
	retryWithAnotherPeer              bool
	cacheIsEnabled                    bool
	preferredChunks                   bool
	adjustableThreshold               bool
	edgeLock                          bool
	sameOriginator                    bool
	precomputeRespNodes               bool
	writeRoutesToFile                 bool
	writeStatesToFile                 bool
	iterationMeansUniqueChunk         bool
	debugPrints                       bool
	debugInterval                     int
	numRoutingGoroutines              int
	epoch                             int
	meanRewardPerForward              bool
	averageNumberOfHops               bool
	averageFractionOfTotalRewardsK8   bool
	averageFractionOfTotalRewardsK16  bool
	rewardFairnessForForwardingAction bool
	rewardFairnessForStoringAction    bool
	rewardFairnessForAllActions       bool
}

var constants = constant{
	runs:                              1,
	bits:                              16,
	networkSize:                       10000,
	binSize:                           16,
	rangeAddress:                      65536, // 2 * *Bits
	originators:                       1000,  // int(0.001 * NetworkSize)
	refreshRate:                       8,
	threshold:                         16,
	randomSeed:                        123456789,
	maxProximityOrder:                 16,
	price:                             1,
	chunks:                            10000,
	requestsPerSecond:                 12500,   // 12500
	thresholdEnabled:                  true,    // The maximum limit of debt an edge can have in one direction
	forgivenessEnabled:                true,    // Edge debt gets forgiven some amount on an interval (amortized)
	forgivenessDuringRouting:          true,    // If the forgiveness should happen before threshold is checked or after in updateGraph
	paymentEnabled:                    false,   // Nodes pay if they Threshold fail
	maxPOCheckEnabled:                 true,    // Used to find the proper variable called "omega" in the python paper
	onlyOriginatorPays:                false,   // Only the originator will pay, others will threshold fail or wait
	payOnlyForCurrentRequest:          false,   // Only pay for current request or the full debt on the edge
	payIfOrigPays:                     false,   // Only pay if the originator pays -- NOT NEEDED
	forwarderPayForceOriginatorToPay:  false,   // If Threshold fails, forces all the nodes in the route to pay for the current request
	waitingEnabled:                    true,    // When Threshold fails, will wait before trying to traverse same route
	retryWithAnotherPeer:              true,    // The Route to the chunk will try to take many paths to find the chunk
	cacheIsEnabled:                    false,   // Cache, which stores previously looked after chunks on the nodes
	preferredChunks:                   false,   // Fits well with cache, where some chunkIds are chosen more often
	adjustableThreshold:               false,   // The Threshold limit of an edge is determined based on the XOR distance
	edgeLock:                          true,    // Should always be true when using concurrency
	sameOriginator:                    false,   // For testing the usefulness of locking the edges
	precomputeRespNodes:               true,    // Precompute the responsible nodes for every possible chunkId
	writeRoutesToFile:                 false,   // Write the routes to file during run
	writeStatesToFile:                 false,   // Write a subset of the states to file during the run
	iterationMeansUniqueChunk:         false,   // If a single iteration means all unique chunks or include chunks we look for again relating to waiting/retry
	debugPrints:                       true,    // Prints out many useful debug prints during the run
	debugInterval:                     1000000, // How many iterations between each debug print
	numRoutingGoroutines:              25,      // 25 seems to currently be the sweet spot
	epoch:                             1,       // Defined as timeStep / requestsPerSecond, updated by requestWorker
	meanRewardPerForward:              true,    // If the mean reward per forward should be calculated
	averageNumberOfHops:               true,    // If the average number of hops should be calculated
	averageFractionOfTotalRewardsK8:   true,    // If the average fraction of total rewards should be calculated for k=8
	averageFractionOfTotalRewardsK16:  true,    // If the average fraction of total rewards should be calculated for k=16
	rewardFairnessForForwardingAction: true,    // If the reward fairness should be calculated for the forwarding action
	rewardFairnessForStoringAction:    true,    // If the reward fairness should be calculated for the storing action
	rewardFairnessForAllActions:       true,    // If the reward fairness should be calculated for all actions
}

func SetNumRoutingGoroutines(num int) int {
	//num-- // fot the outputWorker
	//if IsWriteStatesToFile() {
	//	num--
	//}
	//if IsWriteRoutesToFile() {
	//	num--
	//}
	num-- // for the outputWorker
	num-- // for the requestWorker
	constants.numRoutingGoroutines = num
	return num
}

// func CreateRangeAddress(c *constant){
// 	c.rangeAddress = 2 * c.bits
// }

// func (c *constant) CreateOriginators(){
// 	c.originators = int(0.001 * float64(c.networkSize))
// }

func GetMeanRewardPerForward() bool {
	return constants.meanRewardPerForward
}

func GetAverageNumberOfHops() bool {
	return constants.averageNumberOfHops
}

func GetAverageFractionOfTotalRewardsK8() bool {
	return constants.averageFractionOfTotalRewardsK8
}

func GetAverageFractionOfTotalRewardsK16() bool {
	return constants.averageFractionOfTotalRewardsK16
}

func GetRewardFairnessForForwardingAction() bool {
	return constants.rewardFairnessForForwardingAction
}

func GetRewardFairnessForStoringAction() bool {
	return constants.rewardFairnessForStoringAction
}

func GetRewardFairnessForAllActions() bool {
	return constants.rewardFairnessForAllActions
}

func SetProximityOrder(po int) {
	constants.maxProximityOrder = po
}

func IsAdjustableThreshold() bool {
	return constants.adjustableThreshold
}

func IsForgivenessEnabled() bool {
	return constants.forgivenessEnabled
}

func IsForgivenessDuringRouting() bool {
	return constants.forgivenessDuringRouting
}

func IsCacheEnabled() bool {
	return constants.cacheIsEnabled
}

func IsPreferredChunksEnabled() bool {
	return constants.preferredChunks
}

func IsRetryWithAnotherPeer() bool {
	return constants.retryWithAnotherPeer
}

func IsForwarderPayForceOriginatorToPay() bool {
	return constants.forwarderPayForceOriginatorToPay
}

func IsPayIfOrigPays() bool {
	return constants.payIfOrigPays
}

func IsPayOnlyForCurrentRequest() bool {
	return constants.payOnlyForCurrentRequest
}

func IsOnlyOriginatorPays() bool {
	return constants.onlyOriginatorPays
}

func IsWaitingEnabled() bool {
	return constants.waitingEnabled
}

func GetMaxPOCheckEnabled() bool {
	return constants.maxPOCheckEnabled
}

func GetThresholdEnabled() bool {
	return constants.thresholdEnabled
}

func GetPaymentEnabled() bool {
	return constants.paymentEnabled
}

func GetRequestsPerSecond() int {
	return constants.requestsPerSecond
}

func GetChunks() int {
	return constants.chunks
}

func GetBits() int {
	return constants.bits
}

func GetNetworkSize() int {
	return constants.networkSize
}

func GetBinSize() int {
	return constants.binSize
}

func GetSimulationRuns() int {
	return 125000
}

func GetRangeAddress() int {
	return constants.rangeAddress
}

func GetOriginators() int {
	return constants.originators
}

func GetRefreshRate() int {
	return constants.refreshRate
}

func GetThreshold() int {
	return constants.threshold
}

func GetRandomSeed() int64 {
	return constants.randomSeed
}

func GetMaxProximityOrder() int {
	return constants.maxProximityOrder
}

func GetPrice() int {
	return constants.price
}

func GetSameOriginator() bool {
	return constants.sameOriginator
}

func GetEdgeLock() bool {
	return constants.edgeLock
}

func IsPrecomputeRespNodes() bool {
	return constants.precomputeRespNodes
}

func IsWriteRoutesToFile() bool {
	return constants.writeRoutesToFile
}

func IsWriteStatesToFile() bool {
	return constants.writeStatesToFile
}

func IsIterationMeansUniqueChunk() bool {
	return constants.iterationMeansUniqueChunk
}

func IsDebugPrints() bool {
	return constants.debugPrints
}

func GetDebugInterval() int {
	return constants.debugInterval
}

func GetNumRoutingGoroutines() int {
	return constants.numRoutingGoroutines
}

func GetEpoch() int {
	return constants.epoch
}
