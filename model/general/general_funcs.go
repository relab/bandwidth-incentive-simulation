package general

import (
	"fmt"
	. "go-incentive-simulation/model/parts/types"
	"math/rand"
)

func Choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	for i := 0; i < k; i++ {
		res = append(res, nodes[rand.Intn(len(nodes))])
	}
	return res
}

func GetNodeById(nodeId int) *Node {
	nodes := &Node{}
	fmt.Println(nodes)

	var res *Node

	if nodes.Id == nodeId {
		res = nodes
	}
	fmt.Println(res)
	return res
}


