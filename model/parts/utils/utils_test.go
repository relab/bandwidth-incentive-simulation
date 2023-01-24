package utils

import (
	ct "go-incentive-simulation/model/constants"
	"go-incentive-simulation/model/parts/types"
	"gotest.tools/assert"
	"testing"
)

func TestCreateGraphNetwork(t *testing.T) {
	// fileName := "input_test.txt"
	fileName := "nodes_data_8_10000.txt"

	graph, err := CreateGraphNetwork(fileName)
	/*	for i, _ := range graph.edges {
		for _, edge := range graph.edges[i] {
			fmt.Print(edge)
			fmt.Print("\n")
		}
	}*/
	assert.Equal(t, err, nil)
	assert.Equal(t, len(graph.nodes), 10000)
	assert.Equal(t, len(graph.edges), 10000)
}

func TestCreateDowloaderList(t *testing.T) {
	// Create a network
	network := types.Network{}
	// Load data to network
	network.load("nodes_data_8_10000.txt")

	// Get number of originators used in the func
	c := ct.Constants.GetOriginators()

	// Create a list of downloaders
	l := network.CreateDowloadersList()

	// Check if the length of the list is equal to the number of originators specified
	assert.Equal(t, len(l), c)
}

func TestChoice(t *testing.T) {
	// List of nodes
	nodes := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Originators
	k := 2

	c := choice(nodes, k)

	assert.Equal(t, len(c), k)
}

func TestGetNext(t *testing.T) {

}
