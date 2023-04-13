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
	//var request types.Request
	var stateSubset types.StateSubset
	var requestResult types.RequestResult
	for {
		select {
		case <-pauseChan:
			continueChan <- true

		case request, open := <-requestChan:
			if !open {
				return
			}

			found, route, thresholdFailed, accessFailed, paymentsList := utils.ConsumeTask(request, globalState.Graph, globalState.RerouteStruct, globalState.CacheStruct)

			requestResult = types.RequestResult{
				Found:                found,
				Route:                route,
				ThresholdFailedLists: thresholdFailed,
				AccessFailed:         accessFailed,
				PaymentList:          paymentsList,
			}

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			//curTimeStep := update.Timestep(globalState)

			curTimeStep := request.TimeStep
			output := update.Graph(globalState, requestResult, curTimeStep)

			pendingStruct := update.PendingMap(globalState, requestResult, request.Epoch)
			rerouteStruct := update.RerouteMap(globalState, requestResult, request.Epoch)
			cacheStruct := update.CacheMap(globalState, requestResult)

			// TODO: originatorIndex is now updated by the requestWorker
			//originatorIndex := UpdateOriginatorIndex(globalState)

			// sending the "output" to the outputWorker
			successfulFound := update.SuccessfulFound(globalState, requestResult)
			failedRequestThreshold := update.FailedRequestsThreshold(globalState, requestResult)
			failedRequestAccess := update.FailedRequestsAccess(globalState, requestResult)
			//routeLists := update.RouteListAndFlush(globalState, requestResult, curTimeStep)

			// sending the "output" to the outputWorker
			if constants.IsDebugPrints() {
				if curTimeStep%constants.GetDebugInterval() == 0 {
					fmt.Println("outputChan length: ", len(outputChan))
				}
			}
			outputChan <- output

			if constants.IsWriteRoutesToFile() {
				if constants.IsDebugPrints() {
					if curTimeStep%constants.GetDebugInterval() == 0 {
						fmt.Println("routeChan length: ", len(routeChan))
					}
				}
				routeChan <- types.RouteData{TimeStep: int32(curTimeStep), Route: route}
			}

			if constants.IsWriteStatesToFile() {
				// TODO: Decide on what subset of values we actually would like to store
				stateSubset = types.StateSubset{
					OriginatorIndex:         int32(request.OriginatorIndex),
					PendingMap:              int32(len(pendingStruct.PendingMap)),
					RerouteMap:              int32(len(rerouteStruct.RerouteMap)),
					CacheStruct:             int32(len(cacheStruct.CacheMap)),
					SuccessfulFound:         successfulFound,
					FailedRequestsThreshold: failedRequestThreshold,
					FailedRequestsAccess:    failedRequestAccess,
					TimeStep:                int32(curTimeStep),
				}
				if constants.IsDebugPrints() {
					if curTimeStep%constants.GetDebugInterval() == 0 {
						fmt.Println("stateChan length: ", len(stateChan))
					}
				}
				stateChan <- stateSubset
			}
		}
	}
}
