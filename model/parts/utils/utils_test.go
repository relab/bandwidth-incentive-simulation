package utils

import (
	"testing"
)

func TestNetwork(t *testing.T) {
	network := Network{}
	bits, bin, nodes := network.load("input_test.txt")
	t.Log("Bits:", bits)
	t.Log("Bin:", bin)
	t.Log("Nodes:", nodes)

	t.Log("Nodes[1]:", nodes[1])
}

func TestPushSync(t *testing.T) {
	network := Network{}
	network.load("input_test.txt")
	network.PushSync("input_test.txt", []string{"file1", "file2"})
}
