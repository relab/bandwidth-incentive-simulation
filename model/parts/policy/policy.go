package policy

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/general"
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	"math/rand"
	"sort"
)

func findResponsibleNodes(nodesId []int, chunkAdd int) []int {
	var distances []int
	var distance int
	nodesMap := make(map[int]int)
	returnNodes := make([]int, 4)

	closestNodes := BinarySearchClosest(nodesId, chunkAdd, 10)

	for _, nodeId := range closestNodes {
		distance = nodeId ^ chunkAdd
		// fmt.Println(distance, nodeId)
		distances = append(distances, distance)
		nodesMap[distance] = nodeId
	}

	sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })

	for i := 0; i < 4; i++ {
		distance = distances[i]
		returnNodes[i] = nodesMap[distance]
	}
	return returnNodes
}

func SendRequest(prevState *State) (bool, Route, [][]Threshold, bool, []Payment) {
	// Gets one random chunkId from the range of addresses
	chunkId := rand.Intn(Constants.GetRangeAddress())
	var random int

	if Constants.IsCacheEnabled() == true {
		random = rand.Intn(1)
		if float32(random) < 0.5 {
			chunkId = rand.Intn(1000)
		} else {
			chunkId = rand.Intn(Constants.GetRangeAddress()-1000) + 0
		}
	}
	responsibleNodes := findResponsibleNodes(prevState.NodesId, chunkId)
	originator := prevState.Originators[prevState.OriginatorIndex]

	if _, ok := prevState.PendingMap[originator]; ok {
		chunkId = prevState.PendingMap[originator]
		responsibleNodes = findResponsibleNodes(prevState.NodesId, chunkId)
	}
	if _, ok := prevState.RerouteMap[originator]; ok {
		chunkId = prevState.RerouteMap[originator][len(prevState.RerouteMap[originator])-1]
		responsibleNodes = findResponsibleNodes(prevState.NodesId, chunkId)
	}

	originatorNode := prevState.Graph.GetNode(originator)

	request := Request{Originator: originatorNode, ChunkId: chunkId}

	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.Graph, responsibleNodes, prevState.RerouteMap, prevState.CacheListMap)

	return found, route, thresholdFailed, accessFailed, paymentsList
}
