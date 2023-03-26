package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Cache(state *types.State, requestResult types.RequestResult) int32 {
	var cacheCounter int32
	if constants.IsCacheEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId

		if requestResult.Found {
			for _, nodeId := range route {
				node := state.Graph.GetNode(nodeId)
				node.CacheStruct.AddToCache(chunkId)
			}

			if requestResult.FoundByCaching {
				cacheCounter = atomic.AddInt32(&state.CacheHits, 1)
			}

		} else {
			cacheCounter = atomic.LoadInt32(&state.CacheHits)
		}

	}
	return cacheCounter
}
