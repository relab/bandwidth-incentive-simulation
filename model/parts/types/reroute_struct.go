package types

import (
	"go-incentive-simulation/model/constants"
	"sync"
)

type RouteStruct struct {
	Reroute Route
	ChunkId int
	Epoch   int
}

type RerouteMap map[int]RouteStruct

type RerouteStruct struct {
	RerouteMap          RerouteMap
	RerouteMutex        *sync.Mutex
	TotalRerouteCounter int
}

func (r *RerouteStruct) GetRerouteMap(originator int) RouteStruct {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	reroute, ok := r.RerouteMap[originator]
	if ok {
		return reroute
	}
	return RouteStruct{}
}

func (r *RerouteStruct) DeleteReroute(originator int) {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	delete(r.RerouteMap, originator)
}

func (r *RerouteStruct) AddNewReroute(originator int, nodeId int, chunkId int) bool {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	_, ok := r.RerouteMap[originator]
	if !ok {
		r.TotalRerouteCounter++
		r.RerouteMap[originator] = RouteStruct{
			Reroute: []int{nodeId},
			ChunkId: chunkId,
			Epoch:   constants.GetEpoch(),
		}
		return true
	}
	return false
}

func (r *RerouteStruct) AddNodeToReroute(originator int, nodeId int) bool {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	routeStruct, ok := r.RerouteMap[originator]
	if ok {
		r.TotalRerouteCounter++
		routeStruct.Reroute = append(routeStruct.Reroute, nodeId)
		r.RerouteMap[originator] = routeStruct
		return true
	}
	return false
}

func (r *RerouteStruct) UpdateEpoch(originator int) int {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	routeStruct, ok := r.RerouteMap[originator]
	if ok {
		routeStruct.Epoch++
		r.RerouteMap[originator] = routeStruct
		return routeStruct.Epoch
	}
	return -1

}
