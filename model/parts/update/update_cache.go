package update

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"sync/atomic"
)

func Cache(state *types.State, requestResult types.RequestResult) int64 {
	var cacheHits int64 = 0
	if config.IsCacheEnabled() {
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
		}

		if cacheHits == 0 {
			cacheHits = atomic.LoadInt64(&state.CacheHits)
		}

	}
	return cacheHits
}
