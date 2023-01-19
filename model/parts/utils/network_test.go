package utils

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	path := "nodes_data_8_10000.txt"
	network := Network{}
	bits, bin, nodes := network.load(path)

	t.Log("Bits:", bits)
	t.Log("Bin:", bin)
	//print the nodes map
	for k, v := range nodes {
		t.Log("Nodes:", k, *v)
	}

	t.Log("Nodes[12381]:", *nodes[12381])

	for _, bucket := range nodes[12381].adj {
		for _, node := range bucket {
			t.Log("Nodes[12381].adj:", node.id)
		}
		t.Log("\n")
	}
}
