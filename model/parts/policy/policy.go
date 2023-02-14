package policy

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"math/rand"
)

func SendRequest(prevState *State, index int) (bool, Route, [][]Threshold, bool, []Payment) {
	// Gets one random chunkId from the range of addresses
	chunkId := rand.Intn(Constants.GetRangeAddress() - 1)
	var random float32

	if Constants.IsCacheEnabled() == true {
		numPreferredChunks := 1000
		random = rand.Float32()
		if float32(random) <= 0.5 {
			chunkId = rand.Intn(numPreferredChunks)
		} else {
			chunkId = rand.Intn(Constants.GetRangeAddress()-numPreferredChunks) + numPreferredChunks
		}
	}

	responsibleNodes := prevState.Graph.FindResponsibleNodes(chunkId)
	originatorId := prevState.Originators[prevState.OriginatorIndex]
	//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]

	if _, ok := prevState.PendingMap[originatorId]; ok {
		chunkId = prevState.PendingMap[originatorId]
		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	}
	if _, ok := prevState.RerouteMap[originatorId]; ok {
		chunkId = prevState.RerouteMap[originatorId][len(prevState.RerouteMap[originatorId])-1]
		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	}

	request := Request{OriginatorId: originatorId, ChunkId: chunkId}

	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.Graph, responsibleNodes, prevState.RerouteMap, prevState.CacheStruct.CacheMap)

	return found, route, thresholdFailed, accessFailed, paymentsList
}
