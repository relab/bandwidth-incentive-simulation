package types

//import (
//	"go-incentive-simulation/model/general"
//	"sync"
//)
//
//type PendingN struct {
//	NodeIds        []int
//	PendingCounter int32
//}
//
//type PendingM map[int]PendingN
//
//type PendingS struct {
//	PendingM     PendingM
//	PendingMutex *sync.Mutex
//}
//
//func (p *PendingS) Check(pendingN []int, chunkId int) int {
//	p.PendingMutex.Lock()
//	defer p.PendingMutex.Unlock()
//	if general.Contains(pendingN, chunkId) {
//		for i, v := range pendingN {
//			if v == chunkId {
//				return i
//			}
//		}
//	}
//	return -1
//}
//
//func (p *PendingS) Get(originator int) PendingN {
//	p.PendingMutex.Lock()
//	defer p.PendingMutex.Unlock()
//	pendingN, ok := p.PendingM[originator]
//	if ok {
//		return pendingN
//	}
//	return PendingN{NodeIds: []int{-1}}
//}
//
//func (p *PendingS) Add(originator int, nodeId int) {
//	p.PendingMutex.Lock()
//	pendingN := p.PendingM[originator]
//	pendingN.NodeIds = append(pendingN.NodeIds, nodeId)
//	pendingN.PendingCounter = 1
//	p.PendingM[originator] = pendingN
//	p.PendingMutex.Unlock()
//}
//
//func (p *PendingS) Delete(originator int, nodeIdIndex int) {
//	p.PendingMutex.Lock()
//	pendingN := p.PendingM[originator]
//	pendingN.NodeIds = append(pendingN.NodeIds[:nodeIdIndex], pendingN.NodeIds[nodeIdIndex+1:]...)
//	p.PendingMutex.Unlock()
//}

// for delete
// if len(p) > 0 or just check the first index is not -1
// for the global delete delet the enitre key - values
// if

//func noe() {
//	pending := PendingS{PendingM: make(PendingM), PendingMutex: &sync.Mutex{}}
//	pending.PendingM[1] = PendingN{NodeIds: []int{1, 2, 3}, PendingCounter: 1}
//	pending.PendingM[2] = PendingN{NodeIds: []int{1, 2, 3}, PendingCounter: 1}
//
//	p := pending.Get(1).NodeIds[0]
//
//}
