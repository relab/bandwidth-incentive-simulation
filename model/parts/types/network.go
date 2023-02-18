package types

import (
	"encoding/json"
	"fmt"
	. "go-incentive-simulation/model/general"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"time"
)

type Network struct {
	Bits     int
	Bin      int
	NodesMap map[int]*Node
}

type Node struct {
	Network    *Network
	Id         int
	AdjIds     [][]int
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
	} `json:"Nodes"`
}

func (network *Network) Load(path string) (int, int, map[int]*Node) {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	decoder := json.NewDecoder(file)

	var test Test
	err := decoder.Decode(&test)
	if err != nil {
		return 0, 0, nil
	}

	network.Bits = test.Bits
	network.Bin = test.Bin
	network.NodesMap = make(map[int]*Node)

	for _, node := range test.Nodes {
		node1 := network.node(node.ID)
		sort.Ints(node.Adj)
		for _, adj := range node.Adj {
			node2 := network.node(adj)
			node1.add(node2)
		}
	}

	return network.Bits, network.Bin, network.NodesMap
}

func (network *Network) node(value int) *Node {
	if value < 0 || value >= (1<<network.Bits) {
		panic("address out of range")
	}
	res := Node{
		Network:    network,
		Id:         value,
		AdjIds:     make([][]int, network.Bits),
		storageSet: []int{0},
		cacheSet:   []int{0},
		canPay:     true,
	}
	if len(network.NodesMap) == 0 {
		network.NodesMap = make(map[int]*Node)
	}
	if _, ok := network.NodesMap[value]; !ok {
		network.NodesMap[value] = &res
		return &res
	}
	return network.NodesMap[value]

}

func (network *Network) Generate(count int) []*Node {
	nodeIds := generateIds(count, (1<<network.Bits)-1)
	nodes := make([]*Node, 0)
	for _, nodeId := range nodeIds {
		node := network.node(nodeId)
		nodes = append(nodes, node)
	}
	fmt.Println("NodesMap:", nodes)
	pairs := make([][2]*Node, 0)
	for i, node1 := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			node2 := nodes[j]
			pairs = append(pairs, [2]*Node{node1, node2})
		}
	}
	shufflePairs(pairs)
	for _, nodes := range pairs {
		nodes[0].add(nodes[1])
	}
	return nodes
}

func (network *Network) Dump(path string) error {
	type NetworkData struct {
		Bits  int `json:"bits"`
		Bin   int `json:"bin"`
		Nodes []struct {
			ID  int   `json:"id"`
			Adj []int `json:"adj"`
		} `json:"nodes"`
	}
	data := NetworkData{network.Bits, network.Bin, make([]struct {
		ID  int   `json:"id"`
		Adj []int `json:"adj"`
	}, 0)}
	for _, node := range network.NodesMap {
		var result []int
		for _, list := range node.AdjIds {
			result = append(result, list...)
		}
		data.Nodes = append(data.Nodes, struct {
			ID  int   `json:"id"`
			Adj []int `json:"adj"`
		}{node.Id, result})
	}
	file, _ := json.Marshal(data)
	err := ioutil.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}
	return nil

}

func shufflePairs(pairs [][2]*Node) {
	rand.Shuffle(len(pairs), func(i, j int) {
		pairs[i], pairs[j] = pairs[j], pairs[i]
	})
}

func generateIds(totalNumbers int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	generatedNumbers := make(map[int]bool)
	for len(generatedNumbers) < totalNumbers {
		num := rand.Intn(maxValue + 1)
		generatedNumbers[num] = true
	}

	result := make([]int, 0, totalNumbers)
	for num := range generatedNumbers {
		result = append(result, num)
	}
	return result
}

func (node *Node) add(other *Node) bool {
	if node.Network == nil || node.Network != other.Network || node == other {
		return false
	}
	if node.AdjIds == nil {
		node.AdjIds = make([][]int, node.Network.Bits)
	}
	if other.AdjIds == nil {
		other.AdjIds = make([][]int, other.Network.Bits)
	}
	bit := node.Network.Bits - BitLength(node.Id^other.Id)
	if bit < 0 || bit >= node.Network.Bits {
		return false
	}
	isDup := Contains(node.AdjIds[bit], other.Id) || Contains(other.AdjIds[bit], node.Id)
	if len(node.AdjIds[bit]) < node.Network.Bin && len(other.AdjIds[bit]) < node.Network.Bin && !isDup {
		node.AdjIds[bit] = append(node.AdjIds[bit], other.Id)
		other.AdjIds[bit] = append(other.AdjIds[bit], node.Id)
		return true
	}
	return false
}

func (node *Node) IsNil() bool {
	return node.Id == 0
}
