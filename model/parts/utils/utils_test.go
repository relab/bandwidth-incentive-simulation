package utils

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"testing"
	"time"

	"gotest.tools/assert"
)

const path = "../../../data/nodes_data_8_10000_0.txt"

func TestCreateGraphNetwork(t *testing.T) {
	// fileName := "input_test.txt"
	network := &types.Network{}
	network.Load(path)
	graph, err := CreateGraphNetwork(network)

	edge := graph.GetEdge(49584, 0)
	graph.LockEdge(49584, 0)
	graph.UnlockEdge(49584, 0)
	node := graph.GetNode(0)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(graph.Nodes), 10000)
	assert.Equal(t, len(graph.Edges), 10000)
	assert.Check(t, *edge != types.Edge{})
	assert.Check(t, node != nil)
}

// TODO: not working right now
func TestCreateDowloaderList(t *testing.T) {
	// Create a network
	network := &types.Network{}
	// Load data to network
	network.Load(path)
	// Creates graph
	graph, _ := CreateGraphNetwork(network)
	// Get number of originators used in the func
	c := config.GetOriginators()

	// Create a list of downloaders
	l := CreateDownloadersList(graph)

	// Check if the length of the list is equal to the number of originators specified
	assert.Equal(t, len(l), c)
}

func TestIsThresholdFailed(t *testing.T) {

	// firstNodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// secondsNodes := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// chunkIds := []int{21, 22, 23, 24, 25, 26, 27, 28, 29, 30}

	// graph := Graph{}
}

func TestGetNext(t *testing.T) {

}

func TestPrecomputeRespNodes(t *testing.T) {
	network := &types.Network{}
	network.Load(path)
	sortedNodeIds := SortedKeys(network.NodesMap)
	loops := 10
	start := time.Now()
	for i := 0; i < loops; i++ {
		_ = PrecomputeRespNodes(sortedNodeIds)
	}
	fmt.Println(time.Since(start).Seconds())
}
