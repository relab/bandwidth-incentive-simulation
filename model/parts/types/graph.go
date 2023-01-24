package types

import (
	"fmt"
)

// Graph structure, node Ids in array and edges in map
type Graph struct {
	nodes []*Node
	Edges map[int][]*Edge
}

// Edge that connects to nodes with attributes about the connection
type Edge struct {
	FromNodeId int
	ToNodeId int
	Attrs    EdgeAttrs
}

// EdgeAttrs Edge attributes structure,
// "a2b" show how much this node asked from other node,
// "last" is for the last forgiveness time
type EdgeAttrs struct {
	A2b  int
	Last int
}

// Nodes Returns all nodes
func (g *Graph) Nodes() []*Node {
	return g.nodes
}

// AddNode will add a Node to a graph
func (g *Graph) AddNode(node *Node) error {
	if ContainsNode(g.nodes, node) {
		err := fmt.Errorf("node %d already exists", node.Id)
		return err
	} else {
		g.nodes = append(g.nodes, node)
		return nil
	}
}

// AddEdge will add an edge from a node to a node
func (g *Graph) AddEdge(edge *Edge) error {
	toNode := g.getNode(edge.ToNodeId)
	fromNode := g.getNode(edge.FromNodeId)
	if toNode == nil || fromNode == nil {
		return fmt.Errorf("not a valid edge from %d ---> %d", fromNode.Id, toNode.Id)
	} else if containsEdge(g.Edges[fromNode.Id], edge) {
		return fmt.Errorf("edge from node %d ---> %d already exists", fromNode.Id, toNode.Id)
	} else {
		newEdges := append(g.Edges[fromNode.Id], edge)
		g.Edges[fromNode.Id] = newEdges
		return nil
	}
}

func (g *Graph) GetEdgeData(fromNodeId int, toNodeId int) EdgeAttrs {
	return g.Edges[fromNodeId][toNodeId].Attrs
}

// getNode will return a node point if exists or return nil
func (g *Graph) getNode(nodeId int) *Node {
	for i, v := range g.nodes {
		if v.Id == nodeId {
			return g.nodes[i]
		}
	}
	return nil
}

func ContainsNode(v []*Node, node *Node) bool {
	for _, v := range v {
		if v.Id == node.Id {
			return true
		}
	}
	return false
}

func containsEdge(v []*Edge, edge *Edge) bool {
	for _, v := range v {
		if v.ToNodeId == edge.ToNodeId {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, v := range g.nodes {
		fmt.Printf("%d : ", v.Id)
		for _, i := range v.Adj {
			for _, v := range i {
				fmt.Printf("%d ", v.Id)
			}
		}
		fmt.Println()
	}
}
