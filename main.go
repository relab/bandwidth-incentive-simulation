package main

import (
	"fmt"
	. "go-incentive-simulation/model/parts/policy"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
	. "go-incentive-simulation/model/parts/utils"
	//. "go-incentive-simulation/model/variables"
)

func MakeInitialState() State {
	// Initialize the state
	fmt.Println("start of make initial state")
	path := "nodes_data_8_10000.txt"
	network := Network{}
	network.Load(path)
	graph, err := CreateGraphNetwork(&network)
	if err != nil {
		fmt.Println("create graph network returned an error: ", err)
	}
	initialState := State{
		Graph:                   graph,
		Originators:             CreateDownloadersList(&network),
		NodesId:                 CreateNodesList(&network),
		RouteLists:              []Route{},
		PendingMap:              make(PendingMap, 0),
		RerouteMap:              make(RerouteMap, 0),
		CacheListMap:            make(CacheListMap, 0),
		OriginatorIndex:         0,
		SuccessfulFound:         0,
		FailedRequestsThreshold: 0,
		FailedRequestsAccess:    0,
		TimeStep:                0,
	}
	return initialState
}

func MakePolicyOutput(state State) Policy {
	fmt.Println("start of make initial policy")
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
	state := MakeInitialState()
	stateArray := []State{state}
	constant := 100
	for i := 0; i < constant; i++ {
		policyOutput := MakePolicyOutput(state)
		state = UpdatePendingMap(state, policyOutput)
		state = UpdateRerouteMap(state, policyOutput)
		state = UpdateOriginatorIndex(state, policyOutput)
		state = UpdateSuccessfulFound(state, policyOutput)
		state = UpdateFailedRequestsThreshold(state, policyOutput)
		state = UpdateRouteListAndFlush(state, policyOutput)
		state = UpdateNetwork(state, policyOutput)
		stateArray = append(stateArray, state)
		PrintState(state)
	}
	fmt.Print("end of main")
}

func PrintState(state State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	fmt.Println("PendingMap: ", state.PendingMap)
	fmt.Println("RerouteMap: ", state.RerouteMap)
	fmt.Println("RouteLists: ", state.RouteLists)
	fmt.Println("CacheListMap: ", state.CacheListMap)
}
