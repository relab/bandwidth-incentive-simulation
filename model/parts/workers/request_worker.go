package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RequestWorker(pauseChan chan bool, continueChan chan bool, requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup, iterations int) {

	defer wg.Done()
	var requestQueueSize = 10
	var originatorIndex int32 = 0
	var timeStep = 0
	var counter = 0
	var responsibleNodes [4]types.NodeId
	var curEpoch = constants.GetEpoch()
	var chunkId types.ChunkId
	var pickedFromWaiting = 0
	var pickedFromRetry = 0

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			timeStep = update.Timestep(globalState)
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)
			if constants.TimeForDebugPrints(timeStep) {
				fmt.Println("TimeStep is currently:", timeStep)
			}

			if constants.TimeForNewEpoch(timeStep) {
				curEpoch = update.Epoch(globalState)

				for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
					pauseChan <- true
				}
				for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
					<-continueChan
				}
			}

			//if constants.IsDebugPrints() && timeStep%(iterations/2) == 0 {
			//	if constants.IsWaitingEnabled() {
			//		fmt.Println("PendingMap is currently:", globalState.PendingStruct.PendingMap)
			//	}
			//	if constants.IsRetryWithAnotherPeer() {
			//		fmt.Println("RerouteMap is currently:", globalState.RerouteStruct.RerouteMap)
			//	}
			//}

			originatorIndex = update.OriginatorIndex(globalState, timeStep)
			originatorId := globalState.GetOriginatorId(originatorIndex)
			originator := globalState.Graph.GetNode(originatorId)

			chunkId = -1

			if constants.IsRetryWithAnotherPeer() {
				rerouteStruct := originator.RerouteStruct
				//if routeStruct.LastEpoch < curEpoch {

				if len(rerouteStruct.Reroute.CheckedNodes) > 0 {
					//fmt.Println(len(rerouteStruct.Reroute.CheckedNodes))
					chunkId = rerouteStruct.Reroute.ChunkId
					responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
					pickedFromRetry++
				}
				//globalState.RerouteStruct.UpdateEpoch(originatorId, curEpoch)
			}

			if constants.IsWaitingEnabled() && chunkId == -1 { // No valid chunkId in reroute
				pendingStruct := originator.PendingStruct

				if pendingStruct.PendingQueue != nil {
					queuedChunk, ok := pendingStruct.GetChunkFromQueue(curEpoch)
					if ok {
						chunkId = queuedChunk.ChunkId
						responsibleNodes = globalState.Graph.FindResponsibleNodes(queuedChunk.ChunkId)
						pickedFromWaiting++
					}
				}
			}

			//if constants.IsRetryWithAnotherPeer() {
			//
			//	routeStruct := globalState.RerouteStruct.GetRerouteMap(originatorId)
			//	//if routeStruct.LastEpoch < curEpoch {
			//	if routeStruct.CheckedNodes != nil {
			//		newChunkSet = true
			//		chunkId = routeStruct.ChunkId
			//		responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			//	}
			//	//globalState.RerouteStruct.UpdateEpoch(originatorId, curEpoch)
			//}
			//
			//if constants.IsWaitingEnabled() && !newChunkSet { // No valid chunkId in reroute
			//	pending, ok := globalState.PendingStruct.GetPending(originatorId)
			//
			//	if ok && len(pending.ChunkQueue) > 0 {
			//		chunkStruct, _ := globalState.PendingStruct.GetChunkFromQueue(originatorId)
			//		if chunkStruct.LastEpoch < curEpoch {
			//			newChunkSet = true
			//			chunkId = chunkStruct.ChunkId
			//			responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkStruct.ChunkId)
			//			globalState.PendingStruct.UpdateEpoch(originatorId, chunkId, curEpoch)
			//		}
			//	}
			//}

			if constants.IsIterationMeansUniqueChunk() {
				if chunkId == -1 {
					counter++
				}
			} else {
				counter++
			}

			if chunkId == -1 && timeStep <= iterations { // No waiting and no retry, and qualify for unique chunk
				chunkId = utils.GetChunkId()

				if constants.IsPreferredChunksEnabled() {
					chunkId = utils.GetPreferredChunkId()
				}
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			if chunkId != -1 {
				request := types.Request{
					TimeStep:        timeStep,
					Epoch:           curEpoch,
					OriginatorIndex: originatorIndex,
					OriginatorId:    originatorId,
					ChunkId:         chunkId,
					RespNodes:       responsibleNodes,
				}
				//fmt.Println(request)
				requestChan <- request
			}
		}
	}
	fmt.Println("Number of requests chunks picked from Pending: ", pickedFromWaiting)
	fmt.Println("Number of requests chunks picked from Retry: ", pickedFromRetry)
}
