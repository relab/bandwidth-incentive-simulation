package policy

import (
	"fmt"
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

func (prevState *State) SendRequest() {
	// chunkAddr := g.Choice(ct.Constants.GetRangeAddress(), 1) //GetRangeAddress er int? no need for choice
	// chunkAddr := ct.Constants.GetRangeAddress()
	random := []int{}

	rand.Seed(time.Now().UnixNano())
	if ct.Constants.IsCacheEnabled() == true {
		random = append(random, rand.Intn(1-0) + 0)
		fmt.Println(random)
		if float32(random[0]) < 0.5 {
			c := rand.Intn(1000-0) + 0
			fmt.Println(c)
		}
	}
}
