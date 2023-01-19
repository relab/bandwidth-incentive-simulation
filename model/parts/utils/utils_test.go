package utils

import (
	"fmt"
	"testing"
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

func TestPushSync(t *testing.T) {
	network := Network{}
	network.load("input_test.txt")
	network.PushSync("input_test.txt", []string{"file1", "file2"})
}
