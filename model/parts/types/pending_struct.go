package types

import (
	"sync"
)

type PendingNode struct {
	NodeIds        []int
	PendingCounter int32
	EpokeDecrement int32
}

type PendingMap map[int]PendingNode

type PendingStruct struct {
	PendingMap   PendingMap
	PendingMutex *sync.Mutex
	Counter      int32
}

func (p *PendingStruct) IsEmpty(originator int) bool {
	pending := p.GetPending(originator)
	if len(pending.NodeIds) > 0 {
		return false
	}
	return true
}

func (p *PendingStruct) GetPending(originator int) PendingNode {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNode, ok := p.PendingMap[originator]
	if ok {
		return pendingNode
	}
	return PendingNode{NodeIds: make([]int, 0, 10000), PendingCounter: 0}
}

func (p *PendingStruct) IncrementPending(originator int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.PendingCounter++
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) AddPending(originator int, chunkId int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeIds = append(pendingNode.NodeIds, chunkId)
	pendingNode.PendingCounter = 1
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) AddToPendingQueue(originator int, chunkId int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeIds = append(pendingNode.NodeIds, chunkId)
	pendingNode.PendingCounter++
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) DeletePendingNodeId(originator int, pendingNodeIdIndex int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeIds = append(pendingNode.NodeIds[:pendingNodeIdIndex])
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}
func (p *PendingStruct) DeletePending(originator int) {
	p.PendingMutex.Lock()
	delete(p.PendingMap, originator)
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) GetPendingIndex(originator int, chunkId int) int {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNodes := p.PendingMap[originator].NodeIds
	for i, v := range pendingNodes {
		if v == chunkId {
			return i
		}
	}
	return -1
}
