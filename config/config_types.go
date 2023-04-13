package config

type Yml struct {
	ConfOptions ConfVariables   `yaml:"ConfOptions"`
	Experiment  Experiment      `yaml:"Experiment"`
	Custom      CustomVariables `yaml:"Custom"`
}

type Experiment struct {
	ExperimentName string `yaml:"ExperimentName"`
}

type ConfVariables struct {
	Iterations                int   `yaml:"Iterations"`
	Bits                      int   `yaml:"Bits"`
	NetworkSize               int   `yaml:"NetworkSize"`
	BinSize                   int   `yaml:"BinSize"`
	RangeAddress              int   `yaml:"RangeAddress"`
	Originators               int   `yaml:"Originators"`
	RefreshRate               int   `yaml:"RefreshRate"`
	Threshold                 int   `yaml:"Threshold"`
	RandomSeed                int64 `yaml:"RandomSeed"`
	MaxProximityOrder         int   `yaml:"MaxProximityOrder"`
	Price                     int   `yaml:"Price"`
	RequestsPerSecond         int   `yaml:"RequestsPerSecond"`
	EdgeLock                  bool  `yaml:"EdgeLock"`
	SameOriginator            bool  `yaml:"SameOriginator"`
	PrecomputeRespNodes       bool  `yaml:"PrecomputeRespNodes"`
	WriteRoutesToFile         bool  `yaml:"WriteRoutesToFile"`
	WriteStatesToFile         bool  `yaml:"WriteStatesToFile"`
	IterationMeansUniqueChunk bool  `yaml:"IterationMeansUniqueChunk"`
	DebugPrints               bool  `yaml:"DebugPrints"`
	DebugInterval             int   `yaml:"DebugInterval"`
	NumGoroutines             int   `yaml:"NumGoroutines"`
}

type CustomVariables struct {
	ThresholdEnabled                  bool `yaml:"ThresholdEnabled"`
	ForgivenessEnabled                bool `yaml:"ForgivenessEnabled"`
	ForgivenessDuringRouting          bool `yaml:"ForgivenessDuringRouting"`
	PaymentEnabled                    bool `yaml:"PaymentEnabled"`
	MaxPOCheckEnabled                 bool `yaml:"MaxPOCheckEnabled"`
	OnlyOriginatorPays                bool `yaml:"OnlyOriginatorPays"`
	PayOnlyForCurrentRequest          bool `yaml:"PayOnlyForCurrentRequest"`
	ForwardersPayForceOriginatorToPay bool `yaml:"ForwardersPayForceOriginatorToPay"`
	WaitingEnabled                    bool `yaml:"WaitingEnabled"`
	RetryWithAnotherPeer              bool `yaml:"RetryWithAnotherPeer"`
	CacheIsEnabled                    bool `yaml:"CacheIsEnabled"`
	PreferredChunks                   bool `yaml:"PreferredChunks"`
	AdjustableThreshold               bool `yaml:"AdjustableThreshold"`
	PayIfOrigPays                     bool `yaml:"PayIfOrigPays"`
}

type VariablesType struct {
	Iterations                        int
	Bits                              int
	NetworkSize                       int
	BinSize                           int
	RangeAddress                      int
	Originators                       int
	RefreshRate                       int
	Threshold                         int
	RandomSeed                        int64
	MaxProximityOrder                 int
	Price                             int
	RequestsPerSecond                 int
	ThresholdEnabled                  bool
	ForgivenessEnabled                bool
	ForgivenessDuringRouting          bool
	PaymentEnabled                    bool
	MaxPOCheckEnabled                 bool
	WaitingEnabled                    bool
	OnlyOriginatorPays                bool
	PayOnlyForCurrentRequest          bool
	PayIfOrigPays                     bool
	ForwardersPayForceOriginatorToPay bool
	RetryWithAnotherPeer              bool
	CacheIsEnabled                    bool
	PreferredChunks                   bool
	AdjustableThreshold               bool
	EdgeLock                          bool
	SameOriginator                    bool
	PrecomputeRespNodes               bool
	WriteRoutesToFile                 bool
	WriteStatesToFile                 bool
	IterationMeansUniqueChunk         bool
	DebugPrints                       bool
	DebugInterval                     int
	NumGoroutines                     int
}
