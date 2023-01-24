package policy

import (
	. "go-incentive-simulation/model/parts/utils"
	"gotest.tools/assert"
	"testing"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{64132, 49693, 45280, 42779, 41852, 43812, 47987, 43377, 41471}
	chunkAdd := 11
	values := findResponisbleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}

//func TestSendRequest(t *testing.T) {
//	state := State{}
//	state.nodesId = []int{64132, 49693, 45280, 42779, 41852, 43812, 47987, 43377, 41471}
//
//	state.originators = Node
//	state.originatorIndex = 2
//	state.SendRequest()
//}
