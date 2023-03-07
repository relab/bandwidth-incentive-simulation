package types

import "sync"

type PendingMap map[int]int

type PendingStruct struct {
	PendingMap     PendingMap
	PendingCounter int
	PendingMutex   *sync.Mutex
}

func (p *PendingStruct) GetPending(originator int) int {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pendingNodeId, ok := p.PendingMap[originator]
	if ok {
		return pendingNodeId
	}
	return -1
}

func (p *PendingStruct) DeletePending(originator int) {
	p.PendingMutex.Lock()
	delete(p.PendingMap, originator)
	p.PendingMutex.Unlock()
}

func (p *PendingStruct) AddPending(originator int, route int) {
	p.PendingMutex.Lock()
	p.PendingMap[originator] = route
	p.PendingMutex.Unlock()
}
