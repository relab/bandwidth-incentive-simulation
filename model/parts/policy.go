package policy

import (
	"sort"
	ut "go-incentive-simulation/model/parts/utils"
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

func (prevState *State) sendRequest() {
	

}
