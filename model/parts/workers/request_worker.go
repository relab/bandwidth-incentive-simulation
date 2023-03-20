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
	var curEpoch = constants.GetEpoch()
	var PickedFromWaiting int

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
				curEpoch = update.Epoch(globalState)

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

			chunkId := -1

			if constants.IsRetryWithAnotherPeer() {

				routeStruct := globalState.RerouteStruct.GetRerouteMap(originatorId)
				//if routeStruct.LastEpoch < curEpoch {
				if routeStruct.Reroute != nil {
					chunkId = routeStruct.ChunkId
					responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
				}
				//globalState.RerouteStruct.UpdateEpoch(originatorId, curEpoch)
			}

			if constants.IsWaitingEnabled() && chunkId == -1 { // No valid chunkId in reroute
				pending, ok := globalState.PendingStruct.GetPending(originatorId)

				if ok && len(pending.ChunkQueue) > 0 {
					chunkStruct, _ := globalState.PendingStruct.GetChunkFromQueue(originatorId)
					if chunkStruct.LastEpoch < curEpoch {
						chunkId = chunkStruct.ChunkId
						responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkStruct.ChunkId)
						globalState.PendingStruct.UpdateEpoch(originatorId, chunkId, curEpoch)
						PickedFromWaiting++
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
					Epoch:           curEpoch,
					ChunkId:         chunkId,
					RespNodes:       responsibleNodes,
				}
			}
		}
	}
	fmt.Println("Number of requests chunks picked from Pending: ", PickedFromWaiting)
}
