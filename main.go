package main

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/workers"
	"go-incentive-simulation/model/state"
	"runtime"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	network := fmt.Sprintf("./data/nodes_data_%d_10000.txt", constants.GetBinSize())
	globalState := state.MakeInitialState(network)

	const iterations = 10_000_000
	numTotalGoRoutines := runtime.NumCPU()
	numRoutingGoroutines := constants.SetNumRoutingGoroutines(numTotalGoRoutines)
	//numLoops := iterations / numGoroutines

	wgMain := &sync.WaitGroup{}
	wgOutput := &sync.WaitGroup{}
	requestChan := make(chan types.Request, numRoutingGoroutines)
	outputChan := make(chan types.Output, 100000)
	routeChan := make(chan types.RouteData, 10000)
	stateChan := make(chan types.StateSubset, 10000)
	pauseChan := make(chan bool, numRoutingGoroutines)
	continueChan := make(chan bool, numRoutingGoroutines)

	if constants.IsWriteRoutesToFile() {
		wgOutput.Add(1)
		go workers.RouteFlushWorker(routeChan, wgOutput)
	}
	if constants.IsWriteStatesToFile() {
		wgOutput.Add(1)
		go workers.StateFlushWorker(stateChan, wgOutput)
	}

	go workers.RequestWorker(pauseChan, continueChan, requestChan, &globalState, wgMain, iterations)
	wgMain.Add(1)

	go workers.OutputWorker(outputChan, wgOutput)
	wgOutput.Add(1)

	for i := 0; i < numRoutingGoroutines; i++ {
		wgMain.Add(1)
		go workers.RoutingWorker(pauseChan, continueChan, requestChan, outputChan, routeChan, stateChan, &globalState, wgMain)
	}

	wgMain.Wait()
	close(outputChan)
	close(stateChan)
	close(routeChan)
	wgOutput.Wait()

	fmt.Println("")
	fmt.Println("end of main: ")
	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
	fmt.Println("Number of Iterations: ", iterations)
	fmt.Println("Number of Total Goroutines: ", numTotalGoRoutines)
	fmt.Println("Number of Routing Goroutines: ", numRoutingGoroutines)
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
	// 	fmt.Printf("RequestResult: %v\n", route.GetWaypoints())
	// 	fmt.Printf("Length: %d\n", route.GetLength())
	// 	count++
	// }

}

func PrintState(state types.State) {
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound)
	fmt.Println("FailedRequestsThreshold: ", state.FailedRequestsThreshold)
	fmt.Println("FailedRequestsAccess: ", state.FailedRequestsAccess)
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	fmt.Println("CacheHits:", state.CacheHits)
	fmt.Println("UniqueRetryCounter: ", state.UniqueRetryCounter)
	fmt.Println("UniqueWaitingCounter: ", state.UniqueWaitingCounter)
	//fmt.Println("PendingMap: ", state.PendingStruct.PendingMap, state.PendingStruct.Counter)
	//fmt.Println("RerouteMap: ", state.RerouteStruct.RerouteMap)
	//fmt.Println("RouteLists: ", state.RouteLists)
}
