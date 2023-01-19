package utils

import (
	"encoding/json"
	"math"
	"os"
)

type Network struct {
	bits  int
	bin   int
	nodes map[int]*Node
}

type Node struct {
	network    *Network
	id         int
	adj        [][]*Node
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

func (network *Network) load(path string) (int, int, map[int]*Node) {
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
	network.nodes = make(map[int]*Node)
	for _, data := range test.Nodes {
		node1 := network.node(data.ID)
		for _, adj := range data.Adj {
			node2 := network.node(adj)
			node1.add(node2)
		}
	}

	return network.bits, network.bin, network.nodes
}

func (network *Network) node(value int) *Node {
	if value < 0 || value >= (1<<network.bits) {
		panic("address out of range")
	}
	res := Node{
		network:    network,
		id:         value,
		adj:        make([][]*Node, network.bits),
		storageSet: []int{0},
		cacheSet:   []int{0},
		canPay:     true,
	}
	if _, ok := network.nodes[value]; !ok {
		network.nodes[value] = &res
		return &res
	}
	return network.nodes[value]

}

func (network *Network) generate(count int) {
	// TODO: implement
}

func (node *Node) add(other *Node) bool {
	if node.network == nil || node.network != other.network || node == other {
		return false
	}
	if node.adj == nil {
		node.adj = make([][]*Node, node.network.bits)
	}
	if other.adj == nil {
		other.adj = make([][]*Node, other.network.bits)
	}
	bit := node.network.bits - int(math.Ceil(math.Log2(float64(node.id^other.id))))
	if bit < 0 || bit >= node.network.bits {
		return false
	}
	isDup := contains(node.adj[bit], other.id) || contains(other.adj[bit], node.id)
	if len(node.adj[bit]) < node.network.bin && len(other.adj[bit]) < node.network.bin && !isDup {
		node.adj[bit] = append(node.adj[bit], other)
		other.adj[bit] = append(other.adj[bit], node)
		return true
	}
	return false
}

func contains(slice []*Node, value int) bool {
	for _, item := range slice {
		if item.id == value {
			return true
		}
	}
	return false
}
