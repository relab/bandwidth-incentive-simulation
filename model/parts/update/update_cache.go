package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/general"
	"go-incentive-simulation/model/parts/types"
)

func CacheMap(state *types.State, policyInput types.RequestResult) types.CacheStruct {
	chunkId := 0

	if constants.Constants.IsCacheEnabled() {
		route := policyInput.Route
		if general.Contains(route, -3) {
			// -3 means found by caching
			state.CacheStruct.CacheHits++
			chunkId = route[len(route)-2]
		} else {
			chunkId = route[len(route)-1]
		}
		if !general.Contains(route, -1) && !general.Contains(route, -2) {
			if general.Contains(route, -3) {
				for i := 0; i < len(route)-3; i++ {
					nodeId := route[i]
					state.CacheStruct.AddToCache(nodeId, chunkId)
					node := state.Graph.GetNode(nodeId)
					node.Mutex.Lock()
					cacheMap := node.CacheMap
					if cacheMap != nil {
						if _, ok := cacheMap[chunkId]; ok {
							cacheMap[chunkId]++
						} else {
							cacheMap[chunkId] = 1
						}
					} else {
						node.CacheMap = map[int]int{node.Id: 1}
					}
					node.Mutex.Unlock()
				}
			} else {
				for i := 0; i < len(route)-2; i++ {
					nodeId := route[i]
					state.CacheStruct.AddToCache(nodeId, chunkId)
					node := state.Graph.GetNode(nodeId)
					node.Mutex.Lock()
					cacheMap := node.CacheMap
					if cacheMap != nil {
						if _, ok := cacheMap[chunkId]; ok {
							cacheMap[chunkId]++
						} else {
							cacheMap[chunkId] = 1
						}
					} else {
						node.CacheMap = map[int]int{node.Id: 1}
					}
					node.Mutex.Unlock()
				}
			}
		}
	}
	//state.CacheStruct = cacheStruct
	return state.CacheStruct
}
