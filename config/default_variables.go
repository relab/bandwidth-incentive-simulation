package config

func GetDefaultVariables() VariablesType {
	return defaultVariables
}

var defaultVariables = VariablesType{
	confOptions: confOptions{
		Iterations:                10_000_000, // 10_000_000
		Bits:                      16,         // 16
		NetworkSize:               10000,      // 10000
		BinSize:                   16,         // 16
		RangeAddress:              65536,      // 2 ** Bits
		Originators:               1000,       // 0.01 * NetworkSize
		RefreshRate:               8,          // 8
		Threshold:                 16,         // 16
		RandomSeed:                123456789,  // 123456789
		MaxProximityOrder:         16,         // 16
		Price:                     1,          // 1
		RequestsPerSecond:         12500,      // 12500
		EdgeLock:                  true,       // false
		SameOriginator:            false,      // false
		PrecomputeRespNodes:       true,       // false
		WriteRoutesToFile:         false,      // false
		WriteStatesToFile:         false,      // false
		IterationMeansUniqueChunk: false,      // false
		DebugPrints:               false,      // false
		DebugInterval:             1000000,    // 1000000
		NumGoroutines:             -1,         // -1 means gets overwritten by numCPU
		OutputEnabled:             false,      // false
		OutputOptions: outputOptions{
			MeanRewardPerForward:              false,     // false
			AverageNumberOfHops:               false,     // false
			AverageFractionOfTotalRewardsK8:   false,     // false
			AverageFractionOfTotalRewardsK16:  false,     // false
			RewardFairnessForForwardingAction: false,     // false
			RewardFairnessForStoringAction:    false,     // false
			RewardFairnessForAllActions:       false,     // false
			NegativeIncome:                    false,     // false
			ComputeWorkFairness:               false,     // false
			ExpeimentId:                       "default", // default
		},
	},
	experimentOptions: experimentOptions{
		ThresholdEnabled:                  true,  // true
		ForgivenessEnabled:                true,  // true
		PaymentEnabled:                    false, // false
		MaxPOCheckEnabled:                 false, // false
		OnlyOriginatorPays:                false, // false
		PayOnlyForCurrentRequest:          false, // false
		PayIfOrigPays:                     false, // false
		ForwardersPayForceOriginatorToPay: false, // false
		WaitingEnabled:                    false, // false
		RetryWithAnotherPeer:              false, // false
		CacheIsEnabled:                    false, // false
		PreferredChunks:                   false, // false
		AdjustableThreshold:               false, // false
	},
}
