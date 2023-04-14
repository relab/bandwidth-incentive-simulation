package config

type Yml struct {
	ConfOptions      confOptions       `yaml:"ConfOptions"`
	Experiment       experiment        `yaml:"Experiment"`
	CustomExperiment experimentOptions `yaml:"CustomExperiment"`
}

type experiment struct {
	Name string `yaml:"Name"`
}

type VariablesType struct {
	confOptions       confOptions
	experimentOptions experimentOptions
}

type confOptions struct {
	Iterations                int           `yaml:"Iterations"`
	Bits                      int           `yaml:"Bits"`
	NetworkSize               int           `yaml:"NetworkSize"`
	BinSize                   int           `yaml:"BinSize"`
	RangeAddress              int           `yaml:"RangeAddress"`
	Originators               int           `yaml:"Originators"`
	RefreshRate               int           `yaml:"RefreshRate"`
	Threshold                 int           `yaml:"Threshold"`
	RandomSeed                int64         `yaml:"RandomSeed"`
	MaxProximityOrder         int           `yaml:"MaxProximityOrder"`
	Price                     int           `yaml:"Price"`
	RequestsPerSecond         int           `yaml:"RequestsPerSecond"`
	EdgeLock                  bool          `yaml:"EdgeLock"`
	SameOriginator            bool          `yaml:"SameOriginator"`
	PrecomputeRespNodes       bool          `yaml:"PrecomputeRespNodes"`
	WriteRoutesToFile         bool          `yaml:"WriteRoutesToFile"`
	WriteStatesToFile         bool          `yaml:"WriteStatesToFile"`
	IterationMeansUniqueChunk bool          `yaml:"IterationMeansUniqueChunk"`
	DebugPrints               bool          `yaml:"DebugPrints"`
	DebugInterval             int           `yaml:"DebugInterval"`
	NumGoroutines             int           `yaml:"NumGoroutines"`
	OutputEnabled             bool          `yaml:"OutputEnabled"`
	OutputOptions             outputOptions `yaml:"OutputOptions"`
}

type experimentOptions struct {
	ThresholdEnabled                  bool `yaml:"ThresholdEnabled"`
	ForgivenessEnabled                bool `yaml:"ForgivenessEnabled"`
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

type outputOptions struct {
	MeanRewardPerForward              bool `yaml:"MeanRewardPerForward"`
	AverageNumberOfHops               bool `yaml:"AverageNumberOfHops"`
	AverageFractionOfTotalRewardsK8   bool `yaml:"AverageFractionOfTotalRewardsK8"`
	AverageFractionOfTotalRewardsK16  bool `yaml:"AverageFractionOfTotalRewardsK16"`
	RewardFairnessForForwardingAction bool `yaml:"RewardFairnessForForwardingAction"`
	RewardFairnessForStoringAction    bool `yaml:"RewardFairnessForStoringAction"`
	RewardFairnessForAllActions       bool `yaml:"RewardFairnessForAllActions"`
	NegativeIncome                    bool `yaml:"NegativeIncome"`
}
