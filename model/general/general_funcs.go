package general

import "math/rand"

func Choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	for i := 0; i < k; i++ {
		res = append(res, nodes[rand.Intn(len(nodes))])
	}
	return res
}
