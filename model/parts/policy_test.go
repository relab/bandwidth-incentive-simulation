package policy

import (
	"testing"
	"gotest.tools/assert"
)

func TestResponisbleNodes(t *testing.T) {
	nodesId := []int{64132, 49693, 45280, 42779, 41852, 43812, 47987, 43377, 41471}
	chunkAdd := 11
	values := findResponisbleNodes(nodesId, chunkAdd)

	assert.Equal(t, len(values), 4)
}

// func TestSendRequest(t *testing.T) {
// 	path := "../../../data/nodes_data_8_10000.txt"
// 	network := Network{}
// 	bits, bin, nodes := network.Load(path)
// 	fmt.Println("Bits:", bits)
// 	fmt.Println("Bin:", bin)


// 	state := State{}


// 	// state.originators = Node
// 	// state.originatorIndex = 2
// 	// state.SendRequest()
// }
