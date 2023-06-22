package utils

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/parts/types"
	"math"
	"math/rand"
	"testing"
	"time"

	"gotest.tools/assert"
)

const path = "../../../data/nodes_data_16_10000_0.txt"

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
	edge = graph.GetEdge(27481, 46283)
	assert.Equal(t, edge.Attrs.A2B, 0)
	for node := range graph.Edges {
		for _, edge := range graph.Edges[node] {
			assert.Equal(t, edge.Attrs.A2B, 0)
		}
	}
}

func TestBinSize(t *testing.T) {
	// fileName := "input_test.txt"
	network := &types.Network{}
	network.Load(path)
	graph, err := CreateGraphNetwork(network)

	bins := graph.GetNodeAdj(graph.NodeIds[0])

	assert.Equal(t, err, nil)
	assert.Equal(t, len(bins[0]), network.Bin)
	assert.Equal(t, len(bins[1]), graph.Bin)
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

func TestTimePrecomputeRespNodes(t *testing.T) {
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

func TestDistributionRespNodeswithStorageDepth(t *testing.T) {
	network := &types.Network{}
	network.Load(path)
	addrRange := math.Pow(2, float64(network.Bits))
	sortedNodeIds := SortedKeys(network.NodesMap)
	fmt.Printf("First and last node %d, %d, length %d \n", sortedNodeIds[0], sortedNodeIds[len(sortedNodeIds)-1], len(sortedNodeIds))
	n := len(sortedNodeIds)
	depth := 0
	for n > 8 {
		n = n / 2
		depth++
	}
	hits := make([]int, 100)
	for i := 0; i < 100; i++ {
		chunkId := types.ChunkId(rand.Intn(int(addrRange)-1) + 1)
		for _, id := range sortedNodeIds {
			if getProximityChunk(id, chunkId) >= depth {
				hits[i]++
			}
		}
	}

	fmt.Printf("Depth %d gives responsible groups %v", depth, hits)
}

func TestGini(t *testing.T) {
	values := []int{4, 0, 0, 0}

	assert.Equal(t, Mean(values), float64(1))
	assert.Equal(t, Gini(values), 0.75)
}

func TestAdjustedRefreshrate(t *testing.T) {

	assert.Equal(t, GetAdjustedRefreshrate(15, 16, 8, 2), 8)
	assert.Equal(t, GetAdjustedRefreshrate(14, 16, 8, 2), 7)
	assert.Equal(t, GetAdjustedRefreshrate(13, 16, 8, 2), 6)
	assert.Equal(t, GetAdjustedRefreshrate(12, 16, 8, 2), 5)
	assert.Equal(t, GetAdjustedRefreshrate(11, 16, 8, 2), 4)
	assert.Equal(t, GetAdjustedRefreshrate(10, 16, 8, 2), 4)
	assert.Equal(t, GetAdjustedRefreshrate(9, 16, 8, 2), 3)
	assert.Equal(t, GetAdjustedRefreshrate(8, 16, 8, 2), 2)
	assert.Equal(t, GetAdjustedRefreshrate(7, 16, 8, 2), 2)
	assert.Equal(t, GetAdjustedRefreshrate(6, 16, 8, 2), 2)
	assert.Equal(t, GetAdjustedRefreshrate(5, 16, 8, 2), 1)
	assert.Equal(t, GetAdjustedRefreshrate(4, 16, 8, 2), 1)
	assert.Equal(t, GetAdjustedRefreshrate(3, 16, 8, 2), 1)

	assert.Equal(t, GetAdjustedRefreshrate(15, 16, 8, 3), 7)
	assert.Equal(t, GetAdjustedRefreshrate(14, 16, 8, 3), 6)
	assert.Equal(t, GetAdjustedRefreshrate(13, 16, 8, 3), 5)
	assert.Equal(t, GetAdjustedRefreshrate(12, 16, 8, 3), 4)
	assert.Equal(t, GetAdjustedRefreshrate(11, 16, 8, 3), 3)
	assert.Equal(t, GetAdjustedRefreshrate(10, 16, 8, 3), 2)
	assert.Equal(t, GetAdjustedRefreshrate(9, 16, 8, 3), 2)
	assert.Equal(t, GetAdjustedRefreshrate(8, 16, 8, 3), 1)
	assert.Equal(t, GetAdjustedRefreshrate(7, 16, 8, 3), 1)
	assert.Equal(t, GetAdjustedRefreshrate(3, 16, 8, 3), 1)
	assert.Equal(t, GetAdjustedRefreshrate(15, 16, 2, 3), 2)
	assert.Equal(t, GetAdjustedRefreshrate(14, 16, 2, 3), 2)
	assert.Equal(t, GetAdjustedRefreshrate(13, 16, 2, 3), 2)

}
