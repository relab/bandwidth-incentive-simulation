package types

import "sync"

type CacheMap map[int]map[int]int

type CacheStruct struct {
	CacheHits  int
	CacheMap   CacheMap
	CacheMutex *sync.Mutex
}

func (c *CacheStruct) Contains(nodeId int, chunkId int) bool {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	nodeMap := c.CacheMap[nodeId]
	if nodeMap != nil && nodeMap[chunkId] > 0 {
		return true
	}
	return false
}

func (c *CacheStruct) AddToCache(nodeId int, chunkId int) {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	nodeMap := c.CacheMap[nodeId]
	if nodeMap != nil {
		if _, ok2 := nodeMap[chunkId]; ok2 {
			nodeMap[chunkId]++
		} else {
			nodeMap[chunkId] = 1
		}
	} else {
		c.CacheMap[nodeId] = map[int]int{chunkId: 1}
	}
	return
}
