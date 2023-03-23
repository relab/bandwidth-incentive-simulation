package config

func OmegaExperiment() {
	Variable.ThresholdEnabled = false
	Variable.ForgivenessEnabled = false
	Variable.ForgivenessDuringRouting = false
	Variable.MaxPOCheckEnabled = true
}

func BucketSize20And20pOriginators() {
	Variable.BinSize = 20
	Variable.Originators = 10_000 * 0.2
}

func BucketSize20And100pOriginators() {
	Variable.BinSize = 20
	Variable.Originators = 10_000
}

func BucketSize16And20pOriginators() {
	Variable.BinSize = 16
	Variable.Originators = 10_000 * 0.2
}

func BucketSize16And100pOriginators() {
	Variable.BinSize = 16
	Variable.Originators = 10_000
}

func BucketSize8And20pOriginators() {
	Variable.BinSize = 8
	Variable.Originators = 10_000 * 0.2
}

func BucketSize8And100pOriginators() {
	Variable.BinSize = 8
	Variable.Originators = 10_000
}

func BucketSize4And20pOriginators() {
	Variable.BinSize = 4
	Variable.Originators = 10_000 * 0.2
}

func BucketSize4And100pOriginators() {
	Variable.BinSize = 4
	Variable.Originators = 10_000
}

func ConfOptions(configOptions ConfVariables) {
	Variable.Bits = configOptions.Bits
	Variable.NetworkSize = configOptions.NetworkSize
	Variable.BinSize = configOptions.BinSize
	Variable.RangeAddress = configOptions.RangeAddress
	Variable.Originators = configOptions.Originators
	Variable.RefreshRate = configOptions.RefreshRate
	Variable.Threshold = configOptions.Threshold
	Variable.RandomSeed = configOptions.RandomSeed
	Variable.MaxProximityOrder = configOptions.MaxProximityOrder
	Variable.Price = configOptions.Price
	Variable.Chunks = configOptions.Chunks
	Variable.RequestsPerSecond = configOptions.RequestsPerSecond
	Variable.EdgeLock = configOptions.EdgeLock
	Variable.SameOriginator = configOptions.SameOriginator
	Variable.PrecomputeRespNodes = configOptions.PrecomputeRespNodes
	Variable.WriteRoutesToFile = configOptions.WriteRoutesToFile
	Variable.WriteStatesToFile = configOptions.WriteStatesToFile
	Variable.IterationMeansUniqueChunk = configOptions.IterationMeansUniqueChunk
	Variable.DebugPrints = configOptions.DebugPrints
	Variable.DebugInterval = configOptions.DebugInterval
	Variable.NumRoutingGoroutines = configOptions.NumRoutingGoroutines
	Variable.Epoch = configOptions.Epoch
	Variable.Iterations = configOptions.Iterations
}

func CustomExperiment(customExperiment CustomVariables) {
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
}
