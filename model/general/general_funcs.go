package general

import (
	"math/rand"
	"time"
)

func Choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	rand.Seed(time.Now().UnixMicro())

	for i := 0; i < k; i++ {
		res = append(res, nodes[rand.Intn(len(nodes))])
	}
	return res
}
