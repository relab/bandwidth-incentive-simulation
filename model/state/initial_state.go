package state

import (
	"fmt"
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"math/rand"
	"sync"
)

func MakeInitialState(path string) State {
	// Initialize the state
	fmt.Println("start of make initial state")
	rand.Seed(Constants.GetRandomSeed())
	network := Network{}
	network.Load(path)
	graph, err := CreateGraphNetwork(&network)
	if err != nil {
		fmt.Println("create graph network returned an error: ", err)
	}
	pendingStruct := PendingStruct{PendingMap: make(PendingMap, 0), PendingMutex: &sync.Mutex{}}
	rerouteStruct := RerouteStruct{RerouteMap: make(RerouteMap, 0), RerouteMutex: &sync.Mutex{}}
	cacheStruct := CacheStruct{CacheHits: 0, CacheMap: make(CacheMap), CacheMutex: &sync.Mutex{}}

	initialState := State{
		Graph:                   graph,
		Originators:             CreateDownloadersList(graph),
		NodesId:                 CreateNodesList(graph),
		RouteLists:              make([]Route, 10000),
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
