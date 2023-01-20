package utils

import (
	"fmt"
	ct "go-incentive-simulation/model"
	"math/rand"
	"time"
)

func CreateGraphNetwork(filename string) *Graph {
	fmt.Println("Creating graph network...")
	graph := &Graph{
		edges: make(map[int][]*Edge),
	}
	net := new(Network)
	_, _, nodes := net.load(filename)
	for _, node := range nodes {
		graph.AddNode(node)
	}
	for _, node := range graph.Nodes() {
		nodeAdj := node.adj
		for _, adjItems := range nodeAdj {
			for _, item := range adjItems {
				// "a2b" show how much this node asked from other node,
				// "last" is for the last forgiveness time
				attrs := EdgeAttrs{a2b: 0, last: 0}
				edge := Edge{fromNodeId: node.id, toNodeId: item.id, attrs: attrs}
				graph.AddEdge(&edge)
				// graph.SetEdgeAttributes()
			}
		}
	}
	fmt.Println("Graph network is created.")
	return graph
}

func choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < k; i++ {
		res = append(res, nodes[rand.Intn(len(nodes))])
	}
	return res
}

func MakeFiles() []int {
	fmt.Println("Making files...")
	var filesList []int

	for i := 0; i <= ct.Constants.GetOriginators(); i++ {
		// TODO: fix this, GetChuncks should be a list?
		// chunksList := choice(ct.Constants.GetChunks(), ct.Constants.GetRangeAddress())
		// filesList = append(chunksList)
	}
	fmt.Println("Files made!")
	return filesList
}

func (net *Network) CreateDowloadersList() []int {
	fmt.Println("Creating downloaders list...")

	nodesValue := make([]int, 0, len(net.nodes))
	for i := range net.nodes {
		nodesValue = append(nodesValue, net.nodes[i].id)
	}
	downloadersList := choice(nodesValue, ct.Constants.GetOriginators())

	fmt.Println("Downloaders list create...!")
	return downloadersList
}

// no need for this function
func (net *Network) PushSync(fileName string, files []string) {
	fmt.Println("Pushing sync...")
	if net == nil {
		fmt.Println("Network is nil!")
		return
	}
	nodes := net.nodes
	for i := range nodes {
		fmt.Println(nodes[i].id)
	}

	fmt.Println("Pushing sync finished...")
}
