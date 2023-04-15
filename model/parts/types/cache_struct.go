package types

import "sync"

type CacheMap map[ChunkId]int

type CacheStruct struct {
	Node       *Node
	CacheMap   CacheMap
	CacheMutex *sync.Mutex
}

func (c *CacheStruct) AddToCache(chunkId ChunkId) CacheMap {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()

	if _, ok := c.CacheMap[chunkId]; ok {
		c.CacheMap[chunkId]++
	} else {
		c.CacheMap[chunkId] = 1
	}

	return c.CacheMap
}

func (c *CacheStruct) Contains(chunkId ChunkId) bool {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	cacheMap := c.CacheMap
	if _, ok := cacheMap[chunkId]; ok {
		return true
	}
	return false
}

//type CacheMap map[NodeId]map[ChunkId]int

//func (c *CacheStruct) LenMap() int {
//	c.CacheMutex.Lock()
//	defer c.CacheMutex.Unlock()
//	return len(c.CacheMap)
//}
//
//func (c *CacheStruct) Contains(nodeId NodeId, chunkId ChunkId) bool {
//	c.CacheMutex.Lock()
//	defer c.CacheMutex.Unlock()
//	nodeMap := c.CacheMap[nodeId]
//	if nodeMap != nil && nodeMap[chunkId] > 0 {
//		return true
//	}
//	return false
//}
//
//func (c *CacheStruct) AddToCache(nodeId NodeId, chunkId ChunkId) {
//	c.CacheMutex.Lock()
//	defer c.CacheMutex.Unlock()
//	nodeMap := c.CacheMap[nodeId]
//	if nodeMap != nil {
//		if _, ok2 := nodeMap[chunkId]; ok2 {
//			nodeMap[chunkId]++
//		} else {
//			nodeMap[chunkId] = 1
//		}
//	} else {
//		c.CacheMap[nodeId] = map[ChunkId]int{chunkId: 1}
//	}
//}
