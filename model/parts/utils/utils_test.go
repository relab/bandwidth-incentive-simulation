package utils

import (
	"fmt"
	"testing"
	ct "go-incentive-simulation/model"
	"gotest.tools/assert"
)

func TestCreateGraphNetwork(t *testing.T) {
	// fileName := "input_test.txt"
	fileName := "nodes_data_8_10000.txt"

	graph := CreateGraphNetwork(fileName)
	for i, _ := range graph.edges {
		for _, edge := range graph.edges[i]{
			fmt.Print(edge)
			fmt.Print("\n")
		}
	}
}

func TestCreateDowloaderList(t *testing.T) {
	// Create a network
	network := Network{}
	// Load data to network 
	network.load("nodes_data_8_10000.txt")

	// Get number of originators used in the func
	c := ct.Constants.GetOriginators()

	// Create a list of downloaders
	l:= network.CreateDowloadersList()

	// Check if the length of the list is equal to the number of originators specified
	assert.Equal(t, len(l), c)
}
