package workers

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"sync"
)

func RoutingWorker(requestChan chan Request, policyChan chan Policy, globalState *State, wg *sync.WaitGroup, numLoops int) {

	defer wg.Done()
	var request Request
	for i := 0; i < numLoops; i++ {
		request = <-requestChan

		found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, globalState.Graph, globalState.RerouteMap, globalState.CacheStruct)

		policy := Policy{
			Found:                found,
			Route:                route,
			ThresholdFailedLists: thresholdFailed,
			AccessFailed:         accessFailed,
			PaymentList:          paymentsList,
		}

		policyChan <- policy
	}

}
