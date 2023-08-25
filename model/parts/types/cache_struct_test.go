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
	cache.AddToCache(ChunkId(400), node.Id)
	if len(cache.CacheMap) != 3 {
		t.Errorf("Expected cache size to be 3, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(3000)]; exists {
		t.Errorf("Expected ChunkId 3000 to be removed")
	}

	config.SetCacheModel(2)
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
	cache.AddToCache(ChunkId(5), node.Id)
	if len(cache.CacheMap) != 4 {
		t.Errorf("Expected cache size to be 4 (Unlimited caching) after adding one more entry, but got %d", len(cache.CacheMap))
	}

	if _, exists := cache.CacheMap[ChunkId(5)]; !exists {
		t.Errorf("Expected ChunkId 5 to be added")
	}
}

func TestUpdateCacheMap_Proximity(t *testing.T) {
	config.InitConfig()
	config.SetCacheModel(1)

	cache := CacheStruct{
		Size:       3,
		CacheMap:   make(CacheMap),
		CacheMutex: &sync.Mutex{},
	}

	cache.CacheMap[ChunkId(1)] = CacheData{Proximity: 5}
	cache.CacheMap[ChunkId(2)] = CacheData{Proximity: 2}
	cache.CacheMap[ChunkId(3)] = CacheData{Proximity: 8}

	UpdateCacheMap(&cache, ChunkId(4), 1, 3)

	if _, exists := cache.CacheMap[ChunkId(2)]; exists {
		t.Errorf("Expected ChunkId 2 to be removed")
	}
}

func TestUpdateCacheMap_LeastRecentUsed(t *testing.T) {
	config.InitConfig()
	config.SetCacheModel(2)

	cache := CacheStruct{
		Size:       3,
		CacheMap:   make(CacheMap),
		CacheMutex: &sync.Mutex{},
	}

	now := time.Now()

	cache.CacheMap[ChunkId(1)] = CacheData{LastTimeUsed: now.Add(-time.Hour)}
	cache.CacheMap[ChunkId(2)] = CacheData{LastTimeUsed: now.Add(-time.Minute)}
	cache.CacheMap[ChunkId(3)] = CacheData{LastTimeUsed: now.Add(-time.Second)}

	UpdateCacheMap(&cache, ChunkId(4), 2, 0)

	if _, exists := cache.CacheMap[ChunkId(1)]; exists {
		t.Errorf("Expected ChunkId 1 to be removed")
	}
}

func TestUpdateCacheMap_LeastFrequentlyUsed(t *testing.T) {
	config.InitConfig()
	config.SetCacheModel(3)

	cache := CacheStruct{
		Size:       3,
		CacheMap:   make(CacheMap),
		CacheMutex: &sync.Mutex{},
	}

	cache.CacheMap[ChunkId(1)] = CacheData{Frequency: 5}
	cache.CacheMap[ChunkId(2)] = CacheData{Frequency: 2}
	cache.CacheMap[ChunkId(3)] = CacheData{Frequency: 8}

	UpdateCacheMap(&cache, ChunkId(4), 3, 0)

	if _, exists := cache.CacheMap[ChunkId(2)]; exists {
		t.Errorf("Expected ChunkId 2 to be removed")
	}
}

func TestCacheStruct_Contains(t *testing.T) {
	cache := CacheStruct{
		Size:       3,
		CacheMap:   make(CacheMap),
		CacheMutex: &sync.Mutex{},
	}

	config.InitConfig()
	config.SetCacheModel(0)

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
