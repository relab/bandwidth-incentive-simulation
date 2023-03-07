package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(requestChan chan types.Request, routeChan chan types.RouteData, stateChan chan types.StateSubset, globalState *types.State, wg *sync.WaitGroup, numLoops int) {
	defer wg.Done()
	var request types.Request
	var stateSubset types.StateSubset
	var requestResult types.RequestResult
	for i := 0; i < numLoops; i++ {
		request = <-requestChan

		found, route, thresholdFailed, accessFailed, paymentsList := utils.ConsumeTask(&request, globalState.Graph, globalState.RerouteStruct, globalState.CacheStruct)

		requestResult = types.RequestResult{
			Found:                found,
			Route:                route,
			ThresholdFailedLists: thresholdFailed,
			AccessFailed:         accessFailed,
			PaymentList:          paymentsList,
		}

		// TODO: decide on where we should update the timestep. At request creation or request fulfillment
		//curTimeStep := update.Timestep(globalState)
		curTimeStep := int(request.TimeStep)
		update.Graph(globalState, requestResult, curTimeStep)
		pendingStruct := update.PendingMap(globalState, requestResult)
		rerouteStruct := update.RerouteMap(globalState, requestResult)
		cacheStruct := update.CacheMap(globalState, requestResult)
		// TODO: originatorIndex is now updated by the requestWorker
		//originatorIndex := UpdateOriginatorIndex(globalState)
		successfulFound := update.SuccessfulFound(globalState, requestResult)
		failedRequestThreshold := update.FailedRequestsThreshold(globalState, requestResult)
		failedRequestAccess := update.FailedRequestsAccess(globalState, requestResult)
		//routeLists := update.RouteListAndFlush(globalState, requestResult, curTimeStep)

		if constants.Constants.IsWriteRoutesToFile() {
			if curTimeStep%1000000 == 0 {
				fmt.Println("routeChan length: ", len(routeChan))
			}
			routeChan <- types.RouteData{TimeStep: int32(curTimeStep), Route: route}
		}

		if constants.Constants.IsWriteStatesToFile() {
			// TODO: Decide on what subset of values we actually would like to store
			stateSubset = types.StateSubset{
				OriginatorIndex:         request.OriginatorIndex,
				PendingMap:              int32(len(pendingStruct.PendingMap)),
				RerouteMap:              int32(len(rerouteStruct.RerouteMap)),
				CacheStruct:             int32(len(cacheStruct.CacheMap)),
				SuccessfulFound:         successfulFound,
				FailedRequestsThreshold: failedRequestThreshold,
				FailedRequestsAccess:    failedRequestAccess,
				TimeStep:                int32(curTimeStep),
			}
			if curTimeStep%1000000 == 0 {
				fmt.Println("stateChan length: ", len(stateChan))
			}

			stateChan <- stateSubset
		}

	}
}
