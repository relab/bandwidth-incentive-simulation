package policy

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/parts/utils"
	. "go-incentive-simulation/model/variables"
	"math/rand"
	"sort"
	"time"
	// g "go-incentive-simulation/model/general"
)

type State struct {
	network                 *Graph
	originators             []*Node
	originatorsIndex        int
	nodesId                 []int
	routeList               []int
	pendingMap              map[int]int
	rerouteMap              map[int]int
	cacheMap                map[int]int
	originatorIndex         int
	successfulFound         int
	failedRequestsThreshold int
	failedRequestsAccess    int
	timeStep                int
}

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

func (prevState *State) SendRequest() Response {
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
	responsibleNodes := findResponisbleNodes(prevState.nodesId, chunkId)
	originator := prevState.originators[prevState.originatorIndex]

	for _, v := range prevState.pendingMap {
		if originator.id == v {
			chunkId = prevState.pendingMap[originator]
			responsibleNodes = findResponisbleNodes(prevState.nodesId, chunkId)
		}
	}
	for _, v := range prevState.rerouteMap {
		if originator == v {
			chunkId = prevState.rerouteMap[originator]
			responsibleNodes = findResponisbleNodes(prevState.nodesId, chunkId)
		}
	}
	request := Request{Originator: originator, ChunkId: chunkId}

	found, route, thresholdFailed, accessFailed, paymentsList := ConsumeTask(&request, prevState.network, responsibleNodes, prevState.rerouteDict, prevState.cacheDict)

	res := Response{
		found:           found,
		route:           route,
		thresholdFailed: thresholdFailed,
		accessFailed:    accessFailed,
		paymentsList:    paymentsList,
	}
	return res
}
