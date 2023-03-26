package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(pauseChan chan bool, continueChan chan bool, requestChan chan types.Request, outputChan chan types.Output, routeChan chan types.RouteData, stateChan chan types.StateSubset, globalState *types.State, wg *sync.WaitGroup) {

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

	var stateSubset types.StateSubset
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
			cacheCounter := update.CacheMap(globalState, requestResult)

			// TODO: originatorIndex is now updated by the requestWorker
			//originatorIndex := UpdateOriginatorIndex(globalState)

			successfulFound := update.SuccessfulFound(globalState, requestResult)
			failedRequestThreshold := update.FailedRequestsThreshold(globalState, requestResult)
			failedRequestAccess := update.FailedRequestsAccess(globalState, requestResult)
			//routeLists := update.RouteListAndFlush(globalState, requestResult, curTimeStep)

			// sending the "output" to the outputWorker
			if constants.GetMaxPOCheckEnabled() {
				if constants.IsDebugPrints() && constants.TimeForNewEpoch(curTimeStep) {
					fmt.Println("outputChan length: ", len(outputChan))
				}
				outputChan <- output
			}

			if constants.IsWriteRoutesToFile() {
				if constants.IsDebugPrints() && constants.TimeForNewEpoch(curTimeStep) {
					fmt.Println("routeChan length: ", len(routeChan))
				}
				routeChan <- types.RouteData{
					Epoch:           int32(request.Epoch),
					Route:           requestResult.Route,
					ChunkId:         requestResult.ChunkId,
					AccessFailed:    requestResult.AccessFailed,
					ThresholdFailed: requestResult.ThresholdFailed}
			}

			if constants.IsWriteStatesToFile() {
				// TODO: Decide on what subset of values we actually would like to store
				stateSubset = types.StateSubset{
					OriginatorIndex:         request.OriginatorIndex,
					PendingMap:              waitingCounter,
					RerouteMap:              retryCounter,
					CacheStruct:             cacheCounter,
					SuccessfulFound:         successfulFound,
					FailedRequestsThreshold: failedRequestThreshold,
					FailedRequestsAccess:    failedRequestAccess,
					TimeStep:                int32(curTimeStep),
					Epoch:                   int32(request.Epoch),
				}
				if constants.IsDebugPrints() && constants.TimeForNewEpoch(curTimeStep) {
					fmt.Println("stateChan length: ", len(stateChan))
				}
				stateChan <- stateSubset
			}
		}
	}
}
