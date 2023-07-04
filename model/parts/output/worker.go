package output

import (
	"go-incentive-simulation/config"
	"sync"
)

func Worker(outputChan chan OutputStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	var outputStruct OutputStruct
	counter := 0

	loggers := CreateLoggers()
	logInterval := config.GetEvaluateInterval()
	reset := config.DoReset()

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

func CreateLoggers() []LogResetUpdater {
	loggers := make([]LogResetUpdater, 0)

	if config.GetAverageNumberOfHops() {
		hopInfo := InitHopInfo()
		defer hopInfo.Close()
		loggers = append(loggers, hopInfo)
	}

	if config.GetAverageNumberOfHops() && config.GetPaymentEnabled() {
		hopPaymentInfo := InitHopPaymentInfo()
		defer hopPaymentInfo.Close()
		loggers = append(loggers, hopPaymentInfo)
	}

	if config.GetNegativeIncome() {
		incomeInfo := InitIncomeInfo()
		defer incomeInfo.Close()
		loggers = append(loggers, incomeInfo)
	}

	if config.GetComputeWorkFairness() {
		workInfo := InitWorkInfo()
		defer workInfo.Close()
		loggers = append(loggers, workInfo)
	}

	if config.GetBucketInfo() {
		bucketInfo := InitBucketInfo()
		defer bucketInfo.Close()
		loggers = append(loggers, bucketInfo)
	}

	if config.GetLinkInfo() {
		linkInfo := InitLinkInfo()
		defer linkInfo.Close()
		loggers = append(loggers, linkInfo)
	}

	if config.JustPrintOutPut() {
		outputWriter := InitOutputWriter()
		defer outputWriter.Close()
		loggers = append(loggers, outputWriter)
	}
	return loggers
}
