package policy

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	. "go-incentive-simulation/model/variables"
	"math/rand"
	"sort"
	"time"
)

type Response struct {
	found           bool
	route           Route
	thresholdFailed [][]Threshold
	accessFailed    bool
	paymentsList    []Payment
}

func findResponisbleNodes(nodesId []int, chunkAdd int) []int {
	var v []int
	for i := range nodesId {
		v = append(v, nodesId[i]^chunkAdd)
	}
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
	return v[:4]
}

func SendRequest(prevState *State) (bool, Route, [][]Threshold, bool, []Payment) {
	rand.Seed(time.Now().UnixNano())

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
	responsibleNodes := findResponisbleNodes(prevState.NodesId, chunkId)
	originator := prevState.Originators[prevState.OriginatorIndex]

	if _, ok := prevState.PendingMap[originator]; ok {
		chunkId = prevState.PendingMap[originator]
		responsibleNodes = findResponisbleNodes(prevState.NodesId, chunkId)
	}
	if _, ok := prevState.RerouteMap[originator]; ok {
		chunkId = prevState.RerouteMap[originator][len(prevState.RerouteMap[originator])-1]
		responsibleNodes = findResponisbleNodes(prevState.NodesId, chunkId)
	}

	getNode := GetNodeById(originator)

	request := Request{Originator: getNode, ChunkId: chunkId}

	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.Network, responsibleNodes, prevState.RerouteMap, prevState.CacheListMap)

	return found, route, thresholdFailed, accessFailed, paymentsList
}
