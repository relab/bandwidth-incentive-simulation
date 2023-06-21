package workers

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/output"
	"go-incentive-simulation/model/parts/types"
	"sync"
)

func OutputWorker(outputChan chan types.OutputStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	var outputStruct types.OutputStruct
	counter := 0

	loggers := make([]output.LogResetUpdater, 0)

	if config.GetAverageNumberOfHops() {
		hopInfo := output.InitHopInfo()
		defer hopInfo.Close()
		loggers = append(loggers, hopInfo)
	}

	if config.GetAverageNumberOfHops() && config.GetPaymentEnabled() {
		hopPaymentInfo := output.InitHopPaymentInfo()
		defer hopPaymentInfo.Close()
		loggers = append(loggers, hopPaymentInfo)
	}

	if config.GetNegativeIncome() {
		incomeInfo := output.InitIncomeInfo()
		defer incomeInfo.Close()
		loggers = append(loggers, incomeInfo)
	}

	if config.GetComputeWorkFairness() {
		workInfo := output.InitWorkInfo()
		defer workInfo.Close()
		loggers = append(loggers, workInfo)
	}

	if config.GetBucketInfo() {
		bucketInfo := output.InitBucketInfo()
		defer bucketInfo.Close()
		loggers = append(loggers, bucketInfo)
	}

	if config.GetLinkInfo() {
		linkInfo := output.InitLinkInfo()
		defer linkInfo.Close()
		loggers = append(loggers, linkInfo)
	}

	if config.JustPrintOutPut() {
		outputWriter := output.InitOutputWriter()
		defer outputWriter.Close()
		loggers = append(loggers, outputWriter)
	}

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
