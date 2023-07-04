package config

type Config struct {
	BaseOptions       baseOptions       `yaml:"BaseOptions"`
	Experiment        experiment        `yaml:"Experiment"`
	ExperimentOptions experimentOptions `yaml:"CustomExperiment"`
}

type experiment struct {
	Name string `yaml:"Name"`
}

type baseOptions struct {
	Iterations                int           `yaml:"Iterations"`
	Bits                      int           `yaml:"Bits"`
	NetworkSize               int           `yaml:"NetworkSize"`
	BinSize                   int           `yaml:"BinSize"`
	Originators               int           `yaml:"Originators"`
	RefreshRate               int           `yaml:"RefreshRate"`
	Threshold                 int           `yaml:"Threshold"`
	RandomSeed                int64         `yaml:"RandomSeed"`
	MaxProximityOrder         int           `yaml:"MaxProximityOrder"`
	Price                     int           `yaml:"Price"`
	RequestsPerSecond         int           `yaml:"RequestsPerSecond"`
	EdgeLock                  bool          `yaml:"EdgeLock"`
	SameOriginator            bool          `yaml:"SameOriginator"`
	WriteRoutesToFile         bool          `yaml:"WriteRoutesToFile"`
	WriteStatesToFile         bool          `yaml:"WriteStatesToFile"`
	IterationMeansUniqueChunk bool          `yaml:"IterationMeansUniqueChunk"`
	DebugPrints               bool          `yaml:"DebugPrints"`
	DebugInterval             int           `yaml:"DebugInterval"`
	NumGoroutines             int           `yaml:"NumGoroutines"`
	OutputEnabled             bool          `yaml:"OutputEnabled"`
	OutputOptions             outputOptions `yaml:"OutputOptions"`
	ReplicationFactor         int           `yaml:"ReplicationFactor"`
	AddressRange              int
	StorageDepth              int
}

type experimentOptions struct {
	ThresholdEnabled                  bool `yaml:"ThresholdEnabled"`
	ReciprocityEnabled                bool `yaml:"ReciprocityEnabled"`
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
	MeanRewardPerForward      bool   `yaml:"MeanRewardPerForward"`
	AverageNumberOfHops       bool   `yaml:"AverageNumberOfHops"`
	HopFractionOfTotalRewards bool   `yaml:"HopFractionOfTotalRewards"`
	NegativeIncome            bool   `yaml:"NegativeIncome"`
	IncomeGini                bool   `yaml:"IncomeGini"`
	HopIncome                 bool   `yaml:"HopIncome"`
	DensenessIncome           bool   `yaml:"DensenessIncome"`
	WorkInfo                  bool   `yaml:"WorkInfo"`
	BucketInfo                bool   `yaml:"BucketInfo"`
	LinkInfo                  bool   `yaml:"LinkInfo"`
	ExperimentId              string `yaml:"ExperimentId"`
	Reset                     bool   `yaml:"Reset"`
	EvaluateInterval          int    `yaml:"EvaluateInterval"`
}
