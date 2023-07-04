package output

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"sync"
)

func Worker(outputChan chan types.OutputStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	var outputStruct types.OutputStruct
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

	if config.GetNegativeIncome() ||
		config.GetIncomeGini() ||
		config.GetHopIncome() ||
		config.GetDensnessIncome() {
		incomeInfo := InitIncomeInfo()
		defer incomeInfo.Close()
		loggers = append(loggers, incomeInfo)
	}

	if config.GetWorkInfo() {
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
