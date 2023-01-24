package utils

import (
	. "go-incentive-simulation/model/constants"
	. "go-incentive-simulation/model/parts/types"
	"gotest.tools/assert"
	"testing"
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
	l := CreateDowloadersList(network)

	// Check if the length of the list is equal to the number of originators specified
	assert.Equal(t, len(l), c)
}

func TestGetNext(t *testing.T) {

}
