package workers

import (
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(requestChan chan types.Request, newStateChan chan bool, globalState *types.State, stateList []types.State, wg *sync.WaitGroup, numLoops int) {
	defer wg.Done()
	var request types.Request
	for i := 0; i < numLoops; i++ {
		request = <-requestChan

		found, route, thresholdFailed, accessFailed, paymentsList := utils.ConsumeTask(&request, globalState.Graph, globalState.RerouteStruct, globalState.CacheStruct)

		policyOutput := types.Policy{
			Found:                found,
			Route:                route,
			ThresholdFailedLists: thresholdFailed,
			AccessFailed:         accessFailed,
			PaymentList:          paymentsList,
		}

		//policyChan <- policy

		//curTimeStep := update.Timestep(globalState)
		curTimeStep := int(request.TimeStep)
		graph := update.Graph(globalState, policyOutput, curTimeStep)
		pendingStruct := update.PendingMap(globalState, policyOutput)
		rerouteStruct := update.RerouteMap(globalState, policyOutput)
		cacheStruct := update.CacheMap(globalState, policyOutput)
		//originatorIndex := UpdateOriginatorIndex(globalState)
		successfulFound := update.SuccessfulFound(globalState, policyOutput)
		failedRequestThreshold := update.FailedRequestsThreshold(globalState, policyOutput)
		failedRequestAccess := update.FailedRequestsAccess(globalState, policyOutput)
		//routeLists := update.RouteListAndFlush(globalState, policyOutput, curTimeStep)

		newState := &types.State{
			Graph:       graph,
			Originators: globalState.Originators,
			NodesId:     globalState.NodesId,
			//RouteLists:              routeLists,
			PendingStruct:           pendingStruct,
			RerouteStruct:           rerouteStruct,
			CacheStruct:             cacheStruct,
			OriginatorIndex:         request.OriginatorIndex,
			SuccessfulFound:         successfulFound,
			FailedRequestsThreshold: failedRequestThreshold,
			FailedRequestsAccess:    failedRequestAccess,
			TimeStep:                int32(curTimeStep),
		}

		//update.StateListAndFlush(newState, stateList, curTimeStep)

		stateList = append(stateList, *newState)
		//newStateChan <- true
	}

}
