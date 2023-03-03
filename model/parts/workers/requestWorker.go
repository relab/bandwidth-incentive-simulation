package workers

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
	"math/rand"
)

func RequestWorker(newStateChan chan bool, requestChan chan Request, globalState *State, iterations int32) {

	//curState := globalState
	requestQueueSize := 1
	var originatorIndex int32
	var timeStep int32 = 0
	for timeStep < iterations {
		//if len(requestChan) < Constants.GetNumGoroutines() {
		if len(requestChan) <= requestQueueSize {

			//curState = <-stateChan
			//<-newStateChan
			timeStep = UpdateTimestep(globalState)
			//timeStep = atomic.LoadInt32(&globalState.TimeStep)

			originatorIndex = UpdateOriginatorIndex(globalState, timeStep)

			chunkId := rand.Intn(Constants.GetRangeAddress() - 1)

			if Constants.IsPreferredChunksEnabled() {
				var random float32
				numPreferredChunks := 1000
				random = rand.Float32()
				if float32(random) <= 0.5 {
					chunkId = rand.Intn(numPreferredChunks)
				} else {
					chunkId = rand.Intn(Constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
				}
			}

			responsibleNodes := globalState.Graph.FindResponsibleNodes(chunkId)
			originatorId := globalState.Originators[originatorIndex]
			//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]

			if _, ok := globalState.PendingMap[originatorId]; ok {
				chunkId = globalState.PendingMap[originatorId]
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			if _, ok := globalState.RerouteMap[originatorId]; ok {
				chunkId = globalState.RerouteMap[originatorId][len(globalState.RerouteMap[originatorId])-1]
				responsibleNodes = globalState.Graph.FindResponsibleNodes(chunkId)
			}

			requestChan <- Request{
				OriginatorId: originatorId,
				ChunkId:      chunkId,
				RespNodes:    responsibleNodes,
			}

		}
	}
}
