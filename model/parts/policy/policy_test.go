package policy

//import (
//	. "go-incentive-simulation/model/state"
//	"testing"
//
//	"gotest.tools/assert"
//)
//
//const path = "../../../data/nodes_data_8_10000.txt"
//
//func TestResponisbleNodes(t *testing.T) {
//	//nodesId := []int{8190, 11683, 11211, 16935, 21020, 21725, 39525, 41162, 41471, 41852, 42779, 43377, 43812, 45280, 47987, 49693, 57841, 59951, 64132}
//	chunkAdd := 43000
//	state := MakeInitialState(path)
//	respNodes := state.Graph.FindResponsibleNodes(chunkAdd)
//
//	//values := findResponsibleNodes(chunkAdd)
//	assert.Equal(t, len(respNodes), 4)
//	for _, node := range respNodes {
//		distance := node ^ chunkAdd
//		assert.Assert(t, distance < 50)
//	}
//}
