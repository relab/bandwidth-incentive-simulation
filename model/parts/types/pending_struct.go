package types

import (
	"sync"
)

type PendingNode struct {
	ChunkIds       []int
	PendingCounter int32
	EpokeDecrement int32
}

type PendingMap map[int]PendingNode

type PendingStruct struct {
	PendingMap   PendingMap
	PendingMutex *sync.Mutex
	Counter      int32
}

func (p *PendingStruct) GetPending(originator int) PendingNode {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNode, ok := p.PendingMap[originator]
	if ok {
		return pendingNode
	}
	return PendingNode{ChunkIds: []int{}, PendingCounter: 0}
}

func (p *PendingStruct) AddPendingChunkIdToQueue(originator int, chunkId int) {
	pendingNode := p.GetPending(originator)

	p.PendingMutex.Lock()
	if len(pendingNode.ChunkIds) > 0 {
		// add to the queue
		pendingNode.ChunkIds = append(pendingNode.ChunkIds, chunkId)
		pendingNode.PendingCounter++
	} else {
		// if empty, create a new queue
		pendingNode.ChunkIds = append(pendingNode.ChunkIds, chunkId)
		pendingNode.PendingCounter = 1
	}
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) DeleteChunkIdFromPendingQueue(originator int, chunkId int) {
	pendingNode := p.GetPending(originator)

	if len(pendingNode.ChunkIds) > 0 {
		chunkIdIndex := p.GetPendingIndex(originator, chunkId)
		if chunkIdIndex != -1 {
			p.PendingMutex.Lock()
			pendingNode.ChunkIds = append(pendingNode.ChunkIds[:chunkIdIndex])
			p.PendingMap[originator] = pendingNode
			p.PendingMutex.Unlock()
		}
	}
}

func (p *PendingStruct) GetPendingIndex(originator int, chunkId int) int {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNodes := p.PendingMap[originator].ChunkIds
	for i, v := range pendingNodes {
		if v == chunkId {
			return i
		}
	}
	return -1
}

//func (p *PendingStruct) AddPending(originator int, chunkId int) {
//	p.PendingMutex.Lock()
//	pendingNode := p.PendingMap[originator]
//	pendingNode.ChunkIds = append(pendingNode.ChunkIds, chunkId)
//	pendingNode.PendingCounter = 1
//	p.PendingMap[originator] = pendingNode
//	p.PendingMutex.Unlock()
//}
//
//func (p *PendingStruct) AddToPendingQueue(originator int, chunkId int) {
//	p.PendingMutex.Lock()
//	pendingNode := p.PendingMap[originator]
//	pendingNode.ChunkIds = append(pendingNode.ChunkIds, chunkId)
//	pendingNode.PendingCounter++
//	p.PendingMap[originator] = pendingNode
//	p.PendingMutex.Unlock()
//}
//func (p *PendingStruct) DeletePending(originator int) {
//	p.PendingMutex.Lock()
//	delete(p.PendingMap, originator)
//	p.PendingMutex.Unlock()
//}

//func (p *PendingStruct) IncrementPending(originator int) {
//	p.PendingMutex.Lock()
//	pendingNode := p.PendingMap[originator]
//	pendingNode.PendingCounter++
//	p.PendingMap[originator] = pendingNode
//	p.PendingMutex.Unlock()
//}

//func (p *PendingStruct) IsEmpty(originator int) bool {
//	pending := p.GetPending(originator)
//	if len(pending.ChunkIds) > 0 {
//		return false
//	}
//	return true
//}
//
//func (p *PendingStruct) DeletePendingNodeId(originator int, pendingNodeIdIndex int) {
//	p.PendingMutex.Lock()
//	pendingNode := p.PendingMap[originator]
//	pendingNode.ChunkIds = append(pendingNode.ChunkIds[:pendingNodeIdIndex])
//	p.PendingMap[originator] = pendingNode
//	p.PendingMutex.Unlock()
//}
