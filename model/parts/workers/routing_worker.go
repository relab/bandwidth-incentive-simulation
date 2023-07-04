package workers

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(pauseChan chan bool, continueChan chan bool, requestChan chan types.Request, outputChan chan types.OutputStruct, routeChan chan types.RouteData, stateChan chan types.StateSubset, globalState *types.State, wg *sync.WaitGroup) {

	defer wg.Done()
	openChannel := true
	var request types.Request
	var requestResult types.RequestResult
	var route []types.NodeId
	var paymentList []types.Payment
	var found bool
	var accessFailed bool
	var thresholdFailed bool
	var foundByCaching bool

	for {
		select {
		case <-pauseChan:
			continueChan <- true

		case request, openChannel = <-requestChan:
			if !openChannel {
				return
			}

			route, paymentList, found, accessFailed, thresholdFailed, foundByCaching = utils.FindRoute(request, globalState.Graph)

			requestResult = types.RequestResult{
				Route:           route,
				PaymentList:     paymentList,
				ChunkId:         request.ChunkId,
				Found:           found,
				AccessFailed:    accessFailed,
				ThresholdFailed: thresholdFailed,
				FoundByCaching:  foundByCaching,
			}

			curTimeStep := request.TimeStep
			output := update.Graph(globalState, requestResult, curTimeStep)

			waitingCounter := update.Pending(globalState, requestResult, request.Epoch)
			retryCounter := update.Reroute(globalState, requestResult, request.Epoch)
			cacheHits := update.Cache(globalState, requestResult)

			// sending the "output" to the outputWorker
			successfulFound := update.SuccessfulFound(globalState, requestResult)
			failedRequestThreshold := update.FailedRequestsThreshold(globalState, requestResult)
			failedRequestAccess := update.FailedRequestsAccess(globalState, requestResult)

			// sending the "output" to the outputWorker

			if config.GetMaxPOCheckEnabled() {
				if config.IsDebugPrints() && config.TimeForDebugPrints(curTimeStep) {
					fmt.Println("outputChan length: ", len(outputChan))
				}
				outputChan <- output
			}

			if config.IsWriteStatesToFile() {
				if config.IsDebugPrints() && config.TimeForDebugPrints(curTimeStep) {
					fmt.Println("stateChan length: ", len(stateChan))
				}
				// TODO: Decide on what subset of values we actually would like to store
				stateChan <- types.StateSubset{
					WaitingCounter:          int(waitingCounter),
					RetryCounter:            int(retryCounter),
					CacheHits:               int(cacheHits),
					ChunkId:                 int(request.ChunkId),
					OriginatorIndex:         request.OriginatorIndex,
					SuccessfulFound:         int(successfulFound),
					FailedRequestsThreshold: int(failedRequestThreshold),
					FailedRequestsAccess:    int(failedRequestAccess),
					TimeStep:                curTimeStep,
					Epoch:                   request.Epoch,
				}
			}
		}
	}
}
