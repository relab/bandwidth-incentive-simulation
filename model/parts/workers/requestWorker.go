package workers

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/update"
	"math/rand"
)

func RequestWorker(stateChan chan *State, requestChan chan Request, globalState *State, iterations int) {

	curState := globalState
	for i := 0; i < iterations; i++ {
		//curState = <-stateChan
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

		responsibleNodes := curState.Graph.FindResponsibleNodes(chunkId)
		originatorId := curState.Originators[curState.OriginatorIndex]
		//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]

		if _, ok := curState.PendingMap[originatorId]; ok {
			chunkId = curState.PendingMap[originatorId]
			responsibleNodes = curState.Graph.FindResponsibleNodes(chunkId)
		}

		if _, ok := curState.RerouteMap[originatorId]; ok {
			chunkId = curState.RerouteMap[originatorId][len(curState.RerouteMap[originatorId])-1]
			responsibleNodes = curState.Graph.FindResponsibleNodes(chunkId)
		}

		request := Request{
			OriginatorId: originatorId,
			ChunkId:      chunkId,
			RespNodes:    responsibleNodes,
		}

		requestChan <- request

		UpdateOriginatorIndex(curState)
	}
}
