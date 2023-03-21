package experiments

var Constant = Constants{
	Runs:                             1,
	Bits:                             16,
	NetworkSize:                      10000,
	BinSize:                          16,
	RangeAddress:                     65536,
	Originators:                      1000,
	RefreshRate:                      8,
	Threshold:                        16,
	RandomSeed:                       123456789,
	MaxProximityOrder:                16,
	Price:                            1,
	Chunks:                           10000,
	RequestsPerSecond:                12500,
	ThresholdEnabled:                 true,
	ForgivenessEnabled:               true,
	ForgivenessDuringRouting:         true,
	PaymentEnabled:                   false,
	MaxPOCheckEnabled:                false,
	OnlyOriginatorPays:               false,
	PayOnlyForCurrentRequest:         false,
	PayIfOrigPays:                    false,
	ForwarderPayForceOriginatorToPay: false,
	WaitingEnabled:                   false,
	RetryWithAnotherPeer:             false,
	CacheIsEnabled:                   false,
	PreferredChunks:                  false,
	AdjustableThreshold:              false,
	EdgeLock:                         true,
	SameOriginator:                   false,
	PrecomputeRespNodes:              true,
	WriteRoutesToFile:                false,
	WriteStatesToFile:                false,
	IterationMeansUniqueChunk:        false,
	DebugPrints:                      false,
	DebugInterval:                    1000000,
	NumRoutingGoroutines:             25,
	Epoch:                            1,
}

func OmegaExperiment() {
	Constant.ThresholdEnabled = false
	Constant.ForgivenessEnabled = false
	Constant.ForgivenessDuringRouting = false
	Constant.MaxPOCheckEnabled = true
}

func WaitingAndRetry() {
	Constant.WaitingEnabled = true
	Constant.RetryWithAnotherPeer = true
}

func CustomExperiment(customExperiment YmlConstants) {
	Constant.Runs = customExperiment.Runs
	Constant.Bits = customExperiment.Bits
	Constant.NetworkSize = customExperiment.NetworkSize
	Constant.BinSize = customExperiment.BinSize
	Constant.RangeAddress = customExperiment.RangeAddress
	Constant.Originators = customExperiment.Originators
	Constant.RefreshRate = customExperiment.RefreshRate
	Constant.Threshold = customExperiment.Threshold
	Constant.RandomSeed = customExperiment.RandomSeed
	Constant.MaxProximityOrder = customExperiment.MaxProximityOrder
	Constant.Price = customExperiment.Price
	Constant.Chunks = customExperiment.Chunks
	Constant.RequestsPerSecond = customExperiment.RequestsPerSecond
	Constant.ThresholdEnabled = customExperiment.ThresholdEnabled
	Constant.ForgivenessEnabled = customExperiment.ForgivenessEnabled
	Constant.ForgivenessDuringRouting = customExperiment.ForgivenessDuringRouting
	Constant.PaymentEnabled = customExperiment.PaymentEnabled
	Constant.MaxPOCheckEnabled = customExperiment.MaxPOCheckEnabled
	Constant.OnlyOriginatorPays = customExperiment.OnlyOriginatorPays
	Constant.PayOnlyForCurrentRequest = customExperiment.PayOnlyForCurrentRequest
	Constant.PayIfOrigPays = customExperiment.PayIfOrigPays
	Constant.ForwarderPayForceOriginatorToPay = customExperiment.ForwarderPayForceOriginatorToPay
	Constant.WaitingEnabled = customExperiment.WaitingEnabled
	Constant.RetryWithAnotherPeer = customExperiment.RetryWithAnotherPeer
	Constant.CacheIsEnabled = customExperiment.CacheIsEnabled
	Constant.PreferredChunks = customExperiment.PreferredChunks
	Constant.AdjustableThreshold = customExperiment.AdjustableThreshold
	Constant.EdgeLock = customExperiment.EdgeLock
	Constant.SameOriginator = customExperiment.SameOriginator
	Constant.PrecomputeRespNodes = customExperiment.PrecomputeRespNodes
	Constant.WriteRoutesToFile = customExperiment.WriteRoutesToFile
	Constant.WriteStatesToFile = customExperiment.WriteStatesToFile
	Constant.IterationMeansUniqueChunk = customExperiment.IterationMeansUniqueChunk
	Constant.DebugPrints = customExperiment.DebugPrints
	Constant.DebugInterval = customExperiment.DebugInterval
	Constant.NumRoutingGoroutines = customExperiment.NumRoutingGoroutines
	Constant.Epoch = customExperiment.Epoch
}
