package utils

import (
	"fmt"
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

func isThresholdFailed(firstNode *Node, secondNode *Node, chunkId int, net *Network) {
	if ct.
}

func getNext(firstNode *Node, chunkId int, net *Network, mainOriginator *Node, prevNodePaid *Node, rerouteMap int) {
	var nextNode *Node
	var payNext *Node
	thresholdFailed := false
	// thresholdList := []
	accessFailed := false
	var payment int
	lastDistance := int(firstNode.id ^ chunkId)
	fmt.Println("last distance is : %d, chunk is: %d, first is: %d", lastDistance, chunkId, firstNode.id)
	fmt.Println("which bucket: %d", 16-BitLength(chunkId^firstNode.id))

	currDist := int(lastDistance)
	payDist := int(lastDistance)
	for _, adj := range firstNode.adj {
		fmt.Println("adj: %d", adj)
		for _, node := range adj {
			dist := int(node.id^chunkId)
			if BitLength(dist) >= BitLength(lastDistance) {
				continue
			}
			if !thresholdFailed(firstNode, node, chunkId, net)
		}
	}
}

func getBin(src int, dest int, index int) int {
	distance := src ^ dest
	result := index
	for distance > 0 {
		distance >>= 1
		result -= 1
	}
	return result
}

func whichPowerTwo(rangeAddress int) int {
	return BitLength(rangeAddress) - 1
}
