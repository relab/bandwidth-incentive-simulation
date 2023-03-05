package types

import "sync"

type RerouteMap map[int][]int

type RerouteStruct struct {
	RerouteMap   RerouteMap
	RerouteMutex *sync.RWMutex
}

func (r *RerouteStruct) GetRerouteMap(originator int) []int {
	r.RerouteMutex.RLock()
	defer r.RerouteMutex.RUnlock()
	reroute, ok := r.RerouteMap[originator]
	if ok {
		return reroute
	}
	return nil
}

func (r *RerouteStruct) DeleteReroute(originator int) {
	r.RerouteMutex.Lock()
	defer r.RerouteMutex.Unlock()
	delete(r.RerouteMap, originator)
}
