package types

import (
	"go-incentive-simulation/model/general"
	"sync"
)

type PendingNode struct {
	NodeIds        []int
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
	return PendingNode{NodeIds: []int{-1}, PendingCounter: 0}
}

func (p *PendingStruct) AddPending(originator int, chunkId int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeIds = append(pendingNode.NodeIds, chunkId)
	pendingNode.PendingCounter = 1
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) DeletePending(originator int) {
	p.PendingMutex.Lock()
	delete(p.PendingMap, originator)
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) AddP(originator int, chunkId int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.NodeIds = append(pendingNode.NodeIds, chunkId)
	pendingNode.PendingCounter++
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) IncrementPending(originator int) {
	p.PendingMutex.Lock()
	pendingNode := p.PendingMap[originator]
	pendingNode.PendingCounter++
	p.PendingMap[originator] = pendingNode
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) CheckPending(originator int, chunkId int) int {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNodes := p.PendingMap[originator].NodeIds
	if general.Contains(pendingNodes, chunkId) {
		for i, v := range pendingNodes {
			if v == chunkId {
				return i
			}
		}
	}
	return -1
}

func (p *PendingStruct) DeletePendingNodeId(originator int, pendingNodeIdIndex int) {
	p.PendingMutex.Lock()
	pendingNodeIds := p.PendingMap[originator]
	pendingNodeIds.NodeIds = append(pendingNodeIds.NodeIds[:pendingNodeIdIndex], pendingNodeIds.NodeIds[pendingNodeIdIndex+1:]...)
	p.PendingMap[originator] = pendingNodeIds
	p.PendingMutex.Unlock()
}
