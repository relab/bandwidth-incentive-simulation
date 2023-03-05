package workers

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
	. "go-incentive-simulation/model/parts/utils"
	"sync"
	"sync/atomic"
)

func RoutingWorker(requestChan chan Request, newStateChan chan bool, globalState *State, stateArray []State, wg *sync.WaitGroup, numLoops int) {

	defer wg.Done()
	var request Request
	for i := 0; i < numLoops; i++ {
		request = <-requestChan

		found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, globalState.Graph, globalState.RerouteStruct, globalState.CacheStruct)

		policyOutput := Policy{
			Found:                found,
			Route:                route,
			ThresholdFailedLists: thresholdFailed,
			AccessFailed:         accessFailed,
			PaymentList:          paymentsList,
		}

		//policyChan <- policy

		//curTimeStep := UpdateTimestep(globalState)
		curTimeStep := atomic.LoadInt32(&globalState.TimeStep)
		//fmt.Println(curTimeStep)
		//fmt.Println(" ")
		UpdateNetwork(globalState, policyOutput, int(curTimeStep))
		UpdatePendingMap(globalState, policyOutput)
		UpdateRerouteMap(globalState, policyOutput)
		UpdateCacheMap(globalState, policyOutput)
		//UpdateOriginatorIndex(globalState)
		UpdateSuccessfulFound(globalState, policyOutput)
		UpdateFailedRequestsThreshold(globalState, policyOutput)
		UpdateFailedRequestsAccess(globalState, policyOutput)
		//UpdateRouteListAndFlush(globalState, policyOutput, curTimeStep)

		newState := State{
			Graph:                   globalState.Graph,
			Originators:             globalState.Originators,
			NodesId:                 globalState.NodesId,
			RouteLists:              globalState.RouteLists,
			PendingStruct:           globalState.PendingStruct,
			RerouteStruct:           globalState.RerouteStruct,
			CacheStruct:             globalState.CacheStruct,
			OriginatorIndex:         globalState.OriginatorIndex,
			SuccessfulFound:         globalState.SuccessfulFound,
			FailedRequestsThreshold: globalState.FailedRequestsThreshold,
			FailedRequestsAccess:    globalState.FailedRequestsAccess,
			TimeStep:                globalState.TimeStep,
		}

		stateArray[curTimeStep] = newState

		//newStateChan <- true
	}

}
