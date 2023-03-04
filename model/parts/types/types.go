package types

type Request struct {
	OriginatorId int
	ChunkId      int
	RespNodes    [4]int
}

//type PendingMap map[int]int
//
//type PendingStruct struct {
//	PendingM  PendingMap
//	PendingMutex *sync.Mutex
//}
//
//func (p *PendingStruct) GetPendingMap(originator int) int {
//	p.PendingMutex.Lock()
//	defer p.PendingMutex.Unlock()
//	pendingNodeId, ok := p.PendingMap[originator]
//	if ok {
//		return pendingNodeId
//	}
//	return -1
//}
//
//func (p *PendingStruct) DeletePending(originator int) {
//	p.PendingMutex.Lock()
//	defer p.PendingMutex.Unlock()
//	delete(p.PendingMap, originator)
//}
//
//func (p *PendingStruct) AddPending(originator int, route int) {
//	p.PendingMutex.Lock()
//	defer p.PendingMutex.Unlock()
//	p.PendingMap[originator] = route
//}

//type RerouteMap map[int][]int
//
//type RerouteStruct struct {
//	RerouteM  RerouteMap
//	RerouteMutex *sync.RWMutex
//}
//
//func (r *RerouteStruct) GetRerouteMap(originator int) []int {
//	r.RerouteMutex.RLock()
//	defer r.RerouteMutex.RUnlock()
//	reroute, ok := r.RerouteMap[originator]
//	if ok {
//		return reroute
//	}
//	return nil
//}
//
//func (r *RerouteStruct) DeleteReroute(originator int) {
//	r.RerouteMutex.Lock()
//	defer r.RerouteMutex.Unlock()
//	delete(r.RerouteMap, originator)
//}

//type CacheMap map[int]map[int]int
//
//type CacheStruct struct {
//	CacheHits  int
//	CacheM   CacheMap
//	CacheMutex *sync.Mutex
//}
//
//func (c *CacheStruct) Contains(nodeId int, chunkId int) bool {
//	c.CacheMutex.Lock()
//	defer c.CacheMutex.Unlock()
//	nodeMap := c.CacheMap[nodeId]
//	if nodeMap != nil && nodeMap[chunkId] > 0 {
//		return true
//	}
//	return false
//}
//
//func (c *CacheStruct) AddToCache(nodeId int, chunkId int) {
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
//		c.CacheMap[nodeId] = map[int]int{chunkId: 1}
//	}
//	return
//}

type Route []int

type Payment struct {
	FirstNodeId  int
	PayNextId    int
	ChunkId      int
	IsOriginator bool
}

type Threshold [2]int

type State struct {
	Graph                   *Graph
	Originators             []int
	NodesId                 []int
	RouteLists              []Route
	PendingStruct           PendingStruct
	RerouteStruct           RerouteStruct
	CacheStruct             CacheStruct
	OriginatorIndex         int32
	SuccessfulFound         int32
	FailedRequestsThreshold int32
	FailedRequestsAccess    int32
	TimeStep                int32
}

type Policy struct {
	Found                bool
	Route                Route
	ThresholdFailedLists [][]Threshold
	AccessFailed         bool
	PaymentList          []Payment
}
