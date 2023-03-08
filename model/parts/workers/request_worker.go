package workers

import (
	"fmt"
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
	"sync"
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

			if timeStep == 5000000 {
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

			responsibleNodes := globalState.Graph.FindResponsibleNodes(chunkId)
			originatorId := globalState.Originators[originatorIndex]
			//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]

			pendingNodeId := globalState.PendingStruct.GetPending(originatorId).NodeId
			if pendingNodeId != -1 {
				chunkId = pendingNodeId
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
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
