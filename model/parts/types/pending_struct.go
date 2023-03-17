package types

import (
	"go-incentive-simulation/model/constants"
	"sync"
)

type ChunkStruct struct {
	ChunkId int
	Counter int
}

type Pending struct {
	ChunkStructs []ChunkStruct
	//PendingCounters []int // chunkIdIndex -> counter for chunkId
	EpokeDecrement int32
}

type PendingMap map[int]Pending

type PendingStruct struct {
	PendingMap   PendingMap
	PendingMutex *sync.Mutex
	Counter      int32
}

func (p *PendingStruct) GetPending(originator int) Pending {
	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()
	pending, ok := p.PendingMap[originator]
	if ok {
		return pending
	}
	return Pending{ChunkStructs: []ChunkStruct{}, EpokeDecrement: 0}
}

func (p *PendingStruct) AddPendingChunkId(originator int, chunkId int) {
	pending := p.GetPending(originator)

	chunkStructIndex := p.GetChunkStructIndex(pending.ChunkStructs, chunkId)

	p.PendingMutex.Lock()
	defer p.PendingMutex.Unlock()

	if chunkStructIndex == -1 { // new chunk
		pending.ChunkStructs = append(pending.ChunkStructs, ChunkStruct{
			ChunkId: chunkId,
			Counter: 0,
		})
		p.PendingMap[originator] = pending

	} else { // chunk seen before
		if pending.ChunkStructs[chunkStructIndex].Counter < constants.GetBinSize() {
			pending.ChunkStructs[chunkStructIndex].Counter++
			p.PendingMap[originator] = pending

		} else {
			pending.ChunkStructs = append(pending.ChunkStructs[:chunkStructIndex]) // remove chunkStruct
			if len(pending.ChunkStructs) == 0 {
				delete(p.PendingMap, originator)
			}
		}
	}
}

func (p *PendingStruct) DeletePendingChunkId(originator int, chunkId int) {
	pending := p.GetPending(originator)

	if len(pending.ChunkStructs) > 0 {
		chunkIdIndex := p.GetChunkStructIndex(pending.ChunkStructs, chunkId)
		if chunkIdIndex != -1 {
			p.PendingMutex.Lock()
			defer p.PendingMutex.Unlock()
			pending.ChunkStructs = append(pending.ChunkStructs[:chunkIdIndex])
			//pending.PendingCounters = append(pending.PendingCounters[:chunkIdIndex])
			if len(pending.ChunkStructs) == 0 {
				delete(p.PendingMap, originator)
			} else {
				p.PendingMap[originator] = pending
			}
		}
	}
}

func (p *PendingStruct) GetChunkStructIndex(chunkStructs []ChunkStruct, chunkId int) int {

	for i, v := range chunkStructs {
		if v.ChunkId == chunkId {
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
