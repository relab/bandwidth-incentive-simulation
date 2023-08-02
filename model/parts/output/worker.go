package output

import (
	"go-incentive-simulation/config"
	"sync"
)

func Worker(outputChan chan Route, wg *sync.WaitGroup) {
	defer wg.Done()
	var outputStruct Route
	counter := 0

	loggers := CreateLoggers()
	logInterval := config.GetEvaluateInterval()
	reset := config.DoReset()

	for _, logger := range loggers {
		defer logger.Close()
	}

	for outputStruct = range outputChan {
		counter++

		for _, logger := range loggers {
			logger.Update(&outputStruct)

			if counter%logInterval == 0 {
				logger.Log()
				if reset {
					logger.Reset()
				}
			}
		}
	}

}

func CreateLoggers() []LogResetUpdateCloser {
	loggers := make([]LogResetUpdateCloser, 0)

	successInfo := InitSuccessInfo()
	loggers = append(loggers, successInfo)

	if config.GetAverageNumberOfHops() ||
		config.GetHopFractionOfRewards() ||
		config.GetMeanRewardPerForward() {
		hopInfo := InitHopInfo()
		loggers = append(loggers, hopInfo)
	}

	if config.GetPaymentEnabled() &&
		(config.GetAverageNumberOfHops() ||
			config.GetHopFractionOfRewards() ||
			config.GetMeanRewardPerForward()) {
		hopPaymentInfo := InitHopPaymentInfo()
		loggers = append(loggers, hopPaymentInfo)
	}

	if config.GetNegativeIncome() ||
		config.GetIncomeGini() ||
		config.GetHopIncome() ||
		config.GetDensnessIncome() {
		incomeInfo := InitIncomeInfo()
		loggers = append(loggers, incomeInfo)
	}

	if config.GetWorkInfo() {
		workInfo := InitWorkInfo()
		loggers = append(loggers, workInfo)
	}

	if config.GetBucketInfo() {
		bucketInfo := InitBucketInfo()
		loggers = append(loggers, bucketInfo)
	}

	if config.GetLinkInfo() {
		linkInfo := InitLinkInfo()
		loggers = append(loggers, linkInfo)
	}

	if config.JustPrintOutPut() {
		outputWriter := InitOutputWriter()
		loggers = append(loggers, outputWriter)
	}
	return loggers
}
