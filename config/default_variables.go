package config

var Variable = Variables{
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
	Iterations:                       10_000_000,
}
