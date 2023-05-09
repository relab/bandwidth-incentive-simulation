package types

import "sync"

type CacheMap map[ChunkId]int

type CacheStruct struct {
	Size       uint
	Node       *Node
	CacheMap   CacheMap
	CacheList  []ChunkId
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
	c.CacheList = append(c.CacheList, chunkId)
	if len(c.CacheList) > int(c.Size) {
		firstChunk := c.CacheList[0]
		c.CacheMap[firstChunk]--
		c.CacheList = c.CacheList[1:]
	}

	return c.CacheMap
}

func (c *CacheStruct) Contains(chunkId ChunkId) bool {
	c.CacheMutex.Lock()
	defer c.CacheMutex.Unlock()
	cacheMap := c.CacheMap
	if c, ok := cacheMap[chunkId]; ok && c > 0 {
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
