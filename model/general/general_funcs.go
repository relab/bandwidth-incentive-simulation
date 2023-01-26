package general

import (
	"math"
	"math/rand"
)

func Choice(nodes []int, k int) []int {
	res := make([]int, 0, k)

	var val int
	for i := 0; i < k; i++ {
		val = rand.Intn(len(nodes))
		res = append(res, nodes[val-1])
	}
	return res
}

func BitLength(num int) int {
	return int(math.Ceil(math.Log2(float64(num))))
}

func Contains[T comparable](elems []T, value T) bool {
	for _, item := range elems {
		if item == value {
			return true
		}
	}
	return false
}
