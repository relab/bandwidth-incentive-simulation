package main

import (
	"flag"
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/output"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/workers"
	"go-incentive-simulation/model/state"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	graphId := flag.String("graphId", "", "an Id for the graph, e.g. even")
	count := flag.Int("count", 0, "generate count many networks with ids 0,1,...")
	maxPOs := flag.String("maxPOs", "", "min:max maxPO value")

	flag.Parse()

	min := -1
	max := 0
	var err error
	if len(strings.Split(*maxPOs, ":")) == 2 {
		min, err = strconv.Atoi(strings.Split(*maxPOs, ":")[0])
		if err != nil {
			fmt.Println("MaxPO must be informat min:max")
			return
		}
		if min < 0 {
			fmt.Println("MaxPO must be positive")
			return
		}
		max, err = strconv.Atoi(strings.Split(*maxPOs, ":")[1])
		if err != nil {
			fmt.Println("MaxPO must be informat min:max")
			return
		}

	}

	for maxPO := min; maxPO < max; maxPO++ {
		if *count == 0 {
			run(-1, *graphId, maxPO)
		}
		for i := *count; i > 0; i-- {
			run(*count-i, "", maxPO)
		}
	}

}

func run(iteration int, graphId string, maxPO int) {
	start := time.Now()
	if iteration == -1 {
		config.InitConfigs()
		if maxPO > -1 {
			config.SetMaxPO(maxPO)
		}
	}

	network := fmt.Sprintf("./data/nodes_data_%d_%d.txt", config.GetBinSize(), config.GetNetworkSize())
	if graphId != "" {
		network = fmt.Sprintf("./data/nodes_data_%d_%d_%v.txt", config.GetBinSize(), config.GetNetworkSize(), graphId)
	}
	if iteration > -1 {
		config.InitConfigsWithId(fmt.Sprint(iteration))
		if maxPO > -1 {
			config.SetMaxPO(maxPO)
		}
		network = fmt.Sprintf("./data/nodes_data_%d_%d_%d.txt", config.GetBinSize(), config.GetNetworkSize(), iteration)
	}

	fmt.Println("Running with network: ", network)

	globalState := state.MakeInitialState(network)

	iterations := config.GetIterations()
	numTotalGoRoutines := config.GetNumGoroutines()
	numRoutingGoroutines := config.GetNumRoutingGoroutines()

	wgMain := &sync.WaitGroup{}
	wgOutput := &sync.WaitGroup{}
	requestChan := make(chan types.Request, numRoutingGoroutines)
	outputChan := make(chan types.OutputStruct, 100000)
	routeChan := make(chan types.RouteData, 100000)
	stateChan := make(chan types.StateSubset, 100000)
	pauseChan := make(chan bool, numRoutingGoroutines)
	continueChan := make(chan bool, numRoutingGoroutines)

	if config.IsWriteRoutesToFile() {
		wgOutput.Add(1)
		go workers.RouteFlushWorker(routeChan, wgOutput)
	}
	if config.IsWriteStatesToFile() {
		wgOutput.Add(1)
		go workers.StateFlushWorker(stateChan, wgOutput)
	}

	go workers.RequestWorker(pauseChan, continueChan, requestChan, &globalState, wgMain)
	wgMain.Add(1)

	if config.IsOutputEnabled() {
		go output.Worker(outputChan, wgOutput)
		wgOutput.Add(1)
	}

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
	PrintState(globalState)

}

func PrintState(state types.State) {
	total := float64(state.SuccessfulFound + state.FailedRequestsThreshold + state.FailedRequestsAccess)
	fmt.Println("SuccessfulFound: ", state.SuccessfulFound, "-->", math.Round(float64(state.SuccessfulFound)/total*1000000)/10000, "%")
	fmt.Println("ThresholdFail: ", state.FailedRequestsThreshold, "-->", math.Round(float64(state.FailedRequestsThreshold)/total*1000000)/10000, "%")
	fmt.Println("AccessFail: ", state.FailedRequestsAccess, "-->", math.Round(float64(state.FailedRequestsAccess)/total*1000000)/10000, "%")
	fmt.Println("TimeStep: ", state.TimeStep)
	fmt.Println("OriginatorIndex: ", state.OriginatorIndex)
	fmt.Println("CacheHits:", state.CacheHits)
	fmt.Println("UniqueRetryCounter: ", state.UniqueRetryCounter)
	fmt.Println("UniqueWaitingCounter: ", state.UniqueWaitingCounter)
}
