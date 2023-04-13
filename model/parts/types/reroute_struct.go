package types

import (
	"sync"
)

type RouteStruct struct {
	Reroute   Route
	ChunkId   int
	LastEpoch int
}

type RerouteMap map[int]RouteStruct

type RerouteStruct struct {
	RerouteMap           RerouteMap
	RerouteMutex         *sync.Mutex
	UniqueRerouteCounter int
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

func (r *RerouteStruct) AddNewReroute(originator int, nodeId int, chunkId int, curEpoch int) bool {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	_, ok := r.RerouteMap[originator]
	if !ok {
		r.UniqueRerouteCounter++
		r.RerouteMap[originator] = RouteStruct{
			Reroute:   []int{nodeId},
			ChunkId:   chunkId,
			LastEpoch: curEpoch,
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
		r.UniqueRerouteCounter++
		routeStruct.Reroute = append(routeStruct.Reroute, nodeId)
		r.RerouteMap[originator] = routeStruct
		return true
	}
	return false
}

func (r *RerouteStruct) UpdateEpoch(originator int, curEpoch int) int {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	routeStruct, ok := r.RerouteMap[originator]
	if ok {
		routeStruct.LastEpoch = curEpoch
		r.RerouteMap[originator] = routeStruct
		return routeStruct.LastEpoch
	}
	return -1

}
