package types

import "sync"

type PendingMap map[int]int

type PendingStruct struct {
	PendingMap   PendingMap
	PendingMutex *sync.Mutex
}

func (p *PendingStruct) GetPendingMap(originator int) int {
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
	defer p.PendingMutex.Unlock()
	delete(p.PendingMap, originator)
}

func (p *PendingStruct) AddPending(originator int, route int) {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	p.PendingMap[originator] = route
}
