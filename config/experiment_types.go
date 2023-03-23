package config

type Yml struct {
	Experiment Experiment   `yaml:"Experiment"`
	Custom     YmlVariables `yaml:"Custom"`
}

type Experiment struct {
	ExperimentName       string       `yaml:"ExperimentName"`
	ExperimentSubOptions YmlVariables `yaml:"ExperimentSubOptions"`
}

type YmlVariables struct {
	Runs                             int   `yaml:"Runs"`
	Bits                             int   `yaml:"Bits"`
	NetworkSize                      int   `yaml:"NetworkSize"`
	BinSize                          int   `yaml:"BinSize"`
	RangeAddress                     int   `yaml:"RangeAddress"`
	Originators                      int   `yaml:"Originators"`
	RefreshRate                      int   `yaml:"RefreshRate"`
	Threshold                        int   `yaml:"Threshold"`
	RandomSeed                       int64 `yaml:"RandomSeed"`
	MaxProximityOrder                int   `yaml:"MaxProximityOrder"`
	Price                            int   `yaml:"Price"`
	Chunks                           int   `yaml:"Chunks"`
	RequestsPerSecond                int   `yaml:"RequestsPerSecond"`
	ThresholdEnabled                 bool  `yaml:"ThresholdEnabled"`
	ForgivenessEnabled               bool  `yaml:"ForgivenessEnabled"`
	ForgivenessDuringRouting         bool  `yaml:"ForgivenessDuringRouting"`
	PaymentEnabled                   bool  `yaml:"PaymentEnabled"`
	MaxPOCheckEnabled                bool  `yaml:"MaxPOCheckEnabled"`
	OnlyOriginatorPays               bool  `yaml:"OnlyOriginatorPays"`
	PayOnlyForCurrentRequest         bool  `yaml:"PayOnlyForCurrentRequest"`
	PayIfOrigPays                    bool  `yaml:"PayIfOrigPays"`
	ForwarderPayForceOriginatorToPay bool  `yaml:"ForwarderPayForceOriginatorToPay"`
	WaitingEnabled                   bool  `yaml:"WaitingEnabled"`
	RetryWithAnotherPeer             bool  `yaml:"RetryWithAnotherPeer"`
	CacheIsEnabled                   bool  `yaml:"CacheIsEnabled"`
	PreferredChunks                  bool  `yaml:"PreferredChunks"`
	AdjustableThreshold              bool  `yaml:"AdjustableThreshold"`
	EdgeLock                         bool  `yaml:"EdgeLock"`
	SameOriginator                   bool  `yaml:"SameOriginator"`
	PrecomputeRespNodes              bool  `yaml:"PrecomputeRespNodes"`
	WriteRoutesToFile                bool  `yaml:"WriteRoutesToFile"`
	WriteStatesToFile                bool  `yaml:"WriteStatesToFile"`
	IterationMeansUniqueChunk        bool  `yaml:"IterationMeansUniqueChunk"`
	DebugPrints                      bool  `yaml:"DebugPrints"`
	DebugInterval                    int   `yaml:"DebugInterval"`
	NumRoutingGoroutines             int   `yaml:"NumRoutingGoroutines"`
	Epoch                            int   `yaml:"Epoch"`
	Iterations                       int   `yaml:"Iterations"`
}

type Variables struct {
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
	Iterations                       int
}
