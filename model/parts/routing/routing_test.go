package routing

import (
	"testing"
)

const path = "../../network_data/nodes_data_b16_k16_10000_.txt"

func TestFindRoute(t *testing.T) {
	// network := &types.Network{}
	// network.Load(path)
	// graph, err := utils.CreateGraphNetwork(network)
	// if err != nil {
	// 	panic(err)
	// }

	// var request types.Request
	// request.OriginatorId = 59104
	// request.ChunkId = 2164

	// route, payments, flag1, flag2, flag3, flag4 := FindRoute(request, graph)

	// expectedRoute := []types.NodeId{1, 2, 3, 4}
	// expectedPayments := []types.Payment{}
	// expectedFlag1 := false
	// expectedFlag2 := false
	// expectedFlag3 := false
	// expectedFlag4 := false

	// assert.Equal(t, expectedRoute, route)
	// assert.Equal(t, expectedPayments, payments)
	// assert.Equal(t, expectedFlag1, flag1)
	// assert.Equal(t, expectedFlag2, flag2)
	// assert.Equal(t, expectedFlag3, flag3)
	// assert.Equal(t, expectedFlag4, flag4)
}

func TestGetNext(t *testing.T) {

}
