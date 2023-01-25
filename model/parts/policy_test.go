package policy

import (
	. "go-incentive-simulation/model/parts/types"
	"testing"

	"gotest.tools/assert"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{64132, 49693, 45280, 42779, 41852, 43812, 47987, 43377, 41471}
	chunkAdd := 11
	values := findResponsibleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}

func TestGetNodeById(t *testing.T) {
	n := Network{}
	n.Bin = 8
	n.Bits = 16

	n1 := Node{}
	n1.Id = 111

	n2 := Node{}
	n2.Id = 222

	//GetNodeById(111)
}
