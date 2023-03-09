package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
	"sync"
	"sync/atomic"
)

func RequestWorker(requestChan chan types.Request, globalState *types.State, wg *sync.WaitGroup, iterations int32) {

	defer wg.Done()
	requestQueueSize := 10
	var originatorIndex int32
	var timeStep int32 = 0
	for timeStep < iterations {
		if len(requestChan) <= requestQueueSize {

			// TODO: decide on where we should update the timestep. At request creation or request fulfillment
			timeStep = int32(update.Timestep(globalState))
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)

			if timeStep%(iterations/2) == 0 {
				fmt.Println("PendingMap is currently:", globalState.PendingStruct.PendingMap)
				fmt.Println("RerouteMap is currently:", globalState.RerouteStruct.RerouteMap)
			}

			originatorIndex = update.OriginatorIndex(globalState, timeStep)

			chunkId := rand.Intn(constants.Constants.GetRangeAddress() - 1)

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
			//fmt.Println("counter is:", counter)

			responsibleNodes := globalState.Graph.FindResponsibleNodes(chunkId)
			originatorId := globalState.Originators[originatorIndex]
			//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]
			//
			if constants.Constants.IsWaitingEnabled() {
				pendingNode := globalState.PendingStruct.GetPending(originatorId)

				//var epoke int32 = 50_000
				//if (timeStep-originatorIndex)%epoke == 0 {
				if len(pendingNode.ChunkIds) > 0 {
					pendingNode.EpokeDecrement = int32(len(pendingNode.ChunkIds))
					atomic.AddInt32(&globalState.PendingStruct.Counter, int32(len(pendingNode.ChunkIds)))
				}
				//}

				if pendingNode.EpokeDecrement > 0 {
					pendingNodeIds := pendingNode.ChunkIds
					if !globalState.PendingStruct.IsEmpty(originatorId) {
						chunkId = pendingNodeIds[pendingNode.EpokeDecrement-1]
						responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
						pendingNode.EpokeDecrement--
					}
				}
			}

			//if _, ok := globalState.PendingMap[originatorId]; ok {
			//	chunkId = globalState.PendingMap[originatorId]
			//	responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			//}

			reroute := globalState.RerouteStruct.GetRerouteMap(originatorId)
			if reroute != nil {
				chunkId = reroute[len(reroute)-1]
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			//if _, ok := globalState.RerouteMap[originatorId]; ok {
			//	chunkId = globalState.RerouteMap[originatorId][len(globalState.RerouteMap[originatorId])-1]
			//	responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			//}

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
