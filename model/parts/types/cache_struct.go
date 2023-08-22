package types

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"math"
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
	Size       uint
	Node       *Node
	CacheMap   CacheMap
	CacheMutex *sync.Mutex
}

func FindDistance(chunkId ChunkId, nodeId NodeId) int {
	return config.GetBits() - general.BitLength(chunkId.ToInt()^nodeId.ToInt())
}

func (c *CacheStruct) AddToCache(chunkId ChunkId, nodeId NodeId) CacheMap {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	distance := FindDistance(chunkId, nodeId)
	cacheModel := config.GetCacheModel()
	if cacheModel == -1 {
		fmt.Println("No cache model is selected.")
		return nil
	}

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

	if len(c.CacheMap) > int(c.Size) {
		UpdateCacheMap(c, chunkId, cacheModel, distance)
	}

	// fmt.Println("CacheMapSize: ", len(c.CacheMap))
	return c.CacheMap
}

func UpdateCacheMap(c *CacheStruct, newChunkId ChunkId, cacheModel int, distance int) {
	chunkIdToDelete := ChunkId(-1)
	if cacheModel == 1 { // proximity
		minProximity := math.MaxInt32
		for chunkId, cacheData := range c.CacheMap {
			if cacheData.Proximity < minProximity {
				minProximity = cacheData.Proximity
				chunkIdToDelete = chunkId
			}
		}
	}
	if cacheModel == 2 { // leastRecentUsed
		leastRecentUsedTime := time.Now()
		for chunkId, cacheData := range c.CacheMap {
			if cacheData.LastTimeUsed.Before(leastRecentUsedTime) {
				leastRecentUsedTime = cacheData.LastTimeUsed
				chunkIdToDelete = chunkId
			}
		}
	}
	if cacheModel == 3 { // leastFrequentlyUsed
		leastFrequentlyUsed := math.MaxInt32
		for chunkId, cacheData := range c.CacheMap {
			if cacheData.Frequency < leastFrequentlyUsed {
				leastFrequentlyUsed = cacheData.Frequency
				chunkIdToDelete = chunkId
			}
		}
	}
	if int(chunkIdToDelete) != -1 {
		delete(c.CacheMap, chunkIdToDelete)
	}
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
