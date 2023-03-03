package policy

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"math/rand"
)

// TODO: keeping this here to compare with other implmentation later
//func oldFindResponsibleNodes(nodesId []int, chunkAdd int) []int {
//	//numNodes := Constants.GetBits()
//	numNodes := 100
//	distances := make([]int, 0, numNodes)
//	var distance int
//	nodesMap := make(map[int]int)
//	returnNodes := make([]int, 4)
//
//	closestNodes := BinarySearchClosest(nodesId, chunkAdd, numNodes)
//
//	for _, nodeId := range closestNodes {
//		distance = nodeId ^ chunkAdd
//		// fmt.Println(distance, nodeId)
//		distances = append(distances, distance)
//		nodesMap[distance] = nodeId
//	}
//
//	sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })
//
//	for i := 0; i < 4; i++ {
//		distance = distances[i]
//		returnNodes[i] = nodesMap[distance]
//	}
//	return returnNodes
//}

func SendRequest(prevState *State, index int) (bool, Route, [][]Threshold, bool, []Payment) {
	// Gets one random chunkId from the range of addresses
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

	responsibleNodes := prevState.Graph.FindResponsibleNodes(chunkId)
	originatorId := prevState.Originators[prevState.OriginatorIndex]
	//originatorId := prevState.Originators[rand.Intn(Constants.GetOriginators())]

	pendingNodeId := prevState.PendingStruct.GetPendingMap(originatorId)
	if pendingNodeId != -1 {
		chunkId = prevState.PendingStruct.GetPendingMap(originatorId)
		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	}
	//if _, ok := prevState.PendingMap[originatorId]; ok {
	//	chunkId = prevState.PendingMap[originatorId]
	//	responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	//}
	reroute := prevState.RerouteStruct.GetRerouteMap(originatorId)
	if reroute != nil {
		chunkId = reroute[len(reroute)-1]
		responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	}
	//if _, ok := prevState.RerouteMap[originatorId]; ok {
	//	chunkId = prevState.RerouteMap[originatorId][len(prevState.RerouteMap[originatorId])-1]
	//	responsibleNodes = prevState.Graph.FindResponsibleNodes(chunkId)
	//}

	request := Request{OriginatorId: originatorId, ChunkId: chunkId, RespNodes: responsibleNodes}

	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.Graph, prevState.RerouteStruct, prevState.CacheStruct)

	return found, route, thresholdFailed, accessFailed, paymentsList
}
