package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
)

func CacheMap(state *types.State, requestResult types.RequestResult) types.CacheStruct {
	if constants.IsCacheEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId

		if requestResult.FoundByCaching {
			state.CacheStruct.CacheHits++
		}

		if requestResult.Found {
			for _, nodeId := range route {
				//state.CacheStruct.AddToCache(nodeId, chunkId)
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
					node.CacheMap = map[types.ChunkId]int{chunkId: 1}
				}
				node.Mutex.Unlock()
			}
		}
	}
	//state.CacheStruct = cacheStruct
	return state.CacheStruct
}
