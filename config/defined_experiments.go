package config

// These functions modify the respective fields based changes from default

func OmegaExperiment() {
	variables.experimentOptions.ThresholdEnabled = false
	variables.experimentOptions.ForgivenessEnabled = false
	variables.experimentOptions.MaxPOCheckEnabled = true
}

func CustomExperiment(customExperiment experimentOptions) {
	variables.experimentOptions = customExperiment
}
