package main

import (
	"fmt"
	. "go-incentive-simulation/model/parts/policy"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
	. "go-incentive-simulation/model/state"
	"sync"
	"time"
)

func MakePolicyOutput(state State) Policy {
	//fmt.Println("start of make initial policy")
	found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(&state)
	policy := Policy{
		Found:                found,
		Route:                route,
		ThresholdFailedLists: thresholdFailed,
		OriginatorIndex:      state.OriginatorIndex,
		AccessFailed:         accessFailed,
		PaymentList:          paymentsList,
	}
	return policy
}

func main() {
	start := time.Now()
	state := MakeInitialState("./data/nodes_data_16_10000.txt")
	//stateArray := []State{state}
	iterations := 10000000
	const numRotines = 10000
	// for i := 0; i < iterations; i++ {
	//     policyOutput := MakePolicyOutput(state)
	//     state = UpdatePendingMap(state, policyOutput)
	//     state = UpdateRerouteMap(state, policyOutput)
	//     state = UpdateCacheMap(state, policyOutput)
	//     state = UpdateOriginatorIndex(state, policyOutput)
	//     state = UpdateSuccessfulFound(state, policyOutput)
	//     state = UpdateFailedRequestsThreshold(state, policyOutput)
	//     state = UpdateFailedRequestsAccess(state, policyOutput)
	//     state = UpdateRouteListAndFlush(state, policyOutput)
	//     state = UpdateNetwork(state, policyOutput)
	// }
	for i := 0; i < iterations/numRotines; i++ {
		//fmt.Println("Start of lop ", time.Since(start))
		var policyOutputs [numRotines]Policy
		var wg sync.WaitGroup
		for j := 0; j < numRotines; j++ {
			wg.Add(1)
			go func(index int) {
				defer wg.Done()
				policyOutputs[index] = MakePolicyOutput(state)
			}(j)
		}
		wg.Wait()
		//fmt.Println("end of lop ", time.Since(start))
		for j := 0; j < numRotines; j++ {
			state = UpdatePendingMap(state, policyOutputs[j])
			state = UpdateRerouteMap(state, policyOutputs[j])
			state = UpdateCacheMap(state, policyOutputs[j])
			state = UpdateOriginatorIndex(state, policyOutputs[j])
			state = UpdateSuccessfulFound(state, policyOutputs[j])
			state = UpdateFailedRequestsThreshold(state, policyOutputs[j])
			state = UpdateFailedRequestsAccess(state, policyOutputs[j])
			state = UpdateRouteListAndFlush(state, policyOutputs[j])
			state = UpdateNetwork(state, policyOutputs[j])
		}
	}
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
	// allReq, thresholdFails, requestsToBucketZero, rejectedBucketZero, rejectedFirstHop := ReadRoutes("routes.json")
	// fmt.Println("allReq: ", allReq)
	// fmt.Println("thresholdFails: ", thresholdFails)
	// fmt.Println("requestsToBucketZero: ", requestsToBucketZero)
	// fmt.Println("rejectedBucketZero: ", rejectedBucketZero)
	// fmt.Println("rejectedFirstHop: ", rejectedFirstHop)
	fmt.Print("end of main: ")
	PrintState(state)
}

func PrintState(state State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("CacheHits:", state.CacheStruct.CacheHits)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	fmt.Println("PendingMap: ", state.PendingMap)
	fmt.Println("RerouteMap: ", state.RerouteMap)
	//fmt.Println("RouteLists: ", state.RouteLists)
	//fmt.Println("CacheMap: ", state.CacheStruct.CacheMap)
}
