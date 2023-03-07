package types

import (
	"sync"
)

type PendingNode struct {
	NodeId         int
	PendingCounter int32
}

type PendingMap map[int]PendingNode

type PendingStruct struct {
	PendingMap   PendingMap
	PendingMutex *sync.Mutex
}

func (p *PendingStruct) GetPending(originator int) PendingNode {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNode, ok := p.PendingMap[originator]
	if ok {
		return pendingNode
	}
	return PendingNode{NodeId: -1}
}

func (p *PendingStruct) DeletePending(originator int) {
	p.PendingMutex.Lock()
	delete(p.PendingMap, originator)
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) AddPending(originator int, chunkId int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeId = chunkId
	pendingNode.PendingCounter = 1
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) Increment(originator int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.PendingCounter++
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}
