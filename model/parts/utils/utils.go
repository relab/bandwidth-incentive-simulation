package utils

import (
	"fmt"
)

func CreateGraphNetwork(filename string) {
	fmt.Println("Creating graph network...")
	graph := new(Graph)
	net := new(Network)
	_,_, nodes := net.load(filename)
	for _, node := range nodes {
		graph.AddNode(node.id)
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
}

// package utils

// import (
// 	"gonum.org/v1/gonum/graph"
// 	"gonum.org/v1/gonum/graph/simple"
// )

// type Edge struct {
// 	from, to int
// 	a2b      float64
// 	last     float64
// }

// func (e Edge) From() graph.Node {
// 	return simple.Node(e.from)
// }

// func (e Edge) To() graph.Node {
// 	return simple.Node(e.to)
// }

// func createGraph(filename string) *simple.DirectedGraph {
// 	g := simple.NewDirectedGraph()
// 	// load network from file
// 	// net := Network.Load(filename)
// 	// nodes := []Node{net.nodes.values()}
// 	nodes := []Node{}

// 	for _, node := range nodes {
// 		graph.AddNode(node.id)
// 		graph.NewNode

// 	nodes := []Node{}
// 	for _, node := range nodes {
// 		g.AddNode(simple.Node(node.ID))
// 		for _, adj := range node.adj {
// 			e := Edge{from: node.id, to: adj.id, a2b: 0, last: 0}
// 			g.SetEdge(e)
// 		}
// 	}
// 	return g
// }
