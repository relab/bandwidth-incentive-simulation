package config

func OmegaExperiment() {
	Variable.ThresholdEnabled = false
	Variable.ForgivenessEnabled = false
	Variable.ForgivenessDuringRouting = false
	Variable.MaxPOCheckEnabled = true
}

func BucketSize20And20pOriginators() {
	Variable.BinSize = 20
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000 * 0.2
}

func BucketSize20And100pOriginators() {
	Variable.BinSize = 20
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000
}

func BucketSize16And20pOriginators() {
	Variable.BinSize = 16
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000 * 0.2
}

func BucketSize16And100pOriginators() {
	Variable.BinSize = 16
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000
}

func BucketSize8And20pOriginators() {
	Variable.BinSize = 8
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000 * 0.2
}

func BucketSize8And100pOriginators() {
	Variable.BinSize = 8
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000
}

func BucketSize4And20pOriginators() {
	Variable.BinSize = 4
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000 * 0.2
}

func BucketSize4And100pOriginators() {
	Variable.BinSize = 4
	Variable.Iterations = 10_000_000
	Variable.Originators = 10_000
}
func WaitingEnabled() {
	Variable.WaitingEnabled = true
}

func RetryEnabled() {
	Variable.RetryWithAnotherPeer = true
}

func WaitingAndRetryEnabled() {
	Variable.WaitingEnabled = true
	Variable.RetryWithAnotherPeer = true
}

func DebugPrints() {
	Variable.DebugPrints = true
}

func CustomExperiment(customExperiment YmlVariables) {
	Variable.Runs = customExperiment.Runs
	Variable.Bits = customExperiment.Bits
	Variable.NetworkSize = customExperiment.NetworkSize
	Variable.BinSize = customExperiment.BinSize
	Variable.RangeAddress = customExperiment.RangeAddress
	Variable.Originators = customExperiment.Originators
	Variable.RefreshRate = customExperiment.RefreshRate
	Variable.Threshold = customExperiment.Threshold
	Variable.RandomSeed = customExperiment.RandomSeed
	Variable.MaxProximityOrder = customExperiment.MaxProximityOrder
	Variable.Price = customExperiment.Price
	Variable.Chunks = customExperiment.Chunks
	Variable.RequestsPerSecond = customExperiment.RequestsPerSecond
	Variable.ThresholdEnabled = customExperiment.ThresholdEnabled
	Variable.ForgivenessEnabled = customExperiment.ForgivenessEnabled
	Variable.ForgivenessDuringRouting = customExperiment.ForgivenessDuringRouting
	Variable.PaymentEnabled = customExperiment.PaymentEnabled
	Variable.MaxPOCheckEnabled = customExperiment.MaxPOCheckEnabled
	Variable.OnlyOriginatorPays = customExperiment.OnlyOriginatorPays
	Variable.PayOnlyForCurrentRequest = customExperiment.PayOnlyForCurrentRequest
	Variable.PayIfOrigPays = customExperiment.PayIfOrigPays
	Variable.ForwarderPayForceOriginatorToPay = customExperiment.ForwarderPayForceOriginatorToPay
	Variable.WaitingEnabled = customExperiment.WaitingEnabled
	Variable.RetryWithAnotherPeer = customExperiment.RetryWithAnotherPeer
	Variable.CacheIsEnabled = customExperiment.CacheIsEnabled
	Variable.PreferredChunks = customExperiment.PreferredChunks
	Variable.AdjustableThreshold = customExperiment.AdjustableThreshold
	Variable.EdgeLock = customExperiment.EdgeLock
	Variable.SameOriginator = customExperiment.SameOriginator
	Variable.PrecomputeRespNodes = customExperiment.PrecomputeRespNodes
	Variable.WriteRoutesToFile = customExperiment.WriteRoutesToFile
	Variable.WriteStatesToFile = customExperiment.WriteStatesToFile
	Variable.IterationMeansUniqueChunk = customExperiment.IterationMeansUniqueChunk
	Variable.DebugPrints = customExperiment.DebugPrints
	Variable.DebugInterval = customExperiment.DebugInterval
	Variable.NumRoutingGoroutines = customExperiment.NumRoutingGoroutines
	Variable.Epoch = customExperiment.Epoch
	Variable.Iterations = customExperiment.Iterations
}
