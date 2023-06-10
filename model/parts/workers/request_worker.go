package workers

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"go-incentive-simulation/model/parts/utils"
	"sync"
)

func RequestWorker(pauseChan chan bool, continueChan chan bool, requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup) {

	defer wg.Done()
	var requestQueueSize = 10
	var counter = 0
	var curEpoch = 0
	var chunkId types.ChunkId
	// var respNodes [4]types.NodeId
	iterations := config.GetIterations()
	numRoutingGoroutines := config.GetNumRoutingGoroutines()

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			timeStep := update.TimeStep(globalState)

			if config.TimeForNewEpoch(timeStep) {
				curEpoch = update.Epoch(globalState)

				waitForRoutingWorkers(pauseChan, continueChan, numRoutingGoroutines)
			}

			originatorIndex := int(update.OriginatorIndex(globalState, timeStep))
			originatorId := globalState.GetOriginatorId(originatorIndex)
			originator := globalState.Graph.GetNode(originatorId)

			// Needed for checks waiting and retry
			chunkId = -1

			if config.IsRetryWithAnotherPeer() {
				rerouteStruct := originator.RerouteStruct

				if len(rerouteStruct.Reroute.RejectedNodes) > 0 {
					chunkId = rerouteStruct.Reroute.ChunkId
					// respNodes = globalState.Graph.FindResponsibleNodes(chunkId)
				}
			}

			if config.IsWaitingEnabled() && chunkId == -1 { // No valid chunkId in reroute
				pendingStruct := originator.PendingStruct

				if pendingStruct.PendingQueue != nil {
					queuedChunk, ok := pendingStruct.GetChunkFromQueue(curEpoch)
					if ok {
						chunkId = queuedChunk.ChunkId
						// respNodes = globalState.Graph.FindResponsibleNodes(queuedChunk.ChunkId)
					}
				}
			}

			if config.IsIterationMeansUniqueChunk() {
				if chunkId == -1 { // Only increment the counter chunk is not chosen from waiting or retry
					counter++
				}
			} else {
				counter++ // Increment all iterations
			}

			if chunkId == -1 { // No waiting and no retry, and qualify for unique chunk
				chunkId = utils.GetNewChunkId()

				if config.IsPreferredChunksEnabled() {
					chunkId = utils.GetPreferredChunkId()
				}

				// respNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			if chunkId != -1 { // Should always be true, but just in case
				request := types.Request{
					TimeStep:        timeStep,
					Epoch:           curEpoch,
					OriginatorIndex: originatorIndex,
					OriginatorId:    originatorId,
					ChunkId:         chunkId,
					// RespNodes:       respNodes,
				}
				requestChan <- request
			}

			if config.TimeForDebugPrints(timeStep) {
				fmt.Println("TimeStep is currently:", timeStep)
			}
		}
	}
}
