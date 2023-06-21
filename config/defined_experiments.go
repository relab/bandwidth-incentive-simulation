package config

// These functions modify the respective fields based changes from default

func OmegaExperiment() {
	Variables.experimentOptions.ThresholdEnabled = false
	Variables.experimentOptions.ForgivenessEnabled = false
	Variables.experimentOptions.MaxPOCheckEnabled = true
}

func CustomExperiment(customExperiment experimentOptions) {
	Variables.experimentOptions = customExperiment
}
