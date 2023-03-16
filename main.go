package main

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/workers"
	"go-incentive-simulation/model/state"
	"go-incentive-simulation/results"
	"strings"
	"sync"
	"time"
)

//func MakePolicyOutput(state *types.State, index int) types.RequestResult {
//	//fmt.Println("start of make initial policy")
//
//	//found, route, thresholdFailed, accessFailed, paymentsList := SendRequest(&state)
//	found, route, thresholdFailed, accessFailed, paymentsList := policy.SendRequest(state, index)
//

//	p := types.RequestResult{
//		Found:                found,
//		Route:                route,
//		ThresholdFailedLists: thresholdFailed,
//		AccessFailed:         accessFailed,
//		PaymentList:          paymentsList,
//	}
//	return p
//}

func main() {
	var hop1 []float64
	var hop2 []float64
	var hop3 []float64
	for i := 5; i < 21; i++ {
		constants.SetProximityOrder(i)
		start := time.Now()
		network := fmt.Sprintf("./data/nodes_data_%d_10000.txt", constants.GetBinSize())
		globalState := state.MakeInitialState(network)

		const iterations = 250_000
		numGoroutines := constants.GetNumGoroutines()
		//numLoops := iterations / numGoroutines

		wgMain := &sync.WaitGroup{}
		wgFlush := &sync.WaitGroup{}
		requestChan := make(chan types.Request, numGoroutines)
		outputChan := make(chan types.Output, numGoroutines)
		routeChan := make(chan types.RouteData, numGoroutines)
		stateChan := make(chan types.StateSubset, 10000)

		if constants.IsWriteRoutesToFile() {
			wgFlush.Add(1)
			go workers.RouteFlushWorker(routeChan, wgFlush)
		}
		if constants.IsWriteStatesToFile() {
			wgFlush.Add(1)
			go workers.StateFlushWorker(stateChan, wgFlush)
		}

		go workers.RequestWorker(requestChan, &globalState, wgMain, iterations)
		wgMain.Add(1)

		go workers.OutputWorker(outputChan)

		for i := 0; i < numGoroutines; i++ {
			wgMain.Add(1)
			go workers.RoutingWorker(requestChan, outputChan, routeChan, stateChan, &globalState, wgMain)
		}

		wgMain.Wait()
		close(stateChan)
		close(routeChan)
		wgFlush.Wait()

		fmt.Println("")
		fmt.Println("end of main: ")
		elapsed := time.Since(start)
		fmt.Println("Time taken:", elapsed)
		fmt.Println("Number of Iterations: ", iterations)
		fmt.Println("Number of Goroutines: ", numGoroutines)
		// allReq, thresholdFails, requestsToBucketZero, rejectedBucketZero, rejectedFirstHop := ReadRoutes("routes.json")
		// fmt.Println("allReq: ", allReq)
		// fmt.Println("thresholdFails: ", thresholdFails)
		// fmt.Println("requestsToBucketZero: ", requestsToBucketZero)
		// fmt.Println("rejectedBucketZero: ", rejectedBucketZero)
		// fmt.Println("rejectedFirstHop: ", rejectedFirstHop)
		PrintState(globalState)

		// TODO: Add this to another function in another file?
		// buf, err := ioutil.ReadFile("states.bin")
		// if err != nil {
		// 	panic(err)
		// }
		// stateSubsets := &protoGenerated.StateSubsets{}
		// err = proto.Unmarshal(buf, stateSubsets)
		// if err != nil {
		// 	panic(err)
		// }
		// // Access the subset field
		// count := 0
		// for _, subset := range stateSubsets.Subset {
		// 	count++
		// 	if count > 10 {
		// 		break
		// 	}
		// 	fmt.Printf("OriginatorIndex: %d\n", subset.OriginatorIndex)
		// 	fmt.Printf("PendingMap: %d\n", subset.PendingMap)
		// 	fmt.Printf("RerouteMap: %d\n", subset.RerouteMap)
		// 	fmt.Printf("CacheStruct: %d\n", subset.CacheStruct)
		// 	fmt.Printf("SuccessfulFound: %d\n", subset.SuccessfulFound)
		// 	fmt.Printf("FailedRequestsThreshold: %d\n", subset.FailedRequestsThreshold)
		// 	fmt.Printf("FailedRequestsAccess: %d\n", subset.FailedRequestsAccess)
		// 	fmt.Printf("TimeStep: %d\n", subset.TimeStep)
		// }
		// // read the binary protobuf message from the file
		// buf, err := ioutil.ReadFile("routes.bin")
		// if err != nil {
		// 	panic(err)
		// }

		// // unmarshal the binary protobuf message into a RouteData struct
		// routeData := &protoGenerated.RouteData{}
		// err = proto.Unmarshal(buf, routeData)
		// if err != nil {
		// 	panic(err)
		// }

		// // print the RouteData struct
		// fmt.Printf("TimeStep: %d\n", routeData.GetTimeStep())
		// count := 0
		// routedata := routeData.GetRoutes()
		// fmt.Println("length", len(routedata))
		// for _, route := range routeData.GetRoutes() {
		// 	if count == 10 {
		// 		break
		// 	}
		// 	fmt.Printf("Route: %v\n", route.GetWaypoints())
		// 	fmt.Printf("Length: %d\n", route.GetLength())
		// 	count++
		// }
		res, _ := results.AvgFractionOfTotalRewards("./results/output.txt")
		hop1 = append(hop1, res.Hop1)
		hop2 = append(hop2, res.Hop2)
		hop3 = append(hop3, res.Hop3)
	}
	strhop1 := make([]string, len(hop1))
	strhop2 := make([]string, len(hop2))
	strhop3 := make([]string, len(hop3))
	for i, v := range hop1 {
		strhop1[i] = fmt.Sprintf("%v", v)
	}
	for i, v := range hop2 {
		strhop2[i] = fmt.Sprintf("%v", v)
	}
	for i, v := range hop3 {
		strhop3[i] = fmt.Sprintf("%v", v)
	}
	fmt.Println(strings.Join(strhop1, ", "))
	fmt.Println(strings.Join(strhop2, ", "))
	fmt.Println(strings.Join(strhop3, ", "))
}

func PrintState(state types.State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("CacheHits:", state.CacheStruct.CacheHits)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	//fmt.Println("PendingMap: ", state.PendingStruct.PendingMap, state.PendingStruct.Counter)
	//fmt.Println("RerouteMap: ", state.RerouteStruct.RerouteMap)
	//fmt.Println("RouteLists: ", state.RouteLists)
}
