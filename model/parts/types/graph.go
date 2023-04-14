package types

import (
	"fmt"
	"go-incentive-simulation/config"
	"go-incentive-simulation/model/general"
	"sort"
	"sync"
)

// Graph structure, node Ids in array and edges in map
type Graph struct {
	*Network
	CurState  State
	Nodes     []*Node
	NodeIds   []int
	Edges     map[int]map[int]*Edge
	RespNodes [][4]int
	Mutex     sync.Mutex
}

// Edge that connects to NodesMap with attributes about the connection
type Edge struct {
	FromNodeId int
	ToNodeId   int
	Attrs      EdgeAttrs
	Mutex      *sync.Mutex
}

// EdgeAttrs Edge attributes structure,
// "a2b" show how much this node asked from other node,
// "last" is for the last forgiveness time,
// "threshold" is for the adjustable threshold limit.
type EdgeAttrs struct {
	A2B               int
	Last              int
	EpokeLastForgiven int
	Threshold         int
}

//func (g *Graph) FindResponsibleNodes(chunkId int) [4]int {
//	return g.RespNodes[chunkId]
//}

func (g *Graph) FindResponsibleNodes(chunkId int) [4]int {
	if config.IsPrecomputeRespNodes() {
		return g.RespNodes[chunkId]

	} else {
		if g.RespNodes[chunkId][0] != 0 {
			return g.RespNodes[chunkId]

		} else {
			numNodesSearch := config.GetBits()
			closestNodes := general.BinarySearchClosest(g.NodeIds, chunkId, numNodesSearch)
			distances := make([]int, len(closestNodes))
			result := [4]int{}

			for i, nodeId := range closestNodes {
				distances[i] = nodeId ^ chunkId
			}

			sort.Slice(distances, func(i, j int) bool { return distances[i] < distances[j] })

			for i := 0; i < 4; i++ {
				result[i] = distances[i] ^ chunkId // this results in the nodeId again
			}
			g.RespNodes[chunkId] = result

			return result
		}
	}
}

// AddNode will add a Node to a graph
func (g *Graph) AddNode(node *Node) error {
	if ContainsNode(g.Nodes, node) {
		err := fmt.Errorf("node %d already exists", node.Id)
		return err
	} else {
		g.Nodes = append(g.Nodes, node)
		return nil
	}
}

func (g *Graph) GetNodeAdj(nodeId int) [][]int {
	return g.GetNode(nodeId).AdjIds
}

// AddEdge will add an edge from a node to a node
func (g *Graph) AddEdge(fromNodeId int, toNodeId int, attrs EdgeAttrs) error {
	toNode := g.GetNode(toNodeId)
	fromNode := g.GetNode(fromNodeId)
	if toNode == nil || fromNode == nil {
		return fmt.Errorf("not a valid edge from %d ---> %d", fromNode.Id, toNode.Id)
	} else if g.EdgeExists(fromNodeId, toNodeId) {
		//if ContainsEdge(g.Edges[edge.FromNodeId], edge) {
		return fmt.Errorf("edge from node %d ---> %d already exists", fromNodeId, toNodeId)
	} else {
		//newEdges := append(g.Edges[edge.FromNodeId], edge)
		//g.Edges[edge.FromNodeId] = newEdges
		mutex := &sync.Mutex{}
		if g.EdgeExists(toNodeId, fromNodeId) {
			mutex = g.GetEdge(toNodeId, fromNodeId).Mutex
		}
		newEdge := &Edge{FromNodeId: fromNodeId, ToNodeId: toNodeId, Attrs: attrs, Mutex: mutex}
		g.Edges[fromNodeId][toNodeId] = newEdge
		return nil
	}
}

func (g *Graph) LockEdge(nodeA int, nodeB int) {
	edge := g.GetEdge(nodeA, nodeB)
	edge.Mutex.Lock()
}

func (g *Graph) UnlockEdge(nodeA int, nodeB int) {
	edge := g.GetEdge(nodeA, nodeB)
	edge.Mutex.Unlock()
}

func (g *Graph) GetEdge(fromNodeId int, toNodeId int) *Edge {
	if g.EdgeExists(fromNodeId, toNodeId) {
		return g.Edges[fromNodeId][toNodeId]
	}
	return &Edge{}
}

func (g *Graph) GetEdgeData(fromNodeId int, toNodeId int) EdgeAttrs {
	if g.EdgeExists(fromNodeId, toNodeId) {
		return g.GetEdge(fromNodeId, toNodeId).Attrs
	}
	return EdgeAttrs{}
}

func (g *Graph) EdgeExists(fromNodeId int, toNodeId int) bool {
	if _, ok := g.Edges[fromNodeId][toNodeId]; ok {
		return true
	}
	return false
}

func (g *Graph) SetEdgeData(fromNodeId int, toNodeId int, edgeAttrs EdgeAttrs) bool {
	if g.EdgeExists(fromNodeId, toNodeId) {
		g.Edges[fromNodeId][toNodeId].Attrs = edgeAttrs
		return true
	}
	return false
}

// GetNode getNode will return a node point if exists or return nil
func (g *Graph) GetNode(nodeId int) *Node {
	//g.Mutex.Lock()
	//defer g.Mutex.Unlock()
	node, ok := g.NodesMap[nodeId]
	if ok {
		return node
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

func (g *Graph) Print() {
	for _, v := range g.Nodes {
		fmt.Printf("%d : ", v.Id)
		for _, i := range v.AdjIds {
			for _, v := range i {
				fmt.Printf("%d ", v)
			}
		}
		fmt.Println()
	}
}
