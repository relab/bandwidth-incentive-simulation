package state

import (
	"fmt"
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"math/rand"
	"os"
)

func MakeInitialState(path string) State {
	// Initialize the state
	fmt.Println("start of make initial state")
	rand.Seed(Constants.GetRandomSeed())
	network := Network{}
	pwd, _ := os.Getwd()
	fmt.Println("current dir: ", pwd)
	network.Load(path)
	graph, err := CreateGraphNetwork(&network)
	if err != nil {
		fmt.Println("create graph network returned an error: ", err)
	}
	initialState := State{
		Graph:                   graph,
		Originators:             CreateDownloadersList(graph),
		NodesId:                 CreateNodesList(graph),
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
