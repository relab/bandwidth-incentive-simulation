package config

// These functions modify the respective fields based changes from default

func OmegaExperiment() {
	Variables.ThresholdEnabled = false
	Variables.ForgivenessEnabled = false
	Variables.ForgivenessDuringRouting = false
	Variables.MaxPOCheckEnabled = true
}

func CustomExperiment(customExperiment CustomVariables) {
	Variables.ThresholdEnabled = customExperiment.ThresholdEnabled
	Variables.ForgivenessEnabled = customExperiment.ForgivenessEnabled
	Variables.ForgivenessDuringRouting = customExperiment.ForgivenessDuringRouting
	Variables.PaymentEnabled = customExperiment.PaymentEnabled
	Variables.MaxPOCheckEnabled = customExperiment.MaxPOCheckEnabled
	Variables.OnlyOriginatorPays = customExperiment.OnlyOriginatorPays
	Variables.PayOnlyForCurrentRequest = customExperiment.PayOnlyForCurrentRequest
	Variables.ForwardersPayForceOriginatorToPay = customExperiment.ForwardersPayForceOriginatorToPay
	Variables.WaitingEnabled = customExperiment.WaitingEnabled
	Variables.RetryWithAnotherPeer = customExperiment.RetryWithAnotherPeer
	Variables.CacheIsEnabled = customExperiment.CacheIsEnabled
	Variables.PreferredChunks = customExperiment.PreferredChunks
	Variables.AdjustableThreshold = customExperiment.AdjustableThreshold
	Variables.PayIfOrigPays = customExperiment.PayIfOrigPays
}
