package workers

import (
	"go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/update"
	"math/rand"
)

func RequestWorker(newStateChan chan bool, requestChan chan types.Request, globalState *types.State, iterations int32) {

	//curState := globalState
	requestQueueSize := 10
	var originatorIndex int32
	var timeStep int32 = 0
	for timeStep < iterations {
		//if len(requestChan) < Constants.GetNumGoroutines() {
		if len(requestChan) <= requestQueueSize {

			//curState = <-stateChan
			//<-newStateChan
			timeStep = int32(update.Timestep(globalState))
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)

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

			pendingNodeId := globalState.PendingStruct.GetPending(originatorId)
			if pendingNodeId != -1 {
				chunkId = globalState.PendingStruct.GetPending(originatorId)
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
