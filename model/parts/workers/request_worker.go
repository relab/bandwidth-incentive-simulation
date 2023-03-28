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
	var counter = 0
	var curEpoch = constants.GetEpoch()
	var chunkId types.ChunkId
	var respNodes [4]types.NodeId
	var pickedFromWaiting = 0
	var pickedFromRetry = 0

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			timeStep := update.Timestep(globalState)

			if constants.TimeForNewEpoch(timeStep) {
				curEpoch = update.Epoch(globalState)

				waitForRoutingWorkers(pauseChan, continueChan)
			}

			originatorIndex := update.OriginatorIndex(globalState, timeStep)
			originatorId := globalState.GetOriginatorId(originatorIndex)
			originator := globalState.Graph.GetNode(originatorId)

			chunkId = -1

			if constants.IsRetryWithAnotherPeer() {
				rerouteStruct := originator.RerouteStruct

				if len(rerouteStruct.Reroute.RejectedNodes) > 0 {
					chunkId = rerouteStruct.Reroute.ChunkId
					respNodes = globalState.Graph.FindResponsibleNodes(chunkId)
					pickedFromRetry++
				}
			}

			if constants.IsWaitingEnabled() && chunkId == -1 { // No valid chunkId in reroute
				pendingStruct := originator.PendingStruct

				if pendingStruct.PendingQueue != nil {
					queuedChunk, ok := pendingStruct.GetChunkFromQueue(curEpoch)
					if ok {
						chunkId = queuedChunk.ChunkId
						respNodes = globalState.Graph.FindResponsibleNodes(queuedChunk.ChunkId)
						pickedFromWaiting++
					}
				}
			}

			if constants.IsIterationMeansUniqueChunk() {
				if chunkId == -1 {
					counter++
				}
			} else {
				counter++
			}

			if chunkId == -1 { // No waiting and no retry, and qualify for unique chunk
				chunkId = utils.GetChunkId()

				if constants.IsPreferredChunksEnabled() {
					chunkId = utils.GetPreferredChunkId()
				}
				respNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			if chunkId != -1 {
				request := types.Request{
					TimeStep:        timeStep,
					Epoch:           curEpoch,
					OriginatorIndex: originatorIndex,
					OriginatorId:    originatorId,
					ChunkId:         chunkId,
					RespNodes:       respNodes,
				}
				requestChan <- request
			}

			if constants.TimeForDebugPrints(timeStep) {
				fmt.Println("TimeStep is currently:", timeStep)
			}
		}
	}
	fmt.Println("chunks picked from Pending: ", pickedFromWaiting)
	fmt.Println("chunks picked from Retry: ", pickedFromRetry)
}
