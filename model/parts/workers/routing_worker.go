package workers

import (
	"encoding/json"
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(requestChan chan types.Request, routeChan chan types.Route, stateChan chan []byte, newStateChan chan bool, globalState *types.State, stateList []types.StateSubset, wg *sync.WaitGroup, numLoops int) {
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
		update.Graph(globalState, policyOutput, curTimeStep)

		pendingStruct := update.PendingMap(globalState, policyOutput)
		rerouteStruct := update.RerouteMap(globalState, policyOutput)
		cacheStruct := update.CacheMap(globalState, policyOutput)
		//originatorIndex := UpdateOriginatorIndex(globalState)
		successfulFound := update.SuccessfulFound(globalState, policyOutput)
		failedRequestThreshold := update.FailedRequestsThreshold(globalState, policyOutput)
		failedRequestAccess := update.FailedRequestsAccess(globalState, policyOutput)

		//routeLists := update.RouteListAndFlush(globalState, policyOutput, curTimeStep)
		if constants.Constants.IsWriteRoutesToFile() {
			if curTimeStep%1000000 == 0 {
				fmt.Println("routeChan: ", len(routeChan))
			}
			routeChan <- policyOutput.Route
		}

		//newState := types.State{
		//	Graph:       graph,
		//	Originators: globalState.Originators,
		//	NodesId:     globalState.NodesId,
		//	//RouteLists:              routeLists,
		//	PendingStruct:           pendingStruct,
		//	RerouteStruct:           rerouteStruct,
		//	CacheStruct:             cacheStruct,
		//	OriginatorIndex:         request.OriginatorIndex,
		//	SuccessfulFound:         successfulFound,
		//	FailedRequestsThreshold: failedRequestThreshold,
		//	FailedRequestsAccess:    failedRequestAccess,
		//	TimeStep:                int32(curTimeStep),
		//}

		if constants.Constants.IsWriteStatesToFile() {
			newState := types.StateSubset{
				OriginatorIndex:         request.OriginatorIndex,
				PendingMap:              pendingStruct.PendingMap,
				RerouteMap:              rerouteStruct.RerouteMap,
				CacheStruct:             cacheStruct,
				SuccessfulFound:         successfulFound,
				FailedRequestsThreshold: failedRequestThreshold,
				FailedRequestsAccess:    failedRequestAccess,
				TimeStep:                int32(curTimeStep),
			}
			if curTimeStep%1000000 == 0 {
				fmt.Println("stateChan: ", len(stateChan))
			}
			encodedData, _ := json.Marshal(types.StateData{TimeStep: int(newState.TimeStep), State: newState})
			stateChan <- encodedData
		}

		//update.StateListAndFlush(newState, stateList, curTimeStep)
		//stateList = append(stateList, newState)
		//newStateChan <- true
	}
}
