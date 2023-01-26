package utils

import (
	. "go-incentive-simulation/model/parts/types"
	. "go-incentive-simulation/model/variables"
	"testing"

	"gotest.tools/assert"
)

func TestCreateGraphNetwork(t *testing.T) {
	// fileName := "input_test.txt"
	fileName := "../../../data/nodes_data_8_10000.txt"

	graph, err := CreateGraphNetwork(fileName)
	/*	for i, _ := range graph.edges {
		for _, edge := range graph.edges[i] {
			fmt.Print(edge)
			fmt.Print("\n")
		}
	}*/
	assert.Equal(t, err, nil)
	assert.Equal(t, len(graph.Nodes), 10000)
	assert.Equal(t, len(graph.Edges), 10000)
}

// TODO: not working right now
func TestCreateDowloaderList(t *testing.T) {
	// Create a network
	network := &Network{}
	// Load data to network
	network.Load("../../../data/nodes_data_8_10000.txt")

	// Get number of originators used in the func
	c := Constants.GetOriginators()

	// Create a list of downloaders
	l := CreateDownloadersList(network)

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
