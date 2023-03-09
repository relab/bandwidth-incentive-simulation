package workers

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
	"sync"
	"sync/atomic"
)

func RequestWorker(requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup, iterations int32) {

	defer wg.Done()
	var requestQueueSize = 10
	var originatorIndex int32 = 0
	var timeStep int32 = 0
	var counter int32 = 0

	defer close(requestChan)

	for counter < iterations {
		if len(requestChan) <= requestQueueSize {

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			timeStep = int32(update.Timestep(globalState))
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)

			//if timeStep%(iterations/2) == 0 {
			//	fmt.Println("PendingMap is currently:", globalState.PendingStruct.PendingMap)
			//	fmt.Println("RerouteMap is currently:", globalState.RerouteStruct.RerouteMap)
			//}

			originatorIndex = update.OriginatorIndex(globalState, timeStep)

			originatorId := globalState.Originators[originatorIndex]

			chunkId := -1
			responsibleNodes := [4]int{}

			if constants.Constants.IsWaitingEnabled() {
				pendingNode := globalState.PendingStruct.GetPending(originatorId)

				var epoke int32 = 50_000
				if (timeStep-originatorIndex)%epoke == 0 || timeStep > iterations {
					if len(pendingNode.NodeIds) > 0 {
						pendingNode.EpokeDecrement = int32(len(pendingNode.NodeIds))
						atomic.AddInt32(&globalState.PendingStruct.Counter, int32(len(pendingNode.NodeIds)))

					}
				}

				if pendingNode.EpokeDecrement > 0 {
					pendingNodeIds := pendingNode.NodeIds
					if !globalState.PendingStruct.IsEmpty(originatorId) {
						chunkId = pendingNodeIds[pendingNode.EpokeDecrement-1]
						responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
						pendingNode.EpokeDecrement--
					}
				}
			}

			if constants.Constants.IsRetryWithAnotherPeer() {
				reroute := globalState.RerouteStruct.GetRerouteMap(originatorId)
				if reroute != nil {
					chunkId = reroute[len(reroute)-1]
					responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
				}
			}

			if constants.Constants.IsIterationMeansUniqueChunk() {
				if chunkId == -1 { // No waiting and no retry
					counter++
				}
			} else {
				counter++
			}

			if chunkId == -1 { // No waiting and no retry
				chunkId = rand.Intn(constants.Constants.GetRangeAddress() - 1)

				if constants.Constants.IsPreferredChunksEnabled() {
					var random float32
					numPreferredChunks := 1000
					random = rand.Float32()
					if float32(random) <= 0.5 {
						chunkId = rand.Intn(numPreferredChunks)
					} else {
						chunkId = rand.Intn(constants.Constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
					}
				}
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			//if timeStep%(iterations/10) == 0 {
			//	fmt.Println("TimeStep is currently:", timeStep)
			//}

			requestChan <- types.Request{
				OriginatorIndex: originatorIndex,
				OriginatorId:    originatorId,
				TimeStep:        timeStep,
				ChunkId:         chunkId,
				RespNodes:       responsibleNodes,
			}
		}
	}
}
