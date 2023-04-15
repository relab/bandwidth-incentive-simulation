package config

// These functions modify the respective fields based changes from default

func OmegaExperiment() {
	Variables.experimentOptions.ThresholdEnabled = false
	Variables.experimentOptions.ForgivenessEnabled = false
	Variables.experimentOptions.MaxPOCheckEnabled = true
}

func CustomExperiment(customExperiment experimentOptions) {
	Variables.experimentOptions.ThresholdEnabled = customExperiment.ThresholdEnabled
	Variables.experimentOptions.ForgivenessEnabled = customExperiment.ForgivenessEnabled
	Variables.experimentOptions.PaymentEnabled = customExperiment.PaymentEnabled
	Variables.experimentOptions.MaxPOCheckEnabled = customExperiment.MaxPOCheckEnabled
	Variables.experimentOptions.OnlyOriginatorPays = customExperiment.OnlyOriginatorPays
	Variables.experimentOptions.PayOnlyForCurrentRequest = customExperiment.PayOnlyForCurrentRequest
	Variables.experimentOptions.ForwardersPayForceOriginatorToPay = customExperiment.ForwardersPayForceOriginatorToPay
	Variables.experimentOptions.WaitingEnabled = customExperiment.WaitingEnabled
	Variables.experimentOptions.RetryWithAnotherPeer = customExperiment.RetryWithAnotherPeer
	Variables.experimentOptions.CacheIsEnabled = customExperiment.CacheIsEnabled
	Variables.experimentOptions.PreferredChunks = customExperiment.PreferredChunks
	Variables.experimentOptions.AdjustableThreshold = customExperiment.AdjustableThreshold
	Variables.experimentOptions.PayIfOrigPays = customExperiment.PayIfOrigPays
}
