package types

import (
	"sync"
)

type Reroute struct {
	CheckedNodes []NodeId
	ChunkId      ChunkId
	LastEpoch    int
}

//type RerouteMap map[NodeId]Reroute

type RerouteStruct struct {
	Reroute      Reroute
	RerouteMutex *sync.Mutex
}

//type RerouteStruct struct {
//	RerouteMap           RerouteMap
//	RerouteMutex         *sync.Mutex
//	UniqueRerouteCounter int
//}

func (r *RerouteStruct) GetReroute() Reroute {
	return r.Reroute
}

//func (r *RerouteStruct) GetRerouteMap() Reroute {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	reroute, ok := r.RerouteMap[originator]
//	if ok {
//		return reroute
//	}
//	return Reroute{}
//}

//func (r *RerouteStruct) DeleteReroute(originator NodeId) {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	delete(r.RerouteMap, originator)
//}

func (r *RerouteStruct) AddNewReroute(node *Node, nodeId NodeId, chunkId ChunkId, curEpoch int) {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()

	newReroute := Reroute{
		CheckedNodes: []NodeId{nodeId},
		ChunkId:      chunkId,
		LastEpoch:    curEpoch,
	}
	node.RerouteStruct.Reroute = newReroute
}

//func (r *RerouteStruct) AddNewReroute(originator NodeId, nodeId NodeId, chunkId ChunkId, curEpoch int) bool {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	_, ok := r.RerouteMap[originator]
//	if !ok {
//		r.RerouteMap[originator] = Reroute{
//			CheckedNodes: []NodeId{nodeId},
//			ChunkId:      chunkId,
//			LastEpoch:    curEpoch,
//		}
//		return true
//	}
//	return false
//}

func (r *RerouteStruct) AddNodeToCheckedNodes(node *Node, nodeId NodeId) {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	newCheckedNodes := append(r.Reroute.CheckedNodes, nodeId)

	node.RerouteStruct.Reroute.CheckedNodes = newCheckedNodes
}

//func (r *RerouteStruct) AddNodeToReroute(originator NodeId, nodeId NodeId) bool {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	routeStruct, ok := r.RerouteMap[originator]
//	if ok {
//		routeStruct.CheckedNodes = append(routeStruct.CheckedNodes, nodeId)
//		r.RerouteMap[originator] = routeStruct
//		return true
//	}
//	return false
//}

//func (r *RerouteStruct) UpdateEpoch(originator NodeId, curEpoch int) int {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	routeStruct, ok := r.RerouteMap[originator]
//	if ok {
//		routeStruct.LastEpoch = curEpoch
//		r.RerouteMap[originator] = routeStruct
//		return routeStruct.LastEpoch
//	}
//	return -1
//
//}
