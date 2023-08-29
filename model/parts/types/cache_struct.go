package types

import (
	"fmt"
	"github.com/zavitax/sortedset-go"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"sync"
	"time"
)

type CacheMap map[ChunkId]CacheData

type CacheData struct {
	Proximity    int
	LastTimeUsed time.Time
	Frequency    int
}

type CacheStruct struct {
	Size           int
	Node           *Node
	CacheMap       CacheMap
	CacheMutex     *sync.Mutex
	EvictionPolicy CachePolicy
}

type CachePolicy interface {
	UpdateCacheMap(c *CacheStruct, newChunkId ChunkId, distance int)
}

type (
	proximityPolicy struct {
		ChunkSet *sortedset.SortedSet[ChunkId, int, CacheData]
	}

	lruPolicy struct {
		ChunkSet *sortedset.SortedSet[ChunkId, int, CacheData]
	}

	lfuPolicy struct {
		ChunkSet *sortedset.SortedSet[ChunkId, int, CacheData]
	}
)

func GetCachePolicy() CachePolicy {
	policy := config.GetCacheModel()
	if policy == -1 {
		fmt.Println("No cache model is selected.")
		return nil
	}

	if policy == 1 {
		return &proximityPolicy{ChunkSet: sortedset.New[ChunkId, int, CacheData]()}
	} else if policy == 2 {
		return &lruPolicy{ChunkSet: sortedset.New[ChunkId, int, CacheData]()}
	} else if policy == 3 {
		return &lfuPolicy{ChunkSet: sortedset.New[ChunkId, int, CacheData]()}
	} else {
		return nil
	}
}

func FindDistance(chunkId ChunkId, nodeId NodeId) int {
	return config.GetBits() - general.BitLength(chunkId.ToInt()^nodeId.ToInt())
}

func (c *CacheStruct) AddToCache(chunkId ChunkId, nodeId NodeId) CacheMap {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	distance := FindDistance(chunkId, nodeId)

	if _, ok := c.CacheMap[chunkId]; ok {
		currData := c.CacheMap[chunkId]
		currData.Frequency++
		currData.Proximity = distance
		c.CacheMap[chunkId] = currData
	} else {
		newCacheData := CacheData{
			Proximity:    distance,
			LastTimeUsed: time.Now(),
			Frequency:    1,
		}
		c.CacheMap[chunkId] = newCacheData
	}

	c.EvictionPolicy.UpdateCacheMap(c, chunkId, distance)

	return c.CacheMap
}

func (p *proximityPolicy) UpdateCacheMap(c *CacheStruct, newChunkId ChunkId, distance int) {
	p.ChunkSet.AddOrUpdate(newChunkId, distance, CacheData{distance, time.Now(), 1})

	if len(c.CacheMap) <= int(c.Size) {
		return
	}

	chunkIdToDelete := p.ChunkSet.PopMin().Key()
	delete(c.CacheMap, chunkIdToDelete)
}

func (p *lruPolicy) UpdateCacheMap(c *CacheStruct, newChunkId ChunkId, distance int) {
	p.ChunkSet.AddOrUpdate(newChunkId, time.Now().Nanosecond(), CacheData{distance, time.Now(), 1})

	if len(c.CacheMap) <= int(c.Size) {
		return
	}

	chunkIdToDelete := p.ChunkSet.PopMin().Key()
	delete(c.CacheMap, chunkIdToDelete)
}

func (p *lfuPolicy) UpdateCacheMap(c *CacheStruct, newChunkId ChunkId, distance int) {
	prev := p.ChunkSet.GetByKey(newChunkId)
	freq := 1
	if prev != nil {
		freq = prev.Value.Frequency + 1
	}

	p.ChunkSet.AddOrUpdate(newChunkId, freq, CacheData{distance, time.Now(), freq})

	if len(c.CacheMap) <= int(c.Size) {
		return
	}

	chunkIdToDelete := p.ChunkSet.PopMin().Key()
	delete(c.CacheMap, chunkIdToDelete)
}

func (c *CacheStruct) Contains(chunkId ChunkId) bool {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	if _, ok := c.CacheMap[chunkId]; ok {
		cacheData := c.CacheMap[chunkId]
		cacheData.LastTimeUsed = time.Now()
		cacheData.Frequency = cacheData.Frequency + 1
		c.CacheMap[chunkId] = cacheData
		return true
	}

	return false
}