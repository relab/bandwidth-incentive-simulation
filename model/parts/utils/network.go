package utils

import (
	"encoding/json"
	"os"
)

type Network struct {
	bits  int
	bin   int
	nodes map[int]Node
}

type Node struct {
	network    *Network
	id         int
	adj        [][]int
	storageSet []int
	cacheSet   []int
	canPay     bool
}

type Test struct {
	Bits  int `json:"bits"`
	Bin   int `json:"bin"`
	Nodes []struct {
		ID  int   `json:"id"`
		Adj []int `json:"adj"`
	} `json:"nodes"`
}

func (network *Network) load(path string) (int, int, map[int]Node) {
	file, _ := os.Open(path)
	defer file.Close()
	decoder := json.NewDecoder(file)
	var test Test
	decoder.Decode(&test)

	network.bits = test.Bits
	network.bin = test.Bin
	nodes := make(map[int][]int)
	for _, node := range test.Nodes {
		nodes[node.ID] = node.Adj
	}
	//network.nodes = nodes
	network.nodes = make(map[int]Node)
	for _, data := range test.Nodes {
		network.node(data.ID)
	}

	return network.bits, network.bin, network.nodes
}

func (network *Network) node(value int) Node {
	if value < 0 || value >= (1<<network.bits) {
		panic("address out of range")
	}
	res := Node{
		network:    network,
		id:         value,
		adj:        make([][]int, 0),
		storageSet: make([]int, 0),
		cacheSet:   make([]int, 0),
		canPay:     true,
	}
	if _, ok := network.nodes[value]; !ok {
		network.nodes[value] = res
		return res
	}
	return network.nodes[value]

}

// func (node *Node) add(other *Node) bool {
// 	if node.network == nil || node.network != other.network || node == other {
// 		return false
// 	}
// 	bit := node.network.bits - (node.id ^ other.id).bitLen()	
// }

// func BitLen(x int) int {
// 	bitLen := 0
// 	for n > 0 {
// 		n >>= 1
// 		bitLen++
// 	}
// }