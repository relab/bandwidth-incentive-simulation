package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
	"sync"
)

func RequestWorker(pauseChan chan bool, continueChan chan bool, requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup, iterations int) {

	defer wg.Done()
	var requestQueueSize = 10
	var originatorIndex = 0
	var timeStep = 0
	var counter = 0
	var responsibleNodes [4]int
	var chunkId int
	var curEpoke = constants.GetEpoke()
	var newEpokeCounter = 0

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			timeStep = update.Timestep(globalState)
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)
			if constants.IsDebugPrints() {
				if timeStep%constants.GetDebugInterval() == 0 {
					fmt.Println("TimeStep is currently:", timeStep)
				}
			}

			if timeStep%constants.GetRequestsPerSecond() == 0 {
				curEpoke = constants.UpdateEpoke()
				newEpokeCounter = constants.GetOriginators()

				for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
					pauseChan <- true
				}
				for i := 0; i < constants.GetNumRoutingGoroutines(); i++ {
					<-continueChan
				}
			}

			if constants.IsDebugPrints() {
				if timeStep%(iterations/2) == 0 {
					if constants.IsWaitingEnabled() {
						fmt.Println("PendingMap is currently:", globalState.PendingStruct.PendingMap)
					}
					if constants.IsRetryWithAnotherPeer() {
						fmt.Println("RerouteMap is currently:", globalState.RerouteStruct.RerouteMap)
					}
				}
			}

			originatorIndex = int(update.OriginatorIndex(globalState, timeStep))

			originatorId := globalState.Originators[originatorIndex]

			//originator := globalState.Graph.GetNode(originatorIndex)

			chunkId = -1

			if constants.IsWaitingEnabled() {
				pendingNode := globalState.PendingStruct.GetPending(originatorId)

				if newEpokeCounter > 0 || timeStep > iterations {
					if len(pendingNode.ChunkIds) > 0 {
						pendingNode.EpokeDecrement = int32(len(pendingNode.ChunkIds))
						//atomic.AddInt32(&globalState.PendingStruct.Counter, int32(len(pendingNode.ChunkIds)))
					}
				}
				if pendingNode.EpokeDecrement > 0 {
					pendingNodeIds := pendingNode.ChunkIds
					if len(pendingNodeIds) > 0 {
						//if !globalState.PendingStruct.IsEmpty(originatorId) {
						chunkId = pendingNodeIds[pendingNode.EpokeDecrement-1]
						responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
						pendingNode.EpokeDecrement--
					}
				}
				newEpokeCounter--
			}

			if constants.IsRetryWithAnotherPeer() {
				reroute := globalState.RerouteStruct.GetRerouteMap(originatorId)
				if reroute != nil {
					chunkId = reroute[len(reroute)-1]
					responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
				}
			}

			if constants.IsIterationMeansUniqueChunk() {
				if chunkId == -1 {
					counter++
				}
			} else {
				counter++
			}

			if chunkId == -1 && timeStep <= iterations { // No waiting and no retry, and qualify for unique chunk
				chunkId = rand.Intn(constants.GetRangeAddress() - 1)

				if constants.IsPreferredChunksEnabled() {
					var random float32
					numPreferredChunks := 1000
					random = rand.Float32()
					if float32(random) <= 0.5 {
						chunkId = rand.Intn(numPreferredChunks)
					} else {
						chunkId = rand.Intn(constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
					}
				}
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			if chunkId != -1 {
				requestChan <- types.Request{
					OriginatorIndex: originatorIndex,
					OriginatorId:    originatorId,
					TimeStep:        timeStep,
					Epoke:           curEpoke,
					ChunkId:         chunkId,
					RespNodes:       responsibleNodes,
				}
			}
		}
	}
}
