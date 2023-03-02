package main

import (
	"fmt"
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/policy"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/workers"
	. "go-incentive-simulation/model/state"
	"sync"
	"time"
)

func MakePolicyOutput(state *State, index int) Policy {
	//fmt.Println("start of make initial policy")

	//found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(&state)
	found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(state, index)

	policy := Policy{
		Found:                found,
		Route:                route,
		ThresholdFailedLists: thresholdFailed,
		AccessFailed:         accessFailed,
		PaymentList:          paymentsList,
	}
	return policy
}

func main() {
	start := time.Now()
	globalState := MakeInitialState("./data/nodes_data_16_10000.txt")

	const iterations = 10000000
	numGoroutines := Constants.GetNumGoroutines()

	numLoops := iterations / numGoroutines
	stateArray := make([]State, iterations+1)
	stateArray[0] = globalState

	wg := &sync.WaitGroup{}
	policyChan := make(chan Policy, numGoroutines)
	newStateChan := make(chan bool, numGoroutines)
	requestChan := make(chan Request, iterations+1)

	go RequestWorker(newStateChan, requestChan, &globalState, iterations)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go RoutingWorker(requestChan, policyChan, &globalState, wg, numLoops)
	}

	for j := 0; j < numGoroutines/4; j++ {
		//wg.Add(1)
		go UpdateWorker(newStateChan, policyChan, &globalState, stateArray, wg, iterations)
	}
	//newStateChan <- true
	wg.Wait()

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

func PrintState(state State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("CacheHits:", state.CacheStruct.CacheHits)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	//fmt.Println("PendingMap: ", state.PendingMap)
	//fmt.Println("RerouteMap: ", state.RerouteMap)
	//fmt.Println("RouteLists: ", state.RouteLists)
	//fmt.Println("CacheMapArray: ", state.CacheStruct.CacheMapArray)
}
