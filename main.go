package main

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/policy"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/workers"
	"go-incentive-simulation/model/state"
	"sync"
	"time"
)

func MakePolicyOutput(state *types.State, index int) types.Policy {
	//fmt.Println("start of make initial policy")

	//found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(&state)
	found, route, thresholdFailed, accessFailed, paymentsList := policy.SendRequest(state, index)

	p := types.Policy{
		Found:                found,
		Route:                route,
		ThresholdFailedLists: thresholdFailed,
		AccessFailed:         accessFailed,
		PaymentList:          paymentsList,
	}
	return p
}

func main() {
	start := time.Now()
	globalState := state.MakeInitialState("./data/nodes_data_16_10000.txt")

	const iterations = 10000000
	numGoroutines := constants.Constants.GetNumGoroutines()
	numLoops := iterations / numGoroutines

	stateList := make([]types.StateSubset, 1)
	stateList[0] = types.StateSubset{
		OriginatorIndex:         globalState.OriginatorIndex,
		PendingMap:              globalState.PendingStruct.PendingMap,
		RerouteMap:              globalState.RerouteStruct.RerouteMap,
		CacheStruct:             globalState.CacheStruct,
		SuccessfulFound:         globalState.SuccessfulFound,
		FailedRequestsThreshold: globalState.FailedRequestsThreshold,
		FailedRequestsAccess:    globalState.FailedRequestsAccess,
		TimeStep:                globalState.TimeStep,
	}

	wg := &sync.WaitGroup{}
	//policyChan := make(chan Policy, numGoroutines)
	newStateChan := make(chan bool, numGoroutines)
	requestChan := make(chan types.Request, numGoroutines)
	routeChan := make(chan types.Route, numGoroutines)
	stateChan := make(chan []byte, 100000)

	if constants.Constants.IsWriteRoutesToFile() {
		wg.Add(1)
		go workers.RouteFlushWorker(routeChan, &globalState, wg, iterations)
	}
	if constants.Constants.IsWriteStatesToFile() {
		wg.Add(1)
		go workers.StateFlushWorker(stateChan, &globalState, stateList, wg, iterations)
	}

	go workers.RequestWorker(newStateChan, requestChan, &globalState, iterations)
	//newStateChan <- true

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go workers.RoutingWorker(requestChan, routeChan, stateChan, newStateChan, &globalState, stateList, wg, numLoops)
	}
	wg.Wait()
	close(routeChan)
	close(stateChan)
	//for j := 0; j < numGoroutines/4; j++ {
	//	//wg.Add(1)
	//	go UpdateWorker(newStateChan, policyChan, &globalState, stateList, wg, iterations)
	//}
	//newStateChan <- true
	//wg.Wait()

	//for j := 0; j < numGoroutines; j++ {
	//	wg.Add(1)
	//	go func(index int) {
	//		curState := &globalState
	//		for i := 0; i < numLoops; i++ {
	//			policyChan <- MakePolicyOutput(curState, index)
	//			// waiting for new state from UpdateWorker
	//			curState = <-stateChan
	//		}
	//		wg.Done()
	//	}(j)
	//}
	//wg.Wait()
	fmt.Println("end of main: ")
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
	fmt.Println("Number of iterations: ", iterations)
	fmt.Println("Number of Goroutines: ", numGoroutines)
	// allReq, thresholdFails, requestsToBucketZero, rejectedBucketZero, rejectedFirstHop := ReadRoutes("routes.json")
	// fmt.Println("allReq: ", allReq)
	// fmt.Println("thresholdFails: ", thresholdFails)
	// fmt.Println("requestsToBucketZero: ", requestsToBucketZero)
	// fmt.Println("rejectedBucketZero: ", rejectedBucketZero)
	// fmt.Println("rejectedFirstHop: ", rejectedFirstHop)
	PrintState(globalState)
}

func PrintState(state types.State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("CacheHits:", state.CacheStruct.CacheHits)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	fmt.Println("PendingMap: ", state.PendingStruct.PendingMap)
	fmt.Println("RerouteMap: ", state.RerouteStruct.RerouteMap)
	//fmt.Println("RouteLists: ", state.RouteLists)
}
