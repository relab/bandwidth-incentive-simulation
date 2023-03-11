package state

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
	"math/rand"
	"sync"
)

func MakeInitialState(path string) types.State {
	// Initialize the state
	fmt.Println("start of make initial state")
	rand.Seed(constants.GetRandomSeed())
	network := types.Network{}
	network.Load(path)
	graph, err := utils.CreateGraphNetwork(&network)
	if err != nil {
		fmt.Println("create graph network returned an error: ", err)
	}
	pendingStruct := types.PendingStruct{PendingMap: make(types.PendingMap, 0), PendingMutex: &sync.Mutex{}}
	rerouteStruct := types.RerouteStruct{RerouteMap: make(types.RerouteMap, 0), RerouteMutex: &sync.Mutex{}}
	cacheStruct := types.CacheStruct{CacheHits: 0, CacheMap: make(types.CacheMap), CacheMutex: &sync.Mutex{}}

	initialState := types.State{
		Graph:                   graph,
		Originators:             utils.CreateDownloadersList(graph),
		NodesId:                 utils.CreateNodesList(graph),
		RouteLists:              make([]types.Route, 10000),
		PendingStruct:           pendingStruct,
		RerouteStruct:           rerouteStruct,
		CacheStruct:             cacheStruct,
		OriginatorIndex:         0,
		SuccessfulFound:         0,
		FailedRequestsThreshold: 0,
		FailedRequestsAccess:    0,
		TimeStep:                0,
	}
	return initialState
}
