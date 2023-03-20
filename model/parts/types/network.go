package types

import (
	"encoding/json"
	"go-incentive-simulation/model/general"
	"math/rand"
	"os"
	"sort"
	"sync"
	"time"
)

type Network struct {
	Bits     int
	Bin      int
	NodesMap map[NodeId]*Node
}

type NodeId int32

func (n NodeId) ToInt32() int32 {
	return int32(n)
}

type ChunkId int32

func (c ChunkId) ToInt32() int32 {
	return int32(c)
}

type Node struct {
	Network    *Network
	Id         NodeId
	AdjIds     [][]NodeId
	CacheMap   map[ChunkId]int // int = counter
	Mutex      *sync.Mutex
	storageSet []int
	cacheSet   []int
	canPay     bool
}

type jsonFormat struct {
	Bits  int `json:"bits"`
	Bin   int `json:"bin"`
	Nodes []struct {
		Id  int   `json:"id"`
		Adj []int `json:"adj"`
	} `json:"Nodes"`
}

func (network *Network) Load(path string) (int, int, map[NodeId]*Node) {
	file, _ := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)
	decoder := json.NewDecoder(file)

	var test jsonFormat
	err := decoder.Decode(&test)
	if err != nil {
		return 0, 0, nil
	}

	network.Bits = test.Bits
	network.Bin = test.Bin
	network.NodesMap = make(map[NodeId]*Node)

	for _, node := range test.Nodes {
		node1 := network.node(NodeId(node.Id))
		sort.Ints(node.Adj)
		for _, adj := range node.Adj {
			node2 := network.node(NodeId(adj))
			node1.add(node2)
		}
	}

	return network.Bits, network.Bin, network.NodesMap
}

func (network *Network) node(nodeId NodeId) *Node {
	if nodeId < 0 || nodeId >= (1<<network.Bits) {
		panic("address out of range")
	}
	res := Node{
		Network:    network,
		Id:         nodeId,
		AdjIds:     make([][]NodeId, network.Bits),
		CacheMap:   make(map[ChunkId]int),
		Mutex:      &sync.Mutex{},
		storageSet: []int{0},
		cacheSet:   []int{0},
		canPay:     true,
	}
	if len(network.NodesMap) == 0 {
		network.NodesMap = make(map[NodeId]*Node)
	}
	if _, ok := network.NodesMap[nodeId]; !ok {
		network.NodesMap[nodeId] = &res
		return &res
	}
	return network.NodesMap[nodeId]

}

func (network *Network) Generate(count int) []*Node {
	nodeIds := generateIds(count, (1<<network.Bits)-1)
	nodes := make([]*Node, 0)
	for _, i := range nodeIds {
		node := network.node(NodeId(i))
		nodes = append(nodes, node)
	}
	pairs := make([][2]*Node, 0)
	for i, node1 := range nodes {
		for j := i + 1; j < len(nodes); j++ {
			node2 := nodes[j]
			pairs = append(pairs, [2]*Node{node1, node2})
		}
	}
	shufflePairs(pairs)
	for _, pair := range pairs {
		pair[0].add(pair[1])
	}
	return nodes
}

func (network *Network) Dump(path string) error {
	type NetworkData struct {
		Bits  int `json:"bits"`
		Bin   int `json:"bin"`
		Nodes []struct {
			Id  int   `json:"id"`
			Adj []int `json:"adj"`
		} `json:"nodes"`
	}
	data := NetworkData{network.Bits, network.Bin, make([]struct {
		Id  int   `json:"id"`
		Adj []int `json:"adj"`
	}, 0)}
	for _, node := range network.NodesMap {
		var result []int
		for _, list := range node.AdjIds {
			for _, ele := range list {
				result = append(result, int(ele))
			}
			//result = append(result, list...)
		}
		data.Nodes = append(data.Nodes, struct {
			Id  int   `json:"id"`
			Adj []int `json:"adj"`
		}{Id: int(node.Id), Adj: result})
	}
	file, _ := json.Marshal(data)
	err := os.WriteFile(path, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Choice(nodes []NodeId, k int) []NodeId {
	res := make([]NodeId, k)

	var val int
	for i := 0; i < k; i++ {
		val = rand.Intn(len(nodes)) - 1
		res[i] = nodes[val]
	}
	return res
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
		node.AdjIds = make([][]NodeId, node.Network.Bits)
	}
	if other.AdjIds == nil {
		other.AdjIds = make([][]NodeId, other.Network.Bits)
	}
	bit := node.Network.Bits - general.BitLength(node.Id.ToInt32()^other.Id.ToInt32())
	if bit < 0 || bit >= node.Network.Bits {
		return false
	}
	isDup := general.Contains(node.AdjIds[bit], other.Id) || general.Contains(other.AdjIds[bit], node.Id)
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
