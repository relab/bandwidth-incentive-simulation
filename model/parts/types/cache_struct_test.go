package types

import (
	"fmt"
	"go-incentive-simulation/config"
	"sync"
	"testing"
	"time"
)

const path = "../../../"

func TestAddToChche(t *testing.T) {
	cache := CacheStruct{
		Size:       3,
		CacheMap:   make(CacheMap),
		CacheMutex: &sync.Mutex{},
	}

	config.InitConfigWithPath(path)
	config.SetCacheModel(0)
	cache.EvictionPolicy = GetCachePolicy()

	network := &Network{}
	network.Bits = config.GetBits()
	node := network.node(NodeId(1))

	cache.AddToCache(ChunkId(1), node.Id)
	cache.AddToCache(ChunkId(200), node.Id)
	cache.AddToCache(ChunkId(3000), node.Id)

	if len(cache.CacheMap) != 3 {
		t.Errorf("Expected cache size to be 3, but got %d", len(cache.CacheMap))
	}

	config.SetCacheModel(1)
	cache.EvictionPolicy = GetCachePolicy()

	cache.AddToCache(ChunkId(400), node.Id)
	if len(cache.CacheMap) != 3 {
		t.Errorf("Expected cache size to be 3, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(3000)]; exists {
		t.Errorf("Expected ChunkId 3000 to be removed")
	}

	config.SetCacheModel(2)
	cache.EvictionPolicy = GetCachePolicy()

	now := time.Now()
	dummyCounter := 1
	for id, cacheData := range cache.CacheMap {
		dummyCounter++
		cacheData.LastTimeUsed = now.Add(-time.Hour * time.Duration(dummyCounter*int(id)))
		cache.CacheMap[id] = cacheData
		fmt.Println(id, cache.CacheMap[id].LastTimeUsed)
	}
	cache.AddToCache(ChunkId(4), node.Id)
	if len(cache.CacheMap) != 3 {
		t.Errorf("Expected cache size to be 3, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(400)]; exists {
		t.Errorf("Expected ChunkId 400 to be removed")
	}

	config.SetCacheModel(3)
	cache.EvictionPolicy = GetCachePolicy()

	for nodeId, cacheData := range cache.CacheMap {
		dummyCounter++
		cacheData.Frequency = dummyCounter * int(nodeId)
		cache.CacheMap[nodeId] = cacheData
	}
	cache.AddToCache(ChunkId(4000), node.Id)
	if len(cache.CacheMap) != 3 {
		t.Errorf("Expected cache size to be 3, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(4000)]; exists {
		t.Errorf("Expected ChunkId 4000 to be removed")
	}

	config.SetCacheModel(0)
	cache.EvictionPolicy = GetCachePolicy()

	cache.AddToCache(ChunkId(5), node.Id)
	if len(cache.CacheMap) != 4 {
		t.Errorf("Expected cache size to be 4 (Unlimited caching) after adding one more entry, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(5)]; !exists {
		t.Errorf("Expected ChunkId 5 to be added")
	}
}

func TestUpdateCacheMap_Proximity(t *testing.T) {
	config.InitConfigWithPath(path)
	config.SetCacheModel(1)

	cache := CacheStruct{
		Size:           3,
		CacheMap:       make(CacheMap),
		CacheMutex:     &sync.Mutex{},
		EvictionPolicy: GetCachePolicy(),
	}

	// 79 = 64 + 8 + 4 + 2 + 1,
	cache.AddToCache(ChunkId(80), NodeId(79)) // 80 = 64 + 16,		dist = 21 (26 - 5)
	cache.AddToCache(ChunkId(76), NodeId(79)) // 76 = 64 + 8 + 4,	dist = 24 (26 - 2)
	cache.AddToCache(ChunkId(00), NodeId(79)) // 0,				dist = 19 (26 - 7)
	cache.AddToCache(ChunkId(64), NodeId(79)) // 64,				dist = 22 (26 - 4)

	if _, exists := cache.CacheMap[ChunkId(76)]; exists {
		t.Errorf("Expected ChunkId 76 to be removed")
	}
}

func TestUpdateCacheMap_LeastRecentUsed(t *testing.T) {
	config.InitConfigWithPath(path)
	config.SetCacheModel(2)

	cache := CacheStruct{
		Size:           3,
		CacheMap:       make(CacheMap),
		CacheMutex:     &sync.Mutex{},
		EvictionPolicy: GetCachePolicy(),
	}

	cache.AddToCache(ChunkId(1), NodeId(1))
	cache.AddToCache(ChunkId(2), NodeId(1))
	cache.AddToCache(ChunkId(3), NodeId(1))
	cache.AddToCache(ChunkId(4), NodeId(1))

	cache.EvictionPolicy.UpdateCacheMap(&cache, ChunkId(4), 0)

	if _, exists := cache.CacheMap[ChunkId(1)]; exists {
		t.Errorf("Expected ChunkId 1 to be removed")
	}
}

func TestUpdateCacheMap_LeastFrequentlyUsed(t *testing.T) {
	config.InitConfigWithPath(path)
	config.SetCacheModel(3)

	cache := CacheStruct{
		Size:           3,
		CacheMap:       make(CacheMap),
		CacheMutex:     &sync.Mutex{},
		EvictionPolicy: GetCachePolicy(),
	}

	for i := 0; i < 5; i++ {
		cache.AddToCache(ChunkId(1), NodeId(1))
	}
	for i := 0; i < 2; i++ {
		cache.AddToCache(ChunkId(2), NodeId(1))
	}
	for i := 0; i < 8; i++ {
		cache.AddToCache(ChunkId(8), NodeId(1))
	}

	cache.AddToCache(ChunkId(4), NodeId(1))

	// 4 should be removed, but there's not way for it to make it into the cache
	if _, exists := cache.CacheMap[ChunkId(2)]; exists {
		t.Errorf("Expected ChunkId 2 to be removed")
	}
}

func TestCacheStruct_Contains(t *testing.T) {
	config.InitConfigWithPath(path)
	config.SetCacheModel(0)

	cache := CacheStruct{
		Size:           3,
		CacheMap:       make(CacheMap),
		CacheMutex:     &sync.Mutex{},
		EvictionPolicy: GetCachePolicy(),
	}

	network := &Network{}
	network.Bits = 1
	node := network.node(NodeId(1))

	cache.AddToCache(ChunkId(1), node.Id)
	cache.AddToCache(ChunkId(2), node.Id)
	cache.AddToCache(ChunkId(3), node.Id)

	if !cache.Contains(ChunkId(2)) {
		t.Errorf("Expected ChunkId 2 to be in the cache, but it's not")
	}

	if cache.Contains(ChunkId(4)) {
		t.Errorf("Expected ChunkId 4 to not be in the cache, but it is")
	}
}
