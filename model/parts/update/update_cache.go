package update

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Cache(state *types.State, requestResult types.RequestResult) int64 {
	var cacheHits int64
	if constants.IsCacheEnabled() {
		route := requestResult.Route
		chunkId := requestResult.ChunkId

		if requestResult.Found {
			for _, nodeId := range route {
				node := state.Graph.GetNode(nodeId)
				node.CacheStruct.AddToCache(chunkId)
			}

			if requestResult.FoundByCaching {
				cacheHits = atomic.AddInt64(&state.CacheHits, 1)
			}

		} else {
			cacheHits = atomic.LoadInt64(&state.CacheHits)
		}

	}
	return cacheHits
}
