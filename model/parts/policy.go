package policy

import (
	ct "go-incentive-simulation/model"
	ut "go-incentive-simulation/model/parts/utils"
	"math/rand"
	"sort"
	"time"
	// g "go-incentive-simulation/model/general"
)

type State struct {
	network                 *ut.Graph
	originators             []int
	originatorsIndex        int
	nodesId                 []int
	routeList               []int
	pendingDict             map[int]int
	rerouteDict             map[int]int
	cacheDict               map[int]int
	originatorIndex         int
	successfulFound         int
	failedRequestsThreshold int
	failedRequestsAccess    int
	timeStep                int
}

func findResponisbleNodes(nodesId []int, chunkAdd int) []int {
	v := []int{}
	for i := range nodesId {
		v = append(v, nodesId[i]^chunkAdd)
	}
	sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })

	return v[:4]
}

func (prevState *State) SendRequest() map[string]int {
	random := []int{}
	chunkId := ct.Constants.GetRangeAddress()

	rand.Seed(time.Now().UnixNano())
	if ct.Constants.IsCacheEnabled() == true {
		random = append(random, rand.Intn(1-0) + 0)
		if float32(random[0]) < 0.5 {
			chunkId = rand.Intn(1000-0) + 0
		} else {
			chunkId = rand.Intn(1000-ct.Constants.GetRangeAddress()) + ct.Constants.GetRangeAddress()
		}
	}
	responisbleNodes := findResponisbleNodes(prevState.nodesId, chunkId)
	originator := prevState.originators[prevState.originatorIndex]

	for _, value := range prevState.pendingDict {
		if originator == value {
			chunkId = prevState.pendingDict[originator]
			responisbleNodes = findResponisbleNodes(prevState.nodesId, chunkId)
		}
	}

	for _, value := range prevState.rerouteDict {
		if originator == value {
			chunkId = prevState.rerouteDict[originator]
			responisbleNodes = findResponisbleNodes(prevState.nodesId, chunkId)
		}
	}

	request := []int{originator, chunkId}

	found, route, thresholdFailed, accessFailed, paymentsList := ut.ConsumeTask(request, prevState.network, responisbleNodes, prevState.rerouteDict, prevState.cacheDict)
	
	res := map[string]int{
		"found": found,
		"route": route,
		"thresholdFailed": thresholdFailed,
		"originatorIndex": prevState.originatorIndex,
		"accessFailed": accessFailed,
		"paymentsList": paymentsList,
	}
	return res
}
