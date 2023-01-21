package utils

import (
	"fmt"
)

// Graph structure, node Ids in array and edges in map
type Graph struct {
	nodes []*Node
	edges map[int][]*Edge
}

// Edge that connects to nodes with attributes about the connection
type Edge struct {
	fromNodeId int
	toNodeId   int
	attrs      EdgeAttrs
}

// EdgeAttrs Edge attributes structure,
// "a2b" show how much this node asked from other node,
// "last" is for the last forgiveness time
type EdgeAttrs struct {
	a2b  int
	last int
}

// Nodes Returns all nodes
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// AddNode will add a Node to a graph
func (g *Graph) AddNode(node *Node) error {
	if containsNode(g.nodes, node) {
		err := fmt.Errorf("node %d already exists", node.id)
		return err
	} else {
		g.nodes = append(g.nodes, node)
		return nil
	}
}

// AddEdge will add an edge from a node to a node
func (g *Graph) AddEdge(edge *Edge) error {
	toNode := g.getNode(edge.toNodeId)
	fromNode := g.getNode(edge.fromNodeId)
	if toNode == nil || fromNode == nil {
		return fmt.Errorf("not a valid edge from %d ---> %d", fromNode.id, toNode.id)
	} else if containsEdge(g.edges[fromNode.id], edge) {
		return fmt.Errorf("edge from node %d ---> %d already exists", fromNode.id, toNode.id)
	} else {
		newEdges := append(g.edges[fromNode.id], edge)
		g.edges[fromNode.id] = newEdges
		return nil
	}
}
func (g *Graph) GetEdgeData(fromNode *Node, toNode *Node) EdgeAttrs {
	return g.edges[fromNode.id][toNode.id].attrs
}

// getNode will return a node point if exists or return nil
func (g *Graph) getNode(nodeId int) *Node {
	for i, v := range g.nodes {
		if v.id == nodeId {
			return g.nodes[i]
		}
	}
	return nil
}

func containsNode(v []*Node, node *Node) bool {
	for _, v := range v {
		if v.id == node.id {
			return true
		}
	}
	return false
}

func containsEdge(v []*Edge, edge *Edge) bool {
	for _, v := range v {
		if v.toNodeId == edge.toNodeId {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.nodes {
		fmt.Printf("%d : ", v.id)
		for _, i := range v.adj {
			for _, v := range i {
				fmt.Printf("%d ", v.id)
			}
		}
		fmt.Println()
	}
}
