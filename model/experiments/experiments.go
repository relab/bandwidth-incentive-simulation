package experiments

type Constants struct {
	Runs                             int
	Bits                             int
	NetworkSize                      int
	BinSize                          int
	RangeAddress                     int
	Originators                      int
	RefreshRate                      int
	Threshold                        int
	RandomSeed                       int64
	MaxProximityOrder                int
	Price                            int
	Chunks                           int
	RequestsPerSecond                int
	ThresholdEnabled                 bool
	ForgivenessEnabled               bool
	ForgivenessDuringRouting         bool
	PaymentEnabled                   bool
	MaxPOCheckEnabled                bool
	WaitingEnabled                   bool
	OnlyOriginatorPays               bool
	PayOnlyForCurrentRequest         bool
	PayIfOrigPays                    bool
	ForwarderPayForceOriginatorToPay bool
	RetryWithAnotherPeer             bool
	CacheIsEnabled                   bool
	PreferredChunks                  bool
	AdjustableThreshold              bool
	EdgeLock                         bool
	SameOriginator                   bool
	PrecomputeRespNodes              bool
	WriteRoutesToFile                bool
	WriteStatesToFile                bool
	IterationMeansUniqueChunk        bool
	DebugPrints                      bool
	DebugInterval                    int
	NumRoutingGoroutines             int
	Epoch                            int
}

type Experiments map[int]Constants

var ex1 = Constants{
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
	WaitingEnabled:                   true,
	RetryWithAnotherPeer:             true,
	CacheIsEnabled:                   false,
	PreferredChunks:                  false,
	AdjustableThreshold:              false,
	EdgeLock:                         true,
	SameOriginator:                   false,
	PrecomputeRespNodes:              true,
	WriteRoutesToFile:                false,
	WriteStatesToFile:                false,
	IterationMeansUniqueChunk:        false,
	DebugPrints:                      true,
	DebugInterval:                    1000000,
	NumRoutingGoroutines:             25,
	Epoch:                            1,
}

var defaultExperiment = Constants{
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
	WaitingEnabled:                   true,
	RetryWithAnotherPeer:             true,
	CacheIsEnabled:                   false,
	PreferredChunks:                  false,
	AdjustableThreshold:              false,
	EdgeLock:                         true,
	SameOriginator:                   false,
	PrecomputeRespNodes:              true,
	WriteRoutesToFile:                false,
	WriteStatesToFile:                false,
	IterationMeansUniqueChunk:        false,
	DebugPrints:                      true,
	DebugInterval:                    1000000,
	NumRoutingGoroutines:             25,
	Epoch:                            1,
}

var Experiment = Experiments{
	1: ex1,
	2: defaultExperiment,
}
