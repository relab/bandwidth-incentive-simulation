package utils

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	network := Network{}
	bits, bin, nodes := network.load("nodes_data_8_10000.txt")
	t.Log("Bits:", bits)
	t.Log("Bin:", bin)
	t.Log("Nodes:", nodes)

	t.Log("Nodes[1]:", nodes[1])
}
