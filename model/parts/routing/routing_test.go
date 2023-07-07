package routing

import (
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"go-incentive-simulation/model/parts/utils"
	"testing"

	"gotest.tools/assert"
)

const path = "../../../network_data/nodes_data_b16_k16_10000_.txt"

// TODO: there should be a complete test for the FindRoute and getNext functions.
func TestFindRoute(t *testing.T) {
	config.InitConfig()
	network := &types.Network{}
	network.Load(path)
	graph, err := utils.CreateGraphNetwork(network)
	if err != nil {
		panic(err)
	}

	var request types.Request
	request.OriginatorId = 56914
	request.ChunkId = 18239
	route, _, found, _, _, _ := FindRoute(request, graph)
	expectedRoute := []types.NodeId{56914, 18005, 18303, 18232}
	// expectedPayments := []types.Payment{}
	expectedFound := true
	// expectedFlag2 := false
	// expectedFlag3 := false
	// expectedFlag4 := false
	assert.DeepEqual(t, expectedRoute, route)
	// assert.DeepEqual(t, expectedPayments, paymentList)
	assert.Equal(t, expectedFound, found)
	// assert.Equal(t, expectedFlag2, accessFailed)
	// assert.Equal(t, expectedFlag3, thresholdFailed)
	// assert.Equal(t, expectedFlag4, foundByCaching)

	request.OriginatorId = 9443
	request.ChunkId = 12040
	route, _, found, _, _, _ = FindRoute(request, graph)
	expectedRoute = []types.NodeId{9443, 12124, 12046}
	expectedFound = true
	assert.DeepEqual(t, expectedRoute, route)
	assert.Equal(t, expectedFound, found)

	request.OriginatorId = 38801
	request.ChunkId = 49401
	route, _, found, _, _, _ = FindRoute(request, graph)
	expectedRoute = []types.NodeId{38801, 49917, 49373, 49403}
	expectedFound = true
	assert.DeepEqual(t, expectedRoute, route)
	assert.Equal(t, expectedFound, found)
}

func TestGetNext(t *testing.T) {

}
