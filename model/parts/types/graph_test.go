package types

import (
	"testing"
)

func testContainsNode(t *testing.T) {

}

func testAddNode(t *testing.T) {

}

func testGetNode(t *testing.T) {
	// graph := Graph{}

}

func TestAddEdge(t *testing.T) {

	graph := &Graph{}

	edgeAttrs := EdgeAttrs{10, 20}

	edge1 := &Edge{FromNodeId: 1, ToNodeId: 2, Attrs: edgeAttrs}
	err1 := graph.AddEdge(edge1)
	if err1 != nil {
		t.Error("addEdge function returned an error message: ", err1)
	}

	edge2 := &Edge{FromNodeId: 1, ToNodeId: 2, Attrs: edgeAttrs}
	err2 := graph.AddEdge(edge2)
	if err2 == nil {
		t.Error("addedge should have returned an error for trying to add the same edge twice")
	}

	edge3 := &Edge{FromNodeId: 2, ToNodeId: 3, Attrs: edgeAttrs}
	err3 := graph.AddEdge(edge3)
	if err3 == nil {
		t.Error("addedge should have returned an error for trying to add the same edge twice")
	}

}

func TestGetEdgeData(t *testing.T) {
	// fromNodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// toNodes := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// graph := Graph{}
}
