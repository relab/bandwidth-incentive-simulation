package types

import (
	"sync"
)

type RouteStruct struct {
	Reroute   []NodeId
	ChunkId   ChunkId
	LastEpoch int
}

type RerouteMap map[NodeId]RouteStruct

type RerouteStruct struct {
	RerouteMap           RerouteMap
	RerouteMutex         *sync.Mutex
	UniqueRerouteCounter int
}

func (r *RerouteStruct) GetRerouteMap(originator NodeId) RouteStruct {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	reroute, ok := r.RerouteMap[originator]
	if ok {
		return reroute
	}
	return RouteStruct{}
}

func (r *RerouteStruct) DeleteReroute(originator NodeId) {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	delete(r.RerouteMap, originator)
}

func (r *RerouteStruct) AddNewReroute(originator NodeId, nodeId NodeId, chunkId ChunkId, curEpoch int) bool {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	_, ok := r.RerouteMap[originator]
	if !ok {
		r.UniqueRerouteCounter++
		r.RerouteMap[originator] = RouteStruct{
			Reroute:   []NodeId{nodeId},
			ChunkId:   chunkId,
			LastEpoch: curEpoch,
		}
		return true
	}
	return false
}

func (r *RerouteStruct) AddNodeToReroute(originator NodeId, nodeId NodeId) bool {
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

func (r *RerouteStruct) UpdateEpoch(originator NodeId, curEpoch int) int {
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
